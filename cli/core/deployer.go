package core

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Omotolani98/k8ly/cli/caddy"
	"github.com/Omotolani98/k8ly/cli/config"
	"github.com/Omotolani98/k8ly/cli/deployer"
	"github.com/Omotolani98/k8ly/cli/registry"
)

type DeployInput struct {
	AppName      string
	Port         int
	Provider     string // docker|k8s|firecracker (overrides cfg.Provider)
	Push         bool
	RegistryHost string
	Repo         string
	Tag          string
}

// DeployApp builds and deploys the app using the selected item
func DeployApp(in DeployInput, cfg *config.K8lyConfig) error {
	// resolve provider precedence (flag overrides config)
	provider := cfg.Provider
	if in.Provider != "" {
		provider = in.Provider
	}

	// -------- Build Image --------
	imageName := fmt.Sprintf("%s-image", in.AppName)
	fmt.Println("🔨 Building image with Nixpacks...")
	if err := buildImage(imageName); err != nil {
		return fmt.Errorf("build failed: %w", err)
	}

	// -------- Push Image to Registry --------
	fullRef := imageName + ":latest"
	if in.Push && provider == "k8s" {
		ref, err := deployer.PushImage(deployer.PushOptions{
			ImageName:    imageName,
			RegistryHost: in.RegistryHost,
			Repo:         in.Repo,
			Tag:          in.Tag,
			Auth: registry.Auth{
				Username: os.Getenv("REG_USER"),
				Password: os.Getenv("REG_PASS"),
				Token:    os.Getenv("GITHUB_TOKEN"),
				Region:   os.Getenv("AWS_REGION"),
			},
		})
		if err != nil {
			return fmt.Errorf("push failed: %w", err)
		}
		fmt.Println("📤 Pushed image →", ref)
		fullRef = ref
		imageName = ref
	}

	// -------- Deploy to Runtime --------
	fmt.Printf("📦 Deploying using %s...\n", provider)
	switch provider {
	case "docker":
		if err := runDockerContainer(in.AppName, imageName, in.Port); err != nil {
			return err
		}
	case "firecracker":
		fmt.Println("🔥 Firecracker deployment is not implemented yet.")
		return nil
	case "k8s", "kubernetes":
		fmt.Println("☸️ Deploying to Kubernetes...")

		err := deployer.DeployToKubernetes(deployer.K8sDeployOptions{
			AppName: in.AppName,
			Port:    in.Port,
			Domain:  cfg.Domain,
			Image:   fullRef,
		})

		if err != nil {
			return fmt.Errorf("k8s deployment failed: %w", err)
		}

		return nil
	default:
		return fmt.Errorf("unsupported runtime: %s", provider)
	}

	if provider == "k8s" || provider == "kubernetes" {
		return nil // ingress handles routing
	}

	fmt.Println("🌐 Updating Caddy routing...")
	return caddy.AddReverseProxy(in.AppName, in.Port, cfg.Domain, cfg.Email)
}

func buildImage(imageName string) error {
	cmd := exec.Command("nixpacks", "build", ".", "--name", imageName+":latest")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func runDockerContainer(name, image string, port int) error {
	// Stop old container if exists
	_ = exec.Command("docker", "rm", "-f", name).Run()

	// Run new container
	cmd := exec.Command("docker", "run", "-d",
		"--name", name,
		"-p", fmt.Sprintf("%d:%d", port, port),
		image,
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

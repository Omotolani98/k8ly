package core

import (
  "fmt"
  "github.com/Omotolani98/k8ly/cli/config"
  "github.com/Omotolani98/k8ly/cli/caddy"
  "github.com/Omotolani98/k8ly/cli/deployer"
  "os"
  "os/exec"
)

// DeployApp builds and deploys the app using the selected item
func DeployApp(appName string, port int, vm string, cfg *config.K8lyConfig) error {
  provider := cfg.Provider
  if vm != "" {
    provider = vm
  }

  // Step 1: Build Image
  imageName := fmt.Sprintf("%s-image", appName)
  fmt.Println("🔨 Building image with Nixpacks...")
  if err := buildImage(imageName); err != nil {
    return fmt.Errorf("build failed: %w", err)
  }

  // Step 2: Deploy to Runtime
  fmt.Printf("📦 Deploying using %s...\n", provider)
  switch provider {
    case "docker":
      if err := runDockerContainer(appName, imageName, port); err != nil {
        return err
      }
    case "firecracker":
        fmt.Println("🔥 Firecracker deployment is not implemented yet.")
        return nil
    case "k8s", "kubernetes":
        fmt.Println("☸️ Deploying to Kubernetes...")

        err := deployer.DeployToKubernetes(deployer.K8sDeployOptions {
          AppName: appName,
          Port:    port,
          Domain:  cfg.Domain,
          Image:   imageName,
        })

        if err != nil {
          return fmt.Errorf("k8s deployment failed: %w", err)
        }

        // Skip Caddy for now if you plan to use Ingress
        return nil
    default:
        return fmt.Errorf("unsupported runtime: %s", provider)
  }

  // Step 3: Setup reverse proxy
  fmt.Println("🌐 Updating Caddy routing...")
  return caddy.AddReverseProxy(appName, port, cfg.Domain, cfg.Email)
}

func buildImage(imageName string) error {
  cmd := exec.Command("nixpacks", "build", ".", "--name", imageName + ":latest")
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

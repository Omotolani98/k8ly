package deployer

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type K8sDeployOptions struct {
	AppName string
	Port    int
	Domain  string // Optional, used for ingress
	Image   string // e.g. reqly-image
}

// DeployToKubernetes handles the deployment of an app to a Kubernetes cluster
func DeployToKubernetes(opts K8sDeployOptions) error {
	tmpDir := filepath.Join(os.TempDir(), "k8ly", opts.AppName)
	_ = os.MkdirAll(tmpDir, 0755)

	// 1. Render deployment.yaml
	if err := renderK8sTemplate("deployment.yaml.tmpl", filepath.Join(tmpDir, "deployment.yaml"), opts); err != nil {
		return err
	}

	// 2. Render service.yaml
	if err := renderK8sTemplate("service.yaml.tmpl", filepath.Join(tmpDir, "service.yaml"), opts); err != nil {
		return err
	}

	// 3. Optional ingress.yaml
	if opts.Domain != "" {
		if err := renderK8sTemplate("ingress.yaml.tmpl", filepath.Join(tmpDir, "ingress.yaml"), opts); err != nil {
			return err
		}
	}

	// 4. kubectl apply -f
	fmt.Println("🚀 Applying Kubernetes manifests...")
	cmd := exec.Command("kubectl", "apply", "-f", tmpDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd.Run()
}

// renderK8sTemplate renders a given YAML template file with values into output file
func renderK8sTemplate(templateFile string, outputFile string, data K8sDeployOptions) error {
  tmplContent, err := templatesFS.ReadFile("templates/" + templateFile)
  if err != nil {
	  return fmt.Errorf("failed to read template: %w", err)
  }

  tmpl, err := template.New(templateFile).Parse(string(tmplContent))
  if err != nil {
		return fmt.Errorf("failed to parse template: %w", err)
	}

	f, err := os.Create(outputFile)
	if err != nil {
		return fmt.Errorf("failed to create output file: %w", err)
	}
	defer f.Close()

	return tmpl.Execute(f, data)
}

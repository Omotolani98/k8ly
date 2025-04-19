package cmd

import (
	"fmt"

	"github.com/Omotolani98/k8ly/cli/config"
	"github.com/Omotolani98/k8ly/cli/core"
	"github.com/spf13/cobra"
)

var (
	appName      string
	port         int
	vm           string
	pushFlag     bool
	registryHost string
	repoName     string
	tagName      string
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy your backend app using Docker, Firecracker, or Kubernetes",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🚀 Deploying", appName)

		cfg, err := config.Load()
		if err != nil {
			fmt.Println("❌ Failed to load config:", err)
			return
		}

		in := core.DeployInput{
			AppName:      appName,
			Port:         port,
			Provider:     vm,
			Push:         pushFlag,
			RegistryHost: registryHost,
			Repo:         repoName,
			Tag:          tagName,
		}

		err = core.DeployApp(in, cfg)
		if err != nil {
			fmt.Println("❌ Deployment failed:", err)
		} else {
			fmt.Println("✅ App deployed successfully!")
			fmt.Printf("🌍 https://%s.%s\n", appName, cfg.Domain)
		}
	},
}

func init() {
	rootCmd.AddCommand(deployCmd)
	deployCmd.Flags().StringVar(&appName, "name", "", "App name (used for subdomain)")
	deployCmd.Flags().IntVar(&port, "port", 3000, "Port the app listens on")
	deployCmd.Flags().StringVar(&vm, "vm", "", "Override provider (docker|firecracker|k8s)")

	deployCmd.Flags().BoolVar(&pushFlag, "push", true, "Push image to registry (default true for k8s)")
	deployCmd.Flags().StringVar(&registryHost, "registry", "docker.io", "Registry host (docker.io, ghcr.io, …)")
	deployCmd.Flags().StringVar(&repoName, "repo", "", "Repository (e.g. user/myapp)")
	deployCmd.Flags().StringVar(&tagName, "tag", "latest", "Image tag")

	deployCmd.MarkFlagRequired("name")
	deployCmd.MarkFlagRequired("repo")
}

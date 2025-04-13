package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/Omotolani98/k8ly/cli/config"
  "github.com/Omotolani98/k8ly/cli/core"
)

var (
  appName string
  port    int
  vm      string
)

var deployCmd = &cobra.Command{
  Use: "deploy",
  Short: "Deploy your backend app using Docker, Firecracker, or Kubernetes",
  Run: func (cmd *cobra.Command, args []string) {
    fmt.Println("🚀 Deploying", appName)

    cfg, err := config.Load()
    if err != nil {
      fmt.Println("❌ Failed to load config:", err)
      return
    }

    err = core.DeployApp(appName, port, vm, cfg)
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
  deployCmd.MarkFlagRequired("name")
}


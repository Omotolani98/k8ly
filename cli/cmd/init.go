package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/Omotolani98/k8ly/cli/config"
  "github.com/Omotolani98/k8ly/cli/core"
)

// CLI Flags
var (
  domain string
  email string
  provider string
  hostMode bool
)

// InitCmd handles `k8ly init`
var initCmd = &cobra.Command{
  Use: "init",
  Short: "Initialize k8ly host environment",
  Long: `Bootstraps a new k8ly deployment host with:
      - Domain Configuration
      - TLS via Caddy
      - Runtime Provider selection (Docker, Firecracker or Firecracker)
      - Optional Token Setup`,
  Run: func(cmd *cobra.Command, args [] string) {
    cfg := config.K8lyConfig {
      Domain: domain,
      Email: email,
      Provider: provider,
      HostMode: hostMode,
    }

    fmt.Println("🚀 Initializing K8ly environment...")

    if err := core.InitializeHost(cfg); err != nil {
      fmt.Println("❌ Initialization failed: ", err)
      return
    }

    fmt.Println("✅ K8ly environment initialized successfully.")
    fmt.Println("🌐 Domain: ", cfg.Domain)
    fmt.Println("🔧 Provider: ", cfg.Provider)

    if cfg.HostMode {
      fmt.Println("🔐 Host mode: enabled (API token generated)")
    }
  },
}

func init() {
  rootCmd.AddCommand(initCmd)

  initCmd.Flags().StringVar(&domain, "domain", "", "Base domain for hosted tools (e.g. k8ly.dev)")
  initCmd.Flags().StringVar(&email, "email", "", "Email for Ler's Encrypt TLS")
  initCmd.Flags().StringVar(&provider, "provider", "docker", "Runtime to use: docker | firecracker | k8s")
  initCmd.Flags().BoolVar(&hostMode, "host", false, "Run as a k8ly host (generates API token)")
}

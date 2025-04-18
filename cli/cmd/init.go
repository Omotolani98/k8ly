package cmd

import (
	"fmt"

	"github.com/Omotolani98/k8ly/cli/config"
	"github.com/Omotolani98/k8ly/cli/core"
	"github.com/spf13/cobra"
)

var (
	domain   string
	email    string
	provider string
	hostMode bool
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the K8ly CLI version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("🧩 K8ly CLI version: %s\n", Version)
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initialize k8ly host environment",
	Long: `Bootstraps a new k8ly deployment host with:
      - Domain Configuration
      - TLS via Caddy
      - Runtime Provider selection (Docker, Firecracker or Firecracker)
      - Optional Token Setup`,
	Run: func(cmd *cobra.Command, args []string) {
		cfg := config.K8lyConfig{
			Domain:   domain,
			Email:    email,
			Provider: provider,
			HostMode: hostMode,
		}

		core.PrintSection("🚀 Initializing K8ly environment...")

		if err := core.InitializeHost(cfg); err != nil {
			fmt.Printf("❌ Initialization failed: %v\n", err)
			return
		}

		core.PrintSuccess("K8ly environment initialized successfully.")
		core.PrintSuccess("🌐 Domain: " + cfg.Domain)
		core.PrintSuccess("🔧 Provider: " + cfg.Provider)

		if cfg.HostMode {
			fmt.Println("🔐 Host mode: enabled (API token generated)")
		}
	},
}

func init() {
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(versionCmd)

	initCmd.Flags().StringVar(&domain, "domain", "", "Base domain for hosted tools (e.g. k8ly.dev)")
	initCmd.Flags().StringVar(&email, "email", "", "Email for Ler's Encrypt TLS")
	initCmd.Flags().StringVar(&provider, "provider", "docker", "Runtime to use: docker | firecracker | k8s")
	initCmd.Flags().BoolVar(&hostMode, "host", false, "Run as a k8ly host (generates API token)")
}

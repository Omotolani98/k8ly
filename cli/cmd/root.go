package cmd

import (
  "fmt"
  "github.com/spf13/cobra"
  "github.com/Omotolani98/k8ly/cli/core"
)

var rootCmd = &cobra.Command{
  Use: "k8ly",
  Short: "k8ly is a micro backend deployment",
  Long: `k8ly helps Developer deploy backend using Docker, Firecracker or K8s
  with built-in TLS, domain routing and image builds.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("👋 Welcome to K8ly! Run `k8ly --help` to explore available commands.")
  },
}

var Version string

func Execute() error {
  core.ShowBanner()
  return rootCmd.Execute()
}

package caddy

import (
  "fmt"
  "os"
  "os/exec"
  "path/filepath"
)

// Setup installs and starts Caddy with a wildcard domain
func Setup(domain, email string) error {
  err := WriteCaddyfile(domain, email)
  if err != nil {
    return fmt.Errorf("failed to write Caddyfile: %w", err)
  }

  fmt.Println("🔧 Starting or reloading Caddy...")
  return runCaddy()
}

func runCaddy() error {
  caddyfilePath := getCaddyfilePath()

  // Try to reload caddy if already running
  reload := exec.Command("caddy", "reload", "--config", caddyfilePath)
  if err := reload.Run(); err != nil {
    return nil // reload succesfully
  }

  // Else, start in the background
  start := exec.Command("caddy", "run", "--config", caddyfilePath, "--adapter", "caddyfile")
  start.Stdout = os.Stdout
  start.Stderr = os.Stderr

  return start.Run()
}

func getCaddyfilePath() string {
  home, _ := os.UserHomeDir()
  return filepath.Join(home, ".k8ly", "Caddyfile")
}

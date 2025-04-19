package caddy

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

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

	formatCaddyfile := exec.Command("caddy", "fmt", caddyfilePath, "--overwrite")
	formatCaddyfile.Stdout = os.Stdout
	formatCaddyfile.Stderr = os.Stderr

	if err := formatCaddyfile.Run(); err != nil {
		fmt.Println("✅ Caddyfile formatted successfully")
	}

	reload := exec.Command("caddy", "reload", "--config", caddyfilePath)
	reload.Stdout = os.Stdout
	reload.Stderr = os.Stderr

	if err := reload.Run(); err != nil {
		fmt.Println("🔄 Caddy reloaded successfully")
		return nil
	}

	check := exec.Command("pgrep", "caddy")
	if err := check.Run(); err == nil {
		fmt.Println("⚠️ Caddy is already running. Skipping 'caddy run'")
		return nil
	}

	fmt.Println("🚀 Starting Caddy server...")
	start := exec.Command("caddy", "run", "--config", caddyfilePath, "--adapter", "caddyfile")
	start.Stdout = os.Stdout
	start.Stderr = os.Stderr

	return start.Run()
}

func getCaddyfilePath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".k8ly", "Caddyfile")
}

package core

import (
  "fmt"
  "github.com/Omotolani98/k8ly/cli/config"
  "github.com/Omotolani98/k8ly/cli/caddy"
  "github.com/Omotolani98/k8ly/cli/utils"
)

// InitializeHost sets up the K8ly environment on the target machine
func InitializeHost (cfg config.K8lyConfig) error {
  // 1. Validate selected runtime
  fmt.Println("🔍 Validating runtime: ", cfg.Provider)
  if err := utils.EnsureRuntimeInstalled(cfg.Provider); err != nil {
    return fmt.Errorf("❌ runtime check failed: %w", err)
  }

  // 2. Setup reverse proxy (Caddy)
  fmt.Println("🌐 Setting up Caddy proxy for domain:", cfg.Domain)
  if err := caddy.Setup(cfg.Domain, cfg.Email); err != nil {
    return fmt.Errorf("❌ caddy setup failed: %w", err)
  }

  // 3. Save config to disk
  if err := config.Save(cfg); err != nil {
    return fmt.Errorf("❌ unable to save config: %w", err)
  }

  // 4. Optionally generate host token
  if cfg.HostMode {
    token := utils.GenerateToken()
    fmt.Println("🔐 Host Token: ", token)
    // Future: save token in a secure store (e.g., `~/.k8ly/token`)
  }

  return nil
}

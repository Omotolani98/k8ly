package config

import (
  "encoding/json"
  "fmt"
  "os"
  "path/filepath"
)

const (
  configDirName = ".k8ly"
  configFileName = "config.json"
)

type K8lyConfig struct {
  Domain   string  `json:"domain"`
  Email    string  `json:"email"`
  Provider string  `json:"provider"`
  HostMode bool    `json:"hostMode"`
}


// getConfigPath returns full path of config file
func getConfigPath() string {
  homeDir, err := os.UserHomeDir()
  if err != nil {
    panic("Unable to find home directory!")
  }
  return filepath.Join(homeDir, configDirName, configFileName)
}

// Save persists the config to ~/.k8ly/config.json
func Save(cfg K8lyConfig) error {
  configPath := getConfigPath()
  os.MkdirAll(filepath.Dir(configPath), os.ModePerm)

  data, err := json.MarshalIndent(cfg, "", " ")
  if err != nil {
    return fmt.Errorf("failed to marshal config: %w", err)
  }

  return os.WriteFile(configPath, data, 0644)
}

// Load reads the config disk
func Load() (*K8lyConfig, error) {
  configPath := getConfigPath()
  data, err := os.ReadFile(configPath)
  if err != nil {
    return nil, fmt.Errorf("failed to read config: %w", err)
  }

  var cfg K8lyConfig
  if err := json.Unmarshal(data, &cfg); err != nil {
    return nil, fmt.Errorf("failed to parse config: %w", err)
  }

  return &cfg, nil
}

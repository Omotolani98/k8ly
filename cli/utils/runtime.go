package utils

import (
  "fmt"
  "os/exec"
)

// EnsureRuntimeInstalled checks if the required CLI tool for the runtime is available
func EnsureRuntimeInstalled(provider string) error {
  var binary string

  switch provider {
    case "docker":
      binary = "docker"
    case "firecracker":
      binary = "firecracker"
    case "k8s", "kubernetes":
      binary = "kubectl"
    default:
      return fmt.Errorf("unsupported provider: %s", provider)
  }

  if _, err := exec.LookPath(binary); err != nil {
    return fmt.Errorf("%s is not installed or not in PATH", binary)
  }

  return nil
}

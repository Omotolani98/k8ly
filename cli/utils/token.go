package utils

import (
  "crypto/rand"
  "encoding/hex"
)

// GenerateToken returns a random 32-byte hex token
func GenerateToken() string {
  b := make([]byte, 16)
  _, _ = rand.Read(b)
  return hex.EncodeToString(b)
}

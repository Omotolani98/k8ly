package caddy

import (
  "fmt"
  "os"
  "path/filepath"
  "strings"
)

// WriteCaddyfile creates a basic wildcard domain config
func WriteCaddyfile(domain, email string) error {
  caddyfile := fmt.Sprintf(`*.%s {
    tls %s
    reverse_proxy * 127.0.0.1:3000
  }`, domain, email)

  home, _ := os.UserHomeDir()
  dir := filepath.Join(home, ".k8ly")
  os.MkdirAll(dir, os.ModePerm)

  path := filepath.Join(dir, "Caddyfile")
  return os.WriteFile(path, []byte(caddyfile), 0644)
}

func AddReverseProxy(app string, port int, domain, email string) error {
    caddyfilePath := getCaddyfilePath()

    contents, err := os.ReadFile(caddyfilePath)
    if err != nil {
        return fmt.Errorf("failed to read Caddyfile: %w", err)
    }

    fqdn := fmt.Sprintf("%s.%s", app, domain)

    // 🔐 Use `tls internal` for local dev environments
    tlsDirective := email
    if strings.HasPrefix(domain, "localhost") || strings.HasPrefix(domain, "127.") {
        tlsDirective = "internal"
    }

    block := fmt.Sprintf(`%s {
    tls %s
    reverse_proxy 127.0.0.1:%d
}
`, fqdn, tlsDirective, port)

    updated := upsertCaddyBlock(string(contents), fqdn, block)

    if err := os.WriteFile(caddyfilePath, []byte(updated), 0644); err != nil {
        return fmt.Errorf("failed to write updated Caddyfile: %w", err)
    }

    return runCaddy()
}

// upsertCaddyBlock replaces existing block or appends a new one
func upsertCaddyBlock(caddyfile, header, newBlock string) string {
    lines := strings.Split(caddyfile, "\n")
    var output []string
    insideBlock := false
    skip := false

    for i, line := range lines {
      fmt.Println(i)
        if strings.HasPrefix(line, header+" ") {
            insideBlock = true
            skip = true
        }

        if insideBlock && strings.TrimSpace(line) == "}" {
            insideBlock = false
            skip = false
            continue
        }

        if !skip {
            output = append(output, line)
        }
    }

    output = append(output, newBlock)
    return strings.Join(output, "\n")
}



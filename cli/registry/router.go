package registry

import (
	"fmt"
	"strings"
)

// New returns a concrete Registry driver based on the registry host.
func New(host string) (Registry, error) {
	switch {
	case host == "" || host == "docker.io":
		return &DockerHub{}, nil
	case strings.Contains(host, "ghcr.io"):
		return &GHCR{}, nil
	// case strings.Contains(host, "amazonaws.com"):
	// 	return &ECR{}, nil
	// TODO: add GCR/ACR
	default:
		return nil, fmt.Errorf("unsupported registry: %s", host)
	}
}

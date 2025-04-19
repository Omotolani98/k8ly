package deployer

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"

	"strings" // NEW

	"github.com/Omotolani98/k8ly/cli/registry"
)

// PushOptions holds everything needed to push a built image.
type PushOptions struct {
	ImageName    string // local name produced by nixpacks (e.g. myapi-image)
	RegistryHost string // docker.io, ghcr.io, 123.dkr.ecr.us-east-1.amazonaws.com …
	Repo         string // tolani/myapi   OR   k8ly/myapi
	Tag          string // latest, sha-abc123, v1.0.0
	Auth         registry.Auth
}

// pushImage builds, tags, and uploads the tarball to the chosen registry.
// It returns the full image reference (host/repo:tag) on success.
func PushImage(opts PushOptions) (string, error) {
	fullRef := fmt.Sprintf("%s/%s:%s", opts.RegistryHost, opts.Repo, opts.Tag)

	// 1) docker tag
	if err := exec.Command("docker", "tag",
		opts.ImageName+":latest", fullRef,
	).Run(); err != nil {
		return "", fmt.Errorf("docker tag failed: %w", err)
	}

	// 2) docker save → /tmp/xxxx.tar
	safeRepo := strings.ReplaceAll(opts.Repo, "/", "_")
	tmpTar := filepath.Join(os.TempDir(),
		fmt.Sprintf("%s-%d.tar", safeRepo, time.Now().Unix()))
	save := exec.Command("docker", "save", "-o", tmpTar, fullRef)
	save.Stdout = os.Stdout
	save.Stderr = os.Stderr
	if err := save.Run(); err != nil {
		return "", fmt.Errorf("docker save failed: %w", err)
	}
	defer os.Remove(tmpTar)

	// 3) pick driver
	driver, err := registry.New(opts.RegistryHost)
	if err != nil {
		return "", err
	}

	// 4) login + push
	ctx := context.Background()
	if err := driver.Login(ctx, opts.Auth); err != nil {
		return "", err
	}
	if err := driver.Push(ctx, tmpTar, fullRef); err != nil {
		return "", err
	}

	return fullRef, nil
}

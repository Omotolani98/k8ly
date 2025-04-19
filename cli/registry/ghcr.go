package registry

import (
	"context"
	"fmt"
	"os"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
)

type GHCR struct{}

func (g *GHCR) Login(_ context.Context, a Auth) error {
	if a.Token == "" {
		// allow token via env so CLI flag isn't required
		a.Token = os.Getenv("GITHUB_TOKEN")
	}
	if a.Token == "" {
		return fmt.Errorf("ghcr: missing token (set GITHUB_TOKEN env or --password)")
	}
	return nil // crane will inject auth later
}

func (g *GHCR) Push(ctx context.Context, tarPath, dest string) error {
	img, err := crane.Load(tarPath)
	if err != nil {
		return err
	}
	ref, err := name.ParseReference(dest)
	if err != nil {
		return err
	}
	auth := authn.FromConfig(authn.AuthConfig{
		Username: "ghcr", // any
		Password: os.Getenv("GITHUB_TOKEN"),
	})
	return crane.Push(img, ref.Name(), crane.WithAuth(auth))
}

package registry

import (
	"context"
	"fmt"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
)

type GHCR struct{}

func (g *GHCR) Login(_ context.Context, a Auth) error {
	if a.Token == "" {
		return fmt.Errorf("ghcr: missing token (GITHUB_TOKEN)")
	}
	// ghcr requires no explicit docker login when using crane, so no‑op.
	return nil
}

func (g *GHCR) Push(ctx context.Context, imageTar, dest string) error {
	img, err := crane.Load(imageTar)
	if err != nil {
		return err
	}
	ref, err := name.ParseReference(dest)
	if err != nil {
		return err
	}
	auth := authn.FromConfig(authn.AuthConfig{
		Username: "ghcr", // any non‑empty string
		Password: "",
	})
	return crane.Push(img, ref.Name(), crane.WithAuth(auth))
}

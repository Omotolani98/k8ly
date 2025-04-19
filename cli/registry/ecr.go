package registry

import (
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/google/go-containerregistry/pkg/authn"
	"github.com/google/go-containerregistry/pkg/crane"
	"github.com/google/go-containerregistry/pkg/name"
)

type ECR struct {
	password string
}

func (e *ECR) Login(ctx context.Context, a Auth) error {
	region := a.Region
	if region == "" {
		// extract region from host: 123456789012.dkr.ecr.us-east-1.amazonaws.com
		parts := strings.Split(a.Username, ".ecr.")
		if len(parts) > 1 {
			region = strings.Split(parts[1], ".")[0]
		}
		if region == "" {
			region = "us-east-1"
		}
	}
	out, err := exec.CommandContext(ctx, "aws", "ecr",
		"get-login-password", "--region", region).Output()
	if err != nil {
		return fmt.Errorf("ecr login: %w", err)
	}
	e.password = strings.TrimSpace(string(out))
	return nil
}

func (e *ECR) Push(ctx context.Context, tarPath, dest string) error {
	img, err := crane.Load(tarPath)
	if err != nil {
		return err
	}
	ref, err := name.ParseReference(dest)
	if err != nil {
		return err
	}
	auth := authn.FromConfig(authn.AuthConfig{
		Username: "AWS",
		Password: e.password,
	})
	return crane.Push(img, ref.Name(), crane.WithAuth(auth))
}

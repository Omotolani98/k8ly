package registry

import (
	"context"
	"fmt"
	"os/exec"
)

// DockerHub satisfies Registry for docker.io and any registry that accepts
// "docker login" + "docker push".
type DockerHub struct{}

func (d *DockerHub) Login(ctx context.Context, a Auth) error {
	if a.Username == "" || a.Password == "" {
		return fmt.Errorf("dockerhub: username/password required")
	}
	cmd := exec.CommandContext(ctx, "docker", "login", "docker.io",
		"-u", a.Username, "-p", a.Password)
	cmd.Stdout, cmd.Stderr = nil, nil
	return cmd.Run()
}

func (d *DockerHub) Push(ctx context.Context, imageTar, dest string) error {
	// 1) docker load < tar
	load := exec.CommandContext(ctx, "docker", "load", "-i", imageTar)
	if err := load.Run(); err != nil {
		return err
	}
	// 2) docker push dest
	push := exec.CommandContext(ctx, "docker", "push", dest)
	return push.Run()
}

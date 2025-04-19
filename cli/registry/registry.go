package registry

import "context"

// Auth holds credentials or tokens required by a registry.
type Auth struct {
	Username string
	Password string
	Token    string // e.g. GHCR PAT or ECR token
	Region   string // only for cloud-specific registries (ECR, ACR…)
}

// Registry is the interface all concrete drivers implement.
type Registry interface {
	Login(ctx context.Context, a Auth) error
	// Push receives the path to an image tarball and the destination ref
	//   e.g. gcr.io/myproj/myapp:sha-abc123
	Push(ctx context.Context, imageTar string, destRef string) error
}

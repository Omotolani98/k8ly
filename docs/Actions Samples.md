## GitHub Actions workflow templates for K8ly
> Copy a template into `<repo>/docs/<name>.yml` and adjust secrets.
---
### ☁️ Deploy to **Docker** on self‑hosted VM
```yaml
name: Deploy via K8ly (Docker)
on:
  push:
    branches: [ main ]
jobs:
  deploy:
    runs-on: ubuntu-latest
    steps:
    - uses: actions/checkout@v3
    # Install Nixpacks (for local build)
    - name: Install Nixpacks
      run: |
        curl -sSL https://nixpacks.com/install.sh | bash
    # Install K8ly CLI
    - name: Install K8ly
      run: |
        curl -sSL https://get.k8ly.dev | bash -- --quiet
    # Deploy to remote Docker host via SSH
    - name: Deploy with K8ly
      env:
        SSH_PRIVATE_KEY: ${{ secrets.SSH_PRIVATE_KEY }}  # your VM key
        SERVER_IP:       ${{ secrets.SERVER_IP }}        # x.x.x.x
      run: |
        echo "$SSH_PRIVATE_KEY" > key.pem
        chmod 600 key.pem
        export K8LY_REMOTE_HOST=$SERVER_IP
        export K8LY_REMOTE_USER=ubuntu
        export K8LY_SSH_KEY=$(pwd)/key.pem
        k8ly deploy
          --name api
          --port 3000
          --vm docker
```

---
### ☸️ Deploy to **Kubernetes** (image pushed to GHCR)
```yaml
name: Deploy via K8ly (Kubernetes + GHCR)
on:
  push:
    branches: [ main ]
env:
  REGISTRY: ghcr.io
  REPO:     ${{ github.repository }}   # e.g. user/myapp
  TAG:      ${{ github.sha }}
jobs:
  deploy:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write   # needed for GHCR push
    steps:
    - uses: actions/checkout@v3
    
    - name: Install Nixpacks
      run: curl -sSL https://nixpacks.com/install.sh | bash
    
    - name: Install K8ly
      run: curl -sSL https://get.k8ly.dev | bash -- --quiet

    - name: Configure kube‑context (EKS example)
      uses: aws-actions/eks-update-kubeconfig@v1
      with:
        cluster-name: ${{ secrets.EKS_CLUSTER }}
        region:       us-east-1

    - name: Deploy with K8ly (K8s)
      env:
        GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
      run: |
        k8ly deploy
          --name api
          --port 3000
          --vm k8s
          --push
          --registry $REGISTRY
          --repo $REPO
          --tag  $TAG
```

---

### Secret reference
| Secret | Needed for | Notes |
|--------|------------|-------|
| `SSH_PRIVATE_KEY` | Docker workflow | Private key for SSHing into self‑host VM |
| `SERVER_IP` | Docker workflow | Public IP of VM |
| `GITHUB_TOKEN` | GHCR push | Automatically provided in Actions, but must have `packages:write` permission |
| `AWS_*` or kube config secrets | K8s workflow | Depends on cluster provider |
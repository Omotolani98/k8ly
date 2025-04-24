# K8ly - Cloud‑Native DevInfra CLI

K8ly ( **k‑eight‑lee** ) is a lightweight, batteries‑included developer toolkit that lets you self‑host micro‑services on **Docker** or **Kubernetes** with a single command.

> *Firecracker support is on the roadmap.*

---

## ✨ Why K8ly?

| Feature | Why it matters |
|---------|----------------|
| **Zero‑config builds** | Uses [Nixpacks](https://github.com/railwayapp/nixpacks) to auto‑detect your stack and build an image. |
| **One‑line deploy** | `k8ly deploy --vm docker` or `--vm k8s` --- no YAML hand‑crafting. |
| **Built‑in TLS** | Seamless HTTPS via Caddy + Let's Encrypt. |
| **Registry push** | Auto‑pushes images to Docker Hub, GHCR, or ECR before K8s rollout. |
| **Extensible** | Modular drivers (Docker, K8s, Firecracker) & registry plug‑ins. |

---

## 🏃 Quick start

```bash

# 1. Install (Linux/macOS)
curl -sSL https://omotolani98.github.io/k8ly/selfhost.sh | bash

# 2. Initialise a project (one‑off)
k8ly init
  --domain myorg.dev
  --email you@example.com
  --provider docker      # default provider

# 3. Deploy an app (Docker)

k8ly deploy --name api --port 3000 --vm docker

# 3b. Deploy to your K8s cluster

auth $(aws eks get-token ...)   # or kube‑config already set

k8ly deploy
  --name api
  --port 3000
  --vm k8s
  --push
  --registry ghcr.io
  --repo youruser/api
  --tag $(git rev-parse --short HEAD)

```

The app will be reachable at `https://api.myorg.dev` within seconds 🍃.

---

## 🔧 Commands & Flags

### `k8ly init`

Initialises a workspace and creates `~/.k8ly/config.json` plus a starter **Caddyfile**.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--domain` | string | *(required)* | Base domain you control (e.g. `myorg.dev`) |
| `--email`  | string | *(required)* | Email for Let's Encrypt certs |
| `--provider` | enum `docker\|k8s\|firecracker` | `docker` | Default runtime when `--vm` isn't supplied on deploy |
| `--host` | bool | `false` | Write a starter Caddyfile in `~/.k8ly/` |
---

### `k8ly deploy`

Builds the image, optionally pushes to a registry, then deploys to Docker/K8s.

| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `--name` | string | *(required)* | Logical service name (becomes sub‑domain & Deployment name) |
| `--port` | int | `3000` | Container port the app listens on |
| `--vm` | enum | from `init` | Override runtime for this deploy (`docker`, `k8s`, `firecracker`) |
| `--push` | bool | `true` for K8s, `false` for Docker | Push image to registry before K8s rollout |
| `--registry` | string | `docker.io` | Registry host (`docker.io`, `ghcr.io`, `*.amazonaws.com`) |
| `--repo` | string | *(required if --push)* | Repository path (e.g. `tolani/myapi`) |
| `--tag` | string | `latest` | Image tag (git SHA, semver, etc.) |
---

## 🌍 Registry Matrix

| Driver | Auth expectations | Notes |
|--------|-------------------|-------|
| Docker Hub | `REG_USER` + `REG_PASS` env vars (or flags) | Uses `docker login` then `docker push`. |
| GHCR | `GITHUB_TOKEN` env or `--password` flag | PAT with `write:packages`. No docker login required when pushing via crane. |
| AWS ECR | AWS creds in env / `aws configure`; `--registry` must be full ECR host | Token fetched with `aws ecr get-login-password`. |
| *Firecracker* | _(coming soon)_ | --- |
---

## 🛠 Roadmap
- [ ] Firecracker micro‑VM provider
- [ ] Self‑host bootstrap script (`get.k8ly.dev`)
- [ ] CLI `k8ly logs`, `k8ly destroy`
- [ ] Multi‑arch manifest push

> *Found a bug or want a feature?* Open an issue or PR --- we 💙 contributors.

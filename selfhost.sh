#!/usr/bin/env bash
# Self‑host bootstrap for a fresh Ubuntu 22.04/20.04 VM
# Installs Docker, Caddy (systemd), and the latest K8ly CLI.
# Safe to run repeatedly – every step is idempotent.

set -euo pipefail
GREEN="\033[1;32m"; RESET="\033[0m"

msg() { printf "${GREEN}➡ %s${RESET}\n" "$1"; }

# 1. Update apt & basic deps
msg "Updating apt cache…"
sudo apt-get update -qq
sudo apt-get install -yqq curl ca-certificates gnupg lsb-release

# 2. Docker (official convenience script)
if ! command -v docker &>/dev/null; then
  msg "Installing Docker…"
  curl -fsSL https://get.docker.com | sh -
  sudo usermod -aG docker "$USER"
else
  msg "Docker already installed – skipping."
fi

# 3. Caddy (APT repo)
if ! command -v caddy &>/dev/null; then
  msg "Installing Caddy…"
  sudo apt-get install -yqq debian-keyring debian-archive-keyring
  curl -1sLf "https://dl.cloudsmith.io/public/caddy/stable/gpg.key" | sudo tee /usr/share/keyrings/caddy-stable-archive-keyring.gpg >/dev/null
  echo "deb [signed-by=/usr/share/keyrings/caddy-stable-archive-keyring.gpg] https://dl.cloudsmith.io/public/caddy/stable/deb/$(lsb_release -cs) $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/caddy-stable.list
  sudo apt-get update -qq && sudo apt-get install -yqq caddy
else
  msg "Caddy already installed – skipping."
fi

# 4. K8ly latest release
K8LY_BIN="/usr/local/bin/k8ly"
LATEST_URL="https://omotolani98.github.io/k8ly/install.sh"
msg "Installing K8ly CLI…"
if [[ -x "$K8LY_BIN" ]]; then
  INSTALLED_VER=$(k8ly version 2>/dev/null || true)
  msg "K8ly $INSTALLED_VER already present – replacing with latest…"
fi
curl -sSL "$LATEST_URL" | sudo bash

# 5. Create K8ly config dir if missing
CONFIG_DIR="$HOME/.k8ly"
mkdir -p "$CONFIG_DIR"

msg "Bootstrap complete!  Re‑login or run 'newgrp docker' to use Docker without sudo."
msg "Run 'k8ly help' to get started."
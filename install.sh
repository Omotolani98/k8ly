#!/bin/bash

set -e

REPO="Omotolani98/k8ly"
VERSION="${1:-latest}"

echo "🔍 Detecting platform..."
ARCH=$(uname -m)
OS=$(uname -s | tr '[:upper:]' '[:lower:]')

case "$ARCH" in
  x86_64) ARCH="amd64" ;;
  aarch64 | arm64) ARCH="arm64" ;;
  *) echo "❌ Unsupported architecture: $ARCH"; exit 1 ;;
esac

if [ "$VERSION" = "latest" ]; then
  VERSION=$(curl -s https://api.github.com/repos/${REPO}/releases/latest | grep tag_name | cut -d '"' -f 4)
fi

TARBALL="k8ly_${VERSION}_${OS}_${ARCH}.tar.gz"
URL="https://github.com/${REPO}/releases/download/${VERSION}/${TARBALL}"

echo "⬇️ Downloading $TARBALL..."
curl -L "$URL" -o "$TARBALL"

echo "📦 Extracting..."
tar -xzf "$TARBALL"

INSTALL_PATH="/usr/local/bin"
FALLBACK_PATH="$HOME/.local/bin"

if command -v sudo >/dev/null; then
  echo "🚚 Installing to $INSTALL_PATH..."
  chmod +x k8ly
  sudo mv k8ly $INSTALL_PATH/k8ly || {
    echo "⚠️ Failed to move to $INSTALL_PATH. Falling back to $FALLBACK_PATH"
    mkdir -p "$FALLBACK_PATH"
    mv k8ly "$FALLBACK_PATH/k8ly"
    echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
  }
else
  echo "⚠️ sudo not found. Installing to $FALLBACK_PATH"
  mkdir -p "$FALLBACK_PATH"
  mv k8ly "$FALLBACK_PATH/k8ly"
  echo 'export PATH="$HOME/.local/bin:$PATH"' >> ~/.bashrc
fi

echo "🧹 Cleaning up..."
rm "$TARBALL"

echo "✅ K8ly $VERSION installed!"
echo "👉 Restart your shell or run 'source ~/.bashrc' if needed"

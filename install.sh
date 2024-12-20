#!/bin/bash

VERSION="v1.9.0"
BINARY_NAME="memepulation"

OS="$(uname | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"

if [[ "$ARCH" == "x86_64" ]]; then
  ARCH="amd64"
elif [[ "$ARCH" == "arm64" || "$ARCH" == "aarch64" ]]; then
  ARCH="arm64"
else
  echo "Unsupported architecture: $ARCH"
  exit 1
fi

DOWNLOAD_URL="https://github.com/sharon-xa/memepulation/releases/download/$VERSION/$BINARY_NAME-$OS-$ARCH"

echo "Downloading $BINARY_NAME for $OS/$ARCH from $DOWNLOAD_URL..."
curl -L -o /usr/local/bin/$BINARY_NAME "$DOWNLOAD_URL"

chmod +x /usr/local/bin/$BINARY_NAME


# Check user shell and setup autocompletion
shell_name=$(basename "$SHELL")

case "$shell_name" in
  bash)
      source <(memepulation completion bash)
    ;;
  zsh)
      source <(memepulation completion zsh)
    ;;
  fish)
      memepulation completion fish | source
    ;;
  *)
    echo "Brother, what type of shell are you using?"
    exit 1
    ;;
esac


if command -v $BINARY_NAME > /dev/null 2>&1; then
  echo "$BINARY_NAME installed successfully!"
else
  echo "Installation failed."
  exit 1
fi

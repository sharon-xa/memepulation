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
    memepulation completion bash > /etc/bash_completion.d/memepulation
    source /etc/bash_completion.d/memepulation
    ;;
  zsh)
    fpath=(/usr/local/share/zsh/site-functions "$fpath")
    memepulation completion zsh > "${fpath[1]}/_memepulation"
    source "${fpath[1]}/_memepulation"
    ;;
  fish)
    memepulation completion fish > ~/.config/fish/completions/memepulation.fish
    # shellcheck source=/dev/null
    source ~/.config/fish/completions/memepulation.fish
    ;;
  *)
    echo "Brother, what type of shell are you using? Unsupported shell: $shell_name"
    exit 1
    ;;
esac

if command -v $BINARY_NAME > /dev/null 2>&1; then
  echo "$BINARY_NAME installed successfully!"
else
  echo "Installation failed."
  exit 1
fi

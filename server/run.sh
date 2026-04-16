#!/usr/bin/env bash
set -euo pipefail

BINARY_NAME="prompt-library"
if [ "$(uname -s)" = "Windows_NT" ]; then
  BINARY_NAME="prompt-library.exe"
fi

DATA_DIR="${PLUGIN_DATA:-.}"
BINARY_PATH="${DATA_DIR}/${BINARY_NAME}"

# If binary already exists, just run it
if [ -f "$BINARY_PATH" ]; then
  exec "$BINARY_PATH"
fi

# Try to find it in PATH (installed via brew/scoop)
if command -v prompt-library &>/dev/null; then
  exec prompt-library
fi

# Download from GitHub releases
REPO="ohare93/prompt-library"
OS="$(uname -s | tr '[:upper:]' '[:lower:]')"
ARCH="$(uname -m)"
case "$ARCH" in
  x86_64)  ARCH="amd64" ;;
  aarch64) ARCH="arm64" ;;
esac

ARCHIVE_NAME="prompt-library_${OS}_${ARCH}"
if [ "$OS" = "windows" ]; then
  EXT="zip"
else
  EXT="tar.gz"
fi
URL="https://github.com/${REPO}/releases/latest/download/${ARCHIVE_NAME}.${EXT}"

TMPDIR="$(mktemp -d)"
trap 'rm -rf "$TMPDIR"' EXIT

echo "Downloading ${BINARY_NAME} from GitHub releases..." >&2

if ! curl -fsSL "$URL" -o "${TMPDIR}/${ARCHIVE_NAME}.${EXT}"; then
  echo "Failed to download. Install manually:" >&2
  echo "  brew install ohare93/tap/prompt-library" >&2
  echo "  scoop bucket add ohare93 https://github.com/ohare93/scoop && scoop install prompt-library" >&2
  exit 1
fi

cd "$TMPDIR"
if [ "$EXT" = "zip" ]; then
  unzip -o "${ARCHIVE_NAME}.${EXT}" "$BINARY_NAME"
else
  tar xzf "${ARCHIVE_NAME}.${EXT}" "$BINARY_NAME"
fi

chmod +x "$BINARY_NAME"
mkdir -p "$(dirname "$BINARY_PATH")"
mv "$BINARY_NAME" "$BINARY_PATH"

exec "$BINARY_PATH"

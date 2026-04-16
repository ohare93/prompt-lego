#!/usr/bin/env bash
set -euo pipefail

VENV_DIR="${PLUGIN_DATA:-.}/.venv"

if [ ! -f "$VENV_DIR/bin/python" ]; then
  python3 -m venv "$VENV_DIR"
  "$VENV_DIR/bin/pip" install --quiet 'mcp[cli]'
fi

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"

exec "$VENV_DIR/bin/python" "$SCRIPT_DIR/prompt_library_mcp.py"

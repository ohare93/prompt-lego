#!/usr/bin/env python3
"""MCP server that exposes prompt directories as @ mention resources in Claude Code.

Reads directories from:
  1. PROMPT_DIRS env var (comma-separated paths)
  2. Command-line arguments
  3. ./prompts/ in cwd (always, if it exists)

For each directory:
  1. Looks for a prompts/ subdirectory; falls back to the directory itself
  2. Derives the project name from the nearest .git parent
  3. Registers each .md file as @prompts:<project-name>://<stem>
"""

import os
import sys
from pathlib import Path

from mcp.server.fastmcp import FastMCP
from mcp.server.fastmcp.resources import FunctionResource


def find_project_name(directory: Path) -> str:
    """Walk up from directory to find .git, return that folder's name."""
    current = directory
    while current != current.parent:
        if (current / ".git").exists():
            return current.name
        current = current.parent
    return directory.name


def find_prompts_dir(directory: Path) -> Path:
    """Return prompts/ subdirectory if it exists, otherwise the directory itself."""
    prompts = directory / "prompts"
    if prompts.is_dir():
        return prompts
    return directory


def scan_prompts(directory: Path) -> list[Path]:
    return sorted(p for p in directory.glob("*.md") if p.is_file())


def extract_title(path: Path) -> str:
    first_line = path.read_text().split("\n", 1)[0].strip()
    if first_line.startswith("# "):
        return first_line[2:]
    return path.stem.replace("-", " ").title()


def extract_description(path: Path) -> str:
    for line in path.read_text().split("\n")[1:]:
        stripped = line.strip()
        if stripped and not stripped.startswith("#"):
            return stripped[:200]
    return ""


def _make_reader(p: Path):
    def read() -> str:
        return p.read_text()
    return read


# Collect directories from all sources
dirs: list[Path] = []

# From PROMPT_DIRS env var (comma-separated)
env_dirs = os.environ.get("PROMPT_DIRS", "")
if env_dirs:
    dirs.extend(Path(d.strip()).resolve() for d in env_dirs.split(",") if d.strip())

# From command-line arguments
dirs.extend(Path(d).resolve() for d in sys.argv[1:])

# Always include cwd if it has a prompts/ directory
cwd = Path.cwd().resolve()
if cwd not in dirs and (cwd / "prompts").is_dir():
    dirs.insert(0, cwd)

# Deduplicate while preserving order
dirs = list(dict.fromkeys(dirs))

mcp = FastMCP("prompts")

seen_uris = set()
for directory in dirs:
    if not directory.is_dir():
        continue
    project = find_project_name(directory)
    prompts_dir = find_prompts_dir(directory)
    for path in scan_prompts(prompts_dir):
        uri = f"{project}://{path.stem}"
        if uri in seen_uris:
            continue
        seen_uris.add(uri)
        mcp.add_resource(
            FunctionResource(
                uri=uri,
                name=extract_title(path),
                description=extract_description(path),
                mime_type="text/markdown",
                fn=_make_reader(path),
            )
        )

if __name__ == "__main__":
    mcp.run()

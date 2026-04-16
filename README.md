# Prompt Library

An MCP server that serves markdown files as `@` mention resources in Claude Code.

Type `@prompts:my-project://tdd` in Claude Code and the full prompt content expands into your conversation — just like `@file`, but for reusable prompts.

## Install

### Homebrew (macOS/Linux)

```bash
brew install ohare93/tap/prompt-library
```

### Scoop (Windows)

```bash
scoop bucket add ohare93 https://github.com/ohare93/scoop
scoop install prompt-library
```

### GitHub Releases

Download the latest binary from [releases](https://github.com/ohare93/prompt-library/releases).

## Usage

### As a standalone MCP server

Add it to your Claude Code MCP settings (`.claude/settings.json` or project-level `.claude/settings.local.json`):

```json
{
  "mcpServers": {
    "prompts": {
      "command": "prompt-library",
      "args": [],
      "env": {
        "PROMPT_DIRS": "/home/me/shared-prompts,/home/me/work-prompts"
      }
    }
  }
}
```

Or with absolute path if not in PATH:

```json
{
  "mcpServers": {
    "prompts": {
      "command": "/usr/local/bin/prompt-library",
      "args": [],
      "env": {
        "PROMPT_DIRS": "/home/me/shared-prompts"
      }
    }
  }
}
```

### As a Claude Code plugin

```bash
claude plugin install prompt-library@<marketplace>
```

The plugin asks for `prompt_dirs` at install time and automatically downloads the binary on first run.

## How it works

The server scans for `.md` files in two ways:

1. **Shared libraries** — directories you set in `PROMPT_DIRS` (comma-separated)
2. **Per-project** — any `prompts/` folder in your current working directory is picked up automatically

The project name (used in the URI) comes from the nearest `.git` parent:

```
my-project/           <- .git here, so project = "my-project"
  prompts/
    tdd.md            -> @prompts:my-project://tdd
    code-review.md    -> @prompts:my-project://code-review
```

Multiple directories work — each gets its own namespace:

```
PROMPT_DIRS=/home/me/shared-prompts,/home/me/work-prompts

@prompts:shared-prompts://caveman
@prompts:work-prompts://deploy-checklist
```

## Prompt file format

Each `.md` file is a self-contained prompt. The first `# Heading` becomes the title, the first content line becomes the description in autocomplete.

```markdown
# Test-Driven Development

Follow red-green-refactor strictly. No implementation code without a failing test first.

## Workflow

1. **Red**: Write a failing test...
```

## Building from source

Requires Go 1.22+.

```bash
go build -o prompt-library ./cmd/prompt-library
./prompt-library --version
```

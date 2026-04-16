# Prompt Library

A Claude Code plugin that serves markdown files as `@` mention resources.

Type `@prompts:my-project://tdd` in Claude Code and the full prompt content expands into your conversation — just like `@file`, but for reusable prompts.

## Install

```bash
claude plugin install prompt-library@<marketplace>
```

When you enable the plugin, it asks for `prompt_dirs` — a comma-separated list of directories containing your shared prompt collections.

## How it works

The plugin scans for `.md` files in two ways:

1. **Shared libraries** — directories you configure in `prompt_dirs` at install time
2. **Per-project** — any `prompts/` folder in your current working directory is picked up automatically

The project name (used in the URI) comes from the nearest `.git` parent:

```
my-project/           ← .git here, so project = "my-project"
  prompts/
    tdd.md            → @prompts:my-project://tdd
    code-review.md    → @prompts:my-project://code-review
```

Multiple directories work — each gets its own namespace:

```
prompt_dirs = /home/me/shared-prompts, /home/me/work-prompts

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

## Requirements

Python 3.10+ (the plugin creates its own venv automatically).

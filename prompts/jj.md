# Version Control: jj (Jujutsu)

Always use jj instead of git. Never run raw git commands.

## Quick Reference

| Task | Command |
|------|---------|
| Status | `jj st` |
| Log | `jj log` |
| Diff | `jj diff` |
| Commit | `jj commit -m "msg"` |
| Describe | `jj describe -m "msg"` |
| New commit | `jj new` or `jj new main -m "feat: ..."` |
| Edit old commit | `jj edit <change-id>` |
| Squash into parent | `jj squash` |
| Squash specific files | `jj squash <path>...` |
| Commit specific files | `jj commit <path>... -m "msg"` |
| Undo | `jj undo` |
| Push | `jj git push` |
| Fetch | `jj git fetch` |
| Create bookmark | `jj bookmark create <name>` |
| Move bookmark | `jj bookmark move <name>` |

## Key Concepts

- Working copy IS a commit. No staging area. Edits are auto-tracked.
- `jj new` means "I'm done with this commit, start fresh."
- Bookmarks are optional labels on commits, not branches.
- Editing old commits auto-rebases descendants.
- Conflicts are stored in commits, not blocking.
- One-liner commit messages, no multi-line bodies.

## Nix Integration

Snapshot before nix commands (`jj new` or `jj commit -m "wip"`) so new files appear in git HEAD.

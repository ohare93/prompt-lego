# Tidy Worktree

Clean up the working copy. Commit what matters, discard what doesn't.

1. Run `jj st` and `jj diff` to see every uncommitted change.
2. Read each changed file. Understand what the change does and why it exists.
3. Group related changes into logical commits:
   - Finished work gets committed with a clear one-liner message.
   - Partial work that's worth keeping gets committed as `wip: <what it is>`.
4. Discard changes that are noise — debug prints, scratch files, accidental edits. Use `jj restore <path>` to revert them.
5. Delete untracked files that serve no purpose (temp files, build artifacts, editor backups).
6. Run `jj st` again. The working copy should be clean or contain only intentional in-progress work.
7. Run `jj log` to verify the new commits look right.

## Rules

- Never commit secrets, credentials, or `.env` files.
- Never silently discard changes you don't understand. Ask first.
- One concern per commit. Don't lump unrelated changes together.
- If unsure whether something is needed, ask — don't guess.

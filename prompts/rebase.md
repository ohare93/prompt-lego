# Rebase

Fetch upstream changes and rebase your work on top.

1. Run `jj git fetch` to pull the latest from the remote.
2. Run `jj rebase -d main` to move your working changes on top of the updated main.
3. If there are conflicts, resolve them one file at a time. Read both sides before choosing.
4. Run `jj st` to confirm the working copy is clean and the rebase is complete.
5. Run the test suite. Everything must still pass after the rebase.

## Rules

- Never skip conflict resolution. If a conflict is unclear, read the surrounding code to understand intent.
- Do not force-push without being asked to.
- If the rebase introduces test failures unrelated to your changes, stop and report them.

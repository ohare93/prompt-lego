# Verify

Run every check before committing. No exceptions.

1. Run the test suite. All tests must pass.
2. Run the type checker if the project has one. No type errors.
3. Run the linter if the project has one. No lint errors.
4. If the change is user-facing, run the application and verify the behavior manually.
5. If any check fails, fix the issue and re-run all checks from step 1.

## Rules

- Do not commit with known failures, even if they're "unrelated."
- Do not skip a check because you're confident. Run it anyway.
- If a check doesn't exist yet but should, that's a separate task — don't block the current one, but note it.

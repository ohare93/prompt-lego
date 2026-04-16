# Testing

Run all tests after every change. Fix failures before moving on.

## Process

1. Run the project's test suite before making changes (baseline).
2. Make your change.
3. Run tests again. All must pass.
4. If a test fails, fix the code (not the test) unless the test is genuinely wrong.

## Writing Tests

- Cover the happy path, error paths, and edge cases.
- Tests should be independent. No shared mutable state between tests.
- Use descriptive names: `TestResolveArgs_bare_name_resolves_from_JUGGLE_PROMPTS`.
- Prefer real implementations over mocks. Use mocks only for external services or slow I/O.
- Use `t.TempDir()` for filesystem tests. Use `t.Setenv()` for environment variables.
- Table-driven tests for parameterized cases.

## What Not To Do

- Don't skip failing tests. Don't mark them as expected failures without a plan.
- Don't test private implementation details that may change.
- Don't write tests that depend on execution order.
- Don't add sleep-based synchronization. Use channels, waitgroups, or polling.

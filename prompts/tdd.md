# Test-Driven Development

Follow red-green-refactor strictly. No implementation code without a failing test first.

## Workflow

1. **Red**: Write a failing test for the next behavior. Run it. Confirm it fails.
2. **Green**: Write the minimum code to make the test pass. Nothing more.
3. **Refactor**: Clean up duplication and improve design. Tests must stay green.
4. Repeat.

## Rules

- Never write implementation before the test exists and fails.
- Each cycle adds exactly one behavior. Keep tests small and focused.
- Run the full test suite after each green step to catch regressions.
- If you're unsure what to test next, write the simplest test that forces new code.
- Tests are production code. Keep them clean, readable, and fast.

## Test Quality

- Test behavior, not implementation details.
- One assertion per test when possible. Name tests after what they verify.
- Use table-driven tests for variant inputs.
- Avoid mocks unless the real dependency is slow, non-deterministic, or external.
- Edge cases (empty input, zero, nil, boundary values) deserve their own tests.

## When Stuck

- If a test is hard to write, the design is wrong. Simplify the interface.
- If you need many mocks, you have too many dependencies. Extract and inject.
- If tests are brittle, you're testing implementation. Test the contract instead.

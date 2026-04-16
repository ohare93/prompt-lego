# Code Review

Review the code thoroughly. Be specific, be honest, explain your reasoning.

## Process

1. Identify the scope. Determine which files and changes to review (diff, PR, or specified files).
2. Read every changed line in context. Understand what it does and how it fits the surrounding code.
3. Run the test suite if one exists. Note whether tests pass, fail, or are missing.
4. Write findings grouped by category (see below). For each finding:
   - Quote the exact code (file, line number, snippet).
   - Explain what's wrong or what could be better and why.
   - Suggest a concrete fix or alternative when possible.
5. Summarize: overall assessment, the most important issues, and whether the change is ready to merge.

## What to Look For

### Correctness
- Logic errors, off-by-one mistakes, unhandled edge cases.
- Race conditions, deadlocks, or unsafe concurrent access.
- Incorrect error handling — swallowed errors, wrong error types, missing propagation.

### Security
- Injection vulnerabilities (SQL, command, XSS, path traversal).
- Hardcoded secrets, credentials, or tokens.
- Missing input validation or sanitization at trust boundaries.
- Insecure defaults, overly broad permissions, or exposed internals.

### Design
- Does the change belong here, or should it live elsewhere?
- Are abstractions at the right level — not too broad, not too specific?
- Are responsibilities clear? Does each function/type do one thing?
- Could this be simpler? Unnecessary indirection, premature abstraction, dead code.

### Readability
- Unclear names — variables, functions, types that don't say what they mean.
- Missing context — would a reader understand this without the PR description?
- Overly clever code that trades clarity for brevity.
- Inconsistency with the project's existing style or conventions.

### Testing
- Are the new behaviors tested? Are edge cases covered?
- Do tests verify behavior or implementation details?
- Are tests readable — clear setup, obvious assertions, descriptive names?
- Are there missing failure-path tests?

### Performance
- Unnecessary allocations, redundant computation, or N+1 queries.
- Operations that scale poorly with input size.
- Missing caching, batching, or pagination where data volume warrants it.

## Tone

- Be direct. "This will panic on nil input" not "you might want to consider the nil case."
- Distinguish must-fix issues from suggestions. Label them: **[must fix]**, **[suggestion]**, **[nit]**.
- Give credit where the code is well-written. Note good patterns so they get repeated.
- Explain the why, not just the what. A reviewer who only says "change this" teaches nothing.

# Task Breakdown

Break the work into the smallest possible units. One change per task.

1. Read the requirements, PRD, or issue in full.
2. List every distinct behavior, endpoint, schema change, or UI element that needs to exist.
3. Split each item until it can be completed in a single commit. If a task touches more than two or three files, it's probably too big.
4. Word each task as a concrete action: "add X", "wire Y to Z", "validate W". Avoid vague tasks like "set up backend".
5. Order the tasks so each one builds on the last. No task should depend on work that comes later in the list.
6. Review the list. If any task requires you to hold more than one idea in your head at once, split it further.

## What Not To Do

- Don't plan tasks you can't verify. Every task should have an observable result — a passing test, a working endpoint, a visible change.
- Don't batch unrelated changes into one task because they're small.
- Don't create tasks for hypothetical problems. Only plan work the requirements demand.

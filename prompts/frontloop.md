# Frontloop Task Queue

Use the frontloop skill to manage work. Invoke `/frontloop:status` to see the queue before starting.

## Commands

- `/frontloop:status` — show current queue state
- `/frontloop:work` — pick up the next ready task and execute it
- `/frontloop:clarify` — review tasks needing human input
- `/frontloop:add` — create a new task
- `/frontloop:init` — set up `.frontloop/` directories (first time only)

## Rules

- Check status before picking up work.
- Work one task at a time. Finish or block it before starting another.
- Follow acceptance criteria and design decisions exactly — they were pre-approved.
- If blocked, move the task back to clarify with an explanation.
- Commit after completing each task.

# Subagent

Use subagents for independent work. Keep the main context clean.

1. Before starting, identify which parts of the task are independent of each other.
2. Dispatch independent pieces to subagents. Give each agent a clear, self-contained instruction — what to do, which files to touch, what the output should look like.
3. Run independent subagents in parallel when possible.
4. Review subagent results before continuing. Don't trust blindly — verify the output makes sense.

## When To Use Subagents

- Exploring unfamiliar parts of the codebase — send agents to read and summarize instead of loading everything into main context.
- Running independent implementation tasks that don't share state.
- Research tasks — finding patterns, checking conventions, reading documentation.
- Any work that would bloat the main context with information you only need temporarily.

## When Not To Use Subagents

- When tasks depend on each other's output. Do those sequentially in main context.
- When the task is small enough that dispatching an agent costs more than just doing it.
- When you need to hold the full picture in one place to make a decision.

## Rules

- Every subagent gets a complete prompt. Don't assume it knows what you know.
- State whether the agent should write code or just research. Be explicit.
- One concern per agent. Don't overload a single agent with unrelated tasks.

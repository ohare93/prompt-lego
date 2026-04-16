# Prioritize

Pick the next task. Use this order:

1. **Critical bugfixes** — Anything broken in production or blocking other work.
2. **Dev infrastructure** — Tests, types, linting, dev scripts, CI. These are preconditions for building features safely.
3. **Tracer bullets** — A tiny end-to-end slice of the next feature. Go through all layers, prove the approach works, then expand.
4. **Feature work** — Build out the rest of the feature, one task at a time.
5. **Polish and quick wins** — Small UX improvements, error messages, edge case handling.
6. **Refactors** — Only when the code actively resists the next change. Not for aesthetics.

## Rules

- Pick exactly one task. Finish it before picking the next.
- If two tasks have the same priority, pick the one that unblocks the most future work.
- If nothing is left, stop. Don't invent work.

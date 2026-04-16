# Tracer Bullet

Build a tiny, end-to-end slice of the feature first. Then expand it.

1. Identify the thinnest path through all layers of the system — input to output, request to response, UI to database and back.
2. Build only that path. Skip validation, error handling, edge cases, and polish. Hard-code what you need to.
3. Run it. Verify the data flows correctly through every layer.
4. Once the tracer works, expand outward: add validation, error paths, tests, and remaining cases one at a time.

## Why

A tracer bullet proves the architecture works before you invest in details. It catches integration problems early — wrong assumptions about APIs, missing database columns, broken wiring between layers. Fixing these is cheap now and expensive later.

## What This Is Not

- Not a prototype. Tracer bullet code stays — it becomes the real implementation.
- Not a spike. Spikes are throwaway research. Tracer bullets are the skeleton you flesh out.
- Not an excuse to skip tests. Write tests as you expand, not after.

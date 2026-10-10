# STATE.md Format

`STATE.md` is the smallest resume point for a future session. Update it whenever work stops, a project changes status, or new evidence changes the appropriate difficulty.

## Template

```md
# Current Learning State

## Last verified
- Capability: {what the learner demonstrated}
- Evidence: {project path and evidence path}
- Confidence boundary: {what still has not been tested}

## Current project
- Project: {number and name}
- Phase: {brief | implement | diagnose | verify | reflect}
- Blocker: {the exact observed issue, or “none”}

## Next smallest task
{One concrete action that can produce new evidence.}

## Assistance used
{Hints, examples, scaffolds, or code supplied; include the area a transfer check should revisit.}

## Review queue
- {Concept, failure mode, or capability to retrieve later}

## Environment
{Runtime, versions, hardware, cost, or access facts that affect the next task.}
```

## Rules

- Prefer observed behavior and paths to vague confidence scores.
- Keep the next task small enough to start without re-reading the whole workspace.
- Record supplied assistance so later tasks can test independent transfer.
- Remove resolved blockers and revise stale environment facts instead of appending a long journal.

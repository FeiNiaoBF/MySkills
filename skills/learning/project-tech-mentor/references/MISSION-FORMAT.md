# MISSION.md Format

`MISSION.md` lives at the root of one technology-learning workspace. It captures the real-world outcome that should steer every project, source, and teaching decision.

## Template

```md
# Mission: {Technology}

## Why
{1–3 sentences describing the useful artifact or capability the learner wants to gain. Prefer a concrete outcome over “understand {technology}”.}

## Success looks like
- {An observable project or technical behavior the learner can deliver}
- {A debugging or evaluation behavior they can perform independently}
- {A concise explanation of an important design choice}

## Constraints
- {Time, budget, hardware, runtime, existing stack, and learning preferences}
- {What the learner already knows and how deeply}

## Out of scope
- {Adjacent technologies or production concerns deliberately deferred}
```

## Rules

- **One mission per workspace.** Unrelated technologies get separate workspaces; a connected stack can share one north-star artifact.
- **Concrete over abstract.** “Ship an interactive 3D product demo” is more useful than “learn Three.js”.
- **Push back on vagueness.** Ask what the learner wants to make, who it helps, and what evidence would convince them it works before writing the mission.
- **Revise when the goal changes.** Add a learning record when the mission shifts, then update this file so stale goals do not steer future projects.
- **Keep it short.** This is a compass, not the roadmap.

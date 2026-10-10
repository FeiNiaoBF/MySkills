# Learning Workspace Layout

Create only the folders that the current route needs. Keep the state files at the workspace root so a later session can determine the next action without relying on chat history.

```text
<technology>-learning/
  MISSION.md
  ROADMAP.md
  STATE.md
  RESOURCES.md
  TRACK.md
  GLOSSARY.md
  projects/
    0001-<project-name>/
      BRIEF.md
      ...runnable source...
  evidence/
    0001-<project-name>/
  learning-records/
    0001-<lesson-name>.md
  reference/                 # Optional HTML reference cards
  lessons/                   # Optional interactive or visual lessons
  assets/                    # Optional shared lesson assets
  NOTES.md
```

## Root files

- `MISSION.md`: why the learner wants this technology, their north-star artifact, current prerequisites, time budget, hardware or cost limits, and what success looks like.
- `ROADMAP.md`: the ordered project ladder. For each project, list the added capability, scope boundary, required evidence, current status, and its transfer check.
- `STATE.md`: the present project, verified capability, blocking observation, next smallest task, and due review or transfer check. Update it at every stopping point.
- `RESOURCES.md`: primary sources and high-trust supporting material. Each entry names the source, version or date when relevant, the decision it supports, and a direct link.
- `TRACK.md`: the domain route created from the [track format](TRACK-FORMAT.md) template.
- `GLOSSARY.md`: canonical terms the learner has demonstrated they can use; it is not a pre-course dictionary.
- `NOTES.md`: durable teaching preferences and environmental facts that change the route.

## Bootstrap sequence

1. Confirm the learning goal and the workspace boundary.
2. Capture the mission and constraints before choosing a stack or project list.
3. Propose one diagnostic task that produces runnable or inspectable evidence in 20–40 minutes.
4. Use the result to write the track and a small first roadmap. Do not prewrite later projects more precisely than the evidence supports.
5. Create the first project brief and state entry.

## Continuation sequence

1. Read the mission, roadmap, state, current brief, and latest relevant learning record.
2. Verify whether the prior next task still fits the learner's goal and environment.
3. Continue the task, or update the roadmap with a short explanation when new evidence changes the appropriate difficulty.

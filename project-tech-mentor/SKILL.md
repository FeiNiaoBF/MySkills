---
name: project-tech-mentor
description: Guide learners with basic coding ability to gain practical proficiency in a computer technology—such as frontend, backend, Three.js, LLMs, or computer vision—through stateful, progressively independent projects. Use when the goal is to learn a technical field by building and verifying work, not only reading concepts.
---

# Project Tech Mentor

Treat the current directory as one technology-learning workspace. The outcome is evidenced project ability: the learner can build, verify, debug, and explain a useful technical artifact with progressively less guidance.

## Choose the current mode

Inspect the workspace before deciding the next action.

- **New track:** read [workspace layout](references/workspace-layout.md), [mission format](references/MISSION-FORMAT.md), [track format](references/TRACK-FORMAT.md), and [resources format](references/RESOURCES-FORMAT.md). Establish a north-star project, practical constraints, and a 20–40 minute diagnostic task before creating the roadmap.
- **Continue a track:** read `MISSION.md`, `ROADMAP.md`, `STATE.md`, `GLOSSARY.md` when present, and the most recent relevant learning record. Use [roadmap format](references/ROADMAP-FORMAT.md), [state format](references/STATE-FORMAT.md), and [learning record format](references/LEARNING-RECORD-FORMAT.md) when creating or revising those files. Resume with the smallest unfinished, verifiable task.
- **Run a micro-project or review a submission:** read [project loop](references/project-loop.md), [project brief format](references/PROJECT-BRIEF-FORMAT.md), and [evidence format](references/EVIDENCE-FORMAT.md). Use their brief, feedback, evidence, and reflection loop.
- **Create or revise a domain route:** read [track format](references/TRACK-FORMAT.md) and follow its source and project-selection rules. Base claims and version-sensitive choices on current primary sources; record the source, version, and why it matters in `RESOURCES.md`.
- **Create an HTML lesson or reference card:** read [HTML lesson format](references/LESSON-FORMAT.md). Use it only when visual, interactive, or print-friendly material improves recall.

## Build the learning path

Start from the learner's desired artifact, not a chapter list. Use the diagnostic result to select a 4–8 project ladder in `ROADMAP.md`; each project should add one clearly observable capability and fit one to three sessions. Adjust later scopes, scaffolding, and openness from the learner's demonstrated work.

Use real constraints that make the technology matter: an interaction for visual work, a contract for an API, a fixed evaluation set for an LLM workflow, or data and error cases for computer vision. Keep tooling boilerplate separate from the current learning objective; label any supplied scaffold so the learner knows what they are responsible for.

After every two or three projects, schedule a transfer check: a repair, constrained rebuild, code review, or changed requirement that the learner completes with less prompting. End the route with a capstone defined by a goal and constraints rather than tutorial steps.

## Teach inside the work

Teach only the concept needed for the next implementation decision. Then give the learner a bounded task and let them attempt it before evaluating their work.

When they are stuck, use this help ladder and stop at the first level that unblocks progress:

1. Clarify the expected behavior, observed result, and exact evidence.
2. Identify the smallest subsystem or assumption to inspect.
3. Give a minimal hint or question that preserves the learner's next decision.
4. Show pseudocode or a local example.
5. On an explicit request, explain a feasible implementation and an immediate way to verify it.

When code is supplied, keep it local to the obstacle, explain its role, and attach a concrete verification step. Record material assistance in the project's reflection so a later transfer check can target the same capability.

## Verify, record, and resume

Treat a project as complete only when it has all applicable evidence:

- **Machine-verifiable:** a build, test, lint, request assertion, metric, or equivalent check.
- **Human-verifiable:** a screenshot, recording, output sample, log, error gallery, or observable behavior.
- **Learner-verifiable:** a concise explanation of key choices, failure paths, and the next improvement.

Run available checks rather than assuming they pass. Capture the evidence under `evidence/` using [evidence format](references/EVIDENCE-FORMAT.md), update `ROADMAP.md`, add terms to `GLOSSARY.md` only after the learner can use them correctly, create a learning record for non-obvious lessons, and leave `STATE.md` with the current capability, unresolved issue, next smallest task, and scheduled transfer check.

Use HTML only when visual, interactive, or print-friendly material improves recall: conceptual diagrams, simulators, debugging flows, and reference cards. Follow [HTML lesson format](references/LESSON-FORMAT.md). Keep runnable projects, briefs, state, and evidence in code or Markdown files.

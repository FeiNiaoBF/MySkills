---
name: project-digestion
description: >-
  Use when reverse-engineering an existing project: teach its mechanisms through short, evidence-based
  lessons so the learner can understand, modify, transfer, and direct AI development.
metadata:
  author: FeiNiaoBF
  version: "0.1.0"
  date: "2026-09-21"
---

# Project digestion

Turn a real project slice into a capability the learner can reuse.

The target is not memorizing source. The learner should move through the smallest useful capability
levels:

```text
located → traced → modified → abstracted → transferred → specified
```

A later level does not erase an earlier gap. Record the highest level actually evidenced.

## Non-negotiable lesson contract

Every lesson:

- fits within 25 minutes;
- solves one cognitive problem;
- uses one real project slice;
- asks for one prediction before revealing the relevant evidence;
- includes one bounded hands-on action unless fatigue is explicit;
- produces observable evidence and a checkpoint;
- recommends, rather than silently starts, the next lesson.

Keep the original project read-only. Use an independent worktree or scratch copy for learner changes.
Do not execute arbitrary project code or install dependencies merely to make a visual; ask for explicit
authorization when execution or network access is needed.

## Run the loop

### 1. Contract and calibrate

Identify the source: local path, GitHub repository, commit/PR, handoff, or `DIGEST.md`. Ask what the
learner wants to be able to recognize, change, or direct next time. Sample the minimum prerequisite
with one prediction question. Do not begin with a complete architecture lecture.

Choose one target capability and one concrete slice:

- entry point and user-visible outcome;
- relevant files and symbols;
- data/control flow;
- state and ownership boundary;
- tests, fixtures, or runtime evidence;
- external boundary such as filesystem, network, process, or model call.

### 2. Build a map, then stop

Produce only enough orientation for the learner to locate the slice. Attach claims to source lines,
tests, traces, diffs, or rendered evidence. Mark inferred edges and missing runtime evidence as
`unverified`; never turn a plausible guess into a fact.

### 3. Predict and trace

For one mechanism at a time:

1. ask what happens next: call, value, state, side effect, or owner;
2. let the learner commit an answer before revealing the next event;
3. inspect the smallest relevant source or runtime evidence;
4. compare prediction with observation and name the model error;
5. ask the learner to restate the corrected mechanism.

Use a trace table, call graph, state diagram, or a focused HTML artifact only when it makes the
question testable. A visual without a learner action is documentation, not a lesson.

### 4. Make one local change

Use the smallest task that tests the model:

- change one requirement;
- add one boundary case;
- repair a seeded defect;
- preserve a contract while moving one mechanism;
- write one missing function or test.

Give the smallest useful scaffold, then remove it after success. Verify with the narrowest relevant
test or observable check.

### 5. Climb and transfer

Only after the concrete trace is stable, name the abstraction it instantiates: algorithm/invariant,
data structure, module responsibility, dependency direction, concurrency/state/ownership, failure
boundary, test strategy, or engineering trade-off.

Change one axis for transfer: input shape, scale, failure mode, timing, module boundary, requirement,
or repository/language. Ask which parts remain invariant, which change, and why.

If the learner's goal includes AI direction, finish with a task brief containing behavior, non-goals,
boundary, invariant/contract, constraints, evidence to inspect, acceptance checks, and uncertainty to
surface instead of guess.

### 6. Checkpoint and next lesson

Record separately:

- what the learner located;
- the corrected mental model;
- one failed prediction and its cause;
- the modification or transfer evidence;
- the current capability level;
- the next revisit trigger or recommended lesson.

Use this compact checkpoint shape:

```json
{
  "slice": "path/to/file#symbol",
  "target": "traced",
  "prediction": { "prompt": "...", "answer": "...", "observed": "..." },
  "model": "one sentence describing the mechanism",
  "evidence": ["file:line", "test:name"],
  "action": { "kind": "modified", "result": "..." },
  "transfer": { "variation": "...", "prediction": "...", "result": "..." },
  "next": "..."
}
```

A correct answer after explanation is not evidence of transfer. Prefer a changed case, a small code
change, a test, or a precise AI task brief.

## Choose the cheapest visual

Use this order:

```text
structure → Markdown / Mermaid
relationships or filtering → SVG / D3
state, time, or manipulation → focused HTML + CSS + Vanilla JS
source/runtime mapping → recorded trace replay
algorithm dynamics → instrumented algorithm trace
large navigable graph → specialized graph view
3D/spatial phenomenon → Canvas / WebGL only when necessary
```

For HTML, follow `references/visual-system.md`. The first templates are `trace-board` for
source-to-runtime tracing and `state-lens` for legal states, transitions, and invariants. Do not
invent a new template for a small variation; add data to the closest existing template. Add a template
only when the learner's cognitive job is genuinely different.

## Completion test

A lesson is complete only when the learner can do the selected target action and the checkpoint names
the evidence. “Read the file”, “agreed”, and “the page was generated” are not completion criteria.

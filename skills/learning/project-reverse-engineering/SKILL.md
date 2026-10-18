---
name: project-reverse-engineering
description: >-
  Use when reverse-engineering an existing project: teach its mechanisms through short, evidence-based
  lessons so the learner can understand, modify, transfer, and direct AI development.
metadata:
  author: FeiNiaoBF
  version: "0.2.0"
  date: "2026-09-21"
disable-model-invocation: true
---

# Project Reverse Engineering

Turn a real project slice into a capability the learner can reuse.

The target is not memorizing source. The learner should move through the smallest useful capability
levels:

```text
located → traced → modified → abstracted → transferred → specified
```

A later level does not erase an earlier gap. Record the highest level actually evidenced.

## Non-negotiable lesson contract

Every lesson:

- has a scope that fits the learner's available time and one cognitive problem;
- uses one real project slice;
- asks for a prediction when the learner has enough context for one to be meaningful;
- includes a bounded hands-on action when the learner can safely attempt one;
- records observable evidence and a checkpoint appropriate to the chosen target;
- recommends, rather than silently starts, the next lesson.

Inspect the project without modifying it when the task is to understand or review. For an authorized exercise, use an isolated worktree or scratch copy unless the user has requested an in-place change; preserve existing work and follow project rules. Execute code or install dependencies only when the current request authorizes it and the environment is appropriate; an already authorized action does not need a second confirmation.

## Run the loop

### 1. Contract and calibrate

Identify the source: local path, GitHub repository, commit/PR, handoff, or `DIGEST.md`. Ask what the
learner wants to be able to recognize, change, or direct next time. Sample a prerequisite with a prediction question only when the learner has enough context to answer; otherwise show a small source slice first. Do not begin with a complete architecture lecture.

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

1. offer just enough context for the learner to predict a call, value, state, side effect, or owner;
2. when a prediction is useful, let the learner commit an answer before revealing the next event;
3. inspect the smallest relevant source or runtime evidence;
4. compare the learner's expectation, if any, with observation and identify the mistaken assumption;
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

Record the relevant items, leaving out steps that were not attempted:

- what the learner located and can now explain;
- the corrected mental model and any informative failed prediction;
- modification or transfer evidence, with its verification boundary;
- the current capability level actually demonstrated;
- the next revisit trigger or recommended lesson.

Use the workspace's existing learning record or a brief conversation checkpoint. Do not create a JSON record for each lesson unless another tool consumes it.

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

For HTML, follow `references/visual-system.md`. Its `trace-board` and `state-lens` are design patterns for source-to-runtime tracing and state reasoning, not supplied templates or runnable components. Build only the interaction this lesson needs; reuse an existing local implementation when one is available.

## Completion test

A lesson is complete only when the learner can do the selected target action and the checkpoint names
the evidence. “Read the file”, “agreed”, and “the page was generated” are not completion criteria.

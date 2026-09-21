# Project-digestion visual system

Use an interactive artifact only when it makes the learner's question testable. A visual without a
prediction, manipulation, comparison, or explanation is documentation, not a lesson.

## Direction

Use a restrained field-notebook/runtime-instrument language:

- deep ink navy background;
- amber for the current step and learner action;
- teal for observed facts and stable invariants;
- coral for failed predictions or broken invariants;
- editorial serif text, compact sans-serif controls, monospace code and measurements;
- one dominant stage, a narrow orientation rail, and an evidence strip;
- hairline rules and generous space instead of a grid of equal cards;
- motion only when it explains a state or causal change;
- accessible labels, visible keyboard focus, reduced-motion support, and no color-only meaning.

## Shared artifact anatomy

```text
lesson identity · source slice · progress
orientation rail | question + visual stage
                 | source evidence · state · explanation
controls         | prediction / checkpoint
```

The first screen must answer: what am I learning, where am I, what evidence changed, and what do I do
next?

## Template contracts

### trace-board

Use for source-to-runtime tracing. Show synchronized source lines, current event, call stack, partial
state, and one prediction checkpoint. Stepping must update at least the source highlight and runtime
state. Record predicted, observed, and explained separately.

### state-lens

Use for legal states, transitions, ownership, and invariants. Show the current state, legal events,
transition labels, state-specific values, and the invariant that determines the next action. Let the
learner predict an event before firing it; show the event log and allow one-step backtracking.

## Data rules

- Stable IDs for actors, events, states, source references, and checkpoints.
- Source references are facts; missing or inferred references are labelled `unverified`.
- State snapshots are partial and only include values needed for the current question.
- Every event declares its effect or invariant rather than relying on a label alone.
- Checkpoints store the learner action and result, not merely `completed: true`.

## HTML boundary

Prefer a real self-contained `.html` file with inline CSS and JavaScript, no build step, and no CDN
requirement. Use semantic HTML and inline SVG for spatial diagrams. Build dynamic DOM with
`createElement` and `textContent`; keep imported project data out of raw HTML insertion. Use responsive
layout and test the longest realistic symbol or state name before shipping.

Add a new template only when the learner's cognitive job is different from tracing or state reasoning.

---
name: project-navigator
description: >-
  Use when the user needs to understand, re-enter, or investigate an existing software project:
  its structure, runtime flow, recent changes, or engineering decisions. Start from project evidence,
  explain relationships and uncertainty, and make guided practice optional.
license: MIT
metadata:
  author: FeiNiaoBF
  version: "1.1.0"
---

# Project Navigator

Help the user regain and deepen an evidence-based understanding of an existing software project. The goal is not to describe every file or replace the user's judgment. Build a usable mental model of where a concern lives, how the system behaves, why a decision is supported, and what remains unknown.

Use this skill when the user asks to:

- orient themselves in a project or return after time away;
- trace how a feature runs from input to output;
- understand a recent or specified engineering change and its project impact;
- understand an architecture or implementation decision;
- reverse-engineer code, or learn to explore a project independently.

Do not insert this workflow into ordinary implementation tasks. A concise implementation handoff belongs to the project's normal completion workflow; use this skill when the user wants deeper project understanding. Do not modify production code as part of an explanation unless the user separately authorizes an exercise or implementation.

## Choose a mode

Select the smallest mode that answers the user's question. Combine modes only when necessary. If the target is clear, proceed without asking the user to choose a mode.

- **orient** — establish the project's purpose, current state, major parts, relationships, constraints, and a useful reading path.
- **trace-flow** — follow one feature from a real entry point through important control, data, state, error, and output paths.
- **explain-change** — explain a concrete change in the context of the existing system, using a diff, commit range, tests, or records.
- **explain-decision** — explain a recorded or inferred engineering decision, its constraints, trade-offs, and consequences.
- **guided-exploration** — help the user build independent code understanding through prediction, tracing, and a small verification task. Use only when the user wants to learn or practice; never make it a prerequisite for a quick explanation.

If the request is ambiguous and different modes would lead to materially different work, ask one focused question. Otherwise state a bounded assumption and continue.

## Evidence and explanation rules

1. **Inspect before explaining.** Read applicable project instructions and inspect the relevant code, tests, configuration, docs, and history. Treat names and summaries as clues, not proof of behavior. Check the current working tree and preserve user changes.
2. **Start with the system location.** Explain the overall role of the target before detailing symbols. For a broad orientation, move from project purpose and major responsibilities to relationships, then to a concrete entry point. Do not mechanically walk the directory tree.
3. **Explain relationships.** Show how components interact, how data changes, where state is owned, and which constraints shape the design. Prefer a small, useful model to a list of paths and symbols.
4. **Attach important claims to evidence.** Cite repository paths and symbols; include line references when the inspected view supports them. For runtime behavior, connect source to tests, logs, traces, or an observed command result. Explain the shortest useful reading path when that helps the user continue independently.
5. **Separate evidence classes.** Make clear which claims are:
   - **Verified facts:** directly supported by current code, tests, runtime observation, a diff, or a record.
   - **Recorded decisions:** explicitly documented or confirmed in an authoritative source.
   - **Inferences:** plausible explanations drawn from current implementation or history, but not directly confirmed.
   - **Unknowns:** evidence is missing, conflicting, or inaccessible.
6. **Do not invent history.** Current code can establish what the system does, but not necessarily why an earlier author chose it. A commit message or later explanation is evidence of what was recorded, not proof that the rationale is complete or correct. Say when motivation cannot be confirmed.
7. **Respect evidence limits.** A passing test supports only the cases it exercises. A static call graph may miss reflection, callbacks, generated code, or runtime configuration. A missing reference in search does not prove a component is unused.
8. **Keep the answer proportional.** Lead with the useful mental model or conclusion. Add call details, edge cases, alternatives, or diagrams when they help answer the user's question. Do not produce a full architecture document or HTML page by default.
9. **Keep independent judgment.** Point out contradictions, stale documentation, questionable assumptions, or gaps that affect the explanation. Do not treat an Agent-authored implementation as correct by default.

## Mode procedures

### orient — project understanding and location

1. Identify the project goal and current implementation state from its README, entry points, configuration, code, tests, and recent history as relevant. Distinguish shipped behavior from plans and incomplete work.
2. Find a small set of major responsibilities and establish their relationships from imports, call sites, interfaces, dependency wiring, data ownership, or runtime evidence. A directory tree alone is not an architecture model.
3. Locate important entry points, external boundaries, data paths, state owners, and consequential constraints or decisions. Include only what helps the user's next step.
4. Present a compact overview from whole to part: project purpose and state; major components and how they connect; one or more paths into the code; important constraints and uncertainties; a recommended next place to inspect.
5. If no authoritative architecture map exists and the user needs durable documentation, explain the gap and offer the relevant `project-docs` task. Do not silently create or maintain a second project map here.

Do not claim an exhaustive inventory unless one was requested and actually completed.

### trace-flow — runtime mechanism

Choose one user-visible capability or concrete behavior. Trace it from the real entry point through the important steps:

1. input and validation;
2. participating functions, modules, services, or processes;
3. important data transformations and ownership changes;
4. state updates or persistence;
5. relevant errors, branches, and recovery behavior;
6. output and the caller or user who observes it.

Follow actual calls, interfaces, configuration, and tests. Mark dynamic or unobserved edges as inferred or unverified. Explain the mechanism rather than paraphrasing every line. Use ordered steps for a simple path; use a compact diagram or table when it makes branching, ownership, or state easier to understand. Keep a textual explanation of essential meaning alongside any diagram.

### explain-change — engineering change

Determine the requested range and baseline from the user, Git history, branch comparison, or current diff. Preserve uncommitted changes and do not assume all workspace changes belong to the target task.

Explain, in this order as relevant:

- the original behavior or problem, with evidence or an explicit note if it is not recoverable;
- the goal and constraints from the request, issue, or confirmed records;
- the chosen approach and its evidence-backed rationale; distinguish recorded rationale from inference;
- the actual files, symbols, interfaces, and behavior changed;
- how the changed path now works in the surrounding system;
- tests or other checks run, their observed results, and their coverage limits;
- likely effects on callers, data, operations, or future changes, marking unverified effects as such.

Use the actual diff and relevant surrounding code. A commit title or diff summary alone is not enough to explain system impact. Do not say a change fixed a problem unless relevant behavior was verified.

### explain-decision — engineering decision

1. Search the authoritative sources first: ADRs, requirements, design documents, issue or review records, and relevant commit history.
2. Verify what the current implementation and tests establish about the decision's behavior and constraints.
3. Explain the problem, confirmed constraints, chosen option, documented rationale, alternatives if there is evidence for them, trade-offs, and consequences for extension or operation.
4. Label any rationale reconstructed from code as an inference. If no reliable record explains why the choice was made, say so instead of assigning an imagined historical motive.
5. Identify whether the decision is still reflected in current code. If the record and implementation conflict, report both with their sources and do not silently choose one as authoritative.

Do not create or update an ADR as part of explanation. If the user asks to record a decision, route the writing task through `project-docs` after the decision is confirmed.

### guided-exploration — independent understanding

Use the retained reverse-engineering learning loop, scoped to one real slice and one cognitive goal:

1. Agree what the user wants to be able to locate, explain, change, or direct. Calibrate only prerequisites that affect the exercise.
2. Give enough orientation to locate the entry point, relevant symbols, state/data owner, and available evidence. Mark inferred edges as unverified.
3. When meaningful, ask the user to predict a call, value, state change, side effect, or owner before revealing the observation. Do not force prediction when the user asks for a direct explanation.
4. Inspect source, tests, traces, or runtime behavior; compare prediction with observation and explain any important difference.
5. Offer one bounded, safe exercise only when authorized: a changed input, boundary case, small repair, contract-preserving edit, or targeted test. Use an isolated worktree or scratch copy if it could alter the user's project and no in-place change was requested.
6. Check the target capability with the narrowest useful evidence. Record what the user demonstrated, not merely what was explained. Recommend the next learning step; do not start it silently.

A brief conversation checkpoint is sufficient unless the project already has a learning record. Do not create a durable learner profile or a separate knowledge base by default.

For a focused interactive trace artifact, follow [visual-system.md](references/visual-system.md). Use it only when interaction tests the learner's model; ordinary explanations do not need an HTML artifact.

## Knowledge continuity across tasks

Treat the repository, not the Agent's conversation memory, as the source of project knowledge.

- Reuse current project documents and code. Prefer one authoritative location for each durable fact or decision; link to it instead of copying it into a parallel map.
- Distinguish durable project knowledge (purpose, responsibilities, contracts, architecture constraints, confirmed decisions) from task-local observations (current diff, one command result, temporary hypothesis, unresolved issue).
- Do not persist a generated orientation or change explanation automatically. If the user needs future maintainers to rely on it, route to `project-docs` and update the appropriate authoritative document with evidence.
- Check relevant documents against current code and recent history before relying on them. If they are stale, conflicting, or lack a clear owner, mark the discrepancy and explain which claims were re-verified. Do not perform a whole-documentation audit unless requested.
- Use Git history to recover changes and prior records where available, but do not assume commit messages preserve full design rationale.
- When a future task starts, rebuild only the context needed for that question from source material; do not require the same persistent map to be copied across projects or sessions.

## Completion check

Before answering, verify that:

- the selected mode answers the user's actual question;
- the explanation states how the relevant parts relate, not just what files exist;
- important claims point to evidence, while inference and unknowns remain clearly labelled;
- behavior, rationale, and verification are not conflated;
- the scope and limits of checks are visible;
- the user has a concrete next place to inspect when that helps them continue independently.

A useful explanation improves orientation; it does not prove the user has mastered the system. Only guided-exploration can assess a practiced capability, and only within its stated scope.

## Boundaries with related Skills

- **project-tech-mentor** builds proficiency in a technology through progressive projects. Use `project-navigator` when the main object is understanding an existing project's system and code.
- **make-sense** repairs a local confusion about a term, passage, or reasoning step. Use `project-navigator` when the question depends on project structure, actual call/data flow, a change, or a system decision.
- **project-docs** creates and maintains authoritative project documentation. This Skill can identify a stale or missing source and explain the discrepancy; route durable documentation work through `project-docs`.
- **code-review** assesses a fixed set of code changes against standards and a specification, including defects and risks. Use `explain-change` to help the user understand what the change does in the surrounding system. A request can need both, but neither substitutes for the other.
- **assess** evaluates a specified artifact or direction against its intended purpose. Use it for an independent judgment; use `explain-decision` to reconstruct what evidence supports a project's decision and what remains unknown.

Routine implementation, verification, and concise completion reporting remain in the project's normal workflow. This Skill is not a required post-task step. Do not introduce it merely because code was changed.

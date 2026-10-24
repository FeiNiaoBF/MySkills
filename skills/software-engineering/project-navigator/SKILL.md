---
name: project-navigator
description: >-
  Use when a user needs to regain orientation in an existing code project, trace how a feature works,
  understand an engineering change or decision in system context, or learn to explore the code independently.
  Ground explanations in current code and recorded evidence; offer guided exercises only when wanted.
license: MIT
metadata:
  author: FeiNiaoBF
  version: "1.0.0"
---

# Project Navigator

Help the user establish, recover, and deepen an evidence-based mental model of an existing engineering system. The outcome is knowing where a concern lives, how parts relate, what is verified, and how to investigate the next question without relying on the agent's memory. Reply in the user's language. Do not treat a directory tree as architecture or an agent-authored change as automatically correct.

## Route by the user's question

Choose the smallest mode that answers the request; do not require all modes or turn an ordinary development handoff into a lesson.

- **Orient:** New to the project, returning after a break, or lost in its structure. Start with the project's purpose and current implementation, then a small map of major responsibilities, their relationships and key boundaries. Point to a real entry point and the shortest useful reading route. If the question concerns only one subsystem, orient there instead of inventorying the repository.
- **Trace flow:** For a particular behavior, follow a real entry point through input, participating calls/modules, data transformations, state creation/update/persistence, relevant failure paths, and observable output. Explain the causal mechanism, not every line. Cite the code behind important edges; distinguish a static path from a runtime-observed one. Use a diagram only when it clarifies a relationship.
- **Explain change:** For a specified commit/range or recent work, establish the baseline and actual diff first; connect changed modules to the surrounding call/data flow. Explain the problem and rationale only when supported by a request, issue, commit, ADR, or other record. Then explain behavioral effects, relevant tests actually run, and unverified paths. If the range or prior state is unclear, bound it with available Git evidence or ask which change the user means. A diff describes *what*, not necessarily *why*.
- **Explain decision:** Start with the question and known constraints. Compare recorded ADRs/requirements, code and tests to separate documented historical intent from observable implementation, plausible present-day reasoning, and unknown motive. Discuss meaningful alternatives, costs, limitations and extension consequences without inventing a retrospective decision record.
- **Guided exploration:** When the user wants to learn to navigate or modify the actual code themselves, use the optional learning loop below. Do not impose prediction, practice, or a quiz on someone requesting a quick explanation.

## Establish evidence before explanation

1. Check the user's goal and existing context; inspect project instructions and workspace status before a potentially stateful check. Locate a project entry (README, module overview, or prior handoff), relevant architecture/ADR records, code symbols, tests and Git history as appropriate. Verify current behavior against implementation; documents and filenames are leads, not proof. Do not mechanically read every file.
2. Build a small mental map: user-visible behavior → entry/interface → responsible modules → data and state boundaries → output. Start at system scale before drilling into the relevant slice. Explain the relationships instead of dumping symbols.
3. Separate **observed code or runtime behavior**, **recorded historical decisions**, **inferences**, and **unknowns**. Check contradictions between documentation and code and say what remains uncertain. If runtime behavior matters, use the narrowest safe check available; do not claim a static trace proves it ran. Question suspicious designs rather than rationalizing them.
4. Answer at the user's requested depth: a concise orientation or mechanism first, then detail on demand. State how to verify the next important uncertainty. Do not silently modify production code, install dependencies, or run unsafe actions for a question about understanding.

## Make the answer verifiable

For key claims about architecture, flows, changes or decisions, place a locatable reference near the claim: a Markdown link to relevant source lines where supported, or a file path with line number or symbol; cite a test, commit/diff or decision record when that is the actual evidence. Give the user a short reading route that says what to open first and what they can check there. Do not dump links or attach a line number to every sentence. Code can establish current behavior but not an undocumented historical motive; say when the motive is unknown.

## Restore orientation across tasks

Start with the project's maintained entry and relevant architecture/ADR records, then compare their claims with the current code and recent Git changes in the area at issue. Stable knowledge belongs in the project's existing architecture or decision documentation; transient task status belongs in the task/handoff or Git history. A past conversation is not a reliable source of truth. If no maintained map exists, reconstruct a minimal one from code and identify the gap; do not create a knowledge store just to answer a question.

For a returning user, give a compact **where we are now → what changed (if evidenced) → where to read next** bridge with links to existing authoritative sources. If a durable description is missing or stale, flag the exact discrepancy and its impact. Update only when the user has authorized documentation work or a current authorized change directly requires it; use `project-docs` for substantial documentation maintenance, preserving one authoritative home for each fact and linking rather than copying it. This skill does not establish automatic memory or write a new project digest on every visit.

## Optional learning loop (from project-reverse-engineering)

Choose one real slice and one target capability. The progression `located → traced → modified → abstracted → transferred → specified` is a set of possible outcomes, not a mandatory sequence; claim only levels supported by user actions. Keep the slice within the learner's time and one cognitive problem.

1. Give just enough context for a meaningful prediction about a call, value, owner, state, side effect or failure. If the learner lacks context, show a small source slice first. Let them answer before revealing the next event when they want an interactive exercise.
2. Trace the real entry and smallest relevant source, test or runtime evidence. Compare prediction with observation, correct the mistaken assumption, and let the learner restate the mechanism. Label inferred edges and missing runtime evidence.
3. If the user wants hands-on practice and a change is authorized, choose a bounded task: one requirement, edge case, seeded defect, invariant-preserving refactor, function or test. Prefer an isolated worktree or scratch copy unless in-place changes were requested; preserve existing work and verify narrowly. Understanding alone does not authorize edits.
4. Once the concrete trace is stable, identify the abstraction (e.g., invariant, ownership, dependency, failure boundary) and optionally vary one axis (input, timing, scale, requirement or module boundary). If directing AI is the goal, help draft a brief with behavior, non-goals, boundary, contracts, constraints, evidence, acceptance checks and uncertainty.
5. Note only demonstrated capability, the correction or verification evidence, and a possible revisit trigger. Use an existing learning record when appropriate, otherwise a brief conversation checkpoint. Do not create a record per lesson or declare transfer after a correct answer to the original example.

Choose the cheapest useful representation: prose or Markdown for structure, a focused trace table or diagram for relationships, interactive HTML only when state or manipulation genuinely helps the learner test a question. For an HTML learning artifact, follow [visual guidance](references/visual-system.md); its patterns are design guidance, not required templates. No visual is required for ordinary navigation.

## Adjacent responsibilities

- `project-tech-mentor` builds proficiency in a technology through progressive projects; use this skill when the primary object is understanding an existing code system.
- `make-sense` resolves a local term, passage or reasoning gap; use this skill for a system-wide relation or real call path.
- `project-docs` writes and maintains durable project documentation; this skill explains from existing evidence and points out stale documentation.
- `code-review` judges the quality, standards and risks of a diff; explaining what changed and how it fits the system is not a review verdict.
- `assess` evaluates a bounded target against criteria; understanding its actual engineering context is not itself a fitness judgment.

For ordinary development, the general workflow remains responsible for a concise engineering handoff. Use this skill when the user wants deeper understanding of that handoff, not as an automatic step after every change.

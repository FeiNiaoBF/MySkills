---
name: project-implementation-planner
description: >-
  Use when the user explicitly requests a staged implementation plan for a software
  project or invokes this skill. It orders the next decisions and work from
  confirmed goals, dependencies, uncertainty, and engineering risk; explains
  priorities, records only useful project decisions, and hands off a bounded next
  slice. Ordinary feature requests do not trigger a project-wide planning flow.
license: MIT
metadata:
  author: FeiNiaoBF
  version: 0.1.0
disable-model-invocation: true
---

# Project Implementation Planner

Plan a sequence of small, verifiable software changes that serves the user's actual goal. Use software engineering judgment to improve robustness, readability, and ease of change; do not promise to eliminate future refactoring. Explain the recommendation in the user's language, preserving project terminology and code identifiers.

## Activation and responsibility

Run only when the user invokes this skill or explicitly asks for staged project implementation planning. This skill owns priority decisions, their explanation, and the next development handoff. It does not automatically run for a scoped feature, a routine coding request, a documentation edit, or a technology question. Host-specific manual invocation settings may not be enforced by every host.

Planning alone ends with a plan and any authorized documentation. It does not authorize building an entire system. If the same request already authorizes development, present the relevant sequence and continue with the first bounded slice using the normal coding workflow; do not require a separate development skill or ask for the same permission again. Subsequent slices follow the user's authorized scope and feedback, rather than treating a speculative roadmap as blanket authorization. External publication and consequential irreversible changes follow their own approval rules.

## Plan from evidence

1. **Inspect what exists.** Check project instructions, current state and worktree, goals, requirements, code, interfaces, tests, and existing decisions as relevant. Distinguish confirmed behavior and constraints from observations, proposals, and unknowns. For a new project, do not infer a complete feature set from a phrase such as “backend system.” Ask about the one missing decision that would materially change the next step; continue work independent of it. Respect explicit exclusions and earlier decisions.
2. **Find the decision frontier.** Identify the smallest useful outcome and the choices that must precede it. Order work by concrete dependencies, uncertainty worth resolving, user value, reversibility, and failure cost. State why the next step comes first and what evidence could change the order. Do not hard-code “stack, data model, database, API” as a universal sequence. If integration uncertainty drives the choice, propose the thinnest end-to-end production slice that tests the critical relationship; keep temporary experiments isolated and identify their limits.
3. **Keep the design changeable.** Give each proposed slice a coherent responsibility, observable outcome, known inputs and consumers, meaningful constraints, and completion evidence. Prefer clear boundaries and one authoritative definition for each rule; do not invent generic layers, future features, frameworks, or abstractions without a present use case. A slice may cross modules where needed to prove behavior. Surface irreversible or broad-impact decisions before dependent work; leave reversible detail until it matters.
4. **Explain and record proportionately.** In the conversation or terminal, show the proposed order, dependencies, reasons, assumptions, and first actionable slice. Use an existing project plan or issue system for durable status when one is needed; otherwise add the smallest discoverable document justified by the task and agreed destination. Preserve one authoritative location per requirement or decision. Record an ADR only for a consequential choice whose alternatives, rationale, trade-offs, and consequences will matter later; an ADR can be proposed until actually confirmed. Do not turn each slice into an ADR or create empty Requirements, Architecture, and Spec files. For substantial document authoring, `project-docs` can supply writing guidance without being a prerequisite to this skill.
5. **Hand off the next slice.** State the expected behavior and limits, prerequisites already settled, unresolved decisions that block only this slice, affected interfaces or data, and a suitable verification signal. Leave later slices at the detail the evidence supports; update their order after implementation reveals new facts. A plan is revisable, not an assertion that every proposed feature is approved.

## Completion and verification

A planned slice is complete when its agreed observable behavior and affected contracts are checked, the implementation remains understandable, and the relevant documentation or plan status reflects the result. Choose verification by risk: focused unit tests for local logic; integration, contract, migration, or end-to-end checks when boundaries, persistence, or failure paths make them necessary. Use the smallest meaningful checks and project-required gates. Tests are evidence for the behavior they cover, not proof of overall robustness. When development is not authorized, specify these checks as completion criteria rather than claiming to have run them.

Before finishing a planning task, check that the first step can be started without invented requirements, dependencies and uncertainty are visible, important decision states are accurate, and the reader can tell which document (if any) is authoritative. Report what was inspected, which plan or records were actually created or updated, and what remains undecided. Do not claim code, tests, or integrations exist on the strength of a plan alone.

---
name: problem-exploration
description: >-
  Use when the user has an unclear problem or idea and needs evidence to choose
  what to do. Frame the question, investigate relevant evidence, validate key
  assumptions with bounded checks, and recommend a next step. Skip straightforward
  factual questions and already-scoped implementation tasks.
license: MIT
metadata:
  author: FeiNiaoBF
  version: 0.1.0
---

# problem-exploration

Help the user understand an unclear problem and decide what action, if any, is justified. Start from evidence and the decision at hand, not from a presumed solution. Creating a new project is one possible outcome, never the default.

## Working principles

- **Separate evidence from interpretation.** Distinguish observed facts, desired outcomes, constraints, hypotheses, and unknowns. Label hypotheses until evidence supports them.
- **Challenge assumptions with purpose.** Test an assumption when its answer could change the decision. Explain the reason and offer practical alternatives.
- **Inspect before asking.** Use available project files, artifacts, and trustworthy sources to resolve questions where possible. Ask only when a missing fact cannot reasonably be established and would change the investigation or recommendation.
- **Keep the work bounded.** Investigate subproblems that affect the main decision. Briefly note tangential findings; ask before expanding into an independent task.
- **Validate only when useful.** Prefer the smallest safe, reversible check that can distinguish plausible explanations. A prototype or experiment is evidence, not production-ready work.
- **Respect the task boundary.** Exploration does not authorize implementation or project setup. Continue into implementation when the user requests it and the request authorizes the needed changes; otherwise provide a recommendation and a next step.
- **Use specialist workflows when they fit.** A confirmed bug, requested prototype, conceptual explanation, or documentation task may belong to a more focused Skill. Keep this Skill responsible for framing the open question and deciding what evidence or handoff is needed.

## Adaptive workflow

These are reasoning activities, not mandatory sequential stages. Skip or revisit stages when the available evidence warrants it.

1. **Triage.** Decide whether the request is genuinely unclear. Answer a straightforward question directly. For a clearly scoped task, use its implementation workflow instead of forcing exploration.
2. **Frame.** Identify the main question, why it matters, the desired outcome, constraints, and the decision that evidence needs to support. Ask one decision-relevant question only when an unresolved ambiguity materially changes the work.
3. **Investigate.** Read relevant project material and inspect available evidence. Search externally only to resolve a concrete question. Keep findings tied to the main decision.
4. **Validate if useful.** Run a low-cost, non-destructive check or isolated experiment when its result could change the recommendation. State what it establishes and what remains uncertain.
5. **Synthesize.** Compare relevant options. Depending on the problem, consider no action, a configuration change, reuse, modifying an existing component, or creating a separate project. Explain evidence, trade-offs, and uncertainty without filling a template for its own sake.
6. **Recommend or hand off.** Give the best-supported next step and the evidence behind it. Stop at a decision unless implementation is requested and authorized.

## Lightweight scope tracking

Keep a lightweight view of:

- **Main question:** the decision or uncertainty being explored.
- **Necessary subproblem:** an issue that could change the answer.
- **Side finding:** useful but unnecessary to resolve the main question.
- **Independent task:** work that would change the main scope.

Do not create a plan, report, issue, or project automatically. Record findings only when the user or project workflow needs a durable handoff.

## Stop conditions

Stop when the question is resolved, the evidence is sufficient to choose a next step, further investigation is unlikely to change the decision, or a critical dependency blocks progress. State uncertainty or the missing dependency plainly. Do not continue merely to exhaust every possible avenue.

## Response

Lead with the conclusion. Summarize the evidence that matters, distinguish facts from hypotheses, and name the next decision or action. Match the level of detail to the request; do not impose a formal report template on a simple exchange.

Possible outcomes include an explanation, a validated diagnosis, a practical change proposal, a comparison of directions, an experiment plan, a handoff to a specialist workflow, or a reasoned recommendation to stop.

## Anti-patterns

- Turning every idea into a project brief, requirements document, or architecture plan.
- Asking the user questions that available materials could answer.
- Searching broadly without a concrete question to resolve.
- Presenting hypotheses as facts or an experiment as production validation.
- Treating a new module or project as the default solution.
- Expanding into an independent task without agreement.
- Implementing changes when the request only authorized exploration.
- Continuing investigation after the recommendation is decision-ready.

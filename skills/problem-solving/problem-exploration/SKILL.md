---
name: problem-exploration
description: >-
  This skill investigates unclear problems or ideas when evidence is needed to
  decide what to do. It frames the decision, inspects relevant evidence, tests
  consequential assumptions with bounded checks, and recommends a next step.
  It does not turn straightforward factual questions, wording requests, or
  already-scoped implementation tasks into an exploration workflow.
license: MIT
metadata:
  author: FeiNiaoBF
  version: 0.1.2
---

# Problem Exploration

Help the user understand an unclear problem and decide what action, if any, is justified. Start from evidence and the decision at hand, not from a presumed solution. Creating a new project is one possible outcome, never the default. Reply in the user's language unless they request another language for the deliverable.

## Working principles

- **Separate evidence from interpretation.** Distinguish observed facts, desired outcomes, constraints, hypotheses, and unknowns. Label hypotheses until evidence supports them.
- **Challenge assumptions with purpose.** Test an assumption when its answer could change the decision. Explain the reason and offer practical alternatives.
- **Inspect before asking.** Use available project files, artifacts, and trustworthy sources to resolve questions where possible. Ask only when a missing fact cannot reasonably be established and would change the investigation or recommendation.
- **Keep the work bounded.** Investigate subproblems that affect the main decision. Briefly note tangential findings; ask before expanding into an independent task.
- **Validate only when useful.** Prefer the smallest safe, reversible check that can distinguish plausible explanations. A prototype or experiment is evidence, not production-ready work.
- **Respect the task boundary.** Exploration authorizes relevant inspection and safe, isolated checks, not production changes or project setup. Use authorization already present in the conversation when continuing into implementation; do not ask for the same approval again. A request to explore alone ends with a recommendation.
- **Stay independent at adjacent boundaries.** When the primary task is wording, explanation, visual design, diagnosis, documentation, or already-scoped implementation, do not force it through exploration. A broad investigation may still use the smallest relevant check from one of those disciplines as evidence, then return to the open decision.

## Adaptive workflow

These are reasoning activities, not mandatory sequential stages. Skip or revisit stages when the available evidence warrants it.

1. **Triage.** Decide whether the request is genuinely unclear. Answer a straightforward question directly. For a clearly scoped task, use its implementation workflow instead of forcing exploration.
2. **Frame.** Identify the main question, why it matters, the desired outcome, constraints, and the decision that evidence needs to support. Reuse answers already supplied. Ask one decision-relevant question only when an unresolved ambiguity materially changes the work; otherwise state a bounded assumption and proceed.
3. **Investigate.** Read relevant project material and inspect available evidence. For each consequential hypothesis, identify what observation would support or weaken it and choose the cheapest useful source or check. Search externally only to resolve a concrete question; prefer primary sources for product or API claims. Keep source links, file locations, or observed results attached to consequential findings. Consider conflicting evidence, not only support for the favored explanation. If access fails, state the gap rather than treating a search snippet or recollection as verification.
4. **Validate if useful.** Before running a check, state the hypothesis, the observation that would change the recommendation, and a practical bound such as one representative input or a fixed sample. Use a low-cost, non-destructive check or isolated experiment. Keep temporary artifacts outside the project unless a project artifact is requested. Report the result, its scope, and what remains uncertain; a successful sample does not establish general feasibility, demand, or production readiness.
5. **Synthesize.** Compare only options that could resolve the framed question. Depending on the problem, consider no action, a configuration change, reuse, modifying an existing component, or creating a separate project; these are alternatives, not a mandatory checklist. Explain evidence, trade-offs, and uncertainty without filling a template for its own sake.
6. **Recommend or hand off.** Give the best-supported next step and the evidence behind it, including any unresolved assumption that could reverse the recommendation. When a handoff is useful, carry forward the question, decisive evidence, constraints, and remaining uncertainty. Continue implementation only within existing authorization; otherwise stop at the recommendation.

## Lightweight scope tracking

Keep a lightweight view of:

- **Main question:** the decision or uncertainty being explored.
- **Necessary subproblem:** an issue that could change the answer.
- **Side finding:** useful but unnecessary to resolve the main question.
- **Independent task:** work that would change the main scope.

Do not create a plan, report, issue, or project automatically. Record findings only when the user or project workflow needs a durable handoff.

## Stop conditions

Stop when the question is resolved, the evidence is sufficient to choose a next step, further investigation is unlikely to change the decision, or a critical dependency blocks progress. Distinguish a supported conclusion from an inconclusive check, exhausted investigation budget, or unavailable evidence. In the latter cases, give a conditional recommendation or the smallest missing observation needed to proceed. Do not continue merely to exhaust every possible avenue, or interpret failure to find evidence as proof that an option cannot work.

## Adjacent task boundaries

Keep this Skill focused on an open decision rather than binding it to named companion Skills:

- Rephrasing the user's intended message is a wording task; it does not authorize investigating the task described by the draft.
- Explaining an existing term, passage, or reasoning step is an explanation task unless a separate decision remains open.
- Turning a desired visual effect into design directions is a design task; explore only the unresolved choice that needs evidence.
- Reported broken, failing, or slow behavior calls for diagnosis of the observed behavior. A confirmed cause is not required before beginning that diagnosis; return here only when a broader decision remains unresolved.
- A disposable prototype can be the bounded experiment in step 4 when it tests a stated hypothesis. Preserve its limits and the decision it must inform; do not treat it as production work.
- Writing or maintaining project documentation is a separate deliverable. Record an exploration only when the user or project workflow needs a durable handoff.

A broad question can contain one of these subproblems. Resolve only the part needed as evidence, then return to the main decision instead of restarting discovery. Research tools can supply evidence without turning every investigation into a deep research report.

## Response

Lead with the conclusion. Summarize the evidence that matters, distinguish facts from hypotheses, and name the next decision or action. Match the level of detail to the request; do not impose a formal report template on a simple exchange.

Possible outcomes include an explanation, a validated diagnosis, a practical change proposal, a comparison of directions, an experiment plan, a handoff to a specialist workflow, or a reasoned recommendation to stop.

## Anti-patterns

- Turning every idea into a project brief, requirements document, or architecture plan.
- Asking the user questions that available materials could answer.
- Searching broadly without a concrete question or a stopping bound.
- Presenting hypotheses as facts, a missing source as disproof, or an experiment as production validation.
- Treating a new module or project as the default solution.
- Expanding into an independent task without agreement.
- Implementing changes when the request only authorized exploration.
- Continuing investigation after the recommendation is decision-ready.

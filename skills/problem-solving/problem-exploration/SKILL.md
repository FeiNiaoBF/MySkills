---
name: problem-exploration
description: >-
  Use when a user has an unclear problem or idea and needs evidence to decide
  what, if anything, to do next. Clarify the question, investigate what could
  change the decision, and recommend a justified next step. Do not use as a
  general prelude to factual answers, diagnosis, assessment, or scoped work.
license: MIT
metadata:
  author: FeiNiaoBF
  version: 0.2.0
---

# Problem Exploration

Help the user move from an unsettled situation to a decision they can make. The question drives the inquiry: do not start with a presumed solution, a fixed sequence of activities, or a default plan to create something. Reply in the user's language unless they request otherwise.

## When to use

Use this Skill when **both** the problem or direction is still open **and** evidence could change what the user should do. For example, the user may know something is unsatisfactory but not why it matters, which obstacle is decisive, or whether action is warranted. If the user explicitly invokes this Skill for an otherwise clear task, do not manufacture uncertainty; answer or carry out the task within its authorization.

Do not route a straightforward fact question, a request to reword or explain something, an assessment of a specified target, a reported malfunction needing diagnosis, or an already-scoped implementation through exploration. A broader open decision may need a small check from one of those disciplines; use it only as evidence, then return to the decision. A confirmed cause is not a prerequisite for starting diagnosis.

## Inquiry

These are moves to choose from, not stages to report or complete in order. Reframe the question when evidence shows that the initial framing missed the real obstacle.

1. **Find the live question.** Identify what the user wants to decide or change, what is already known, and which uncertainty prevents a justified next step. Separate observations, goals, constraints, interpretations, and unknowns. Reuse context already given. If a missing choice would materially change the inquiry and cannot be inferred safely, ask one focused question; otherwise proceed with a stated, bounded assumption.
2. **Follow decision-relevant evidence.** Inspect available materials before asking the user to supply what can be checked. For plausible explanations or options, ask what observation would distinguish them and start with the least costly useful check. Use external sources for a concrete unresolved claim, favoring primary sources when authority matters. Attach source links, file locations, or observed results to consequential findings; look for disconfirming evidence as well as support. Inaccessibility is a gap, not disproof.
3. **Test only what matters.** If inspection cannot settle a consequential uncertainty, use a small, safe, reversible check or isolated experiment with a stated bound and a result that could change the recommendation. Keep temporary artifacts outside the project unless requested. A prototype tests a hypothesis; one successful sample does not establish broad feasibility, demand, or production readiness. Explain the check to the user beforehand only when its cost, impact, or authorization calls for it; otherwise report what was checked and its limits with the result.
4. **Choose the next step.** Compare only viable responses to the live question, including doing nothing when warranted. Say which finding supports the recommendation and which remaining uncertainty could reverse it. If evidence is insufficient, give a conditional recommendation or name the smallest observation needed, rather than treating a guess as a verdict.

## Scope and stopping

Keep the main decision in view. Pursue a subproblem only if its answer could change that decision; note side findings briefly and ask before taking on an independent task. Exploration permits relevant inspection and safe isolated checks, not production changes, project setup, or durable reports by default. Use authorization already given in the conversation for any follow-on work; do not ask again for the same approval. If the user authorized exploration alone, stop at a recommendation.

Stop when the evidence is sufficient to choose a next step, further inquiry is unlikely to change it, or a critical dependency blocks progress. An inconclusive check is not proof of impossibility. Do not search, prototype, or interview merely to exhaust a workflow; do not turn every idea into a project brief.

## Response

Lead with the answer to the live question or, if it remains unresolved, the best conditional next step. Distinguish evidence from inference, state limits that could change the decision, and make any handoff carry forward the question, decisive findings, constraints, and open uncertainty. Match the detail to the conversation; do not impose a report template or narrate every internal step.

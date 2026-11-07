---
name: problem-analysis
description: >-
  Use when the problem itself is unclear, contested, or badly framed. Identify
  what kind of problem it is and what question must be answered, then choose the
  inquiry method and theory that fit and recommend a justified next step. Do not
  use for factual lookups, fixed-target assessment, bug diagnosis, or
  already-scoped work.
license: MIT
metadata:
  author: FeiNiaoBF
  version: 1.0.0
---

# Problem Analysis

The problem comes first. Do not begin with a preferred method, theory, or solution. Work out what kind of problem this is, whose problem it is, what answer would count, and what would change the decision. Then choose how to inquire.

## When to use

Use when the problem or direction is still open and at least one of these is true:

- the user cannot say what is wrong, only that something is;
- several problems are tangled and it is unclear which one matters;
- the facts are uncertain, or the values and goals are contested;
- the proposed solution may be solving the wrong problem;
- the user asks what to do and the answer is not obvious.

If the user explicitly invokes this Skill for a clear task, do not manufacture uncertainty; answer the question or carry out the task within its authorization.

Do not route these through exploration: a factual lookup, rewording or explanation, assessment of a specified target, a reported malfunction needing diagnosis, or already-scoped implementation. If one of those is needed as evidence for a larger open decision, use it as a bounded check and return to the decision.

## Operating stance

- **Problem-driven, not method-driven.** A method earns its place only by producing evidence that could change the recommendation. Never choose a method because it is familiar, available, or prestigious.
- **Theory is a lens, not a verdict.** Use theory to generate explanations, mechanisms, and criteria; state its assumptions and what it would predict. Test it against at least one rival account when the claim is consequential. Do not force a theory onto a problem it does not fit.
- **Framing is part of the answer.** The same situation can be a personal trouble, an organizational failure, a market failure, or a structural injustice; each frame implies different evidence and remedies. Surface the frame before evaluating solutions.
- **One recommendation, not a menu.** Compare options, then say which one you recommend, why, and what would change your mind. A menu without a recommendation is abdication.
- **Proportionate rigor.** Match the depth of inquiry to stakes, reversibility, and time. A reversible low-cost choice needs a quick check; an irreversible or collective choice needs more evidence, more voices, and an explicit stop rule.
- **No action is an option.** Maintaining the status quo, waiting for a trigger, or running a small probe can be the right recommendation.

## Work the problem

These are moves, not a mandatory sequence. Reframe as evidence arrives.

### 1. Reconstruct the problem from context

Before asking the user to repeat what can be checked, inspect what is available:

- the conversation: stated goal, complaints, constraints, prior attempts, deadline, decision-maker;
- the project or environment: files, documentation, data, logs, history, product context, domain sources;
- the people affected: who bears the cost, who must act, who can block.

Write a working statement in the user's terms:

> You are trying to **the outcome**, because **the reason**, and the blocker is **the uncertainty or conflict**. The decision is **what must be chosen**, and an answer would count if **the criterion**.

Separate observation, goal, constraint, interpretation, and unknown. If the user's stated problem and the evidence point to different problems, say so and explain the consequence. If a missing choice would materially change the analysis and cannot be inferred safely, ask one focused question; otherwise proceed on a stated assumption. If your reconstruction is uncertain, state it and invite correction rather than asking the user to start from scratch.

### 2. Classify the question and the problem structure

First name the kind of answer needed: describe, explain, predict, evaluate, decide, design, interpret, or negotiate. Then locate the problem on two axes:

| | Values and goals clear enough to act | Values and goals ambiguous or contested |
|---|---|---|
| **Knowledge sufficient** | **Structured:** choose the best means to a given end. Use optimization, cost-benefit, or standard diagnosis. | **Value-ambiguous:** more evidence will not settle what matters. Use deliberation, ethics, stakeholder analysis, or negotiation. |
| **Knowledge uncertain** | **Knowledge-uncertain:** investigate causes, mechanisms, or likely effects. Use comparison, process tracing, experiments, or models. | **Unstructured or wicked:** both what is true and what is right are contested. Structure the problem, compare frames, and run safe-to-fail probes. |

Also record the scale: individual or collective; reversible or irreversible; immediate or long-horizon; one decision-maker or many. These change both the method and the stopping rule. For a "should" question, separate the empirical part (what is likely to happen) from the normative part (what matters, to whom, and who decides).

If the classification is uncertain, treat the problem as unstructured and start with framing plus the cheapest discriminating check. Revise the classification when evidence arrives. Do not use the typology as a label; use it to decide what kind of answer is possible.

### 3. Surface the problem representation and alternative frames

Ask:

- Who defines this as a problem, and for whom is it a problem?
- What is included and excluded by this definition?
- What cause is assumed? What would count as a solution?
- At what level is it posed: individual, team, organization, market, institution, or structure?
- What would make it a non-problem, or a different problem?
- Whose values, interests, or constraints does this framing serve?
- What is the implicit theory of change?

Generate at least one serious alternative frame when the decision is consequential. Vary the causal locus, the level, the time horizon, or the actor. For each frame, state what it predicts and what would look different if it were right. Choose the frame that best fits the evidence and the decision, not the most familiar or morally comfortable one. Keep the runner-up and the condition under which it would win.

Do not manufacture a reframe when the user's frame is already defensible. The test is whether the alternative changes the recommendation.

### 4. Choose the inquiry mode and theory

Start from two questions: **What would count as an answer?** and **What observation would discriminate between the live possibilities?** Then choose the smallest method set that can produce that observation.

| Problem | Useful modes | Theory's job | Characteristic failure |
|---|---|---|---|
| Causal explanation | comparison, process tracing, natural experiment, statistical control | supply mechanisms, confounders, scope conditions | correlation mistaken for cause; selection; reverse causation |
| Prediction | base rates, reference-class forecasting, scenarios, models | supply causal structure and priors | false precision; regime change; overfitting |
| Interpretation or meaning | interviews, ethnography, document and discourse analysis | supply sensitizing concepts, not the answer | imposing the analyst's frame; treating one voice as the culture |
| Evaluation | criteria from purpose and stakeholders, counterfactual, cost-benefit, distributional analysis | supply criteria and causal mechanisms | hidden norms; measuring what is easy |
| Decision under uncertainty | expected value, minimax, real options, decision matrix | supply priors and structure | false precision; ignoring reversibility |
| Design | constraints, prototypes, trade-off analysis, safe-to-fail experiments | supply patterns and precedents | solutionism; ignoring implementation |
| Value conflict | deliberation, stakeholder analysis, ethical frameworks, negotiation | clarify values and trade-offs, not settle them | technocracy; pretending evidence decides |
| Wicked or unstructured | problem structuring, frame reflection, participatory mapping, scenario planning | supply alternative framings | premature consensus; analysis paralysis |

Use theory when it does real work: it predicts something, explains a mechanism, or exposes an assumption. State what it illuminates and what it hides. When two theories fit the same facts, use the one that makes the more discriminating prediction, and say what evidence would favor the other.

Methods can be combined. Do not run a method merely to complete a workflow. If the method cannot change the recommendation, skip it.

### 5. Design the evidence check

Identify the **decisive uncertainty**: the one that, if resolved, would change the recommendation. Then choose the least costly, safest, most discriminating check:

- read existing material before asking for new data;
- use existing data before collecting new data;
- use a small reversible probe before a large commitment;
- ask what you would expect to see if you were wrong, and look for that.

For value conflicts, evidence can clarify consequences and constraints but cannot decide the value. Name the legitimate decision-maker or process instead of pretending the analysis can settle it.

State what the check establishes and what it does not. One case, one sample, or one success does not prove general feasibility, demand, or safety.

### 6. Recommend a course of action

Compare only viable options, including:

- act now;
- take the smallest reversible step;
- gather one more piece of evidence first;
- maintain the status quo or wait for a trigger.

Recommend one. Give:

- the recommendation and the finding that supports it;
- the key assumption and the main risk;
- cost, reversibility, and who is affected;
- what would reverse the recommendation;
- the first concrete step and the checkpoint for reviewing it.

If evidence is insufficient, give a conditional recommendation and name the smallest check that would settle it. If the honest answer is "this is a value choice, not an evidence question", say that and name who must decide. Do not present a guess as a verdict.

If the user has authorized action, carry out the next step within that scope; otherwise stop at the recommendation.

## Boundaries

- **Description is not explanation, and explanation is not a solution.** Naming a problem, critiquing a proposal, or describing a discourse can be necessary, but none of them is a recommendation. Do not stop at the trap of describing victims, critiquing without constructing, or "going meta" when the user needs a way forward.
- **More analysis is not always progress.** If further inquiry cannot change the action, stop. Analysis can become avoidance; inaction is also a choice.
- **Do not mistake a method for a finding or a theory for evidence.** A framework that organizes the discussion has not established the conclusion.
- **Do not assume shared values or shared facts.** Check both. Where values are contested, the answer is a decision process, not a cleverer analysis.
- **Do not promise a definitive solution to a wicked problem.** Offer a better frame, a direction, a safe-to-fail next step, and a review trigger.
- **Urgent or safety-critical situations:** contain the immediate harm and involve the responsible authority or expert first; do not run an open-ended inquiry.
- **Personal crisis or venting:** respond to the person first. Do not convert distress into an analysis unless the user asks for one.

## Completion conditions

Stop when you can state all of the following, or explicitly say which cannot be obtained and why:

1. The working problem statement and the decision it serves.
2. The question type and problem structure, with the evidence for that classification.
3. The chosen frame, at least one alternative, and why the chosen one wins.
4. The inquiry mode and theory, why they fit, and their limits.
5. The decisive evidence, or the smallest check, with its status: verified, inferred, or unknown.
6. A recommendation with alternatives, assumptions, reversal conditions, first step, and checkpoint.
7. The residual uncertainty and what would change the answer.

An explicit "we cannot decide yet because X is a value choice or Y is missing" is a valid completion. A polished answer to the wrong question is not.

## Response

Lead with your read of the problem and your recommendation. Then give only what the user needs to judge it: the problem type and frame, why this approach, the decisive evidence, the alternatives and conditions, and the next step. Mark claims as verified, inferred, or unknown when the difference matters.

Match the detail to the stakes and the conversation. Do not narrate the framework, list every move, or force headings for a small question. Use a compact table only when comparing frames or options. Reply in the user's language unless they request otherwise.

For the source basis and the limits of these heuristics, see [references/foundations.md](references/foundations.md). For behavioral test cases, see [references/evaluation.md](references/evaluation.md).

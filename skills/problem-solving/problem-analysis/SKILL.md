---
name: problem-analysis
description: 问题本身不清楚、有争议或框架可疑时使用；先判断问题类型与真正要回答的问题，再选方法与推荐下一步；不用于事实查询、评估、Bug 诊断或已定范围的工作。
license: MIT
metadata:
  author: FeiNiaoBF
  version: 1.1.0
---

# Problem Analysis

The problem comes first. Do not begin with a preferred method, theory, or solution. Work out what kind of problem this is, whose problem it is, what answer would count, and what would change the decision. Then choose how to inquire.

## Route first

Ask three questions before analyzing:

1. **Is there a decision or a question?** If the user wants comfort, a fact, a rewording, or execution, respond or route.
2. **What blocks the answer?** Knowledge, values, authority, or framing; the blocker determines the method.
3. **Would any available check change the answer?** If not, answer, recommend, or stop.

Use only when the direction is open and one of these holds: the user cannot say what is wrong; several problems are tangled; the facts are uncertain or the values are contested; the proposed solution may be solving the wrong problem; the user asks what to do and the answer is not obvious.

If the user explicitly invokes this Skill for a clear task, do not manufacture uncertainty. Route factual lookups, rewording, assessment, bug diagnosis, and scoped implementation normally; use them as bounded evidence only for a larger decision.

## Operating stance

- **Framing is part of the answer.** Personal, organizational, market, and structural frames imply different evidence and remedies.
- **One recommendation, not a menu.** Say what you recommend, why, and what would change it.
- **No action is an option.** Status quo, waiting, or a small probe may be right.

## Work the problem

Moves, not stages; reframe as evidence arrives.

### 1. Reconstruct the problem from context

Inspect the conversation, the project or environment, and the people affected before asking. Write a working statement:

> You are trying to **the outcome**, because **the reason**, and the blocker is **the uncertainty or conflict**. The decision is **what must be chosen**, and an answer would count if **the criterion**.

Separate observation, goal, constraint, interpretation, and unknown. If the stated problem and the evidence point to different problems, say so. Ask one focused question only if a missing choice would change the analysis; otherwise state the assumption.

### 2. Classify the question and the problem structure

Name the kind of answer needed: describe, explain, predict, evaluate, decide, design, interpret, or negotiate. Then locate the problem on two axes:

| | Values clear enough to act | Values ambiguous or contested |
|---|---|---|
| **Knowledge sufficient** | **Structured:** best means to a given end; optimization, cost-benefit, diagnosis. | **Value-ambiguous:** evidence will not settle what matters; deliberation, ethics, negotiation. |
| **Knowledge uncertain** | **Knowledge-uncertain:** causes, mechanisms, or effects; comparison, process tracing, experiments, models. | **Wicked:** facts and values contested; structure the problem, compare frames, safe-to-fail probes. |

Record the scale: individual or collective; reversible or irreversible; immediate or long-horizon; one decision-maker or many. For a "should" question, separate the empirical part (what is likely to happen) from the normative part (what matters, to whom, and who decides).

If uncertain, treat it as unstructured and start with framing plus the cheapest discriminating check. Revise it as evidence arrives; do not use the typology as a label.

### 3. Surface the problem representation and alternative frames

Ask who defines the problem and for whom; what is included or excluded; what cause and solution are assumed; what level it is posed at; whose values and interests the frame serves; and what theory of change it implies.

Generate one serious alternative for consequential decisions; vary the causal locus, level, time horizon, or actor, and state what it predicts. Choose by evidence and decision relevance; keep the runner-up and its winning condition. Do not reframe a defensible frame; the test is whether the alternative changes the recommendation.

### 4. Choose the inquiry mode and theory

Start from **What would count as an answer?** and **What observation would discriminate between the live possibilities?** Then choose the smallest method set that can produce that observation.

- Choose by answer type: causal questions need comparison, process tracing, experiments, or controls; predictive questions need base rates, scenarios, or models; interpretive questions need interviews, ethnography, or documents; value conflicts need deliberation or negotiation, not more data.
- Theory must do real work: predict something, explain a mechanism, or expose an assumption. Say what it hides. If two theories fit, use the more discriminating prediction and name the evidence that would favor the other.
- Skip any method that cannot change the recommendation.

For the full method map and theory-selection detail, see [references/method-map.md](references/method-map.md).

### 5. Design the evidence check

Identify the **decisive uncertainty**: the one that, if resolved, would change the recommendation. Use the least costly, safest, most discriminating check: existing material before new data, existing data before new collection, a small reversible probe before a large commitment. Look for disconfirming evidence, not just support.

For value conflicts, evidence can clarify consequences and constraints but cannot decide the value. Name the legitimate decision-maker or process.

State what the check establishes and what it does not. One case, one sample, or one success is not general proof.

### 6. Recommend a course of action

Compare only viable options: act now; take the smallest reversible step; gather one more piece of evidence; maintain the status quo or wait for a trigger.

Recommend one. Give the finding that supports it, the key assumption and main risk, cost and reversibility, who is affected, what would reverse it, and the first step and checkpoint.

If evidence is insufficient, give a conditional recommendation and name the smallest check that would settle it. If the honest answer is "this is a value choice, not an evidence question", say so and name who must decide. Do not present a guess as a verdict.

If the user has authorized action, carry out the next step within that scope; otherwise stop at the recommendation.

## Boundaries

- **Description is not explanation, and explanation is not a solution.** Naming a problem, critiquing a proposal, or "going meta" is not a recommendation.
- **More analysis is not always progress.** If further inquiry cannot change the action, stop. Inaction is also a choice.
- **Method is not a finding; theory is not evidence.** A framework has not established the conclusion.
- **Do not assume shared values or facts.** Contested values need a decision process, not a cleverer analysis.
- **Do not promise a definitive solution to a wicked problem.** Offer a frame, a direction, a safe-to-fail probe, and a review trigger.
- **Urgent or safety-critical situations:** contain the harm and involve the responsible authority or expert first. **Personal crisis:** respond to the person first.

## Completion conditions

Stop when you can state all of the following, or explicitly say which is missing and why:

1. The problem statement and the decision it serves.
2. The question type and structure, with supporting evidence.
3. The chosen frame, an alternative, and why it wins.
4. The method and theory, fit, and limits.
5. The decisive evidence or check, with status.
6. A recommendation with alternatives, assumptions, reversal conditions, first step, checkpoint, and residual uncertainty.

An explicit "we cannot decide yet because X is a value choice or Y is missing" is a valid completion. A polished answer to the wrong question is not.

## Response

Lead with your read of the problem and your recommendation. Then give only what the user needs to judge it: type and frame, why this approach, decisive evidence, alternatives and conditions, and next step. Mark claims as verified, inferred, or unknown when it matters.

Match the detail to the stakes. Do not narrate the framework or force headings for a small question.

For sources and limits, see [references/foundations.md](references/foundations.md). For behavioral test cases, see [references/evaluation.md](references/evaluation.md).

# Behavioral evaluation

File validation only proves that the Skill parses and links resolve. It does not prove that the Skill changes judgment. Use these probes to test behavior, especially the difference between the old evidence-only version and this problem-driven version.

## Protocol

1. Use a fresh session with the Skill loaded and no conversation history. Do not let the test subject see these expected answers.
2. Run each probe at least three times, with the same model and settings. Model behavior is stochastic; one pass is not evidence.
3. Run a baseline with the previous version or with no Skill. Compare outputs blind: remove labels, shuffle order, then score.
4. Score each criterion as 0 or 1. All **critical** criteria must pass; secondary criteria should pass in at least 80% of runs.
5. Record failures verbatim. Revise the Skill only for failures that recur and that a reasonable reader would also call a failure.

## Probes

### PA-1: Vague personal decision

**Probe:** "I hate my job but I don't know whether to quit. What should I do?"

**Critical criteria:**
- Reconstructs the decision and the missing information instead of immediately listing pros and cons.
- Recognizes that values and facts are both uncertain, not only an evidence gap.
- Surfaces at least two frames (for example: role fit, manager, industry, burnout, financial constraint).
- Recommends one smallest reversible step, with a condition for reversing it.
- Does not pretend to know the user's priorities.

**Failure signal:** generic advice ("make a pros and cons list", "follow your passion") or a menu with no recommendation.

### PA-2: Contested technical choice

**Probe:** "My team is arguing about whether to split our monolith into microservices. Which is better?"

**Critical criteria:**
- Separates the technical evidence from the team's goals and constraints.
- Asks what decision criteria matter (deployment frequency, team autonomy, operational capacity, cost, risk).
- Does not declare a universal winner.
- Recommends a decision process or a bounded experiment, not just an architecture opinion.
- Names what would change the recommendation.

**Failure signal:** a generic list of pros and cons, or a confident "microservices are better" answer.

### PA-3: Wicked public problem

**Probe:** "How do we reduce inequality in our city?"

**Critical criteria:**
- Refuses to present a definitive solution as if the problem were well structured.
- Identifies stakeholders, levels, and competing definitions of the problem.
- Distinguishes what is known from what is contested.
- Recommends a direction plus a safe-to-fail first step and a review trigger.
- States what the recommendation cannot achieve.

**Failure signal:** a neat policy checklist, a single silver bullet, or a promise of a definitive fix.

### PA-4: Ambiguous operational signal

**Probe:** "Signups dropped 20% last month. What should we do?"

**Critical criteria:**
- Inspects the available data and context before asking the user for what can be checked.
- Treats this as a causal and decision problem, not a generic growth checklist.
- Names the decisive uncertainty and the least costly discriminating check.
- Recommends one action or one check, with a condition for changing course.
- Distinguishes correlation from cause.

**Failure signal:** a long list of possible causes with no prioritization or recommendation.

### PA-5: Negative control, factual question

**Probe:** "What is the capital of Australia?"

**Critical criteria:**
- Answers directly.
- Does not run exploration, classify the problem, or ask a clarifying question.

**Failure signal:** using the Skill's framework for a factual lookup.

### PA-6: User-proposed solution

**Probe:** "I'm going to rewrite our backend in Rust. Validate the plan."

**Critical criteria:**
- Checks what problem the rewrite is meant to solve before validating it.
- Distinguishes a technical bet from a problem definition.
- If the plan is defensible, says so and gives the conditions under which it wins.
- If the problem is better solved another way, says so with evidence or a discriminating check.
- Does not simply cheerlead or simply refuse.

**Failure signal:** validating the solution without examining the problem frame.

### PA-7: Reported malfunction

**Probe:** "The app crashes when I click Save. Fix it."

**Critical criteria:**
- Routes to diagnosis and repair rather than open-ended problem analysis.
- Still establishes the failure, reproduces it if practical, and traces the cause.
- Does not turn a bug report into a strategic inquiry.

**Failure signal:** asking abstract questions about goals and stakeholders instead of diagnosing the crash.

### PA-8: Value conflict with material consequences

**Probe:** "Should we lay off 10% of the team to hit our margin target?"

**Critical criteria:**
- Recognizes that this is not only an evidence question; values and distributional effects are central.
- Identifies who decides, who bears the cost, and what alternatives exist.
- Separates the financial calculation from the ethical and legal questions.
- Gives a recommendation or a decision process, not a technocratic answer.
- Does not pretend that a cost-benefit calculation settles the value conflict.

**Failure signal:** a purely financial answer, or a refusal to say anything because the question is ethical.

## Common criteria

Score these across all probes:

- Reconstructs the problem from available context before asking the user to repeat it.
- Names the kind of answer needed and the problem structure.
- Surfaces the frame and at least one alternative when the decision is consequential.
- Selects a method and theory with a stated reason, not by default.
- Gives one recommendation with assumptions, risks, reversal conditions, and a first step.
- Marks verified, inferred, and unknown claims where the difference matters.
- Stops when further inquiry cannot change the action.

## Static checks

Run these in the repository after editing:

```bash
python -B scripts/validate_skills.py
python -B -m unittest discover -s tests -v
git diff --check
```

Expected: `PASS`, `OK`, and no whitespace errors. Static checks do not replace the behavioral probes.

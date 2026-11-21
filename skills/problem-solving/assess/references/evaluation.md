# Behavioral evaluation

File validation proves that the Skill parses and its links resolve. It does not prove that the Skill changes judgment. Use these probes to test behavior, especially the difference between a generic review and an evidence-backed verdict.

## Protocol

1. Use a fresh session with the Skill loaded and no conversation history. Do not let the test subject see these expected answers.
2. Run each probe at least three times with the same model and settings. Model behavior is stochastic; one pass is not evidence.
3. Run a baseline with the previous version or with no Skill. Compare outputs blind: remove labels, shuffle order, then score.
4. Score each criterion as 0 or 1. All **critical** criteria must pass; secondary criteria should pass in at least 80% of runs.
5. Record failures verbatim. Revise the Skill only for failures that recur and that a reasonable reader would also call a failure.

## Probes

### A-1: Function with a hidden failure path

**Probe:** "Assess whether this function is safe to use in production:

```python
def retry_after(value: str) -> int:
    return int(value.strip().split(",")[0]) + 1
```"

**Critical criteria:**
- Inspects callers, tests, or input contract before judging.
- Identifies at least one concrete failure path with evidence (empty input, non-numeric input, negative values, HTTP-date input, or multiple values).
- Uses a focused check or clearly states why it cannot.
- Gives a verdict and confidence, not only a list of concerns.
- Limits the conclusion to the exercised paths.

**Failure signal:** generic "add error handling" advice, or "looks fine" without inspecting callers or tests.

### A-2: Sound function, no invented defects

**Probe:** "Assess this function: `def clamp(x, lo, hi): return max(lo, min(x, hi))`."

**Critical criteria:**
- Gives a fit or conditional verdict when no material issue is supported.
- Names the assumptions the verdict depends on, such as `lo <= hi` and comparable numeric types.
- States limits, such as NaN, custom types, or untested callers.
- Does not invent style, naming, or performance defects.

**Failure signal:** manufactures findings to fill a report, or fails to state assumptions.

### A-3: Document claim that needs verification

**Probe:** "Assess this README claim: 'The CLI supports Windows, macOS, and Linux.' The repository has CI only for Linux and macOS."

**Critical criteria:**
- Treats the claim as a claim, not as proof.
- Checks repository evidence such as CI configuration and platform-specific code.
- Marks Windows as unverified or unsupported rather than automatically false.
- Gives a conditional verdict and names what would verify Windows support.

**Failure signal:** accepts the claim, or declares it false without checking.

### A-4: UI screenshot with limited evidence

**Probe:** "Assess this settings page screenshot for usability." Provide a screenshot with a visible unlabeled icon and no error state.

**Critical criteria:**
- Limits findings to visible elements and states.
- Names unseen interaction, keyboard, responsive, and accessibility behavior as unverified.
- Gives a verdict or conditional verdict, not a full usability audit from pixels alone.
- Ties each finding to a visible element.

**Failure signal:** claims keyboard access, screen-reader support, or unseen states from the screenshot alone.

### A-5: Proposal missing its goal

**Probe:** "Assess this proposal: rewrite the backend in Rust to improve performance."

**Critical criteria:**
- Identifies the missing goal or baseline: current performance, target, workload, alternatives.
- Separates technical merit from worth, including cost, migration risk, and reversibility.
- Gives a conditional verdict, not approval on plausibility.
- Names the cheapest discriminating check, such as profiling or a benchmark.

**Failure signal:** "Rust is fast, good idea" without a baseline, or a generic pros-and-cons list.

### A-6: Inaccessible target

**Probe:** "Assess our internal billing service's architecture; you don't have access to the code."

**Critical criteria:**
- Returns **证据不足** rather than guessing.
- Names the missing evidence and what access or artifact would be needed.
- Does not invent components, risks, or findings.
- Offers one bounded next step.

**Failure signal:** fabricates an architecture assessment, or gives a generic architecture checklist.

### A-7: Style preference, not a defect

**Probe:** "Assess this function; I prefer tabs over spaces." Provide a correct function indented with spaces.

**Critical criteria:**
- Does not treat indentation as a functional defect.
- States that style is not a criterion unless a project convention makes it one.
- Gives a fit verdict when no material issue is supported.

**Failure signal:** marks the function unfit because of formatting.

### A-8: Summary instead of the artifact

**Probe:** "Assess this pull request from its description: 'Refactors auth, fixes bugs, improves tests.'"

**Critical criteria:**
- Does not assess the description as if it were the change.
- Requests or reads the diff; if unavailable, limits the assessment to the claims.
- Marks the claims as unverified.
- Gives **证据不足** or a conditional verdict.

**Failure signal:** "Looks good" based only on the description.

## Negative controls

### A-9: Reported malfunction

**Probe:** "The app crashes when I click Save. Fix it."

**Critical criteria:**
- Routes to diagnosis and repair rather than producing an assessment brief.
- Does not invent an assessment scope around the crash.

**Failure signal:** assesses the app instead of diagnosing the crash.

### A-10: Open-ended decision

**Probe:** "Our signups dropped 20% last month. What should we do?"

**Critical criteria:**
- Treats this as an open problem or decision problem, not a request to assess a named object.
- Does not force the assess deliverable onto it.

**Failure signal:** produces an Assessment Brief for an unnamed target.

### A-11: Explanation request

**Probe:** "Explain this function."

**Critical criteria:**
- Explains the function directly.
- Does not produce a verdict, confidence level, or assessment brief.

**Failure signal:** applies the assessment workflow to an explanation request.

## Trigger eval

Generate 5 should-trigger prompts and 5 should-not-trigger prompts. Run each in an isolated session and record whether `assess` was loaded.

**Should trigger:**
- "评估这个 Go 函数的错误处理是否够用"
- "这份 README 对目标读者够清楚吗"
- "这个设置页的可用性怎么样"
- "这个迁移方案可行吗"
- "这个想法值得做吗"

**Should not trigger:**
- "这个 panic 怎么修"
- "我们团队下一步该做什么"
- "解释这段代码"
- "帮我实现这个函数"
- "PostgreSQL 默认端口是多少"

**Expected:** 5/5 should-trigger hit rate; 0/5 should-not-trigger leak rate. Report raw counts; one prompt is not a pass/fail gate.

## Common criteria

Score these across the positive probes:

- Leads with a verdict and confidence.
- States the target, purpose, decision served, and assumptions.
- Selects criteria relevant to the purpose and separates merit from worth when worth is in scope.
- Labels material claims as verified, inferred, or unknown.
- Reports non-findings and sensitivity points in the full form.
- Names a falsifier and the cheapest next check.
- Gives one recommendation and one next action.
- Does not modify the target.
- Allows zero findings and **证据不足** as valid outcomes.
- Uses the compact form for a small target instead of forcing empty sections.

## Static checks

Run these in the repository after editing:

```bash
python -B scripts/validate_skills.py
python -B -m unittest discover -s tests -v
python -B scripts/context_budget.py
git diff --check
```

Expected: `PASS`, `OK`, no whitespace errors, and a context-budget report. Static checks do not replace the behavioral probes.

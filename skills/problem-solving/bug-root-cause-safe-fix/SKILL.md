---
name: bug-root-cause-safe-fix
description: >-
  Use when the user reports a software bug, error, regression, or unexpected
  behavior and wants diagnosis or a fix. Trace evidence through logs, code, and
  call paths to identify the root cause and impact; make the smallest safe fix
  and verify it with relevant tests. Explain the cause, change, verification,
  and prevention clearly.
license: MIT
metadata:
  version: 1.0.0
---

# Bug Root Cause and Safe Fix

Fixing the bug is the goal. Do not stop at explaining the theory or hiding the
symptom. Find the cause, understand the affected paths, make a small complete
change, and verify the result.

## Workflow

1. **Establish the failure**
   - Read the reported behavior, expected behavior, reproduction steps, logs,
     stack traces, and error messages.
   - Inspect the relevant implementation, configuration, tests, and call paths.
     Trace how the triggering input or state reaches the failure.
   - Reproduce the issue when practical. If it cannot be reproduced, continue
     with code and evidence-based investigation; state what remains uncertain.

2. **Identify the root cause and scope**
   - Separate the immediate trigger from the underlying defect. Tie conclusions
     to observed evidence and keep hypotheses distinct from confirmed facts.
   - Check callers, shared state, data formats, neighboring paths, and error
     handling to determine which features and inputs may be affected.
   - Consider the proposed change's behavior and regression risks. Where relevant,
     inspect boundary conditions, validation, authorization, resource lifecycle,
     concurrency, and failure paths.

3. **Choose a small, complete fix**
   - Prefer the narrowest change that removes the root cause, not a workaround
     that only hides one example or an unrelated refactor.
   - Preserve existing interfaces and behavior unless changing them is necessary
     to fix the cause; explain any necessary behavior change.
   - If multiple fixes are viable, compare correctness, impact, and regression
     risk, then use the best-supported option.

4. **Implement and verify**
   - Make the fix and add or update regression coverage for the original failure
     and relevant edge cases.
   - Run the focused tests first, then broader tests, static checks, or builds
     when useful and available.
   - Review the diff for unrelated changes and recheck the original failure path
     plus adjacent behavior.
   - If verification fails, investigate and continue fixing. If a check cannot
     run, state why. Report only checks actually performed and their observed
     results; never imply unrun checks passed.

5. **Explain the outcome**
   - Briefly tell the user why the bug occurred, what changed, and what was
     verified and observed.
   - Call out material remaining uncertainty or uncovered risk. Give practical
     prevention advice tied to the cause, such as a specific validation,
     regression test, or monitoring signal.
   - Keep theory in service of the repair; do not let a long explanation replace
     the fix.

## Principles

- Repair first: the desired outcome is a safe, correct, complete fix.
- Root cause over symptom suppression.
- Control risk with the smallest complete change and a clear impact assessment.
- Verify the behavior and relevant regressions with evidence.
- Be direct about uncertainty, test limitations, and remaining risk.

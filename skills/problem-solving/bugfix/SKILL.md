---
name: bugfix
description: 用户报告软件缺陷、报错、回归或异常行为并要诊断或修复时使用；追到根因，做最小修复并验证回归。
license: MIT
metadata:
  version: 1.0.0
---

# Bugfix

The goal is the root cause gone, the smallest complete change, and regression evidence for the original failure. Do not stop at a theory or hide the symptom.

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
   - Verify the original failure path and adjacent behavior; choose test scope by regression risk.
   - Review the diff for unrelated changes.
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

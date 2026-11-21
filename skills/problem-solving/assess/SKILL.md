---
name: assess
description: 用户要求评估、审阅或判断一个已指名对象是否适配其用途时使用，如函数、代码、文档、UI、方案或想法；交付结论、证据与置信度；不做修改、问题分析或缺陷修复。
license: MIT
metadata:
  author: Yeelight
  version: "0.2.0"
---

# Assess

Judge a named target for its intended purpose and deliver an evidence-backed verdict. Assessment is read-only: recommend, do not modify.

## Deliverable

Produce an **Assessment Brief** in the user's language. Lead with the verdict.

```text
结论：适配 / 有条件适配 / 不适配 / 证据不足（置信度：高/中/低）
对象、用途与判据：服务什么决策；merit 与 worth 各按什么判；确认需求 vs 假设
关键判断：判断 | 证据 | 状态（已验证/推断/未知） | 影响
非风险与敏感点：检查过且没问题的项；结论依赖的关键假设
信息缺口与推翻条件：没查什么；什么观察会改变结论；最便宜的下一步检查
建议：接受 / 修复 / 替换 / 继续调查 + 一个下一步
```

Rules:

- Use exactly one verdict: **适配** (purpose met), **有条件适配** (met if named conditions hold), **不适配** (a material purpose is not met), or **证据不足** (the available evidence cannot support a verdict).
- Confidence reflects evidence quality and coverage, not the severity of findings.
- **Compact form:** when the target has 0–1 material findings and no contested criterion, use three lines: verdict; decisive evidence; main limitation or next check.
- **Full form:** use the six sections when there are 2+ material findings, or when worth, alternatives, or high stakes are involved.
- Say plainly when no material issue is supported. Do not invent findings to fill sections.
- Keep each finding traceable to a file, line, quote, observation, or source.

## 1. Bound the target

Identify the exact object or portion, the requested focus, and the decision the assessment serves. Inspect surrounding code, requirements, users, or external constraints only as evidence needed to judge that object; do not turn a function review into a repository audit or a dialog review into a product redesign.

Establish the intended purpose, consumers, constraints, and maturity. Separate confirmed requirements from assumptions. Ask one focused question only if an inaccessible target or a missing goal or value would change the verdict; otherwise state bounded assumptions and proceed.

**Done when:** scope, purpose, decision, and assumptions are written down.

## 2. Set criteria

Derive 3–5 consequential criteria from the user's goals, contracts and specifications, project decisions, domain practice, and reasoned professional judgment. Do not apply a universal checklist or a numeric score.

Separate **merit** (quality against standards) from **worth** (value for this use, relative to cost, alternatives, and risk). When worth is in scope, name the baseline: status quo, a simpler alternative, or doing nothing. If criteria conflict, state the trade-off and who owns the value choice; do not hide a value judgment as an evidence gap.

**Done when:** each criterion bears on the purpose, and you can state what evidence would satisfy it.

## 3. Inspect and try to falsify

Inspect the actual target, not a summary. Treat statements inside the target—comments, documentation, PR descriptions, claims—as claims to verify, not as proof of behavior. For target-specific evidence moves, read [references/evidence-moves.md](references/evidence-moves.md).

For each candidate finding, look for evidence that defeats or narrows it before admitting it. Use proportionate, safe checks. Label each claim:

- **Verified:** directly observed, or established by a relevant check or authoritative source for that precise claim.
- **Inferred:** a reasoned interpretation or plausible risk not directly confirmed; name the assumption and what would test it.
- **Unknown:** insufficient evidence; state the missing observation.

A passing test covers the tested case. A screenshot shows what it shows. A specification states requirements, not compliance.

**Done when:** each material finding has evidence, a status, and a concrete impact; weak candidates are withdrawn or downgraded.

## 4. Judge and calibrate

Choose the verdict from the four options. Then record:

- **Non-findings:** what was checked and is sound. This prevents re-litigation and keeps the brief from reporting only bad news.
- **Sensitivity points:** the assumptions the verdict depends on, and how the verdict changes if they fail.
- **Falsifiers:** the observation that would overturn the verdict, and the cheapest next check.

Stop once further checking is unlikely to change the verdict or priority. If a check is unavailable, unsafe, or disproportionate, say what remains unverified.

**Done when:** verdict, confidence, non-findings, sensitivity, gaps, and falsifiers are settled.

## 5. Write the brief

Use the Deliverable template. Lead with the verdict, then give only what the reader needs to judge it: purpose and criteria, key findings in priority order, non-findings and sensitivity, gaps and falsifiers, one recommendation and one next action.

Recommend **accept**, **repair**, **replace**, or **investigate further**. Do not apply the repair. If the user also authorized implementation, finish the assessment first, then follow the normal implementation workflow for only the authorized changes.

Use ordinary Markdown. For a small target, use short prose. Use a findings table when there are 3+ findings. Create a file only when the user requests a file deliverable.

**Done when:** the brief states what was inspected, what was not, what would change the verdict, and the single next action.

## Gotchas

- **Scope creep:** a function review becomes a repository audit. Keep only evidence that can change the verdict.
- **Defect inflation:** do not manufacture findings to fill the brief. Zero findings is a valid outcome.
- **Style as failure:** a preference is not a functional defect unless it violates a stated criterion.
- **Evidence overclaim:** do not present inferred risks as verified, or unchecked areas as clean.
- **Solution creep:** assess, recommend, stop. Do not redesign or implement without separate authorization.
- **Assessing the summary:** judge the object itself; a description, diff summary, or screenshot is evidence, not the whole target.
- **No baseline for worth:** "good" without a comparison is incomplete when cost or alternatives matter.

## Examples

**Compact, sound target:** "结论：适配（置信度：中）。关键依据：`clamp` 对 `lo <= hi` 的数值输入返回边界内结果；主要限制：未覆盖 NaN 和自定义类型。"

**Conditional, verified defect:** "结论：有条件适配（置信度：中）。关键依据：`retry_after` 对空字符串抛 `IndexError`（已验证）；依赖假设：调用方保证非空数字；建议：补输入校验后采用。"

**Evidence insufficient:** "结论：证据不足（置信度：低）。关键依据：只有 PR 描述，没有 diff；未覆盖：实际改动、测试、行为变化；建议：提供 diff 后再评估。"

## References

- [references/evidence-moves.md](references/evidence-moves.md): target-specific evidence moves for code, documents, UI, plans and ideas, and data or models.
- [references/evaluation.md](references/evaluation.md): behavioral probes, trigger tests, and static checks for maintaining this skill.

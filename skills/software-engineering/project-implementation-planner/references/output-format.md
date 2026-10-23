# Planning output contract

Read this file before delivering or revising an implementation plan. It is the authoritative output format; [the example](output-example.md) illustrates it without adding requirements.

## Format and scope

- Deliver a Markdown plan with the six numbered level-two headings below, in that order. Keep headings, field labels, and table columns stable. Scale the content and number of slices to the task, including a single slice when appropriate; retain empty sections with an explicit `None` or `Not yet known` and its consequence.
- Use the user's language. Chinese and English labels are provided below; for another language, translate once and keep that vocabulary throughout revisions. Keep identifiers and project terminology unchanged.
- Use this format for full plan deliveries and revisions. A narrow follow-up question or interim progress update can be answered directly; it does not require regenerating the plan.
- An explicit user template or mandatory project format takes precedence. Map this contract's information into that format and state the mapping once; do not create a competing plan or silently change formats on later turns.
- By default, deliver in the conversation. Save or update a project document only when authorized or required by project instructions. The saved plan uses the same format. When delivering a saved plan, use a short receipt with exactly three fields: `Plan / 计划` (link), `Changes / 本次变更` (what changed), and `Next / 下一步` (slice ID and any blocker). Do not duplicate the full plan in chat.

## Writing for human readers

Apply these rules to the plan's prose, table cells, and delivery receipt. Preserve the required headings, labels, IDs, and statuses. These are adaptations of Orwell's six rules for technical planning in the user's language, not quotations or rigid word-count limits:

1. Replace stock metaphors and slogans with the action or result they stand for. Explain what “打通链路”, “形成闭环”, or “夯实底座” means in this project; omit the phrase if it adds no meaning.
2. Prefer familiar, concrete words when they express the same meaning. Write “读取文件并汇总金额” instead of “实施数据读取与聚合能力建设”. Keep precise technical terms when a simpler substitute would change the meaning.
3. Cut words and sentences that add no information. Remove ceremonial openings, repeated conclusions, and empty assurances such as “全面提升健壮性”. Preserve reasons, conditions, uncertainty, and failure behavior that affect the decision.
4. Prefer active sentences that name the actor and action: “程序报告错误行” instead of “错误行将被报告”. If the actor is unknown or irrelevant, keep a natural passive sentence; do not invent an owner to satisfy the rule.
5. Use the reader's vocabulary. Explain an unfamiliar term or abbreviation briefly at first use, tied to its role here; for example, a slice is a piece of work that can be implemented and checked on its own. Reuse one term consistently. Preserve API names, code identifiers, paths, and commands exactly.
6. Put accuracy and natural expression ahead of stylistic shortcuts. Keep a longer word, passive sentence, or necessary repetition when it prevents ambiguity. Simplifying prose must not turn a proposal into a commitment, an assumption into a fact, or a planned check into a reported success.

Within each field, put the action or conclusion before its explanation. State dependencies as concrete reasons: “S2 uses the file reader from S1” explains more than “S1 lays a solid foundation”. Describe a risk through its condition, effect, and response. If the needed detail is unknown, say what is missing rather than inventing specificity.

Keep one main point per sentence. Split a sentence when it combines an action, several conditions, a rationale, and a caveat; keep related conditions beside the action they qualify. Vary sentence length naturally instead of imposing a maximum. Use the sequence table for short summaries and section 4 for execution details. Keep the fixed field labels plain; use numbered steps for actions and ordinary prose for explanations. Avoid decorative emphasis, artificial contrasts, and closing slogans.

### Wording examples

These pairs illustrate wording changes using the same underlying facts. They do not prescribe project requirements.

| Draft | Clearer wording |
| --- | --- |
| 先用最小可运行命令贯通读取与汇总。 | 先让命令能读取 CSV 并输出分类总额。 |
| 币种约定仍是未验证假设，影响 S1 的金额语义。 | 尚未确认一个文件是否只含一种币种；这决定了 S1 能否直接相加。 |
| 缺少真实样本，因此当前不具备全面验证银行格式适配能力的条件。 | 尚无真实样本，暂时无法验证银行导出格式。 |

## Canonical template

Replace placeholders with evidence or explicitly bounded proposals. Do not output the template's instructions or invent requirements, paths, commands, dates, estimates, or completed work to fill fields.

```markdown
# Implementation Plan: <project or outcome>

## 1. Goal and recommendation

- Goal: <observable outcome and why it matters>
- Scope: <included work; explicit exclusions and constraints>
- Recommendation: <next slice ID, why it comes first, and planning-only or authorized implementation scope>

## 2. Current state and evidence

- Confirmed: <relevant behavior or constraint, with file, observation, or user-provided source>
- Assumptions: <unverified premise and what changes if it is false; or none>
- Unknowns: <missing evidence, affected slice IDs, and whether it blocks them; or none>

## 3. Implementation sequence

| ID | Outcome | Dependencies | Why this order | Acceptance evidence | Status |
| --- | --- | --- | --- | --- | --- |
| S1 | <observable result> | <slice IDs or none> | <value, dependency, or risk> | <check and expected observation> | Planned |

## 4. Next slice

- Slice: <one ID from section 3>
- Behavior and limits: <expected behavior and what this slice excludes>
- Prerequisites: <settled inputs and any blocking decision; or none>
- Affected areas: <verified paths, interfaces, data, or explicitly proposed locations>
- Actions: <short ordered steps sufficient to start; defer unrelated design details>
- Verification: <check or command, working directory when relevant, and expected observation; mark unverified commands as proposed>
- Stop or replan when: <specific evidence that changes the approach or prevents completion; or no additional condition>

## 5. Risks and open decisions

- <Risk or decision; confirmed/proposed/unresolved; affected slice IDs; mitigation or smallest useful check; or no known material issue>

## 6. Delivery record

- Inspected: <files, user materials, and checks actually examined or run; relevant coverage limits>
- Created or updated: <actual artifacts and their authoritative location; or conversation-only plan, no files changed>
- Change and reason: <initial plan, or what changed in this revision and why>
- Next action: <immediate action and responsible party when known; any one blocking question; no automatic request for approval already granted>
```

## Chinese labels

For Chinese plans use `# 实施计划：<项目或目标>` and these exact headings and labels. These are translations of the same template, not a separate output mode.

| English heading | Chinese heading | Chinese field labels, in order |
| --- | --- | --- |
| 1. Goal and recommendation | 1. 目标与建议 | 目标、范围、建议 |
| 2. Current state and evidence | 2. 现状与依据 | 已确认、假设、未知项 |
| 3. Implementation sequence | 3. 实施顺序 | 表头：编号、可见结果、依赖、排序理由、验收证据、状态 |
| 4. Next slice | 4. 下一切片 | 切片、行为与边界、前置条件、影响范围、执行步骤、验证方式、停止或重排条件 |
| 5. Risks and open decisions | 5. 风险与待决事项 | 每项写明风险或决策、确认状态、影响切片及应对方式；无则注明 |
| 6. Delivery record | 6. 交付记录 | 已检查、本次产物、变更与原因、下一步行动 |

## Sequence and revision rules

- Assign stable IDs (`S1`, `S2`, …). Row order expresses priority; IDs do not. Preserve IDs when reordering, add a new ID for new work, and retain cancelled rows with a reason when revising an existing plan. Dependencies must name existing IDs, be acyclic, and be satisfied before a dependent slice starts. Mark work as parallel only when shared contracts and resources permit it.
- Use only these status labels: `Planned / 待实施`, `In progress / 进行中`, `Blocked / 受阻`, `Done / 已完成`, `Deferred / 暂缓`, `Cancelled / 已取消`. A proposal starts as Planned or Blocked; use Done only with observed completion evidence. A plan or test command written down is not execution evidence.
- Section 3 contains an overview of every currently justified slice. Section 4 expands exactly the next actionable slice without repeating the whole roadmap. If all remaining work is blocked, name the nearest blocked slice and its smallest unblock action; keep independent work actionable. If the plan is complete, retain section 4 with `Slice: None — plan complete` (`切片：无，计划已完成`) and omit its other fields. This is the only completed-plan exception; do not invent another slice.
- On revision, update dependent rows, the next slice, risks, and the delivery record together. Preserve the document location, heading language, and labels. Record the change and its reason; use Git or the existing issue history for older revisions instead of appending duplicate plans.
- Put additional material inside the relevant section: diagrams under sequence, interface details under the next slice, decision links under risks. Add no competing summary, roadmap, or top-level appendix by default. Fixed structure does not impose a stack-first sequence, a minimum number of stages, or mandatory project documents.

## Delivery check

Before sending, compare the output to the template: all six headings in order; all fields and table columns present (except the completed-plan exception above); no unfilled placeholders; evidence distinguished from assumptions; stable IDs and valid dependencies; a startable next slice or explicit blocker; acceptance observations and honest status; accurate artifact and check claims. For a saved plan, check the file and then use the three-field receipt. Resolve format drift before delivery.

Then read it as the intended reader: can they tell what happens next, why it comes first, and how to recognize completion without decoding jargon? Revise abstract claims and overloaded sentences using the writing rules above. Compare the revision with the evidence and draft to ensure no constraint, dependency, qualification, identifier, or acceptance condition was lost or invented. Do not print this editorial checklist in the delivered plan.

## Basis and limits

This contract is a local design choice for predictable planning output, informed by these primary sources:

- [George Orwell: Politics and the English Language (1946)](https://www.orwellfoundation.com/the-orwell-foundation/orwell/essays-and-other-works/politics-and-the-english-language/): six writing rules on stale imagery, word choice, unnecessary words, active voice, jargon, and exceptions. The writing rules above adapt them to technical plans; the paragraph organization and template conventions are local guidance.
- [OpenAI Cookbook: Using PLANS.md for multi-hour problem solving](https://github.com/openai/openai-cookbook/blob/main/articles/codex_exec_plans.md): explicit skeleton, observable acceptance, and living-plan updates. It describes one prompting approach and invites customization. We adopt those practices without its mandatory long-form sections, blanket execution assumptions, or full replication of project documentation.
- [GitHub Spec Kit plan template](https://github.com/github/spec-kit/blob/main/templates/plan-template.md) and [task template](https://github.com/github/spec-kit/blob/main/templates/tasks-template.md): named fields, task IDs, dependency explanations, and independently verifiable checkpoints. We adopt these conventions without copying their predefined infrastructure phases, directory layout, or testing policy.

These sources demonstrate practical conventions; they do not prove that six sections are optimal or guarantee compliance by every model. The format is intended to be checkable. Validate stability with repeated representative runs when evaluating a host or model; repository link and metadata checks alone do not test generated plans.

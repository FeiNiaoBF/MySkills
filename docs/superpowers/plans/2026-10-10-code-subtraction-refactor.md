# Code Subtraction Skill 重构实施计划

> **For agentic workers:** 按任务逐步完成并记录复选框。每步结束后检查该步产物，再继续。

**Goal:** 让 `code-subtraction` 能依据证据分析并安全实施经授权的减法，同时避免固定阈值和不必要的重复流程。

**Architecture:** `SKILL.md` 作为短入口，定义减法任务的边界、决策流程和验证原则。`references/subtraction-playbook.md` 承载场景示例、可选报告模板，以及带出处和适用范围的理论/实证依据。

**Tech Stack:** Markdown、YAML frontmatter、仓库 Python skill validator 与 unittest 校验套件。

**Spec:** [docs/superpowers/specs/2026-10-10-code-subtraction-design.md](../specs/2026-10-10-code-subtraction-design.md)

## Global Constraints

- 兼顾准确安全与精简，不为缩短篇幅牺牲关键判断。
- 分析建议与用户授权的代码实施明确区分。
- 消融仅用于风险较高、行为影响不明或其他证据不足的情况。
- 按影响面、可逆性、兼容性与证据质量评估风险，不以固定次数阈值决定抽象或删除。
- 修改范围限于 `skills/software-engineering/code-subtraction/SKILL.md` 与 `skills/software-engineering/code-subtraction/references/subtraction-playbook.md`；保留 skill 名称、中文和显式调用策略。
- 理论与实证材料须标明来源和适用范围；不得宣称普遍安全删除证明。

## Review Focus

- **触发范围过宽**：“想重构”可能被误解为所有结构重构；检查 description 是否把触发限定在减法意图。
- **分析与实施授权混淆**：建议不应自动变成代码修改；检查入口流程是否明确区分。
- **风险标签变成机械门槛**：检查低/中/高等级是否作为摘要而非唯一判据。
- **验证被过度要求或过度相信**：检查测试不是唯一证据，且消融只用于高风险或证据不足情形。
- **研究结论被过度外推**：检查 Parnas、Linux 内核去重研究和 debloating 预印本均有清楚的范围说明。

---

### Task 1: 精简入口流程与触发定位

**Files:**
- Modify: `skills/software-engineering/code-subtraction/SKILL.md`

**Interfaces:**
- Consumes: 已审阅设计说明；当前入口文档及 `references/subtraction-playbook.md`。
- Produces: 一个保持现有 frontmatter 策略、以减法任务为中心的短流程，供 Task 2 的参考手册承接细节。

- [x] **Step 1: 重写 description 的触发条件**

把“想重构”限定为寻找不必要复杂度、删除或简化的意图；保留死代码、重复逻辑、过度抽象、冗余依赖和功能/API 下线等具体触发场景。保留 `name`、`disable-model-invocation: true`、中文定位和 `agents/openai.yaml` 的显式调用策略，不改其他文件。

**检查：**逐项核对 Review Focus 中“触发范围过宽”。

- [x] **Step 2: 将正文改成主决策流程**

保留减法核心原则；按“明确目标和范围 → 收集证据 → 综合评估风险 → 提出建议或实施获授权的删除 → 按影响验证并报告”组织入口。风险级别只作摘要；验证遵循项目约定的最小充分方法。把场景细节留在 playbook，并保留其相对链接。

**检查：**逐项核对 Review Focus 中“授权混淆”“机械风险门槛”和“验证要求”。

- [x] **Step 3: 手工复核入口边界**

确认概念性问题可以直接解释；未授权分析任务只输出建议；所有改动均在任务明确包含实施时才执行；不得把一般重构都路由到减法流程。

**检查：**阅读更新后的完整 `SKILL.md`，确认措辞简洁且每个阶段都有明确目的。

---

### Task 2: 重写场景手册并加入有边界的研究依据

**Files:**
- Modify: `skills/software-engineering/code-subtraction/references/subtraction-playbook.md`

**Interfaces:**
- Consumes: Task 1 定义的范围和流程。
- Produces: 与入口流程一致的按需参考，包括代码、依赖、模块场景，适用的高影响删除流程、可选报告模板，以及可追溯的理论/实证依据。

- [x] **Step 1: 按任务场景整理现有示例**

保留死代码、条件简化、依赖和模块等有用例子；重写“重复三次以上必须抽象”、半年无调用等固定门槛。重复只作为检查信号，按语义一致性、可读性、变更传播、性能及维护成本决定是否合并。

**检查：**确认示例不会暗示减少行数本身就是成功，也不会把局部内联或合并作为无条件建议。

- [x] **Step 2: 将风险、消融和高影响删除指导改为条件式**

以影响面、可逆性、兼容性和证据质量说明风险。消融写成行为不明或证据不足时的可回滚对照。保留公共 API、线上功能和共享服务的分阶段弃用、观察与回滚流程，并明确其仅适用于高影响情形。

**检查：**确认普通本地删除无需套用发布观察周期；测试不是唯一安全证据。

- [x] **Step 3: 加入研究依据、适用范围和来源**

在 playbook 中新增简短“依据与边界”小节，至少包括：Parnas 1979 软件收缩与依赖结构；2026 Linux 内核维护者去重研究对可读性、集成成本和性能权衡的观察；2026 代码 debloating 预印本关于测试驱动裁剪遗漏必要错误处理和并发代码的结果。使用设计说明中的 DOI/原文链接，标明 debloating 研究针对自动裁剪工具，Linux 研究针对特定内核评审情境；明确这些材料提供判断依据，不构成具体删除的证明。

**检查：**逐项核对 Review Focus 中“研究结论被过度外推”；来源 title、作者（有把握时）、年份、链接和研究类型须准确。

- [x] **Step 4: 保留可选报告模板并检查单一来源**

报告格式只作为需要正式建议时的可选模板；更新其字段以匹配新流程，避免与入口文件重复解释核心原则。

**检查：**确认入口链接仍指向本 skill 内真实文件；参考文档不重复缓存可从环境读取的信息。

---

### Task 3: 全量复核并执行仓库校验

**Files:**
- Review: `skills/software-engineering/code-subtraction/SKILL.md`
- Review: `skills/software-engineering/code-subtraction/references/subtraction-playbook.md`

**Interfaces:**
- Consumes: Task 1 和 Task 2 的更新文档。
- Produces: 通过仓库校验、链接检查与差异检查的聚焦文档改动。

- [x] **Step 1: 对照设计说明做覆盖检查**

逐项确认建议/实施边界、风险判断、按风险验证、可选消融、高影响删除流程、研究依据与限制、范围和调用策略均有对应文本。

**检查：**按 Review Focus 五项复核全文，并核对 `agents/openai.yaml`、README 不在改动范围。

- [x] **Step 2: 运行仓库要求的校验**

```bash
python -B scripts/validate_skills.py
git diff --check
```

**预期：**技能格式/链接校验与 diff whitespace 检查通过。已有或环境性失败单独报告，不放宽规则掩盖失败。

- [x] **Step 3: 检查最终 diff 范围**

确认差异只包含两个授权 skill 文档；核对源链接和引用没有错误，保留其他工作区改动，不顺带修复无关问题。

**检查：**`git status --short` 与完整目标文件 diff。

- [x] **Step 4: 提交聚焦改动**

仅暂存两个目标 skill 文档并提交：

```bash
git add skills/software-engineering/code-subtraction/SKILL.md skills/software-engineering/code-subtraction/references/subtraction-playbook.md
git commit -m "docs: clarify code subtraction skill workflow"
```

**预期：**提交只包含本次 skill 重构文件。

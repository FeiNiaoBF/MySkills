# Code Subtraction Skill 重构设计

日期：2026-10-10

## 目标

重构 `skills/software-engineering/code-subtraction/`，让它能同时指导减法分析和经用户授权后的实际删除；规则准确、风险相称，并减少重复说明。保留“识别并移除不必要复杂度”的聚焦定位。

## 已确认的约束

- 兼顾准确安全与精简，不为缩短篇幅牺牲关键判断。
- 用户明确要求实施时，可以指导执行删除；分析建议与实施任务要区分。
- 消融仅用于风险较高、行为影响不明或其他证据不足的情况。
- 不使用固定次数等阈值决定是否抽象或删除；根据影响面、可逆性和现有证据判断。
- 修改范围限于 `SKILL.md` 和 `references/subtraction-playbook.md`；保留名称、中文及显式调用策略。

## 设计

### 理论与实证依据

重构后的 skill 应给出可追溯、有限定条件的依据；研究不能证明“代码越少越好”，也不能替某个具体项目判定某段代码可删。Playbook 将用简短的“依据与边界”说明支撑关键规则：

- Parnas 将软件的收缩（移除功能子集）作为设计变更问题，指出依赖分散会让移除一种能力牵动多个组件，并讨论用模块边界隔离易变决策。[1] 这支撑检查依赖传播和可移除性，不支撑以减少行数为目标。
- Linux 内核维护者的 2026 实证研究观察到，去重可能因可读性下降、集成成本或性能顾虑而被拒绝；该研究只覆盖 Linux 内核中的特定补丁和评审情境。[2] 这支持把重复视作调查信号，而不是自动抽象阈值。
- 2026 年代码裁剪预印本以 11 个程序建立人工 ground truth，报告自动化、测试驱动的裁剪工具可能删除未被测试触发但仍必要的错误处理和并发代码。[3] 该结果针对自动 debloating，不能直接量化人工删除风险；它说明测试通过是证据之一，不能单独证明行为等价。

因此，skill 的理论定位是“有文献与实证依据的决策框架”，不是对安全删除的形式化证明。实施判断仍依赖目标系统的需求、调用关系、兼容性和验证证据。

**来源**

1. D. L. Parnas, “Designing Software for Ease of Extension and Contraction,” *IEEE Transactions on Software Engineering*, 5(2), 1979, pp. 128–138. [DOI](https://doi.org/10.1109/TSE.1979.234169); [原文 PDF](https://users.ece.utexas.edu/~perry/education/SE-Intro/parnas.pdf)。
2. L. Arcanjo et al., “When Do You Repeat Yourself? Voices from the Trenches of Linux Kernel Maintainers on Code Duplication,” *Proceedings of the 2026 ACM/IEEE International Conference on Technical Debt*, published 2026-07-16. [DOI](https://doi.org/10.1145/3794915.3795783)。
3. “Revisiting Code Debloating with Ground Truth-based Evaluation,” arXiv:2604.17717, 2026 preprint. [arXiv](https://arxiv.org/abs/2604.17717)。

### 入口流程

`SKILL.md` 保留核心原则和简短主流程：明确目标与范围、收集证据、判断风险、提出或实施最小删除、按影响验证并报告。说明减法分析可以只产出建议；只有任务授权包含实施时才修改代码。

风险等级保留为摘要，不作为机械门槛。指导根据影响面、可逆性、兼容性与证据质量综合判断。验证采用项目约定下的最小充分方式，按影响加深；测试不是唯一证据，也不要求每个候选项都进行消融。

### 参考手册

保留代码、依赖和模块层面的实例，以及可选的减法分析报告模板。重写固定次数等绝对规则，修正可能鼓励过度合并的示例。将消融描述为证据不足时的可回滚对照方法。

公共 API、线上功能和共享服务等高影响删除保留分阶段弃用、观察和回滚指导，并清楚限定适用场景；一般本地删除不需要套用发布流程。

### 调用与仓库索引

保持现有 `name`、description 的减法定位、`disable-model-invocation` 和 `agents/openai.yaml` 调用策略。本次不改 README；除非实施时发现对外定位必须同步，否则不扩展范围。

## 验收标准

- Skill 能区分分析建议与获授权的实施。
- 删除决策采用上下文和证据，不依赖重复次数、测试覆盖或消融等单一硬门槛。
- 验证要求与影响相称，风险较高或证据不足时提供更强验证手段。
- 高影响删除仍有适用范围明确的弃用、观察和回滚指导。
- 主入口精简，按需参考集中在 playbook；示例不把经验规则写成普遍定律。
- 减法原则有简短、可追溯的理论和实证依据，并说明证据适用范围与局限，不声称存在普遍适用的安全删除证明。
- Frontmatter、内部链接及仓库校验通过；改动仅涉及授权范围。

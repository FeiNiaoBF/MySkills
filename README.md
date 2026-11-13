# MySkills

公开的可复用 Agent Skills 创作源，采用 [MIT License](LICENSE)。每个 Skill 以 `SKILL.md` 为入口，为支持 Skills 的 Agent 提供特定任务的工作流程和按需参考。课程工程、媒体、缓存和用户环境配置留在各自项目中。

- 找到适合任务的技能：[选择 Skill](#选择-skill)。
- 在 Agent 中使用：[安装与更新](#安装与更新)，安装后[确认并使用](#确认安装并使用)。
- 修改或贡献本仓库：[贡献与验证](CONTRIBUTING.md)；Agent 另见 [工作约定](AGENTS.md)。

## 选择 Skill

### 视频（video）

| Skill | 使用场景 |
| --- | --- |
| [video-workflow](skills/video/video-workflow/SKILL.md) | 一集视频的阶段协调与交接；视频任务入口 |
| [video-edit](skills/video/video-edit/SKILL.md) | 本地转写、精简、字幕对齐与成片校验 |
| [math-review](skills/video/math-review/SKILL.md) | 数学表述、推导、口误与动画候选审校 |
| [manim-authoring](skills/video/manim-authoring/SKILL.md) | 设计并验证讲解中途的 Manim 场景 |
| [manim-assembly](skills/video/manim-assembly/SKILL.md) | 按字幕语义把场景装配为画面替换 |
| [math-recap](skills/video/math-recap/SKILL.md) | 片尾独立、无旁白的数学重温动画 |

中途画面替换使用 `manim-authoring` + `manim-assembly`，保持原声和总时长；片尾独立重温使用 `math-recap`，追加拼接。两条路线不要混用。单集状态由课程项目中的 `workflow.yaml` 维护。

### 设计（design）

| Skill | 使用场景 |
| --- | --- |
| [style-translator](skills/design/style-translator/SKILL.md) | 视觉偏好 → 设计方向、HTML 预览与平台规格 |
| [design-system](skills/design/design-system/SKILL.md) | UI 视觉语言、设计令牌与组件规范 |
| [ui-planner](skills/design/ui-planner/SKILL.md) | 选内容、排布局、出 HTML 初稿与实现提示词 |

### 学习（learning）

| Skill | 使用场景 |
| --- | --- |
| [linear-algebra](skills/learning/linear-algebra/SKILL.md) | 理解线性代数概念、推导与解题 |
| [tech-mentor](skills/learning/tech-mentor/SKILL.md) | 用递进项目掌握一项技术 |

### 沟通（communication）

| Skill | 使用场景 |
| --- | --- |
| [wording-helper](skills/communication/wording-helper/SKILL.md) | 把没说清的想法变成可直接使用的措辞 |
| [make-sense](skills/communication/make-sense/SKILL.md) | 解释卡住的术语、段落或推理 |
| [visual-explainer](skills/communication/visual-explainer/SKILL.md) | 用户点名时生成单文件 HTML 解释页 |

### 问题探索（problem-solving）

| Skill | 使用场景 |
| --- | --- |
| [assess](skills/problem-solving/assess/SKILL.md) | 按用途评估指定对象并给出证据化结论 |
| [problem-analysis](skills/problem-solving/problem-analysis/SKILL.md) | 澄清问题类型与真正要回答的问题 |
| [bugfix](skills/problem-solving/bugfix/SKILL.md) | 追根因、最小修复并验证回归 |

### 软件工程（software-engineering）

| Skill | 使用场景 |
| --- | --- |
| [code-simplification](skills/software-engineering/code-simplification/SKILL.md) | 用证据识别并安全删除多余复杂度 |
| [go-style](skills/software-engineering/go-style/SKILL.md) | 按内置规则直接改写或审查 Go 代码 |
| [implementation-planner](skills/software-engineering/implementation-planner/SKILL.md) | 手动启动的分阶段实施计划 |
| [project-docs](skills/software-engineering/project-docs/SKILL.md) | 编写与维护 README、需求、规格、架构、ADR |
| [project-guide](skills/software-engineering/project-guide/SKILL.md) | 理解现有项目的结构、流程与决策 |
| [skill-discovery](skills/software-engineering/skill-discovery/SKILL.md) | 查找、比较并按授权安装现成 Skill |

### 研究（research）

| Skill | 使用场景 |
| --- | --- |
| [research-report](skills/research/research-report/SKILL.md) | 多来源调查与可追溯的离线 HTML 报告 |

### 网络安全（network-security）

| Skill | 使用场景 |
| --- | --- |
| [network-diagnostics](skills/network-security/network-diagnostics/SKILL.md) | OpenWrt/PassWall/sing-box 只读诊断与最小变更 |

## 安装与更新

需要可用的 Node.js、npm（含 `npx`）和 Git，以及访问 npm 与 GitHub 的网络。安装由 [Vercel skills CLI](https://github.com/vercel-labs/skills) 完成，无需先克隆本仓库。各 Skill 的运行依赖以自身说明为准。

### 选择安装范围

先列出本仓库可安装的 Skill，此命令不会安装：

```bash
npx skills add FeiNiaoBF/MySkills --list
```

在交互式终端中选择 Skill 和目标 Agent：

```bash
npx skills add FeiNiaoBF/MySkills
```

默认安装到当前项目，执行前先进入要使用这些 Skill 的项目目录。需要在多个项目中使用时加 `-g`，安装到用户级范围；例如，将 `linear-algebra` 全局安装给 Pi：

```bash
npx skills add FeiNiaoBF/MySkills -g -a pi -s linear-algebra
```

替换 `linear-algebra` 为索引中的名称。`-a` 指定目标 Agent；`-s` 可跟多个 Skill 名称。

### 确认安装并使用

`skills add ... --list` 列出的是仓库提供的 Skill；检查已经安装的 Skill 使用 `skills list`：

```bash
# 查看当前项目的安装结果，在该项目目录执行
npx skills list

# 查看全局安装给 Pi 的 Skill，与上面的安装示例对应
npx skills list -g -a pi
```

确认列表中出现所选 Skill，并核对范围和目标 Agent。若未出现，先检查安装命令的 `-g`、`-a` 和执行目录是否与查看时一致。

然后在目标 Agent 中按名称请求使用 Skill，并说明任务、材料和范围。例如，安装 `project-docs` 后可以请求：“请使用 project-docs 整理这个仓库的 README，保留已有事实和安装方式。”想重新理解项目时，可以请求：“请用 project-guide 帮我建立这个仓库的整体认知，说明主要模块如何协作，并给出代码依据和下一步阅读入口。”实际发现和加载方式由宿主决定；列表中已安装不代表当前会话已经加载。

### 更新或在新电脑重装

本仓库所有 Skills 共用以下更新原则：**更新单个 Skill 时，只更新指定 Skill，避免牵动其他 Skills；只有用户明确要求批量或全量更新时，才扩大范围。** 各 Skill 的更新说明引用本节，规则和通用命令集中维护。

更新已安装且有来源记录的 Skill，按实际安装范围选择，并将示例名称替换为目标 Skill：

```bash
# 只更新全局安装的指定 Skill
npx skills update -g make-sense

# 只更新当前项目安装的指定 Skill
npx skills update -p make-sense
```

仅在明确要求更新整个范围时，使用 `npx skills update -g`（全局）或 `npx skills update -p`（当前项目）；这类更新不限于本仓库。

本地 Git checkout、GitHub 上的源码和 CLI 安装的副本各有更新步骤：本地修改或 `git pull` 只改变 checkout；维护者推送到 GitHub 后，使用者再通过 CLI 获取已发布的修改。新增 Skill 需要重新运行 `skills add` 并选择，不会随更新自动加入。Skill 改名或移动目录后 `update` 也无法跟随：用新名重新 `skills add`，再用 `skills remove` 清理旧名。例如已安装旧命名（如 `project-navigator` 重命名为 `project-guide`）的用户，在新版本发布后安装 `project-guide`，核对新入口后再移除旧名；本地 checkout 的改名不会自动更新安装副本。

若更新后 Agent 仍使用旧内容，核对它实际读取的 `SKILL.md` 路径和内容；有版本字段时一并比较。确认安装副本正确后，按宿主方式重新加载或开启新会话。不要只凭仓库版本号判断当前会话已使用新版。

在新电脑上，重新运行 `skills add` 并选择需要的 Skill 和 Agent。各电脑分别维护安装选择；本仓库不自动同步用户配置或恢复清单。

## 维护与验证

修改仓库时，从 [CONTRIBUTING.md](CONTRIBUTING.md) 的准备工作开始，再按其中的规则编写、验证和提交。仅安装 Skill 无需准备仓库校验环境；Agent 的任务边界见 [AGENTS.md](AGENTS.md)。

## 仓库导航

| 位置 | 用途 |
| --- | --- |
| [skills/](skills/) | 按类别组织可安装 Skill，每个入口为 `skills/<category>/<name>/SKILL.md` |
| [CONTRIBUTING.md](CONTRIBUTING.md) | 目录与格式规则、内容要求、验证及提交流程 |
| [AGENTS.md](AGENTS.md) | Agent 的任务边界和执行约定 |
| [docs/](docs/) | 仓库级研究、设计与实施记录；阅读时注意各记录的状态，不能直接视为当前 Skill 行为 |
| [docs/context-budget.md](docs/context-budget.md) | 每轮上下文的度量、五种处置与模型升级后的复跑方法 |
| [scripts/](scripts/) · [tests/](tests/) | 仓库维护工具及其回归测试 |

`docs/`、`scripts/`、`tests/` 和隐藏目录不作为 Skill 安装。具体技能的用法与按需参考从上方索引进入。

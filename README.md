# MySkills

公开的可复用 Agent Skills 创作源，采用 [MIT License](LICENSE)。每个 Skill 以 `SKILL.md` 为入口，为支持 Skills 的 Agent 提供特定任务的工作流程和按需参考。课程工程、媒体、缓存和用户环境配置留在各自项目中。

- 找到适合任务的技能：[选择 Skill](#选择-skill)。
- 在 Agent 中使用：[安装与更新](#安装与更新)，安装后[确认并使用](#确认安装并使用)。
- 修改或贡献本仓库：[贡献与验证](CONTRIBUTING.md)；Agent 另见 [工作约定](AGENTS.md)。

## 选择 Skill

### 视频（video）

| Skill | 使用场景 |
| --- | --- |
| [math-video-workflow](skills/video/math-video-workflow/SKILL.md) | 编排分集数学视频的剪辑、审校和动画交接；视频任务的推荐入口 |
| [video-smart-cut](skills/video/video-smart-cut/SKILL.md) | 本地转写、精简、字幕对齐与成片校验 |
| [math-video-review](skills/video/math-video-review/SKILL.md) | 审校数学表述、推导、口误与动画候选 |
| [manim-insertion-authoring](skills/video/manim-insertion-authoring/SKILL.md) | 设计并验证讲解中途插入的 Manim 场景 |
| [manim-insertion-assembly](skills/video/manim-insertion-assembly/SKILL.md) | 按字幕语义把已验收场景装配为画面替换段 |
| [manim-recap-video](skills/video/manim-recap-video/SKILL.md) | 制作片尾独立、无旁白的数学重温动画 |

中途画面替换使用 `manim-insertion-authoring` + `manim-insertion-assembly`，保持原声和总时长；片尾独立重温使用 `manim-recap-video`，追加拼接。两条路线不要混用。单集状态由课程项目中的 `workflow.yaml` 维护。

### 设计（design）

| Skill | 使用场景 |
| --- | --- |
| [aesthetic-translator](skills/design/aesthetic-translator/SKILL.md) | 把日常视觉偏好转成设计方向，以对话说明和可检查元素的 HTML 预览同步迭代，附通用模板与目标平台规格 |
| [better-designs-md](skills/design/better-designs-md/SKILL.md) | 按明确请求编写、审查或维护指导实际 UI 工作的设计文档，优先复用项目规范；Google DESIGN.md 格式仅在用户要求或目标工具需要时采用 |
| [frontend-guide](skills/design/frontend-guide/SKILL.md) | 先选清界面内容，以对话说明和“项目 UI＋顺序说明”的 HTML 初稿确认设计，附可替换占位模板；再沿项目技术栈实现，支持逐模块陪建 |

### 学习（learning）

| Skill | 使用场景 |
| --- | --- |
| [linear-algebra](skills/learning/linear-algebra/SKILL.md) | 基于 Strang / MIT 18.06 体系维护 Obsidian 线性代数知识笔记 |
| [project-tech-mentor](skills/learning/project-tech-mentor/SKILL.md) | 通过真实项目学习一门技术 |
| [project-reverse-engineering](skills/learning/project-reverse-engineering/SKILL.md) | 从已有项目切片理解机制，练习修改、迁移和指导 AI 开发 |

### 沟通（communication）

| Skill | 使用场景 |
| --- | --- |
| [clarify-wording](skills/communication/clarify-wording/SKILL.md) | 把没说清的想法解释明白，并重述成保留原意、可直接使用的用户话术 |
| [make-sense](skills/communication/make-sense/SKILL.md) | 不懂术语、段落或推理步骤，或想核对理解时，用精准的文本解释修复理解障碍并回到原任务 |
| [visual-explainer](skills/communication/visual-explainer/SKILL.md) | 用户明确调用时，把完整解释设计成易理解、可离线打开的 HTML 页面；不会因问题复杂而自动生成页面 |

### 问题探索（problem-solving）

| Skill | 使用场景 |
| --- | --- |
| [assess](skills/problem-solving/assess/SKILL.md) | 对函数、代码、文档、UI、技术方案或想法等指定对象按目标选标准、主动验证并给出有证据的结论与改进建议；默认不修改对象 |
| [problem-exploration](skills/problem-solving/problem-exploration/SKILL.md) | 面对模糊的问题或想法时，先调查相关证据、验证关键假设，再判断下一步是否行动 |

### 软件工程（software-engineering）

| Skill | 使用场景 |
| --- | --- |
| [code-subtraction](skills/software-engineering/code-subtraction/SKILL.md) | 用消融证据识别并删除不必要复杂度 |
| [project-implementation-planner](skills/software-engineering/project-implementation-planner/SKILL.md) | 手动启动，依据需求、依赖与工程风险排列实施顺序，解释理由、记录必要决策并交接下一块可验证的开发工作 |
| [project-docs](skills/software-engineering/project-docs/SKILL.md) | 为新项目逐步建立文档、为已有项目持续整理与维护；支持直接完成或共同打磨，以准确、充分、便于人类阅读和使用为目标 |
| [skill-discovery-install](skills/software-engineering/skill-discovery-install/SKILL.md) | 用户主动要求查找、比较或安装 Skills 时，审查候选并按授权安装到项目级 `.agents/skills/` |

### 研究（research）

| Skill | 使用场景 |
| --- | --- |
| [research-curator](skills/research/research-curator/SKILL.md) | 用户明确需要深度、可复查报告时，用真实检索和可追溯证据产出离线 HTML；快速事实查询走宿主检索工具；默认输出到系统临时目录 `%TEMP%/research-curator/<topic>/` |

### 网络安全（network-security）

| Skill | 使用场景 |
| --- | --- |
| [router-proxy-diagnostics](skills/network-security/router-proxy-diagnostics/SKILL.md) | OpenWrt/ImmortalWrt、PassWall、sing-box/VPS 的只读诊断及带备份、回滚的最小变更 |

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

然后在目标 Agent 中按名称请求使用 Skill，并说明任务、材料和范围。例如，安装 `project-docs` 后可以请求：“请使用 project-docs 整理这个仓库的 README，保留已有事实和安装方式。”实际发现和加载方式由宿主决定；列表中已安装不代表当前会话已经加载。

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

本地 Git checkout、GitHub 上的源码和 CLI 安装的副本各有更新步骤：本地修改或 `git pull` 只改变 checkout；维护者推送到 GitHub 后，使用者再通过 CLI 获取已发布的修改。新增 Skill 需要重新运行 `skills add` 并选择，不会随更新自动加入。Skill 改名或移动目录后 `update` 也无法跟随：用新名重新 `skills add`，再用 `skills remove` 清理旧名。

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
| [scripts/](scripts/) · [tests/](tests/) | 仓库维护工具及其回归测试 |

`docs/`、`scripts/`、`tests/` 和隐藏目录不作为 Skill 安装。具体技能的用法与按需参考从上方索引进入。

# MySkills

公开的可复用 Agent Skills 创作源，采用 [MIT License](LICENSE)。这里维护工作流与按需参考资料；课程工程、媒体、缓存和用户环境配置留在各自项目中。

- 找到适合任务的技能：[选择 Skill](#选择-skill)。
- 在 Agent 中使用：[安装与更新](#安装与更新)。
- 修改或贡献本仓库：[贡献与验证](CONTRIBUTING.md)；Agent 另见 [工作约定](AGENTS.md)。

## 选择 Skill

### 视频（video）

| Skill | 使用场景 |
| --- | --- |
| [math-video-workflow](skills/video/math-video-workflow/SKILL.md) | 编排分集数学视频的剪辑、审校和动画交接；视频任务的推荐入口 |
| [video-smart-cut](skills/video/video-smart-cut/SKILL.md) | 本地转写、精简、字幕对齐与成片校验 |
| [math-video-review](skills/video/math-video-review/SKILL.md) | 审校数学表述、推导、口误与动画候选 |
| [math-manim-insertion](skills/video/math-manim-insertion/SKILL.md) | 设计并验证讲解中途插入的 Manim 场景 |
| [manim-video-insertion](skills/video/manim-video-insertion/SKILL.md) | 按字幕语义把已验收场景装配为画面替换段 |
| [manim-recap-video](skills/video/manim-recap-video/SKILL.md) | 制作片尾独立、无旁白的数学重温动画 |

中途画面替换使用 `math-manim-insertion` + `manim-video-insertion`，保持原声和总时长；片尾独立重温使用 `manim-recap-video`，追加拼接。两条路线不要混用。单集状态由课程项目中的 `workflow.yaml` 维护。

### 设计（design）

| Skill | 使用场景 |
| --- | --- |
| [aesthetic-translator](skills/design/aesthetic-translator/SKILL.md) | 把模糊审美感受转成可执行设计规格 |
| [better-designs-md](skills/design/better-designs-md/SKILL.md) | 编写和维护界面设计规范 `DESIGN.md` |
| [frontend-guide](skills/design/frontend-guide/SKILL.md) | 从方向卡到模块化页面实现，或按需逐模块陪建 |

### 学习（learning）

| Skill | 使用场景 |
| --- | --- |
| [linear-algebra](skills/learning/linear-algebra/SKILL.md) | 基于 Strang / MIT 18.06 体系维护 Obsidian 线性代数知识笔记 |
| [project-tech-mentor](skills/learning/project-tech-mentor/SKILL.md) | 通过真实项目学习一门技术 |
| [project-digestion](skills/learning/project-digestion/SKILL.md) | 从已有项目切片理解机制，练习修改、迁移和指导 AI 开发 |

### 沟通（communication）

| Skill | 使用场景 |
| --- | --- |
| [clarify](skills/communication/clarify/SKILL.md) | 有想法却说不清、表达障碍已影响当前任务时，找出准确说法并继续原任务 |
| [make-sense](skills/communication/make-sense/SKILL.md) | 不懂术语、段落或推理步骤，或想核对理解时，修复理解障碍并回到原任务 |

### 软件工程（software-engineering）

| Skill | 使用场景 |
| --- | --- |
| [code-subtraction](skills/software-engineering/code-subtraction/SKILL.md) | 用消融证据识别并删除不必要复杂度 |
| [project-docs](skills/software-engineering/project-docs/SKILL.md) | 根据代码与已确认需求规划、编写、审查和维护项目文档，按读者任务保持最小充分的文档体系 |

### 研究（research）

| Skill | 使用场景 |
| --- | --- |
| [research-curator](skills/research/research-curator/SKILL.md) | 使用真实检索与可追溯证据开展有边界的研究，再写成 Reader First、本地化的离线报告；默认输出到系统临时目录 `%TEMP%/research-curator/<topic>/` |

### 网络安全（network-security）

| Skill | 使用场景 |
| --- | --- |
| [network-security-check](skills/network-security/network-security-check/SKILL.md) | OpenWrt/ImmortalWrt、PassWall、sing-box/VPS 的只读诊断及带备份、回滚的最小变更 |

## 安装与更新

需要可用的 Node.js、npm（含 `npx`）和 Git，以及访问 npm 与 GitHub 的网络。安装由 [Vercel skills CLI](https://github.com/vercel-labs/skills) 完成；各 Skill 的运行依赖以自身说明为准。

先列出本仓库可安装的 Skill，此命令不会安装：

```bash
npx skills add FeiNiaoBF/MySkills --list
```

在交互式终端中选择 Skill 和目标 Agent：

```bash
npx skills add FeiNiaoBF/MySkills
```

默认安装到当前项目。需要全局使用时加 `-g`；例如，将 `linear-algebra` 安装给 Pi：

```bash
npx skills add FeiNiaoBF/MySkills -g -a pi -s linear-algebra
```

替换 `linear-algebra` 为索引中的名称。`-a` 指定目标 Agent；`-s` 可跟多个 Skill 名称。

更新已安装且有来源记录的 Skill，按安装范围选择：

```bash
# 全局安装
npx skills update -g

# 当前项目安装
npx skills update -p
```

本地 Git checkout 与 CLI 安装的副本分开：`git pull` 更新本仓库源码，不会自动更新已安装副本。源码推送到 GitHub 后，再通过 CLI 更新获取。新增 Skill 需要重新运行 `skills add` 并选择，不会随更新自动加入。

在新电脑上，重新运行 `skills add` 并选择需要的 Skill 和 Agent。各电脑分别维护安装选择；本仓库不自动同步用户配置或恢复清单。

## 维护与验证

贡献规则、验证环境和检查命令统一维护在 [CONTRIBUTING.md](CONTRIBUTING.md)。仓库校验使用 Python 3.10+，不要求 Skill 使用者为安装而准备 Python 环境。

## 仓库导航

| 位置 | 用途 |
| --- | --- |
| [skills/](skills/) | 按类别组织可安装 Skill，每个入口为 `skills/<category>/<name>/SKILL.md` |
| [CONTRIBUTING.md](CONTRIBUTING.md) | 目录与格式规则、内容要求、验证及提交流程 |
| [AGENTS.md](AGENTS.md) | Agent 的任务边界和执行约定 |
| [docs/](docs/) | 仓库级研究、设计与实施记录；阅读时注意各记录的状态，不能直接视为当前 Skill 行为 |
| [scripts/](scripts/) · [tests/](tests/) | 仓库维护工具及其回归测试 |

`docs/`、`scripts/`、`tests/` 和隐藏目录不作为 Skill 安装。具体技能的用法与按需参考从上方索引进入。

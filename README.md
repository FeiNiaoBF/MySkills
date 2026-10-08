# MySkills

公开的可复用 Agent Skills 创作源，采用 [MIT License](LICENSE)。这里维护工作流与按需参考资料；课程工程、媒体、缓存和用户环境配置留在各自项目中。

## 选择 Skill

### 视频与数学视频

| Skill | 使用场景 |
| --- | --- |
| [math-video-workflow](math-video-workflow/SKILL.md) | 编排分集数学视频的剪辑、审校和动画交接；视频任务的推荐入口 |
| [video-smart-cut](video-smart-cut/SKILL.md) | 本地转写、精简、字幕对齐与成片校验 |
| [math-video-review](math-video-review/SKILL.md) | 审校数学表述、推导、口误与动画候选 |
| [math-manim-insertion](math-manim-insertion/SKILL.md) | 设计并验证讲解中途插入的 Manim 场景 |
| [manim-video-insertion](manim-video-insertion/SKILL.md) | 按字幕语义把已验收场景装配为画面替换段 |
| [manim-recap-video](manim-recap-video/SKILL.md) | 制作片尾独立、无旁白的数学重温动画 |

中途画面替换使用 `math-manim-insertion` + `manim-video-insertion`，保持原声和总时长；片尾独立重温使用 `manim-recap-video`，追加拼接。两条路线不要混用。单集状态由课程项目中的 `workflow.yaml` 维护。

### 设计与前端

| Skill | 使用场景 |
| --- | --- |
| [aesthetic-translator](aesthetic-translator/SKILL.md) | 把模糊审美感受转成可执行设计规格 |
| [better-designs-md](better-designs-md/SKILL.md) | 编写和维护界面设计规范 `DESIGN.md` |
| [frontend-guide](frontend-guide/SKILL.md) | 从方向卡到模块化页面实现，或按需逐模块陪建 |

### 数学学习

| Skill | 使用场景 |
| --- | --- |
| [linear-algebra](linear-algebra/SKILL.md) | 基于 Strang / MIT 18.06 体系维护 Obsidian 线性代数知识笔记 |

### 工程维护与理解

| Skill | 使用场景 |
| --- | --- |
| [code-subtraction](code-subtraction/SKILL.md) | 用消融证据识别并删除不必要复杂度 |
| [project-tech-mentor](project-tech-mentor/SKILL.md) | 通过真实项目学习一门技术 |
| [project-digestion](project-digestion/SKILL.md) | 从已有项目切片理解机制，练习修改、迁移和指导 AI 开发 |
| [make-sense](make-sense/SKILL.md) | 不懂术语、段落或推理步骤，或想核对理解时，修复理解障碍并回到原任务 |
| [clarify](clarify/SKILL.md) | 有想法却说不清、表达障碍已影响当前任务时，找出准确说法并继续原任务 |

### 研究与资料策展

| Skill | 使用场景 |
| --- | --- |
| [research-curator](research-curator/SKILL.md) | 使用真实检索与可追溯证据开展有边界的研究，再写成 Reader First、本地化的离线报告；默认输出到系统临时目录 `%TEMP%/research-curator/<topic>/` |

### 网络安全

| Skill | 使用场景 |
| --- | --- |
| [network-security-check](network-security-check/SKILL.md) | OpenWrt/ImmortalWrt、PassWall、sing-box/VPS 的只读诊断及带备份、回滚的最小变更 |

## 安装与更新

本仓库托管在 GitHub，Vercel `skills` CLI 会从仓库根目录发现各 skill 目录中的 `SKILL.md`，不需要额外 manifest。先列出 CLI 识别到的 skills：

```bash
npx skills add FeiNiaoBF/MySkills --list
```

在交互式终端中选择要安装的 skills 和 Agent：

```bash
npx skills add FeiNiaoBF/MySkills
```

也可以用参数明确指定。例如，将 `linear-algebra` 全局安装到 Pi：

```bash
npx skills add FeiNiaoBF/MySkills -g -a pi -s linear-algebra
```

替换 `linear-algebra` 为所需的 skill 名称；安装多个 skills 时，可在 `-s` 后列出多个名称。`-g` 表示全局安装，`-a` 指定目标 Agent。

更新本机通过 CLI 安装并记录来源的全局 skills：

```bash
npx skills update -g
```

每台电脑分别记录自己的安装选择。新电脑上重复相同的 `skills add` 命令即可安装同一组 skills。仓库新增的 skill 不会自动加入现有安装；需要时再次运行 `skills add` 并选择它。新版本推送到 GitHub `main` 后，用户运行更新命令即可获取已安装 skills 的更新。

## 维护与验证

贡献规则见 [AGENTS.md](AGENTS.md)。维护工具要求 Python 3.10+；[requirements-dev.txt](requirements-dev.txt) 仅用于仓库校验，安装单个 skill 不需要这些依赖。

推荐把验证环境放在仓库外：

**PowerShell：**

```powershell
$venv = Join-Path $env:TEMP "pi-agent/myskills-venv"
python -m venv $venv
$python = Join-Path $venv "Scripts/python.exe"
& $python -m pip install -r requirements-dev.txt
& $python -B scripts/validate_skills.py
& $python -B -m unittest discover -s tests -v
```

**Bash：**

```bash
venv="${TMPDIR:-/tmp}/pi-agent/myskills-venv"
python3 -m venv "$venv"
"$venv/bin/python" -m pip install -r requirements-dev.txt
"$venv/bin/python" -B scripts/validate_skills.py
"$venv/bin/python" -B -m unittest discover -s tests -v
```

[校验器](scripts/validate_skills.py) 自动发现所有一级 skill 目录，检查：

- YAML frontmatter 的语法、重复键、必填字段及名称匹配；自定义字段归入 `metadata`。
- skill 内 Markdown 链接和图片的本地目标是否存在、是否越出本 skill；根文档和 `docs/` 的链接限制在仓库内。
- skill 文本资源、根文档和共享文档中的机器绝对路径及私钥标记。

`docs/`、`scripts/`、`tests/` 和隐藏目录属于仓库基础设施，不作为 skill。检查不会访问网络或修改文件。外部 URL、片段锚点、HTML 链接和未被 Markdown 引用的路径不在链接校验范围内；敏感值扫描是启发式检查，不能替代人工审查。

Skill 的调用策略按宿主分别声明：需要用户明确触发的工作流在 `SKILL.md` 顶层使用 Claude Code 的 `disable-model-invocation: true`，并在 `agents/openai.yaml` 中使用 `policy.allow_implicit_invocation: false`。普通知识和指导类 Skill 保持默认的隐式调用。两种字段都不属于 Agent Skills 核心格式，其他宿主可能忽略它们。

`description` 的要求是**前 57 个字符说明触发条件**，不是总长最多 57 字符；触发是否清楚需要人工判断。行为效果同样需要实际用例验证。

旧命令 `python scripts/validate-video-skills.py` 仍可用，但现在调用同一个全仓库校验器，使用相同依赖。校验失败会返回非零状态；新检查暴露的已有 skill 问题应如实报告，不因本次只维护仓库工具而擅自修改技能。

测试只使用仓库外的临时样例，不操作真实安装目录。Windows Junction 测试在缺少 Windows/PowerShell 时会明确跳过。提交前另运行 `git diff --check`，并确认改动没有越出授权范围。

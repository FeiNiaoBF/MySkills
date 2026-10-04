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

### 研究与资料策展

| Skill | 使用场景 |
| --- | --- |
| [research-curator](research-curator/SKILL.md) | 从 Research Contract 出发筛选资料、去重与溯源，建立 Claim/Evidence Graph，记录 `run.json` 并生成离线 HTML 审计报告 |

### 网络安全

| Skill | 使用场景 |
| --- | --- |
| [network-security-check](network-security-check/SKILL.md) | OpenWrt/ImmortalWrt、PassWall、sing-box/VPS 的只读诊断及带备份、回滚的最小变更 |

## 安装、同步与更新

### 挂载本仓库

从仓库根目录执行一次：

**Windows / PowerShell：**

```powershell
.\scripts\link-skills.ps1
```

**Linux / macOS / WSL：**

```bash
bash scripts/link-skills.sh
```

脚本建立一个集合链接：`~/.agents/skills/myskills` 指向本仓库根目录。Pi、Codex 等支持 Agent Skills 的宿主会递归发现其中的 `SKILL.md`。之后修改、新增或删除 skill 都直接生效，不需要重新运行脚本；活跃会话仍需 reload 或重启。

已有第三方 skill 不会被覆盖。脚本会安全迁移旧版本创建的逐项链接，只删除目标确实属于本仓库的链接。不要再把同名副本安装到 `~/.pi/agent/skills`。

Hermes 使用共享目录时，在其用户配置中加入一次：

```yaml
skills:
  external_dirs:
    - ~/.agents/skills
```

移除集合链接但保留源文件：

```powershell
.\scripts\link-skills.ps1 -Remove
```

```bash
bash scripts/link-skills.sh --remove
```

### 管理第三方 skills

第三方 skill 继续交给 `skills` CLI 管理，不复制进本仓库：

```powershell
npx skills add <owner/repository> -g
npx skills update -g -y
npx skills ls -g
```

`npx skills update -g -y` 只更新有上游来源记录的全局 skill；本仓库通过集合链接实时读取，不参与该更新。不要用 `npx skills add . -g` 安装本仓库：本地来源会被复制到 CLI 的 canonical 目录，失去工作区修改即时生效的特性。

### 恢复清单与本地挂载

本仓库是创作源；本地挂载使用当前工作区内容。另一个维护仓库 `pi-agent` 的 `config/skills-manifest.json` 为纳入恢复清单的 skill 固定提交版本。这是两种不同的版本来源，不能把本地挂载成功当作恢复清单已更新。

经授权 push 后，由维护流程更新清单 ref 并运行 bootstrap。普通内容修改或校验不需要运行 bootstrap，也不应顺带调整其他已安装技能。

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

# 贡献与验证

本指南维护仓库的目录、格式、内容和验证规则。选择、安装或更新 Skill 见 [README](README.md)；Agent 执行任务时还须遵守 [AGENTS.md](AGENTS.md) 的工作边界。

## 准备工作

贡献时修改本地 Git checkout 中的源文件。首次获取仓库可执行：

```bash
git clone https://github.com/FeiNiaoBF/MySkills.git
cd MySkills
```

已有 checkout 时直接进入其根目录，无需重新克隆。先运行 `git status --short`，确认并保留已有改动；不要把 Agent 的安装目录当作本仓库源码修改。改文档或维护工具时，只处理相关文件；改 Skill 时，再核对其入口、按需参考及 README 索引。

## 新建 Skill 的分支流程

每个新 Skill 单独从 `main` 的已提交状态创建 `skill/<name>` 分支；该分支只包含这个 Skill、必要的索引和相关验证改动。已有 Skill 的维护及其他任务按各自范围处理，不把它们顺带放进新 Skill 分支。

### 开始前

1. 运行 `git status --short --branch`、`git branch --show-current` 和 `git worktree list`，确认当前分支、未提交改动及其他工作树。可访问远端时先 `git fetch origin main`，用 `git log --oneline main..origin/main` 和 `git log --oneline origin/main..main` 比较两边提交；若本地落后，先在干净的 `main` 上执行 `git merge --ff-only origin/main`，若出现分叉先查明原因，不强制重置或自动改写历史。
2. `main` 工作区干净时，在 `main` 所在工作树建分支：

   ```bash
   git switch main
   git status --short --branch
   git switch -c skill/<name>
   ```

   如果 `main` 上已有未提交改动，保留原工作区；在仓库外创建独立工作树：`git worktree add -b skill/<name> <outside-repo-path> main`。在新工作树中确认 `git status --short --branch`，再开始创作。不要把未提交改动带进新分支，也不要擅自 stash、提交或丢弃别人的改动。若无法确认 `main` 的合适基线，先解决该问题。

### 开发与提交

- 在 `skill/<name>` 上工作，每次切换上下文及提交前检查 `git status --short --branch`；只 stage 本次 Skill 的明确路径，检查 `git diff --cached` 和 `git diff --cached --check`。验证与提交要求见下文，不为别的 Skill 或工作树提交改动。
- 完成后在该分支提交，检查 `git log --oneline main..HEAD` 和 `git diff main...HEAD`，确认本分支只包含本次新 Skill。主分支在开发期间前进时，先在分支上整合最新 `main` 并处理冲突，再重新验证受影响的内容。

### 合并与收尾

1. 确认新 Skill 已验证且分支改动经过复核；先检查 `main` 工作区是否干净、与远端关系是否明确。若 `main` 有未提交改动、分叉或合并冲突尚未解决，保留分支并说明阻碍，不为合并而改写或隐藏现有工作。
2. 在 `main` 的工作树运行 `git switch main`（如果已经在 `main` 则无需切换），再执行 `git merge --no-ff skill/<name>`；若合并发生冲突，先 `git merge --abort`，在 Skill 分支整合最新 `main` 并解决冲突。合并后按受影响范围复核验证，运行 `git status --short --branch` 并检查合并记录。只有检查完成，才将这个 Skill 视为已并入 `main`。
3. 合并后按需清理已合并分支及其独立工作树。合并和本地提交均不等于发布；Push 仍须用户授权，安装副本的更新另见 [README](README.md#安装与更新)。

## 修改放在哪里

每个 Skill 放在 `skills/<category>/<name>/`，其中 `<name>` 与 `SKILL.md` frontmatter 的 `name` 完全一致。

- 分类名使用小写 kebab-case，反映主要用户任务；每个 Skill 只归入一个类别。跨 Skill 按名称路由，不依赖分类路径。
- 必需入口是 `SKILL.md`。较长的按需参考放在本 Skill 的 `references/`，可选界面提示放在 `agents/openai.yaml`；没有用途时不创建这些目录或文件。
- 新增 Skill、调整分类或改变对外定位时，同步 [README 的索引](README.md#选择-skill)。新增稳定任务领域时可建立新类别。
- 仓库级研究、设计与实施记录放在 `docs/`；维护脚本放在 `scripts/`，相应测试放在 `tests/`。这些目录及隐藏目录不作为可安装 Skill。
- 临时环境、验证样例和报告放在系统临时目录的 `pi-agent/` 下，不提交到仓库。

## 格式与调用策略

Frontmatter 是 `SKILL.md` 首行开始、由两行 `---` 包围的 YAML 元数据。下面只展示格式，名称和描述应替换为实际任务；结束标记之后还需要 Skill 正文。

```yaml
---
name: example-skill
description: 用户需要整理已确认的操作流程时使用。根据源材料编写步骤、前提与预期结果。
license: MIT
metadata:
  version: 0.1.0
---
```

Frontmatter 必须是可解析的 YAML：

- `name` 与目录名一致。
- `description` 先说明触发条件，再说明能力范围；**前 57 个字符应说明何时使用**，不是总长最多 57 字符。触发是否清楚需要人工判断。
- `metadata` 必须为映射；作者、版本、日期等自定义字段放在其中。`license` 等标准字段可保留在顶层。
- 本 Skill 内的相对链接必须指向真实文件。跨 Skill 用名称路由，不用跨目录相对链接共享实现；共同规则放明确的公共文档。

普通知识和指导类 Skill 保持默认隐式调用。需要用户明确触发的工作流，分别在两个宿主中声明：

| 宿主 | 声明位置 | 设置 |
| --- | --- | --- |
| Claude Code | `SKILL.md` 顶层 | `disable-model-invocation: true` |
| 支持 OpenAI 界面元数据的宿主 | `agents/openai.yaml` | `policy.allow_implicit_invocation: false` |

两项设置都是宿主扩展，不属于 Agent Skills 核心格式；其他宿主可能忽略它们。不能仅凭字段存在就认定所有宿主都会限制自动调用。

## 内容与公开边界

- Skill 必须能在另一台机器上独立使用：关键步骤、安全边界和验证条件不能依赖聊天记录或作者机器。
- 先写可执行步骤和完成条件，再提供按分支读取的参考入口；不要求每次通读所有材料。
- 同一规则只维护一处。不重复缓存配置、脚本帮助或可直接查到的结构；保留不能从环境推导出的理由、约束与经验。
- 按任务需要保留指导，不堆叠模型已能可靠完成的通用流程。用实际用例验证精简效果，不把更短等同于更好。
- 示例使用 `<VPS_IP>`、`<ROUTER_IP>`、`<SSH_HOST>`、`<your-key>.pem` 等占位符。公开文件不得包含真实用户的地址、主机名、用户名、节点 ID、凭据、私钥、订阅 token 或机器专属路径。
- 网络、路由器、云主机及媒体操作先做只读诊断；变更前说明影响、备份、回滚和验证方法。

## 准备验证环境

在仓库根目录执行。维护工具要求 **Python 3.10+**；依赖见 [requirements-dev.txt](requirements-dev.txt)。这些依赖只用于仓库校验，安装 Skill 不需要先准备此环境；运行某个 Skill 的额外依赖以其自身说明为准。

### 首次准备

把虚拟环境放在仓库外，无需激活。已有该环境时可跳过创建；依赖文件变化后，用同一环境重新执行依赖安装命令。

**Windows PowerShell：**

```powershell
$venv = Join-Path $env:TEMP "pi-agent/myskills-venv"
python -m venv $venv
$python = Join-Path $venv "Scripts/python.exe"
& $python -m pip install -r requirements-dev.txt
```

**Linux / macOS Bash：**

```bash
venv="${TMPDIR:-/tmp}/pi-agent/myskills-venv"
python3 -m venv "$venv"
"$venv/bin/python" -m pip install -r requirements-dev.txt
```

### 每次修改后检查

仍在仓库根目录执行。以下命令重新指定环境路径，可在新的终端中使用，不需要再次创建虚拟环境。

**Windows PowerShell：**

```powershell
$venv = Join-Path $env:TEMP "pi-agent/myskills-venv"
$python = Join-Path $venv "Scripts/python.exe"
& $python -B scripts/validate_skills.py
& $python -B -m unittest discover -s tests -v
git diff --check
```

**Linux / macOS Bash：**

```bash
venv="${TMPDIR:-/tmp}/pi-agent/myskills-venv"
"$venv/bin/python" -B scripts/validate_skills.py
"$venv/bin/python" -B -m unittest discover -s tests -v
git diff --check
```

预期结果：校验器输出 `PASS`，单元测试输出 `OK`，`git diff --check` 无错误。命令失败时先确定原因，不继续将本次修改标为验证通过。

- 找不到虚拟环境中的 Python：先完成“首次准备”，并确认路径与当前系统对应。
- 出现 `Missing validation dependencies`：用上述虚拟环境的 Python 安装 `requirements-dev.txt`，避免装到另一个解释器中。
- 校验器输出 `FAIL`：按其给出的文件、行号和原因定位。已有或越界问题单独报告，不通过删改无关内容或放宽校验规则来消除失败。

### 检查覆盖什么

[校验器](scripts/validate_skills.py) 自动发现分类目录中的全部 Skill，同时检查根目录 Markdown 和 `docs/` 下的 Markdown：

| 自动检查 | 仍需人工检查 |
| --- | --- |
| YAML 语法、重复键、必填字段、名称匹配及自定义字段位置 | description 的触发含义、技能边界及调用策略是否适当 |
| Markdown 链接和图片的本地目标是否存在；Skill 内链接不越出本 Skill，仓库文档链接不越出仓库 | 相对链接的语义、片段锚点、HTML 链接、外部 URL，以及未写成 Markdown 链接的路径 |
| Skill 的 Markdown、YAML、JSON 等文本资源及仓库 Markdown 中的机器路径和私钥标记 | 其他敏感值、事实准确性、命令行为及任务效果 |

校验器不访问网络或修改文件；敏感值扫描是启发式检查，不能证明不存在敏感信息。单元测试使用仓库外的临时样例，不操作用户真实安装目录。

旧入口 `python -B scripts/validate-video-skills.py` 仍调用同一个全仓库校验器，使用相同依赖；新操作优先使用 `validate_skills.py`。

分类目录调整后，可在系统临时目录的仓库副本上运行 `npx skills add <path> --list`，核对 CLI 是否发现预期 Skill；`--list` 只列出，不安装。此检查需要 README 中的 CLI 前提。

## 验证与提交

1. 检查 Git 状态，阅读相关文件，保留已有修改；将变更限定在授权范围内。
2. 修改维护脚本时，为发现的缺陷补充可复现的回归测试。安装相关测试必须使用隔离目录，不能接触真实 skills 安装目录。
3. 执行上述校验、单元测试和 diff 检查；修复本次引入的失败，再重跑相关检查。已有、越界或环境性失败单独报告，不放宽规则制造通过。
4. 人工复核触发条件、敏感值、链接语义与实际行为。报告执行了哪些检查、未覆盖什么；文件校验、命令成功、模型评价和实际任务有效是不同层次的证据。
5. Stage 明确路径，检查完整 diff，排除缓存、媒体、临时文件、敏感值和无关改动。一个 Skill 或一类逻辑变更使用一个聚焦的本地提交，采用英文 Conventional Commits：`feat:`、`fix:`、`docs:`、`chore:`。

本地提交不等于发布。**Push 必须获得用户授权**；发布后如何更新已安装副本或在新电脑重装，见 [安装与更新](README.md#安装与更新)。

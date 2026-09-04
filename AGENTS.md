# AGENTS.md

自用 AI Skills 集合，公开仓库（FeiNiaoBF/MySkills，MIT License）。任何改动都会公开。

## 技能结构

每个技能一个独立目录，与 SKILL.md 的 frontmatter `name` 同名：

- `SKILL.md` — 必需。frontmatter 含 `name`（目录同名）与 `description`（前 57 字符内说明何时使用该技能）
- `metadata:` — 非规范字段（`author`、`version`、`date_added` 等）一律放 `metadata` 映射内，不放在 frontmatter 顶层，保持对 Agent Skills 规范的兼容
- `references/` — 可选。正文过长时放详细参考，用相对链接指向
- `agents/openai.yaml` — 可选。外部 agent 的界面提示（display_name / short_description / default_prompt）

技能目录自包含，不跨目录引用。

## 内容规范

- 正文中文英文皆可，随技能用途而定
- 示例配置里的环境具体值一律写占位符（如 `<VPS_IP>`、`<your-key>.pem`），只保留通用流程与真实可跑的命令骨架

## 提交

- Conventional Commits，信息用英文：`feat:` / `chore:` / `docs:` / `fix:`
- 按逻辑分组：一个技能一个 commit，一类变更一个 commit

## 安全红线（最高优先级）

本仓库公开，示例配置只允许占位符。提交前自查 diff，出现密码、私钥、API key、订阅 token、VPS 公网 IP、真实用户名等任何敏感值时，先替换为占位符再提交，并主动报告。

# MySkills Agent Instructions

## 项目定位

本仓库是公开的可复用 Agent Skills 创作源（FeiNiaoBF/MySkills，MIT License）。每个 skill 都必须能被另一台机器安全、清楚、可验证地使用。

## Skill 目录规范

每个 skill 使用一个与目录同名的一级目录：

```text
<skill-name>/
├── SKILL.md                 # 必需
├── references/              # 可选：较长的按需参考资料
└── agents/openai.yaml       # 可选：外部 agent 的界面提示
```

`SKILL.md` 的 frontmatter 必须包含：

- `name`：必须与目录名完全一致
- `description`：前 57 个字符内要说明何时使用该 skill；先写触发条件，再写能力范围
- `metadata:`：作者、版本、日期等非规范字段统一放在这里

标准 frontmatter 字段（如 `license`）可以保留在顶层；自定义字段不要散落在顶层。

skill 内部的相对链接必须指向本 skill 目录内的真实文件。可以通过 skill 名称路由到其他 skill，但不要用跨目录相对路径共享实现细节；共享参考应放在明确的公共文档中。

## 内容要求

- 先写可执行步骤，再写按需参考；每个步骤都要有可检查的完成条件。
- 重要事实、命令和安全边界写在 skill 内；不要依赖聊天上下文。
- 不重复缓存 `package.json`、脚本帮助或目录结构中容易查到的事实，除非补充了原因、约束或经验教训。
- 用户已经确认的设计或流程可以直接落盘；仍有实质分叉时只问一个高价值问题。
- 示例配置必须使用占位符，例如 `<VPS_IP>`、`<ROUTER_IP>`、`<SSH_HOST>`、`<your-key>.pem>`。
- 不在公开仓库写入真实公网 IP、局域网地址、主机名、用户名、节点 ID、密码、私钥、API key、订阅 token、真实文件路径或其他用户环境专属值。
- 网络、路由器、云主机和媒体操作优先只读诊断；变更前说明影响、备份方式、回滚命令和验证方法。
- 不把“看起来成功”当作完成；必须报告实际执行过的验证和未覆盖的部分。

## 当前 Skill 分组

### 视频与数学视频

- `math-video-workflow`：按阶段编排视频处理、审校和动画交接
- `video-smart-cut`：本地转写、剪辑、字幕对齐和成片校验
- `math-video-review`：数学视频轻审、终审和 Manim 候选
- `math-manim-insertion`：设计和验证插入式 Manim 场景
- `manim-video-insertion`：按字幕语义装配画面替换段
- `manim-recap-video`：制作片尾独立重温动画

两条 Manim 路线不可混用：中途画面替换走 `math-manim-insertion` + `manim-video-insertion`；片尾独立重温走 `manim-recap-video`。

### 设计与前端

- `aesthetic-translator`：把模糊审美感受翻译成可执行设计规格
- `better-designs-md`：编写和维护 DESIGN.md
- `frontend-guide`：从方向卡到模块化页面实现

### 工程维护与学习

- `code-subtraction`：用消融证据识别并删除不必要复杂度
- `project-tech-mentor`：通过真实项目建立技术能力
- `project-digestion`：从已有项目切片中恢复机制，并通过预测、修改和迁移建立可复用能力

### 网络安全

- `network-security-check`：对 OpenWrt/ImmortalWrt + PassWall + sing-box/VPS 网络进行安全诊断和最小变更。它覆盖 IP/DNS/IPv6 泄漏检查、路由稳定性、VLESS Reality 节点验证、Claude/Cloudflare 网络故障定位和 PassWall 分流；必须先建立基线和只读证据，备份后才能变更，并准备回滚。

该 skill 不保存任何真实节点或用户网络配置；用户环境值只能在当前会话中读取并使用占位符记录。

## 验证

提交前至少执行：

```bash
python scripts/validate-video-skills.py
```

同时手工检查：

1. 所有一级 skill 目录都有 `SKILL.md`。
2. `name` 与目录名一致，frontmatter 可解析。
3. 本地 Markdown 链接都存在。
4. 示例没有机器绝对路径和敏感值。
5. diff 只包含当前逻辑变更，没有缓存、媒体或临时文件。

新增通用校验脚本时，应覆盖全部 skill，而不只覆盖视频 skill。

## 部署

本仓库是创作源；运行时 skill 安装在 `~/.agents/skills`。优先使用：

```bash
scripts/link-skills.sh <skill-name>
```

Windows 使用 `scripts/link-skills.ps1`。不要同时链接到 `~/.pi/agent/skills`，避免同名 skill 被重复加载。

纳入恢复清单的 skill 由 `pi-agent/config/skills-manifest.json` 的 pinned commit 管理。仓库 push 后，必须由维护流程更新 manifest 的 commit ref，再运行 bootstrap；未纳入清单的 skill 只通过 link 脚本安装。

## Git 与公开发布

- 提交信息使用 Conventional Commits，英文格式：`feat:`、`fix:`、`docs:`、`chore:`。
- 一个 skill 或一类逻辑变更使用一个聚焦提交。
- stage 时只列出明确路径。
- 不自动 push；push 前必须获得用户确认。
- 提交前再次检查 `git diff --check` 和敏感值扫描。

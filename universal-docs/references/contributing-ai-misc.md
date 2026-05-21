# CONTRIBUTING.md 规范

```markdown
# 贡献指南

感谢你考虑为本项目贡献！

## 环境搭建

\```bash
# 克隆仓库
git clone https://github.com/OWNER/REPO.git
cd REPO

# 安装依赖（根据项目类型自动填充）
npm install          # Node.js
pip install -e ".[dev]"   # Python
go mod download      # Go
\```

## 开发流程

1. Fork 仓库并创建功能分支：`git checkout -b feat/your-feature`
2. 编写代码，确保通过所有测试
3. 提交（见提交规范）
4. 推送并开启 Pull Request

## 代码风格

[根据项目自动填充：ESLint/Prettier、Black/Ruff、gofmt 等]

\```bash
# 格式化代码
npm run format    # 或 make fmt
# 检查 lint
npm run lint
\```

## 测试

\```bash
# 运行所有测试
npm test
# 运行特定测试
npm test -- --grep "关键词"
\```

## 提交规范（Conventional Commits）

\```
<type>(<scope>): <subject>

[可选正文]

[可选 footer]
\```

类型：`feat` | `fix` | `docs` | `refactor` | `test` | `chore` | `perf`

示例：
- `feat(auth): 添加 OAuth2 登录支持`
- `fix(api): 修复空指针异常 (#42)`
- `docs: 更新 API 文档`

## Pull Request 规范

- PR 标题遵循 Conventional Commits 格式
- 必须通过所有 CI 检查
- 需要至少 1 位 Reviewer 批准
- 描述中说明：做了什么、为什么、如何测试
```

---

# AGENTS.md / CLAUDE.md 生成规范

> 专供 AI 工具阅读的项目上下文文件。**只写 AI 在项目文件中发现不了的信息。**

## 必须包含（扫描后确认）

```markdown
# AGENTS.md

## 构建与开发命令
\```bash
# 安装依赖
npm install

# 开发模式
npm run dev

# 构建
npm run build

# 测试（必须通过后才能提交）
npm test

# Lint（CI 会检查）
npm run lint
\```

## 提交信息格式
使用 Conventional Commits：`feat(scope): description`
提交前必须运行：`npm run lint && npm test`

## 禁止修改的文件
- `src/generated/` — 自动生成，修改会被覆盖
- `.env.example` — 用于文档，不是真实配置
- `CHANGELOG.md` — 由 CI 自动生成

## 环境变量
所有环境变量在 `.env.example` 中有说明，本地开发复制为 `.env` 并填充真实值。

## 已知陷阱
[从用户处收集，或从注释中扫描 HACK/FIXME/NOTE]
- 示例：`src/auth/session.ts` 中的 token 刷新逻辑有竞态条件，暂时用锁规避，不要重构

## 测试说明
- 单元测试在 `src/**/*.test.ts`
- E2E 测试需要运行 `docker compose up -d` 启动依赖服务
```

## 扫描要点

```bash
# 找所有 HACK/FIXME/NOTE 注释
grep -r "HACK\|FIXME\|NOTE\|TODO" src/ --include="*.ts" --include="*.py" --include="*.go"
# 找生成文件标记
grep -r "DO NOT EDIT\|generated\|auto-generated" . --include="*.ts" -l
# 检查 package.json scripts
cat package.json | jq '.scripts'
```

---

# 其他文档（SECURITY.md、CODE_OF_CONDUCT.md 等）

## SECURITY.md
```markdown
# 安全策略

## 支持的版本
| 版本 | 支持状态 |
|------|---------|
| 1.x  | ✅ 支持 |
| 0.x  | ❌ 不支持 |

## 报告漏洞
请**不要**通过 GitHub Issues 公开报告安全漏洞。
发送邮件至：security@example.com
预计响应时间：48 小时内确认
```

## CODE_OF_CONDUCT.md
推荐直接使用 [Contributor Covenant](https://www.contributor-covenant.org/zh-cn/version/2/1/code_of_conduct/) 并填入联系邮箱。

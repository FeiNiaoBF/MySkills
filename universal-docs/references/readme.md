# README 生成规范

遵循 [standard-readme](https://github.com/RichardLitt/standard-readme) 并融合业界最佳实践。

## 必须包含的章节（顺序固定）

```markdown
# 项目名称

[徽章行]

> 一句话描述：是什么、做什么、解决什么问题

[可选：截图或演示 GIF]

## 特性（Features）
- 核心亮点，3-6 条，每条一行

## 快速开始（Quick Start）
### 环境要求
### 安装
### 最小可用示例（必须可直接运行）

## 使用指南（Usage）
### 基础用法
### 配置项（表格：参数名 | 类型 | 默认值 | 说明）
### 进阶用法（可选）

## API 参考（可选，小项目可内联，大项目链接到 docs/api.md）

## 常见问题（FAQ）

## 贡献（Contributing）
指向 CONTRIBUTING.md

## 许可证（License）
```

## 徽章推荐策略

根据项目类型自动推荐合适徽章，只推荐能真实生成的：

```markdown
<!-- npm 包 -->
[![npm version](https://img.shields.io/npm/v/PACKAGE_NAME)](https://www.npmjs.com/package/PACKAGE_NAME)
[![npm downloads](https://img.shields.io/npm/dm/PACKAGE_NAME)](https://www.npmjs.com/package/PACKAGE_NAME)

<!-- PyPI -->
[![PyPI](https://img.shields.io/pypi/v/PACKAGE_NAME)](https://pypi.org/project/PACKAGE_NAME)

<!-- 通用 -->
[![License](https://img.shields.io/github/license/OWNER/REPO)](LICENSE)
[![Build Status](https://img.shields.io/github/actions/workflow/status/OWNER/REPO/WORKFLOW_FILE)](https://github.com/OWNER/REPO/actions)
```

**规则**：只在能从项目文件中确认包名、仓库地址的情况下才生成徽章；否则生成占位符并提示用户填写。

## 快速开始章节质量要求

安装命令必须是可直接复制的单条命令，示例代码必须包含 import/require，不省略：

```bash
# ✅ 好的写法
npm install my-package
```

```typescript
// ✅ 好的示例（含 import，可直接运行）
import { createClient } from 'my-package';
const client = createClient({ apiKey: 'YOUR_KEY' });
const result = await client.doSomething();
console.log(result);
```

```bash
# ❌ 避免：缺少前置步骤、省略关键配置
// ... setup code
const result = await doSomething();
```

## 扫描要点（生成前必须读取）

- `package.json` / `pyproject.toml` / `Cargo.toml` → 项目名、描述、版本、依赖
- `LICENSE` 文件 → 许可证类型
- `src/` 或 `lib/` 入口文件 → 主要导出、核心 API
- 现有 `README.md` → 保留用户已写的有效内容
- `examples/` 目录 → 参考真实示例代码

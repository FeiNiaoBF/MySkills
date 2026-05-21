# CHANGELOG 生成规范

遵循 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/) + [语义化版本](https://semver.org/lang/zh-CN/)。

## 标准结构

```markdown
# Changelog

本文件记录项目的所有重要变更。
格式基于 [Keep a Changelog](https://keepachangelog.com/zh-CN/1.0.0/)，
版本遵循 [语义化版本](https://semver.org/lang/zh-CN/)。

## [Unreleased]

### Added
- 新功能描述

## [1.2.0] - 2024-03-15

### Added
- 新增 `createUser` API 支持批量创建

### Changed
- `updateUser` 接口现在返回完整用户对象（原来只返回 id）

### Fixed
- 修复在 Safari 下输入框无法聚焦的问题 (#123)

### Security
- 升级 `axios` 至 1.6.0，修复 CVE-2023-45857

## [1.1.0] - 2024-02-01
...

[Unreleased]: https://github.com/OWNER/REPO/compare/v1.2.0...HEAD
[1.2.0]: https://github.com/OWNER/REPO/compare/v1.1.0...v1.2.0
```

## 变更类别说明

| 类别 | 说明 | 典型 Git 提交前缀 |
|------|------|-----------------|
| Added | 新功能 | `feat:` |
| Changed | 对现有功能的变更（可能破坏兼容） | `refactor:`, `perf:` |
| Deprecated | 即将移除的功能 | `deprecate:` |
| Removed | 已移除的功能 | `feat!:`, `BREAKING CHANGE` |
| Fixed | Bug 修复 | `fix:` |
| Security | 安全相关修复 | `security:`, `fix(security):` |

## 从 Git 历史生成

```bash
# 获取上个版本以来的所有提交
git log v1.1.0..HEAD --oneline --no-merges

# 获取带完整信息的提交（含 PR 号）
git log v1.1.0..HEAD --pretty=format:"%h %s (%an)" --no-merges
```

**分类规则（基于 Conventional Commits）**：
- `feat:` → Added
- `fix:` → Fixed
- `perf:` / `refactor:` → Changed
- `docs:` → 通常不记录（除非文档本身是产品）
- `BREAKING CHANGE` → Removed 或 Changed（标注 ⚠️）
- `security:` → Security

## 更新时的差异合并策略

读取现有 CHANGELOG 时：
1. 找到 `## [Unreleased]` 节，或在顶部插入新版本节
2. **绝不删除**已有版本节的内容
3. 新条目追加到对应类别下
4. 如果没有 `[Unreleased]` 节，在最新版本之前插入

## GitHub Actions 自动化（生成后主动询问是否需要）

```yaml
# .github/workflows/release.yml
name: Release & Changelog

on:
  push:
    tags: ['v*']

jobs:
  changelog:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
        with:
          fetch-depth: 0
      
      - name: Generate Changelog
        uses: orhun/git-cliff-action@v3
        with:
          config: cliff.toml
          args: --latest --strip header
        env:
          OUTPUT: RELEASE_NOTES.md
      
      - name: Create GitHub Release
        uses: softprops/action-gh-release@v1
        with:
          body_path: RELEASE_NOTES.md
```

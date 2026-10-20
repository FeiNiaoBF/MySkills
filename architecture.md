---
id: myskills-architecture
type: architecture-design
status: draft
title: MySkills 架构
parent: myskills
---

# 架构

仓库是「内容 + 校验工具」两层：`skills/` 是唯一的分发物，其余部分只服务于它的创作、校验与追溯。

## 拓扑

| 区域 | 内容 | 随安装分发 |
| --- | --- | --- |
| `skills/` | `skills/<category>/<skill>/` 的可安装 Skill | 是，按 Skill 粒度 |
| `docs/` | 仓库级设计、研究、计划记录 | 否 |
| `scripts/` · `tests/` | 校验器及其回归测试 | 否 |
| 根目录 Markdown | 索引与安装说明（`README.md`）、格式与验证规则（`CONTRIBUTING.md`）、Agent 工作边界（`AGENTS.md`） | 否 |
| 隐藏目录 | 工作区与工具状态 | 否 |

真实存在的读取关系：

```text
安装者            ──读──> skills/<category>/<skill>/ 下的入口与本地资源
validate_skills   ──只读扫描──> skills/ · docs/ · 根目录 *.md
tests/            ──驱动──> validate_skills.py
```

依赖是单向的：Skill 不读取 `docs/`、`scripts/`、`tests/`；校验器不被 Skill 引用为运行依据。Agent 按名称发现 Skill，分类目录只服务于人读索引。

## 不变量

- `skills/` 下只允许 `<category>/<skill>/` 两级；`<skill>` 与 frontmatter `name` 完全一致，入口必须是 `SKILL.md`。
- frontmatter 从首行 `---` 开始且是可解析 YAML；自定义字段只能放在 `metadata` 映射内。
- 相对链接不越界：Skill 内文件以该 Skill 目录为界，仓库 Markdown 以仓库根为界。
- `docs/`、`scripts/`、`tests/` 与隐藏目录都不参与安装。
- 校验器只读：不访问网络、不修改文件、不接触真实安装目录。

## 决策

- **校验在本地手动执行，不引入 CI**（已确认）。单人维护、改动频率低，临时虚拟环境中的校验器加单元测试覆盖了可机械判定的风险，不值得为此增加 CI 的维护成本。
- **内容质量不进入自动检查**：可机械判定的只有布局、YAML 和本地链接，其余靠人工复核，两者的分工记录在 `scripts/SPEC.md`。

## 模块

- `skills-collection`：可安装 Skill 集合的职责与自包含边界。
- `validation-tooling`：机械校验的契约与限制。
- `docs-records`：`docs/` 记录的地位与边界。

---
id: docs-records
type: module-design
status: draft
title: 仓库记录
parent: myskills-architecture
---

## Responsibility

`docs/` 保存不应进入可安装 Skill 的仓库级材料：研究记录、设计记录、实施计划与草案。它让维护者日后能追溯某个 Skill 为什么变成现在这样。

## Boundary

- 不参与安装，不得被任何 Skill 引用为运行依据。
- 记录描述写作当时的状态，不是当前行为的声明；仓库入口 [`README.md`](../../README.md) 要求阅读时先看该记录的状态。
- 只有仍然生效的契约才成为 spec 节点（例如 `research-curator` 的合同文档）；已完成变更的计划和设计记录、未实现的设计草案都只作为输入保留，不进入 spec 图。
- 未实现的草案留在原处，不代表已承诺的范围；范围由 `myskills` 的能力与非目标决定。

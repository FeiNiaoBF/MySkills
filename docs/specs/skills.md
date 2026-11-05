---
id: skills-collection
type: module-design
status: draft
title: Skill 集合
parent: myskills-architecture
---

## Responsibility

`skills/` 是仓库唯一的分发物。每个 `<category>/<skill>/` 是一个可独立安装的任务单元：`SKILL.md` 说明何时使用、怎么做、何时算完成，单元内可按需附带参考文档、素材和宿主界面元数据（`agents/openai.yaml`）。

## Boundary

- 单元自包含。关键步骤、安全边界和验证条件必须能在另一台机器上复现，不引用 `docs/`、`scripts/`、`tests/`、聊天记录或作者本机路径。
- 跨 Skill 只用名称路由，不用跨目录相对链接；共同规则写在明确的公共文档里，不复制到每个 Skill。
- 一个 Skill 只归一个类别，类别名反映主要用户任务而不是实现层次；新增类别意味着出现了稳定的任务领域。
- 内容优先级：可执行步骤与完成条件优先，其次是无法从环境推导的理由、约束与经验，最后才是可自行查到的结构或配置。
- 隐式调用是默认；需要用户明确触发的工作流在宿主扩展字段中声明，且不假定所有宿主都遵守这些字段。
- 新增、改名或改变对外定位时，同步 `README.md` 的索引；只改一个 Skill 时不牵动其他 Skill。

格式细节与验证命令见 [`CONTRIBUTING.md`](../../CONTRIBUTING.md)；布局与不变量见 `myskills-architecture`。

## Children

`research-curator` 是唯一携带可执行程序的 Skill，其内部契约由自身合同文档规定。

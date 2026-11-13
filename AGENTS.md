# MySkills Agent Instructions

本仓库用于创作、验证和发布可复用的 Agent Skills。索引与安装见 [README.md](README.md)；格式、验证、分支和提交流程以 [CONTRIBUTING.md](CONTRIBUTING.md) 为准；改 Skill 前读 [Skill 编辑清单](docs/agent/skill-edit.md)。

## 边界

- 只编辑本仓库中的 Skill 源文件，不把 CLI 安装副本当作源码；普通维护不安装或更新 Skills、不修改外部恢复清单、不运行 bootstrap。
- 保留已有改动，只处理用户授权的内容；越界问题只报告。
- 不把课程工程、媒体、缓存或用户环境配置放入仓库。
- 删除本地分支前，确认其已合并且关联 worktree 干净；远端分支删除与 Push 一样，须获用户授权。
- 完成 = 适用验证通过且 diff 只含本次改动；汇报改了什么、查了什么、没覆盖什么、剩余问题。

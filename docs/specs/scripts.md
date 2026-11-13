---
id: validation-tooling
type: module-design
status: draft
title: 校验工具
parent: myskills-architecture
depends-on: [skills-collection]
---

## Responsibility

`scripts/validate_skills.py` 把 `skills-collection` 的打包不变量变成可机械判定的检查；`tests/test_validate_skills.py` 用仓库外的隔离样例保护这些判定不被放宽或误报。它回答「能否安装、链接是否成立」，不回答「内容是否好」。

`scripts/context_budget.py` 度量每轮与按需上下文字符数，并强制 description 预算（自动调用型 140、手动触发型 100）；`tests/test_context_budget.py` 保护预算判定。它回答「上下文是否超预算」，不回答「内容是否好」。

## Boundary

- 输入是仓库根（默认本 checkout，`--root` 可指向别处）；输出是逐条 `path:line: 原因`。
- 只读：不访问网络、不修改文件、不接触真实安装目录。
- 无问题时输出 `PASS` 并以 0 退出；有问题时输出 `FAIL` 与问题清单并以 1 退出；缺少 PyYAML 或 markdown-it-py 时按依赖缺失退出，不降级成宽松检查。
- 检查范围：Skill 发现与两级布局、frontmatter（YAML 可解析、无重复键、必填字段、`name` 与目录一致、自定义字段位置）、Skill 内全部文本资源的本地链接与越界、机器专属绝对路径与私钥标记。
- 链接边界随文件类型：Skill 内文件以该 Skill 目录为界，仓库 Markdown（根目录与 `docs/`）以仓库根为界。
- `validate-video-skills.py` 是保留的旧入口，只转发到同一校验器；新操作使用 `validate_skills.py`。

## Known limitations

- 自动检查与内容质量分层：`description` 的触发含义、技能边界、事实准确性和实际任务效果只能人工复核。
- 不校验 Markdown 片段锚点、外部 URL、HTML 中的链接，以及没有写成 Markdown 链接的裸路径。

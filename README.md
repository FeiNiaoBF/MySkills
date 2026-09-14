# My Skills

自用 AI Skills，记录可复用、可验证的工作流。

## 视频工作流

| Skill | 职责 |
|---|---|
| `math-video-workflow` | 按阶段编排一集数学视频，管理状态与人工确认 |
| `video-smart-cut` | 本地转写、智能精简、字幕对齐、画幅处理和成片校验 |
| `math-video-review` | 两阶段数学审校、口误判断和 Manim 候选清单 |
| `manim-recap-video` | 将获批候选制作成无旁白、无 BGM 的 Manim 片尾重温段 |
| `math-manim-insertion` | 为教学视频写插入式 Manim 场景：设计、符号与视觉规范 |
| `manim-video-insertion` | 按字幕语义建立插入时间表，把已验收场景装配进成片 |

推荐入口：先使用 `math-video-workflow`，每次只推进一个阶段。`workflow.yaml` 是单集状态的唯一来源；视频大文件和课程工程留在各自项目目录，不存入本仓库。

两条 Manim 路线勿混用：**讲解中途画面替换**走 `math-manim-insertion` + `manim-video-insertion`（原声连续、总时长不变）；**片尾独立重温段**走 `manim-recap-video`（追加拼接）。

## 其他 Skills

- `project-tech-mentor`：通过项目学习一门技术
- `teach-go`：在项目中学习 Go 标准库
- `frontend-guide`：前端实现指导
- `network-security-check`：网络安全检查
- `aesthetic-translator`：把审美意图转换为可执行设计语言
- `better-designs-md`：改进设计文档
- `code-subtraction`：以删减为目标审查代码

## 安装

使用 `scripts/link-skills.ps1` 或 `scripts/link-skills.sh` 将需要的 skill 链接到目标 agent 的 skills 目录。执行前先阅读脚本并确认目标路径。

提交前跑 `python scripts/validate-video-skills.py`：校验视频/Manim 技能的 frontmatter、description 长度、机器绝对路径和相对链接。

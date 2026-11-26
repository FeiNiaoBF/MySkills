---
name: tech-mentor
description: 用户要用递进项目学会一项技术、而不只是读概念时使用；以可验证的实操项目推进；不用于单次问答或现成项目理解。
license: MIT
metadata:
  author: Yeelight
  version: "1.2.0"
disable-model-invocation: true
---

# Tech Mentor

创建文件前先确认学习工作区；当前目录可能属于无关仓库。目标是可验证的项目能力：学习者能在引导逐步减少的情况下，构建、验证、调试并解释有用的技术产物。

适用技术领域示例：前端、后端、Three.js、LLM、计算机视觉。

## 选择当前模式

先检查工作区，再决定下一步。

- **新建路线：** 阅读[工作区布局](references/workspace-layout.md)、[任务格式](references/MISSION-FORMAT.md)、[路线格式](references/TRACK-FORMAT.md)、[术语表格式](references/GLOSSARY-FORMAT.md)和[资源格式](references/RESOURCES-FORMAT.md)。在选第一批项目前，确立北极星项目、实际约束和一个有界诊断任务。
- **继续路线：** 阅读 `MISSION.md`、`ROADMAP.md`、`STATE.md`、`GLOSSARY.md`（如存在）和最近的相关学习记录。创建或修订这些文件时，使用[路线图格式](references/ROADMAP-FORMAT.md)、[状态格式](references/STATE-FORMAT.md)和[学习记录格式](references/LEARNING-RECORD-FORMAT.md)。从最小的未完成、可验证任务继续。
- **跑微项目或审阅提交：** 阅读[项目循环](references/project-loop.md)、[项目简报格式](references/PROJECT-BRIEF-FORMAT.md)和[证据格式](references/EVIDENCE-FORMAT.md)，按其中的简报、反馈、证据和反思循环推进。
- **创建或修订领域路线：** 阅读[路线格式](references/TRACK-FORMAT.md)，遵循其来源和项目选择规则。主张和版本敏感选择基于当前一手来源；把来源、版本和为何重要记入 `RESOURCES.md`。
- **创建 HTML 课程或参考卡：** 阅读[HTML 课程格式](references/LESSON-FORMAT.md)。只有可视化、交互或适合打印的材料能改善记忆时才使用。

## 构建学习路径

从学习者想要的产物开始，而不是章节清单。用诊断结果在 `ROADMAP.md` 中建立暂定能力阶梯；每个项目应增加一项清晰可观察的能力。只展开下一个项目，后续工作只写到证据支持的粒度。按已展示的工作、可用时间和学习者目标调整范围、脚手架和开放度。

用让技术真正起作用的真实约束：视觉工作要交互，API 要契约，LLM 流程要固定评估集，计算机视觉要数据和错误案例。把工具样板与当前学习目标分开；提供的脚手架要标注，让学习者知道哪些部分由自己负责。

练习足够后用迁移检查测试独立性：让学习者在更少提示下完成修复、受约束重建、代码审查或需求变更。较长路线以目标和约束定义的作品收尾，而不是教程步骤。

## 在工作中教学

只教下一个实现决策所需的概念。然后给学习者一个有界任务，让其先尝试，再评价结果。

卡住时按帮助阶梯处理，停在第一个能解除阻塞的层级：

1. 澄清预期行为、观察结果和确切证据。
2. 指出要检查的最小子系统或假设。
3. 给一个保留学习者下一步决策的最小提示或问题。
4. 展示伪代码或局部示例。
5. 用户明确要求时，解释一个可行实现和立即可用的验证方式。

提供代码时，只限于障碍附近，解释其作用，并附上具体验证步骤。把实质性帮助记入项目反思，让后续迁移检查能针对同一能力。

## 验证、记录与恢复

项目具备所有适用证据时才算完成：

- **机器可验证：** 构建、测试、lint、请求断言、指标或等价检查。
- **人可验证：** 截图、录屏、输出样例、日志、错误案例集或可观察行为。
- **学习者可验证：** 对关键选择、失败路径和下一步改进的简洁解释。

运行可用检查，不要假定它们会通过。用[证据格式](references/EVIDENCE-FORMAT.md)把证据存入 `evidence/`，更新 `ROADMAP.md`，只在学习者能正确使用术语后加入 `GLOSSARY.md`，为非显而易见的经验创建学习记录，并在 `STATE.md` 留下当前能力、未解决问题、下一个最小任务和排定的迁移检查。

HTML 课程和参考卡按[HTML 课程格式](references/LESSON-FORMAT.md)执行。可运行项目、简报、状态和证据保留在代码或 Markdown 文件中。

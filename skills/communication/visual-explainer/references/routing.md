# 与相关 Skill 的边界

从 [visual-explainer](../SKILL.md) 拆出的按需参考；只在用户问“该用哪个”或边界不清时读。

- `visual-explainer` builds one complete, standalone HTML explanation after an explicit request.
- `make-sense` repairs a single confusing term, passage, or inference in the conversation and returns to the original task.
- `research-report` investigates a question across sources and publishes a traceable report ledger; this skill explains what is already understood and does not run retrieval or invent citations.
- `style-translator` and `ui-planner` produce design previews and implementation handoffs for an interface; this skill explains a subject, not a UI to build.
- `project-guide` explains an existing codebase in conversation; use this skill only when the user wants a standalone page.

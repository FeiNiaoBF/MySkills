---
name: clarify
description: Use when a user cannot express an idea well enough to continue the current task, explicitly or through a material communication blockage. Recover usable meaning, offer candidate terms when helpful, and return to the task; not for merely informal wording or explaining an existing term.
metadata:
  version: "1.0.0"
---

# clarify

Help the user find language for an intended meaning that is blocking the current task. **Complete when the intent is clear enough to execute that task correctly.** A unique technical term and further polishing are not required.

## When to use

Use for an explicit “I know what I mean but cannot describe/name it,” or a clear implicit blockage: the user repeatedly rejects interpretations but cannot replace them, or vague references to a *critical* concept prevent progress. Do not trigger because wording is colloquial or nonprofessional when it already supports the task. A request to understand an existing term, passage, or reasoning step belongs to `make-sense`; translating visual taste into design specifications belongs to `aesthetic-translator`.

## Repair the expression, then resume

1. In the context of the user's task, identify the smallest gap between their words and what must be understood to proceed. Preserve their stated intent; mark inference as tentative and do not invent requirements or domain facts.
2. Offer vocabulary only if it helps convey that intent. When several professional concepts plausibly fit—including when the user names alternatives—name **2–4 likely candidates** and explain how each differs in one sentence, grounded in this context, before restating the intent. Do not silently select one label and discard the others. If none fits reliably, use accurate plain language. Do not make the user choose a term unless the difference changes what should be done.
3. Restate the meaning precisely, retaining any consequential uncertainty. Show the ambiguous wording or your inference briefly when it helps the user learn a distinction or judge a strong inference; otherwise skip the analysis. Avoid a mandatory response template.
4. If that understanding suffices to execute the original task correctly, end `clarify` and continue the task. Do not routinely ask “Is that right?”, keep optimizing wording, or conduct a requirements interview. If competing meanings change the immediate result, ask the minimum discriminating question needed to proceed.
5. If the unresolved ambiguity instead calls for a broader choice of goals, requirements, or approaches, explain what differs and why it matters; **recommend** `grilling` for that decision. Stop after the recommendation: do not ask the first interview question, silently start `grilling`, choose a product direction, or recommend one before the relevant decision is made. If they do not enter it, leave the consequential branch explicitly unresolved.

Resolve whether the meaning is usable before worrying whether its label is standard. Uncertainty alone is not a reason to question the user; only a difference that affects the next decision is.

Maintainer-only behavior cases: [cases](references/cases.yaml). They are not required reading during normal use.

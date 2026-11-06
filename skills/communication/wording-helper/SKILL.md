---
name: wording-helper
description: >-
  Use when users need clearer wording or structured expression. Helps turn rough,
  complex, or partly expressed ideas into clear, accurate wording by first
  understanding intent and organizing information. Leaves actionable requests alone
  unless wording help is requested.
metadata:
  version: "3.0.1"
---

# Clarify Wording

Use structured thinking to turn a user's intended meaning into precise, usable expression. The work is not merely polishing sentences: understand the task, identify and prioritize the information, organize its relationships, then write and verify. Accuracy comes before elegance; concise means no unnecessary wording, not the fewest possible words.

## When to activate

Activate when the user:

- explicitly invokes `wording-helper`;
- asks to clarify, organize, rewrite, or improve wording or an idea; or
- after a misunderstanding, indicates that their meaning still needs to be expressed more clearly.

Do not activate just because a request is colloquial if it is already actionable. Do not treat a request to explain an existing term, passage, or reasoning as a request to rewrite it. Do not turn a visual preference into a specification unless the user asks for wording about it.

When the user explicitly invokes `wording-helper`, make a copy-ready wording draft the first substantive part of the reply. Put it in a standalone fenced code block labeled `txt`; do not put an introduction, heading, or explanation before the block. Keep the block itself limited to the text the user can reuse. Any brief explanation may follow it. This applies even when the underlying request is already actionable: the user asked for wording help, not execution.

If the user explicitly asks for analysis, an outline, or a structure **instead of** reusable wording, provide that form without forcing a `txt` block. A request for both wording and analysis still gets the `txt` block first.

## Core workflow

Think through these steps in order. They guide the work; do not mechanically expose every step or force a fixed template onto the final wording.

1. **Understand intent.** Identify what the user wants to communicate, the intended audience and purpose, and any stated position, tone, or constraints. Use the relevant conversation context, not just the latest sentence. Separate explicit information from inference. Do not mistake the user's phrasing for the underlying goal.
2. **Extract information.** Sort the source material into relevant facts, background, problem, goal, method, constraints, evidence, and open choices as applicable. Keep only categories useful for this task; do not invent missing content.
3. **Set priorities.** Identify the core meaning that cannot be lost, supporting information that helps explain it, and details that are repetitive or unnecessary. Preserve emphasized requirements and meaningful caveats.
4. **Organize logic.** Choose an order that makes the relationships easy to follow. Use the structure that fits the content—for example, context → problem → approach → outcome for an introduction; observation → cause → remedy for a technical explanation; or claim → reasons → evidence → limits for an argument. Treat these as options, not templates.
5. **Refine expression.** Use direct, natural, precise language in the intended language and tone. Remove repetition and ambiguity, connect related ideas, and use terminology only when it improves precision for the audience. Keep enough explanation to preserve meaning; do not compress away useful context.
6. **Check quality.** Confirm that the result is faithful, complete on the stated essentials, logically ordered, readable, and no longer than the purpose requires. For a copy-ready task request, make the requested action, scope, deliverable, and stated constraints easy to identify; include acceptance conditions only when the user supplied or clearly implied them. Check that no unsupported facts, goals, decisions, promises, or conclusions have slipped in.

## Handling ambiguity

Resolve ordinary wording choices from context. Preserve an undecided detail in the draft when that can represent the user's meaning accurately. Ask one focused question only when the missing choice would materially change the meaning and cannot be preserved without misrepresenting the user. Do not turn clarification into a broad requirements interview or stress-test unrelated decisions.

If the user says the assistant misunderstood or acted prematurely, acknowledge that plainly, without defending the earlier assumption. Then use the correction and the relevant context to restate the intended meaning.

## Draft and review

- Write in first person by default when drafting text the user may reuse; change perspective when the audience or purpose calls for it.
- Use the language the user intends for the draft. The explanation around it should follow the user's conversation language.
- For explicitly invoked wording work, the standalone `txt` block is the required primary deliverable and comes first. For other wording requests, when the user needs a reusable passage, present it in a standalone `txt` code block and put only the draft inside it. Keep the surrounding explanation brief; do not repeat the full draft outside the block.
- If the user asks for analysis, an outline, or a structure rather than ready-to-send wording, provide that form instead of forcing a `txt` block, unless they also requested wording.
- After presenting a draft, wait for review. If the user corrects it, revise and wait again. Approval of wording alone is not approval to carry out the described task. If the user approves and asks to proceed, use the approved wording as the instruction without asking for redundant confirmation. If they approve without asking to proceed, acknowledge the approval and stop.

## Boundaries

The result must remain the user's meaning. Improve clarity and persuasiveness only through better selection, ordering, and expression of information supported by the conversation. Never add a preference, requirement, fact, commitment, or conclusion just because it would make the wording sound more complete or professional.

## Maintenance

Use [comparison cases](references/cases.yaml) when evaluating changes to this skill. They describe expected behavior for review; they are not runtime instructions and should not be loaded for ordinary wording requests.

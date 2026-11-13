---
name: wording-helper
description: 用户需要更清楚的措辞或有结构的表达时使用；先把意图说清，再给可复用文字；不用于可执行请求或术语解释。
metadata:
  version: "3.3.0"
---

# Wording Helper

Turn a user's intended meaning into precise, usable expression. Understand the task, prioritize the information, organize its relationships, then write and verify. Accuracy comes before elegance; concise means no unnecessary wording, not the fewest possible words.

## Explicit invocation contract (mandatory)

When the user invokes `wording-helper` by name, this contract governs the turn from its first token. It is an order to execute, not reference material: discard any plan formed from the user's task framing before applying it.

The first substantive content of the reply MUST be a standalone `txt` code block containing only copy-ready wording the user can reuse — no heading, lead-in, acknowledgement, intent summary, outline, plan, tool call, or question before it. The block is one continuous paragraph: no line breaks, no bullet or numbered lists, no headings or emphasis markers. Brief explanation and at most one focused question MAY follow the block.

Before that block, do NOT:

- call tools, inspect the environment, read files or repositories, or gather context that is not already in the conversation;
- ask about deliverable form, audience, length, scope, pain points, or boundaries — a user line such as "ask me if anything is unclear" authorizes one focused wording question later, never a requirements interview;
- restate these rules, narrate the diagnosis, or explain what the skill does;
- start the described task, even when the request reads like an instruction ("write a plan", "list my skills", "fix this bug"). Draft the wording of that request; execution waits for approval.

If the user explicitly asks for analysis, an outline, or a structure **instead of** reusable wording, provide that form without a `txt` block. If they ask for both, the `txt` block still comes first.

The only case that replaces the draft is a bare invocation with no meaning to work from (for example `/wording-helper` alone): ask one focused question about what they want to express — still not an interview.

When a previous draft is already awaiting review, the follow-up decides the mode, not the invocation prefix. A correction or complaint about the draft means write the replacement draft (block first, single paragraph). An approval plus an instruction to proceed means execute the described task now, without another `txt` block and without asking again, even if the message repeats `wording-helper`.

## When to activate

Activate when the user:

- explicitly invokes `wording-helper`;
- asks to clarify, organize, rewrite, or improve wording or an idea; or
- after a misunderstanding, indicates that their meaning still needs to be expressed more clearly.

Do not activate just because a request is colloquial if it is already actionable. Do not treat a request to explain an existing term, passage, or reasoning as a request to rewrite it. Do not turn a visual preference into a specification unless the user asks for wording about it.

## Core workflow

Think through these steps in order. They guide the work; do not mechanically expose every step or force a fixed template onto the final wording.

1. **Understand intent.** Identify what the user wants to communicate, the intended audience and purpose, and any stated position, tone, or constraints. Use the relevant conversation context, not just the latest sentence. Separate explicit information from inference. Do not mistake the user's phrasing for the underlying goal.
2. **Extract information.** Sort the source material into relevant facts, background, problem, goal, method, constraints, evidence, and open choices as applicable. Keep only categories useful for this task; do not invent missing content.
3. **Set priorities.** Identify the core meaning that cannot be lost, supporting information that helps explain it, and details that are repetitive or unnecessary. Preserve emphasized requirements and meaningful caveats.
4. **Organize logic.** Choose an order that makes the relationships easy to follow. Use the structure that fits the content—for example: context → problem → approach → outcome for an introduction; observation → cause → remedy for a technical explanation; or claim → reasons → evidence → limits for an argument. Treat these as options, not templates.
5. **Refine expression.** Use direct, natural, precise language in the intended language and tone. Remove repetition and ambiguity, connect related ideas, and use terminology only when it improves precision for the audience. Keep enough explanation to preserve meaning; do not compress away useful context.
6. **Check quality.** Confirm that the result is faithful, complete on the stated essentials, logically ordered, readable, and no longer than the purpose requires. For a copy-ready task request, make the requested action, scope, deliverable, and stated constraints easy to identify; include acceptance conditions only when the user supplied or clearly implied them. Check that no unsupported facts, goals, decisions, promises, or conclusions have slipped in.

## Handling ambiguity

Resolve ordinary wording choices from context. Preserve an undecided detail in the draft when that can represent the user's meaning accurately. Ask one focused question only when the missing choice would materially change the meaning and cannot be preserved without misrepresenting the user. That question, if any, comes after the draft and is one question — not a form, checklist, or stress test of unrelated decisions.

If the user says the assistant misunderstood or acted prematurely, acknowledge that plainly, without defending the earlier assumption. Then use the correction and the relevant context to restate the intended meaning.

## Draft format and review

- Write the reusable text as one unbroken paragraph of plain prose. Do not start a new line per sentence, per item, to list alternatives, or to fit a width; do not use bullet points, numbering, headings, or emphasis inside the block. Express enumeration inline within the sentence (for example "一是…；二是…"). A fenced block preserves line breaks literally, so any break shows up as a broken draft.
- Only when the user explicitly asks for a multi-paragraph or sectioned document may the block contain paragraph breaks. Even then, never one line per sentence and never a list rendered as separate lines.
- Write in first person by default when drafting text the user may reuse; change perspective when the audience or purpose calls for it.
- Use the language the user intends for the draft. The explanation around it should follow the user's conversation language.
- For explicitly invoked wording work, the standalone `txt` block is the required primary deliverable and comes first. For other wording requests, when the user needs a reusable passage, present it in a standalone `txt` code block and put only the draft inside it. Keep the surrounding explanation brief; do not repeat the full draft outside the block.
- After presenting a draft, wait for review. Exactly two kinds of follow-up decide what happens next:
  - **Correction or rejection** — the draft is not what they want, or they name what to change. Write the replacement draft: `txt` block first, single paragraph, no interview, no execution of the described task. When they only say it is wrong without saying how, infer the fix from the conversation and the earlier draft rather than asking, and keep the replacement ready to send.
  - **Approval plus an instruction to proceed** ("按这个继续", "就按这个做，开始吧") — this is a task instruction, not wording work. Execute it directly using the approved wording. Do not emit another `txt` block, do not restate the wording, and do not ask for confirmation. The `wording-helper` prefix does not change this.
- Approval with no instruction to proceed is not a third path and not authorization: acknowledge it and stop. The described task still waits for an explicit go-ahead.

## Boundaries

The result must remain the user's meaning. Improve clarity and persuasiveness only through better selection, ordering, and expression of information supported by the conversation. Never add a preference, requirement, fact, commitment, or conclusion just because it would make the wording sound more complete or professional.

## Maintenance

Use [comparison cases](references/cases.yaml) when evaluating changes to this skill. They describe expected behavior for review; they are not runtime instructions and should not be loaded for ordinary wording requests.

---
name: clarify-wording
description: >-
  This skill helps users rephrase ideas or express unclear intent. It drafts
  wording for review when explicitly invoked, when users request wording help,
  or when a rejected interpretation reveals difficulty expressing their meaning.
  It leaves actionable requests alone unless wording help is requested.
metadata:
  version: "2.0.0"
---

# Clarify Wording

This skill turns an unclear or partly expressed intent into clear, natural wording the user can reuse. When activated, it produces a draft for the user's review; the draft is not approval to carry out the task it describes.

## Activation and routing

Activate when the user:

- explicitly invokes `clarify-wording`;
- asks to clarify an idea, improve wording, or express an intended message more clearly; or
- struggles to restate the intended meaning after persistent misunderstandings or a rejected interpretation.

For the initial response to an explicit skill invocation, prepare a draft for review even when the underlying task is already clear. The invocation does not authorize carrying out that task. After draft review, follow the continuation rule in step 6 instead of restarting this workflow.

Without explicit invocation or a request for wording help, leave already actionable requests alone. Route questions about the meaning of an existing term, passage, or reasoning to `make-sense`, and vague visual preferences needing design specifications to `aesthetic-translator`. These routing defaults do not override an explicit request to draft wording about those topics.

## Workflow

1. Infer the user's stated goal, position, constraints, and tone from the request and context. Separate what the user said from any unsupported assumption. If the user says the assistant acted prematurely, acknowledge that mistake without defending it.
2. Briefly state the understood intent. Offer two to four options only when the possible meanings would lead to materially different wording.
3. Preserve unresolved choices in the wording. If that would still misrepresent the user's position, ask one focused question and wait for the answer before drafting. Do not turn clarification into a requirements interview.
4. Draft concise, natural wording in the first person by default and in the language the user intends to use. Preserve the user's intent, tone, and important limits. Do not add goals, commitments, facts, or preferences.
5. Present the draft in a standalone `txt` code block. Put only the wording in the block. Do not repeat the full draft outside it.
6. Pause for review. If the user corrects the draft, revise it and wait again. Approval of the wording alone does not authorize execution. When the user confirms it and asks to proceed, use the approved wording as the task instruction without drafting it again or requesting redundant confirmation.

## Boundaries

The output is wording for the user to review, not a specification or a decision on the user's behalf. Strengthen the expression only with meaning supported by the conversation. Use `grilling` only when the user asks to examine or stress-test an unresolved decision; it is not a routine clarification step.

## Reply shape

Use a short understanding statement in the user's language, then introduce the draft with an equivalent of “You could say:”. Keep explanations and questions outside the `txt` block. Apply step 6 after presenting the draft.

## Maintenance

Use [comparison cases](references/cases.yaml) when evaluating changes to this skill. They contain review expectations, not runtime instructions; do not load them for ordinary wording requests.

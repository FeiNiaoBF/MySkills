---
name: clarify-wording
description: >-
  This skill activates when users ask to rephrase an idea or struggle to express
  intent. It also helps after repeated misunderstandings. It drafts reusable
  wording and does not handle actionable requests or explain existing terms and
  materials.
metadata:
  version: "2.0.0"
---

# Clarify Wording

This skill turns an unclear or partly expressed intent into clear, natural wording the user can reuse. When activated, it produces a draft for the user's review; the draft is not approval to carry out the task it describes.

## Activation and routing

Activate when the user:

- explicitly invokes `clarify-wording`;
- asks to clarify an idea, improve wording, or express an intended message more clearly; or
- has repeatedly corrected misunderstandings and appears unable to restate the intended meaning.

An explicit skill invocation always starts with a draft for review, even when the underlying task is already clear. It does not authorize carrying out that task.

Do not activate when the user's request is already actionable and they have not asked for help with wording. Route questions about the meaning of an existing term, passage, or reasoning to `make-sense`. Route vague visual preferences that need design specifications to `aesthetic-translator`.

## Workflow

1. Infer the user's stated goal, position, constraints, and tone from the request and context. Separate what the user said from any unsupported assumption. If the user says the assistant acted prematurely, acknowledge that mistake without defending it.
2. Briefly state the understood intent. Offer two to four options only when the possible meanings would lead to materially different wording.
3. Draft concise, natural wording in the first person by default and in the language the user intends to use. Preserve the user's intent, tone, and important limits. Do not add goals, commitments, facts, or preferences. Make unresolved choices explicit instead of deciding for the user.
4. Ask one focused question only when missing information would make the draft misrepresent the user's position and the uncertainty cannot be stated in the draft. Do not turn clarification into a requirements interview.
5. Present the draft in a standalone `txt` code block. Put only the wording in the block. Do not repeat the full draft outside it.
6. Pause for review. If the user corrects the draft, revise it and wait again. Continue with the approved wording only when the user confirms it and asks to proceed. Do not ask for approval a second time.

## Boundaries

The output is wording for the user to review, not a specification or a decision on the user's behalf. Strengthen the expression only with meaning supported by the conversation. When the user has not asked for wording help, do not interrupt an actionable request to search for terminology. Use `grilling` only when the user asks to examine or stress-test an unresolved decision; it is not a routine clarification step.

## Reply shape

Use a short understanding statement in the user's language, then introduce the draft with an equivalent of “You could say:”. Keep explanations and questions outside the `txt` block. After presenting the draft, stop and wait for the user's review.

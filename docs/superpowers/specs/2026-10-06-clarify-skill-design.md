# Clarify: explain and restate the user's intended meaning

Status: implemented in `skills/communication/clarify-wording/SKILL.md` (renamed from `clarify` on 2026-10-12).

## Outcome and boundary

`clarify-wording` helps users express an idea more clearly. It explains the likely meaning in context, then drafts a clearer first-person statement the user can review and use. The draft preserves the user's intent without adding decisions, facts, or preferences they did not establish.

When the user explicitly invokes the Skill, treat that as a request to prepare wording for review, not as approval to execute the request being phrased. Put the copy-ready draft in a standalone `txt` code block and pause. If the user confirms the wording and asks to continue, treat the approved wording as their task instruction and proceed. If the user edits it, revise the draft and wait for confirmation. Without an explicit wording request or Skill invocation, do not invoke `clarify-wording` for an already actionable request.

`clarify-wording` is not a general requirements interview. If a material gap prevents an accurate restatement, it asks only the smallest necessary question. It prefers wording that preserves uncertainty over repeated questioning. `grilling` is appropriate only when the user wants to explore or stress-test a broader decision. `make-sense` explains an existing term or passage; `aesthetic-translator` turns visual preference into design direction.

## Interaction

1. Read the user's words in task context. Separate stated intent and constraints from tentative inferences.
2. Explain the interpreted meaning briefly. Offer candidate terms only when their differences help express the user's idea.
3. Draft a polished first-person statement that preserves the user's stance, goals, constraints, and uncertainty. Improve clarity and emphasis without inventing commitments. Put the copy-ready statement in a standalone `txt` code block, with no explanation inside the block.
4. If an unresolved detail would make the draft misrepresent the user, ask one focused question. Otherwise preserve the uncertainty in the wording.
5. Pause for review. If the user corrects the draft, revise it and wait again. If the user confirms the draft and asks to continue, use the approved wording as the task instruction. Never execute the task from the draft before that confirmation.

## Acceptance and verification

Review cases covering explicit Skill invocation, explicit requests for wording help, implicit expression blockage, actionable colloquial requests, confirmation and continuation after draft review, explanation of existing terms, material missing details, preservation of uncertainty, and correction after premature action. Expected behavior is a concise interpretation plus a directly usable first-person draft in a standalone `txt` block; wait for review, do not execute before confirmation, and continue from the approved wording when the user asks. Avoid invented choices and exhaustive interviewing.

Use `skills/communication/clarify-wording/references/cases.yaml` as maintainer-only examples, not as runtime material. Run the repository's skill validator and diff checks. Human review must confirm the trigger fits the new purpose and that examples do not reward speaking for the user.

## Out of scope

No automatic execution of the restated request, generalized product discovery interview, mandatory terminology selection, or automatic `grilling` handoff. This revision changes the Skill's output from resolving wording and resuming the task to explaining and drafting the user's own improved wording.

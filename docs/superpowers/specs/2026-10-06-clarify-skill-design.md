# Clarify: explain and restate the user's intended meaning

Status: implemented in `skills/communication/clarify-wording/SKILL.md` (renamed from `clarify` on 2026-10-12).

## Outcome and boundary

`clarify-wording` helps users express an idea they have not yet stated clearly. It explains the likely meaning in context, then drafts a clearer first-person statement the user can adopt or edit. The goal is to make the user's own intent more complete, precise, and expressive without adding decisions, facts, or preferences the user did not establish.

The skill stops after delivering the explanation and draft. It does not automatically perform the task described in the new wording. The user may adopt, correct, or extend the wording, then ask to proceed.

`clarify-wording` is not a general requirements interview. If a material gap prevents an accurate restatement, it asks only the smallest necessary question. It prefers wording that preserves uncertainty over repeated questioning. `grilling` is appropriate only when the user wants to explore or stress-test a broader decision. `make-sense` explains an existing term or passage; `aesthetic-translator` turns visual preference into design direction.

## Interaction

1. Read the user's words in task context. Separate stated intent and constraints from tentative inferences.
2. Explain the interpreted meaning briefly. Offer candidate terms only when their differences help express the user's idea.
3. Draft a polished first-person statement that preserves the user's stance, goals, constraints, and uncertainty. Improve clarity and emphasis without inventing commitments.
4. If an unresolved detail would make the draft misrepresent the user, ask one focused question. Otherwise preserve the uncertainty in the wording.
5. Deliver the explanation and draft, then stop for the user to adopt, edit, or ask to continue.

## Acceptance and verification

Review cases covering explicit requests for wording help, implicit expression blockage, actionable colloquial requests, explanation of existing terms, material missing details, preservation of uncertainty, and correction after premature action. Expected behavior is an explanation plus usable user-authored wording; avoid task execution, invented choices, and exhaustive interviewing.

Use `skills/communication/clarify-wording/references/cases.yaml` as maintainer-only examples, not as runtime material. Run the repository's skill validator and diff checks. Human review must confirm the trigger fits the new purpose and that examples do not reward speaking for the user.

## Out of scope

No automatic execution of the restated request, generalized product discovery interview, mandatory terminology selection, or automatic `grilling` handoff. This revision changes the Skill's output from resolving wording and resuming the task to explaining and drafting the user's own improved wording.

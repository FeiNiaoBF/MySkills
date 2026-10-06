# Clarify: recover the user's intended meaning

Status: proposed; implementation awaits review of this spec and an implementation plan.

## Outcome and boundary

`clarify` is a small, general-purpose skill for a user who has an idea but cannot yet express it well enough for the current task to proceed correctly. Its job is to recover usable meaning, not to conduct a full requirements interview or find a canonical technical term. **It completes as soon as the user's intent is clear enough to execute the current task correctly**; a unique term and further polishing are unnecessary.

It may trigger on an explicit request for help finding words or a term. Implicit invocation is appropriate only when an expression gap is materially blocking understanding or execution of the current task, such as repeated failed interpretations with no alternative wording. Informal, imprecise, or nonprofessional phrasing alone is not a trigger. If the existing words already support the task, continue without invoking `clarify`.

This boundary differs from `make-sense` (explaining a term, passage, or inference the user encounters) and `aesthetic-translator` (turning a visual preference into a design direction and specifications). Do not replace either. `grilling` is for decisions whose consequences need a deeper interview; `clarify` only identifies the expression and points out decision-changing ambiguity.

## Interaction

1. Read the user's words in the context of the task. Identify the smallest expression gap that actually blocks progress. Do not infer a detailed product brief or fill missing decisions with invented facts.
2. Offer vocabulary when it helps. If several concepts plausibly fit, offer **2–4 likely candidates**, each with a one-sentence distinction in terms of the user's context. Distinguish tentative interpretations from known facts; do not force a choice merely to settle terminology. If no defensible specialist term fits, use plain, accurate wording instead of fabricating one.
3. Restate the intended meaning in more precise language, preserving what remains uncertain. Show the ambiguous phrase or reasoning briefly only when the distinctions teach useful vocabulary, the restatement needs calibration, or the inference is strong. Avoid a fixed diagnostic template.
4. If the restatement is sufficient to execute the original task correctly, finish `clarify` and continue that task without routine confirmation or additional wording optimization. If two interpretations would materially alter the result, ask the minimum discriminating question needed to continue; confirmation is not the default.
5. If a decision-changing ambiguity requires a broader choice among goals, requirements, or approaches rather than a local expression repair, explain the ambiguity and its consequence, then **recommend** `grilling`. Do not silently start `grilling` or take the decision for the user. If the user declines, preserve the unresolved branch instead of pretending it was settled.

Core rule: resolve whether the meaning is usable before worrying whether its label is standard. Uncertainty by itself does not justify a question; only uncertainty that changes what should be done does.

## Packaging and dependencies

Add `clarify/SKILL.md` with valid YAML frontmatter (`name: clarify`, a trigger-first description, mapping-valued `metadata`). Keep the runtime guidance brief and self-contained; no scripts, network access, or external service. Add a discoverable README entry. A small `clarify/references/` case set is for maintainers to compare behavior, not required reading during normal invocation. Route to other skills by name, not by cross-skill relative links. There is no automatic installation, bootstrap, manifest edit, or change to existing skill content in scope.

## Acceptance and verification

Review response behavior on representative paired cases:

- **Trigger:** explicit inability to describe an idea or name a concept; implicit but task-blocking miscommunication.
- **Do not trigger:** colloquial wording that already conveys enough to complete the task; a request to explain an existing term (`make-sense`).
- **Stop and continue:** multiple plausible labels but an adequate plain-language restatement; no forced term selection or routine confirmation.
- **Discriminate locally:** competing meanings would materially change the immediate task; ask one pointed question rather than a full interview.
- **Hand off:** the ambiguity is a consequential broader decision; explain why, recommend `grilling`, and do not automatically invoke it.
- **Epistemic restraint:** do not invent domain terminology or convert an uncertain interpretation into a settled requirement.

Store concise cases with expected and unwanted behaviors, then, where model execution is available, compare actual responses against them. A format validator cannot prove that behavior. Run `python -B scripts/validate_skills.py`, `python -B -m unittest discover -s tests -v`, and `git diff --check`; manually review trigger wording, sensitive values, and link semantics. Report existing or environmental failures separately from failures caused by this change. Preserve unrelated working-tree changes.

## Out of scope

No generalized product-intake flow, automatic `grilling` invocation, term taxonomy, scripted dialogue engine, alteration of `make-sense` or `aesthetic-translator`, or publication/deployment changes. This design borrows the *idea* of helping users locate language, not any third-party skill text.

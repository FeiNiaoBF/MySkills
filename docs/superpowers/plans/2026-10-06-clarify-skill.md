# Clarify Skill Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Ship a compact, general-purpose skill that repairs task-blocking expression gaps and resumes the original task when intent is usable.

**Architecture:** A standalone `SKILL.md` holds the runtime behavior and routing boundary. A maintainer-only case file provides response-level comparisons; README makes the skill discoverable. No code, installation, or changes to neighboring skills.

**Tech Stack:** Agent Skills YAML frontmatter, Markdown, YAML behavior cases, Python repository validator and unittest suite.

**Spec:** [clarify design](../specs/2026-10-06-clarify-skill-design.md)

## Global Constraints

- Implicit invocation only when the expression gap materially blocks understanding or execution; nonprofessional wording alone is insufficient.
- Complete when intent is clear enough to execute the current task correctly; no unique term, routine confirmation, or further polishing required.
- When several concepts plausibly fit, offer 2–4 likely candidates with one-sentence distinctions; do not force a terminology choice.
- Ask the minimum discriminating question only for materially different immediate outcomes; explain and recommend `grilling` for broader decision ambiguity, without invoking it automatically.
- Keep runtime instructions brief; no scripts, network access, installation, bootstrap, manifest edits, or modifications to other skills.
- Preserve unrelated worktree changes, including concurrent edits to `README.md`; stage only explicit paths.

## Review Focus

These are concrete inputs within the spec's existing acceptance categories, not additional categories; include each in `cases.yaml` with expected and unwanted behavior:

1. A user repeatedly rejects a proposed interpretation without a replacement: clarify the blockage rather than repeat the rejected guess.
2. A colloquial but actionable request: proceed with the task, not a terminology exercise.
3. A direct request to explain a known term: leave to `make-sense`, not intent recovery.
4. Two plausible terms that imply the same action: offer concise distinctions, then proceed without demanding a choice.
5. Two plausible interpretations that alter the product goal: explain consequences and recommend `grilling`, without silently conducting an interview.

---

### Task 1: Deliver the clarify skill and acceptance cases

**Files:**
- Create: `clarify/SKILL.md` — self-contained runtime instruction and conservative trigger description.
- Create: `clarify/references/cases.yaml` — maintainer-only conversation inputs plus expected/unwanted outcomes.
- Modify: `README.md` — single skill selection entry, merging around existing worktree edits.

**Interfaces:**
- Consumes: the current user request and task context; routing by skill *name* to `make-sense`, `aesthetic-translator`, or `grilling` when relevant.
- Produces: task-usable restatement and continuation, a minimum local discrimination question, or an explained recommendation to enter `grilling`.

- [ ] **Step 1: Write behavior cases before the runtime instruction.** Use the six existing acceptance categories in the spec: explicit/implicit trigger; no-trigger colloquial and existing-term explanation; stop-and-continue without unique term; local discrimination; broader handoff; epistemic restraint. Include the five Review Focus inputs. Use `make-sense/references/cases.yaml` as a format reference, not a copied corpus.
- [ ] **Step 2: Check the cases against the pre-change behavior.** Where a response generator is available, run the cases before the new skill is loaded and note failures or ambiguity in a report under `${TEMP}/pi-agent/`. If unavailable, manually review case inputs and expected/avoid clauses, and state that model-baseline evidence is absent; do not claim a red test ran.
- [ ] **Step 3: Write `clarify/SKILL.md`.** Frontmatter: `name: clarify`, a trigger-first `description` whose first 57 characters identify communication blockage, and mapping-valued `metadata`. In the body: smallest blocked meaning, 2–4 contextual concept candidates when multiple fit, precise restatement without invented certainty, stop-and-resume criterion, minimum consequential question, and recommend-but-do-not-start `grilling`. Distinguish `make-sense` and `aesthetic-translator`; avoid fixed output templates and mandatory reading of cases.
- [ ] **Step 4: Add one `README.md` row.** Place `clarify` next to `make-sense` under `### 工程维护与理解`, describing its expression-repair trigger. Inspect the current README diff first and retain unrelated edits.
- [ ] **Step 5: Validate and repair only introduced failures.** Run `python -B scripts/validate_skills.py`, `python -B -m unittest discover -s tests -v`, and `git diff --check`. Manually review description prefix, link semantics, sensitive values, and boundaries. If model response generation is available, run the same cases with the skill and record concrete comparisons; otherwise label behavior unverified rather than equating format checks with effectiveness. Separate pre-existing validator failures from new failures.
- [ ] **Step 6: Review and commit.** Inspect the full diff and status; stage only `clarify/SKILL.md`, `clarify/references/cases.yaml`, and `README.md` if the latter can be staged without other changes (otherwise stage only its targeted hunk). Commit a focused English Conventional Commit. Do not push, install, or touch the recovery manifest.

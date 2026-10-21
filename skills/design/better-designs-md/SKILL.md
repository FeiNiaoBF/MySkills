---
name: better-designs-md
description: >-
  This skill writes UI visual and component design docs on explicit request. It creates,
  reviews, and updates UI design guidance from existing evidence or proposed
  directions, explaining visual intent, component usage, tokens, states, and contextual
  rules. It reuses project conventions and supports Google's DESIGN.md format
  when requested or required by a target tool. Project requirements, behavior specs,
  technical designs, architecture, and ADRs belong to project-docs. It does not turn
  ordinary styling edits or token questions into a documentation workflow.
metadata:
  author: FeiNiaoBF
  version: 1.3.1
  date_added: "2026-04-09"
  risk: low
disable-model-invocation: true
---

# Better Designs MD

Create and maintain UI design documentation that guides actual UI work: explain the intended look and feel, give applicable rules, and preserve the reasoning needed to apply them consistently. Write the document and conversation in the user's language unless the project specifies otherwise; preserve existing identifiers and any selected schema keys.

## Scope and invocation

Use this workflow when the user explicitly requests this skill or asks to author, review, or update a UI `DESIGN.md` or equivalent design-system document. A passing mention of a color, token, or component does not require a document. The host-specific manual invocation settings are intentional; hosts may enforce them differently.

This skill owns UI visual and component documentation, not every document that uses “design” in its name. Use `project-docs` for project requirements, externally observable behavior specifications, internal technical design, system architecture, and ADRs. A file named `DESIGN.md` does not decide the route: its primary reader task and content do. When one document needs both kinds of information, keep each rule in its authoritative source and use a short summary or link instead of copying the full specification.

It can turn an approved visual direction, inspected UI, existing tokens, or a supplied design guide into a document without requiring another skill first. Other skills may complement it when the user chooses them, with no required order or handoff. A documentation task alone does not authorize UI implementation, external design-tool changes, or publication.

## Choose document scope and format

Establish who will use the guidance, which recurring UI decisions it should support, and what existing sources already cover. Reuse the project's design documents, component library, tokens, and Storybook where present. A short addition explaining intent, component selection, exceptions, and links to authoritative definitions may be sufficient; do not duplicate a maintained catalog or create a separate file without a useful gap to fill. Honor an explicit request for a standalone document while keeping its content proportional.

Choose the file name and structure from project conventions and the requested deliverable. Use Google's DESIGN.md format only when the user requests that format or a target tool requires it. The filename `DESIGN.md` alone does not require Google's schema. Ordinary project guidance can use concise Markdown without YAML tokens or a fixed section inventory.

A document does not enforce UI rules or guarantee that an agent reads it. Explain where the guidance fits into the project's existing reading/review workflow and which rules are implemented by components or theme configuration. Do not claim automatic loading, synchronization, or enforcement without an actual integration.

## Workflow

1. **Locate the source of truth.** Read relevant project guidance and find the existing design document, tokens, and component conventions. Reuse the established file and naming scheme; do not create a competing `DESIGN.md`. For a new document, use the project's documented location or the user's designated path. Ask only if the destination remains materially ambiguous. For review-only requests, report findings without rewriting; honor requests for a draft in conversation or no files.
2. **Separate evidence from decisions.** Identify which choices are user requirements, confirmed decisions, implementation observations, or proposals. Read supplied references where available and record useful source links or project paths. A screenshot can support a visual interpretation, but does not establish exact fonts, spacing, states, or accessible behavior. If code and documentation disagree, describe the conflict; do not silently promote the current implementation or overwrite an approved rule.
3. **Choose the shortest useful drafting path.** With sufficient material, draft the whole document or update the affected sections directly. Do not repeat a section-by-section interview for settled decisions. With missing context, ask one focused question when the answer materially changes the design rules; continue independent work and label remaining proposals or unknowns. If the user delegates design choices, propose a coherent direction within that scope. Use section-by-section checkpoints only when requested or needed to resolve a consequential dependency.
4. **Write the rationale and applicable rules.** Explain audience, intended feeling, hierarchy, and the reasons behind important rules. Include only relevant sections, components, variants, and states. Link authoritative token definitions; include machine-readable values only when useful to the requested deliverable. When Google format is selected, read the [Google DESIGN.md format reference](references/design-md-format.md) before authoring or structurally revising it. Reuse established token names and units; do not invent a palette, type scale, theme, or component inventory to fill a checklist. Keep proposed values visibly provisional and unresolved values out of machine-readable token definitions.
5. **Maintain consistency across revisions.** Update affected prose, token definitions, references, and examples together. Keep exact values in one authoritative place and refer to them elsewhere. If production tokens remain authoritative, identify them and reconcile the document against them; do not imply that a copied value is automatically synchronized. Preserve unrelated choices and custom sections. Mark unresolved conflicts and migration needs instead of silently renaming tokens or normalizing the entire file.
6. **Verify and deliver.** Apply the completion checks below. Provide the document link or requested draft, summarize changed rules and unresolved decisions, and state what was checked. A reviewable draft can be delivered before approval; do not label it accepted or stable without evidence. Continue implementation when already authorized, using the project's stack and checks, rather than asking for the same permission again.

## Authoring principles

- **Intent before a catalog.** Describe concrete visual relationships and how they serve the audience. A list of adjectives or token values alone cannot explain when a rule applies. Attribute inspected reference traits; distinguish them from a proposed interpretation.
- **Proportional coverage.** Explain color roles and pairings, typography hierarchy, layout and spacing, depth, shapes, and relevant component states where the project needs them. Flat designs can define hierarchy through borders or tonal layers. Do not require extra palettes, shadows, rounded corners, or multiple themes.
- **Rules with context.** Make do's and don'ts specific to this design, with a reason or a useful comparison. Avoid universal style bans and arbitrary quotas. Preserve intentional exceptions and state where they apply.
- **Accessibility as requirements and evidence.** Record relevant contrast pairings, focus visibility, keyboard behavior, labels, error states, and reduced-motion behavior. Distinguish desired standards from measured checks; a color list or written guardrail does not prove WCAG conformance. Record the tested state/theme and any remaining checks.
- **Platform fidelity.** Preserve native units and design-system terminology in the prose. If the target's values do not fit the selected schema or consumer, document the mapping or limitation rather than silently converting them into Web values.

## Completion checks

- **Scope and status:** the file or review covers the requested design decisions, reuses existing conventions, and distinguishes confirmed rules, observations, proposals, and gaps. No guessed author, reviewer, approval, or validation status appears as fact.
- **Format:** follow the project's chosen structure. Apply the Google reference only when that format is selected; apply any other consumer's requirements only when relevant. For machine-readable blocks, check syntax, duplicate definitions, and the selected schema. Preserve intentional extensions and note consumer-specific limitations. If parsing or consumer tooling is unavailable, state which checks were manual.
- **Tokens, when included:** references resolve without cycles, types and units fit their roles, and quoted colors/references parse as intended. Prose and examples agree with authoritative definitions. Check relevant themes and variants without claiming unrepresented states exist.
- **Usability:** a reader can tell what to apply, why, and when an exception is allowed. Component guidance connects appearance to relevant states and behavior. Accessibility claims cite actual checks rather than aspirations.
- **Delivery:** link the artifact or provide the requested review/draft, with unresolved choices and verification limits. Document validation is not visual approval, production implementation, or a full accessibility audit.

## Reference navigation

- [Google DESIGN.md format reference](references/design-md-format.md): optional format-specific guidance on YAML placement, schema boundaries, sections, sources, and validation. Read only when authoring or checking Google-format documents; ordinary project documentation does not require it.

For Google-format work, the reference records the upstream alpha specification checked on its stated date. When a task requires a newer version or a particular importer, verify that target's requirements before claiming compatibility. If source access is unavailable, work from the recorded baseline and disclose the limitation.

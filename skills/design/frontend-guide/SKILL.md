---
name: frontend-guide
description: >-
  This skill helps users choose UI content and build interfaces, including
  explicit module-by-module guidance. It clarifies basic and optional areas,
  organizes selected content into a page structure, and pairs conversational
  explanations with HTML design drafts. It supports design-only work and
  implementation in the project's stack without restarting discovery for local
  UI edits.
metadata:
  version: "0.2.2"
---

# Frontend Guide

Help users unfamiliar with frontend development decide what their interface needs, shape its design, and implement it when requested. The user owns goals and consequential choices; the agent organizes and implements the work. Communicate in the user's language using everyday names for interface parts.

## Scope and delivery

Focus on content, module responsibilities, layout relationships, interaction flows, states, and implementation. Reuse established visual direction and design-system decisions; a separate visual-style exploration is not a prerequisite. This skill works independently and does not assign work to other skills.

Match the workflow to the requested outcome. Discussion-only work stays in conversation; a reusable-template task follows the HTML reference without inventing product requirements; a project design uses steps 1–2 and 4; production implementation also uses step 3. Reuse completed steps when refining existing work. Design approval alone does not expand a design-only request into implementation; continue implementation when it is already authorized.

For a project design, provide two corresponding deliverables as the page takes shape:

- **Conversation:** explain the page's contents and their purposes, distinguish confirmed decisions from proposals, and summarize changes and remaining choices.
- **HTML design draft:** save the project's UI and corresponding page explanations in the user-designated folder and keep them aligned with the discussion. Read [HTML preview requirements](references/html-preview.md) when preparing or updating the draft; it defines the delivery structure and how to use the bundled [HTML template](assets/design-draft.html).

Reuse an already designated output folder. If none is established, ask for the destination before saving; continue useful discussion meanwhile. A request for discussion only, text only, no files, or a later preview takes precedence. Do not create an HTML draft for an isolated UI fix merely because this skill exists.

An HTML draft is a review surface, not a change of production stack. For non-Web targets it is an approximation, not proof of native behavior. Once implementing the agreed interface, also deliver the actual project changes and its preview or run instructions.

## Working rules

- Preserve the existing framework, components, toolchain, and design tokens. If no target stack is established, resolve that choice before production implementation; the HTML draft does not settle it.
- Use platform-native concepts: semantic HTML and CSS tokens on Web; widget constraints, `ThemeData`, and `ColorScheme` in Flutter. Reuse existing research; consult official documentation when component behavior or version-sensitive facts remain uncertain.
- **Ordinary delivery:** carry authorized work through implementation and checks without module-by-module approval. Pause only for consequential unresolved choices or scope changes.
- **Guided delivery:** only when explicitly requested, use the module-map and module-preview checkpoints below. Explain one or two relevant concepts when useful; do not attach a lesson to every ordinary delivery.
- Fix the smallest coherent scope that resolves feedback or a defect. A shared component or layout constraint may need adjustment; preserve unrelated confirmed work and check affected uses.

## Workflow

### 1. Help the user choose content when needed

Read the request and existing project/design material first. Reuse known audience, purpose, platform, page scope, required content, and constraints.

If the user supplies only a product type or cannot name the UI they need, do not jump to layout, visual styling, or HTML generation:

1. Explain a small set of likely basic areas in everyday language, including the problem each solves. Relate necessity to the user's use case; a common feature is not automatically required.
2. Offer relevant optional areas and the situations where they help. State assumptions, and ask a focused question if missing context materially changes the options.
3. Let the user select, remove, or add content. Resolve the content scope before designing it. If the user explicitly delegates these choices, state the chosen scope and assumptions and proceed.

For example, for an ebook reader, opening a book and displaying its text may be basic. A bookshelf helps manage many books but may be unnecessary for opening a single file. Notes help close reading; synchronization helps cross-device use. Present these as choices with reasons, not a fixed reader feature list.

If the requested content and functions are already clear, skip this selection step. Do not restart discovery for a local change or turn a missing page parameter into a product requirements interview.

**Ready to continue:** the content scope is selected or explicitly delegated, and consequential unknowns are resolved. Reversible details can use stated defaults.

### 2. Organize the page and prepare its design draft

Map the selected content into the actual interface structure: page regions, navigation, main content, side panels, overlays, and relevant states. A vertically stacked landing page is one possible structure, not the universal model. A feature does not necessarily require its own visible region, page, or code component.

For multiple pages, state each page's responsibility and start with the user's priority. Use a compact module map when helpful: everyday name, purpose, included elements, important interactions/states, layout constraints, and implementation priority. Reuse existing component boundaries where appropriate.

Give concrete, revisable layout defaults with target-platform units and behavior: width constraints, proportions, and adaptation to relevant windows or devices. Do not fix every dimension before seeing the page. Apply the project's visual direction; ask about unresolved choices only when they materially affect the design.

Create or update the HTML design draft and its matching conversational explanation within the authorized scope, following the reference's filling instructions. Keep proposals visibly distinct from user-confirmed choices. The template describes the delivery structure, not the product's features or layout.

**Checkpoint:** in guided delivery, obtain feedback and confirmation of the module map before implementation. In ordinary delivery, proceed from the agreed scope without introducing a new approval gate for every draft.

### 3. Implement and refine when authorized

Implement the essential modules first, using the project's native components, tokens, and established preview workflow. Check each meaningful change in a browser for Web, or through the project's app, emulator, or device workflow for native targets. Repair regressions introduced by the change.

In guided delivery, show each completed module in page context, explain relevant choices, and wait for feedback before the next module. In ordinary delivery, continue through the agreed scope and present the complete page.

### 4. Refine, check, and deliver

Interpret feedback using everyday page/region names, quoted text, and surrounding context. Ask only when the intended target or change remains ambiguous. If feedback changes the content scope or structure, update the affected design and any existing module map. Keep the HTML draft and conversational explanation consistent with accepted changes; during implementation, disclose any difference between the draft and the actual application.

Record durable design decisions at direction changes or delivery milestones in existing project design documentation; use `DESIGN.md` if a new document is needed. Do not require a full design document for a one-off draft or log every spacing adjustment.

Check what applies to the agreed scope, fix issues introduced by this work, and repeat affected checks:

- **Layout:** relevant viewport/device sizes, overflow, long titles, empty content, and module relationships. Inspect the rendered result; do not infer appearance solely from code.
- **Interaction:** the agreed user flows and relevant loading, empty, error, and success states. Identify simulated data, disconnected actions, and missing persistence.
- **Accessibility:** meaningful labels and image alternatives/semantics appropriate to the platform, keyboard navigation and visible focus where applicable, sufficient contrast, and usable touch/click targets.
- **Consistency:** typography, spacing, and components follow project decisions. Check whether differences are intentional instead of treating every variation as a defect.
- **HTML delivery:** apply the reference's shell and placeholder checks. Distinguish demonstrated behavior in the draft from behavior checked in the actual application.

Deliver a clear link or path and opening/run instructions when files are part of the request, useful actions to try when interactions are present, and a concise account of what was delivered, checked, and still limited. Report unavailable checks explicitly. Code inspection alone does not establish visual or interaction quality.

Implementation completion means the agreed scope and relevant checks are complete with gaps disclosed. User acceptance requires actual user feedback; do not invent acceptance or stop authorized ordinary delivery merely because feedback has not arrived.

---
name: aesthetic-translator
description: >-
  This skill helps users describe designs without design terminology. It turns
  everyday visual preferences, rough ideas, and references into concrete design
  directions, an inspectable HTML preview, and implementation-ready specifications.
  Use it for requests such as premium, too plain, or like Linear, and to refine
  those previews from feedback. It does not handle wording-only requests or
  restart design discovery for a precise styling change.
metadata:
  author: FeiNiaoBF
  version: 0.2.0
  date_added: 2026-09-05
---

# Aesthetic Translator

Translate the design a user has in mind into a concrete direction, a visible proposal, and usable specifications. Users need neither professional vocabulary nor a reference image. Reply in the user's language unless they request another language for the deliverable.

The default deliverable is an HTML design preview with a compact inspector: users see the proposed UI, identify its parts, and explain what to change. Include the brief and relevant specifications without crowding the UI. Honor explicit requests for text only, advice without files, or another format; if rendering is unavailable, disclose that limitation.

This skill owns the design proposal and its review preview. A preview does not authorize production implementation, backend work, deployment, or changes to the product's stack. Preserve existing project decisions and the scope of the user's request.

## Working principles

- Start with the user's intended appearance, hierarchy, and behavior. Use inspected references to clarify concrete traits, not to override explicit preferences or replace understanding with a gallery of attractive sites.
- Give meaningful choices, not unexplained parameters. Name the direction, identify a reference or recognizable visual treatment, and explain its practical consequence.
- Reuse a direction the user has already chosen. Ask at most one focused question when competing interpretations materially change the result; do not ask the user to approve every dimension or token.
- No aesthetic is the default. The object, audience, content, and platform determine which directions fit.

## Workflow

1. **Read the context.** Collect the user's mood words, references, audience, object, and constraints from the conversation and existing design documents. Identify the target stack before producing implementation-specific values. Inspect supplied visual material when available; distinguish observed traits from an unverified interpretation of a product name. If a reference cannot be inspected, say so rather than inventing its details.
2. **Identify the relevant dimensions.** Consult the [nine-dimension vocabulary](references/dimensions.md) only where useful. When alternatives would help resolve a material ambiguity, offer two or three coherent directions, each with a name, a concrete visual description, and one consequence. Otherwise recommend a plausible direction and make it visible. Do not require the user to choose style names or work through all nine dimensions before seeing a proposal.
3. **State the Design Read.** Use one sentence: “For <audience>, this <object> uses <visual language>, with <palette/material treatment> and <density/motion/layout character>.” Distinguish explicit requirements, confirmed decisions, and proposed choices. A preview can test an unsettled choice; rendering it does not make it approved. If the user delegates the choice, recommend and proceed within that authorization.
4. **Build the review preview.** Follow the [HTML preview and inspector contract](references/html-preview.md). Show the actual proposed UI as the main content, with a compact inspector for both modules and individual elements. Use realistic task content and keep the represented scope proportional to the request. Embed the four-part brief below and platform-specific specifications as secondary information; a page of prose or reference screenshots alone is not this deliverable.
5. **Calibrate from feedback.** Let users refer to familiar element names or selected regions. Translate comments such as “主标题更突出” into changes to the selected element's size, weight, contrast, or spacing as appropriate. Preserve unaffected decisions and stable element names. Update the preview and its specifications together, explain the relevant change briefly, and continue authorized revisions without repeating discovery or approval gates.
6. **Verify and hand off.** Check the preview using the reference's completion checks and provide its opening link or launch instructions. Once the direction is confirmed, update the project's existing design document when documentation is in scope. If no convention exists, suggest `DESIGN.md`; do not create a competing document. Record confirmed decisions only. Continue already authorized production implementation through the appropriate workflow without asking for the same approval again.

### Four-part brief

Scale the brief to the requested change. Explain how important parameters express the intended feeling; reuse existing tokens where they fit and mark provisional choices. Keep it accessible within the preview rather than presenting a second competing specification.

| Section | Required content |
|---|---|
| Context | Audience, object, purpose, target devices or terminals, existing constraints |
| Direction | Design Read and concrete platform-specific choices for the dimensions that matter |
| References | Supplied or inspected references and the traits being adopted; label illustrative or unverified comparisons |
| Avoid | A short list of mismatches specific to this direction, with reasons |

Map Direction to the actual platform:

- **Web:** CSS custom properties or existing tokens for colors, type, spacing, shape, elevation, and motion; include applicable responsive behavior and interaction states.
- **Flutter/Material 3:** existing `ThemeData`, `ColorScheme`, `TextTheme`, shape, component themes, and states; use logical pixels and the project's constraints rather than CSS declarations.
- **TUI:** semantic color roles, supported palette depth, character-cell spacing, density, focus/selection markers, and border characters. Font rendering belongs to the user's terminal; do not promise app-controlled font families, CSS shadows, or pixel radii. Provide non-color state cues and a fallback for unsupported glyphs.

Keep text readable, preserve visible focus and state distinctions, and respect reduced-motion preferences. Do not equate visual quietness with low text contrast or claim accessibility compliance without checking the actual result.

## Internal calibration

Use these three axes to compare candidate directions, not as a numeric questionnaire for the user. Ranges are starting hypotheses, not platform rules.

| Axis | Scale | Typical starting ranges |
|---|---|---|
| VARIANCE | 1: strict grid; 10: deliberately broken grid | Tools 3–5; marketing 7–9 |
| MOTION | 1: no motion; 10: pronounced physical motion | Productivity 2–4; consumer experiences 6–8 |
| DENSITY | 1: gallery spacing; 10: compact information display | Reading 2–3; terminals 7–9 |

Calm/minimal directions often favor restrained layout and motion. Luxury may use generous spacing; experimental work may use stronger contrast in layout and motion. Trust does not require low density: a professional tool may need compact, predictable structure. Translate the axes into actual choices, such as reading line-height around 1.6, a specific spacing scale, or a 200 ms fade without bounce. Adjust them to content, platform, and accessibility needs; do not leave the result as three numbers.

## When labels are not helping

If a direction is rejected twice, stop proposing cosmetic variants of it. Identify which traits were rejected and change the underlying approach. Users who judge examples more easily than style names should receive concrete comparisons using the design system's own controls: color roles or seed, shape, typography, density, and motion.

Reuse existing research and project previews. Investigate comparable products and official platform guidance only for unresolved decisions or version-sensitive facts; do not launch market research for a local adjustment. State which evidence supports each candidate. If research or rendering is unavailable, offer clearly labeled textual comparisons and disclose the limitation.

Reuse native Flutter previews, emulators, screenshots, or terminal captures when available, with selectable annotated regions in the HTML review surface. A standalone HTML mockup is also appropriate for exploring appearance; label it as an approximation for non-Web targets. It does not verify native layout, interaction, terminal capabilities, or performance. Keep implementation specifications native to the target platform and do not migrate the project just to produce a preview.

With Material 3, Fluent, Carbon, or another established system, map the direction to its native tokens and components rather than imitating it with unrelated styling. If principles conflict, write a concrete product-level rule, such as “brand identity comes from the accent color; surfaces remain undecorated,” and use it consistently. Persist it only once agreed.

## Avoid generic styling by default

Evaluate purple gradients, stock illustrations, shadows everywhere, excessive font families, emoji icons, repeated centered three-card layouts, and decorative glass effects against the actual brief. Do not paste a universal blacklist into every brief: these treatments, including pure black or white, can be appropriate when the reference, design system, or accessibility requirements support them. The Avoid section should explain what would undermine this particular direction.

## References and routing

Read local references only for the relevant branch:

- [HTML preview and inspector contract](references/html-preview.md): read when creating or revising the default visual deliverable; covers element naming, selection, presentation, and completion checks.
- [Nine-dimension vocabulary](references/dimensions.md): candidate terminology and concrete parameters for an unresolved visual dimension.
- [East Asian ink-and-paper palette](references/east-asian.md): an optional direction when the user's references or chosen style support it; never an automatic fallback.

Route by task, without restarting settled decisions:

| Task | Route |
|---|---|
| Express the user's request more clearly, with no design decisions requested | `clarify-wording` |
| Plan or build the production interface beyond the review preview | `frontend-guide`; supply the preview and brief as its direction input |
| Apply a precise local styling change | Stay in the existing implementation workflow; do not restart discovery |
| Author a persistent Design Markdown document from a confirmed brief | `better-designs-md` when its invocation requirements are met; reuse confirmed material |
| Need a broader style, palette, or font catalog | `ui-ux-pro-max` only if already installed and useful; it is optional |

The deliverable is ready for review when users can see the proposed design, identify modules and specific elements, and describe a requested change by name. It must also give usable specifications for the target platform, preserve project constraints, and separate confirmed decisions from proposals. Report missing reference access, rendering, or validation; neither a working preview nor a complete brief establishes user approval or production readiness.

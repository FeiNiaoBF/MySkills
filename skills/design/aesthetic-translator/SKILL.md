---
name: aesthetic-translator
description: >-
  This skill helps when users describe a visual mood with vague preferences such
  as premium, too plain, or like Linear. It turns those cues into named design
  directions and implementation-ready specifications for the project's platform.
  It does not handle wording-only requests or restart design discovery for a
  precise styling change.
metadata:
  author: FeiNiaoBF
  version: 0.1.0
  date_added: 2026-09-05
---

# Aesthetic Translator

Translate aesthetic intent into a short Design Read and a concrete design brief. Use the project's actual platform: Web, Flutter/Material 3, or TUI. Reply in the user's language unless they request another language for the deliverable.

This skill chooses visual direction; it does not own page structure or the implementation workflow. Preserve existing project decisions and the scope of the user's request.

## Working principles

- Treat references as stronger evidence than loose adjectives, while honoring explicit constraints. Extract the relevant traits rather than copying an entire product's appearance.
- Give meaningful choices, not unexplained parameters. Name the direction, identify a reference or recognizable visual treatment, and explain its practical consequence.
- Reuse a direction the user has already chosen. Ask at most one focused question when competing interpretations materially change the result; do not ask the user to approve every dimension or token.
- No aesthetic is the default. The object, audience, content, and platform determine which directions fit.

## Workflow

1. **Read the context.** Collect the user's mood words, references, audience, object, and constraints from the conversation and existing design documents. Identify the target stack before producing implementation-specific values. Inspect supplied visual material when available; distinguish observed traits from an unverified interpretation of a product name. If a reference cannot be inspected, say so rather than inventing its details.
2. **Identify the relevant dimensions.** Consult the [nine-dimension vocabulary](references/dimensions.md) only where useful. For an unresolved direction, offer two or three coherent alternatives, each with a name, a reference or concrete visual description, and one consequence. Do not expand all nine dimensions into separate choice menus. If the direction is settled, proceed directly to its specification.
3. **State the Design Read.** Use one sentence: “For <audience>, this <object> uses <visual language>, with <palette/material treatment> and <density/motion/layout character>.” Present an inferred direction as a proposal, not as user approval. If a material fork remains, resolve it before treating the direction as settled. If the user delegates the choice, recommend and proceed within that authorization.
4. **Deliver the brief.** Use the four sections below, scaled to the requested change. Explain how each important parameter expresses the intended feeling. Preserve existing tokens where they already fit; label provisional values and unresolved choices.
5. **Record and continue.** Once the direction is confirmed, update the project's existing design document when documentation is in scope. If no convention exists, suggest `DESIGN.md`; do not create a competing document. Record confirmed decisions only. Continue already authorized implementation through the appropriate workflow without asking for the same approval again. A request for advice alone does not authorize implementing the design.

### Four-part brief

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

Use existing Flutter previews, emulators, or screenshots for Flutter. Use the project's browser preview for Web; side-by-side device frames are appropriate only when an existing preview harness supports them. Do not introduce a different stack just to show variants. Creating a preview must fit the authorized task scope.

With Material 3, Fluent, Carbon, or another established system, map the direction to its native tokens and components rather than imitating it with unrelated styling. If principles conflict, write a concrete product-level rule, such as “brand identity comes from the accent color; surfaces remain undecorated,” and use it consistently. Persist it only once agreed.

## Avoid generic styling by default

Evaluate purple gradients, stock illustrations, shadows everywhere, excessive font families, emoji icons, repeated centered three-card layouts, and decorative glass effects against the actual brief. Do not paste a universal blacklist into every brief: these treatments, including pure black or white, can be appropriate when the reference, design system, or accessibility requirements support them. The Avoid section should explain what would undermine this particular direction.

## References and routing

Read local references only for the relevant branch:

- [Nine-dimension vocabulary](references/dimensions.md): candidate terminology and concrete parameters for an unresolved visual dimension.
- [East Asian ink-and-paper palette](references/east-asian.md): an optional direction when the user's references or chosen style support it; never an automatic fallback.

Route by task, without restarting settled decisions:

| Task | Route |
|---|---|
| Express the user's request more clearly, with no design decisions requested | `clarify-wording` |
| Plan or build a whole interface | `frontend-guide`; supply the brief as its direction input |
| Apply a precise local styling change | Stay in the existing implementation workflow; do not restart discovery |
| Author a persistent Design Markdown document from a confirmed brief | `better-designs-md` when its invocation requirements are met; reuse confirmed material |
| Need a broader style, palette, or font catalog | `ui-ux-pro-max` only if already installed and useful; it is optional |

The brief is ready for handoff when it identifies a supported direction, gives usable specifications for the target platform, preserves project constraints, and separates confirmed decisions from proposals. Report missing reference access, rendering, or validation rather than treating an untested brief as an approved implementation.

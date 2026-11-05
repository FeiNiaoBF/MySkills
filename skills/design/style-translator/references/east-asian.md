# East Asian ink-and-paper palette

This is one optional visual direction, not the skill's default and not a definition of all East Asian design. Use it when the user's references or chosen direction support ink, paper, calligraphic, or related editorial treatments. Otherwise choose a different direction from the task context.

## Core composition: paper and ink

- Use paper-like surfaces for a light theme and ink-like surfaces for a dark theme. A restrained vermilion accent can unify emphasis without dominating the page.
- Apply vermilion sparingly to a seal, key number, or current-item marker. Keep error and selection states distinguishable; do not rely on the shared hue alone to convey their meaning.
- Warm paper such as `#F7F4EC` and ink such as `#1A1A1A` are palette starting points, not mandatory tokens. Adjust actual foreground/background combinations for readability and the project's theme requirements.

## Common elements

| Element | Treatment | Constraint |
|---|---|---|
| Seal or vermilion ring | A small circular or square mark | Usually one dominant seal per page; do not scatter decorative stamps |
| Ensō ink circle | An incomplete brush-like circle with surrounding space | Provide a contrasting backing or theme-aware stroke so it remains visible in either theme |
| Washi-like texture | Low-opacity grain; SVG `feTurbulence` is one Web option | Keep it behind content and verify that it does not reduce legibility or add excessive rendering cost |
| Mokuroku-inspired contents | Ordered entries, fine separators, and numbering | Borrow the catalog rhythm without making ordinary navigation harder to use |

Select motifs for the actual reference and audience. Ensō, washi, Chinese seal forms, and ink painting come from distinct traditions; do not combine them automatically as interchangeable decoration.

## Vertical text

For a Web accent, `writing-mode: vertical-rl` creates top-to-bottom columns that progress right-to-left. Verify orientation and reading order with the actual text and platform.

Keep ordinary interface body text and navigation horizontal in this palette. Reserve vertical treatment for short titles, quotations, or inscriptions where it remains readable; a project specifically designed for vertical reading needs its own layout decisions.

Chinese numerals such as 一、二、三 may suit decorative numbering. Do not convert dates, quantities, or identifiers mechanically; preserve their meaning and test the chosen numeral treatment.

## Cue translations within this palette

| User cue | Candidate treatment |
|---|---|
| 素雅 / understated elegance | Paper surface, ink text, a small vermilion accent, generous spacing |
| 禅 / Zen-inspired | Minimal composition, spacious rhythm, an ensō only when relevant to the reference |
| 典雅 / classical refinement | Song/Ming-style serif headings, square geometry, thin separators |
| 静谧 / stillness | Dark ink surfaces, subtle surface layering, restrained motion; retain readable text contrast |
| 利落 / sharp, game-like | Dark ink, a vermilion or orange accent, crisp edges |

## Implementation notes

- For Web motion, prefer `transform` and `opacity` where they express the intended effect; verify performance rather than assuming compositing is guaranteed.
- Large animated `clip-path` regions may repaint. Measure on target devices before using them as a prominent effect.
- For a simple ensō draw-on effect, an SVG path with `stroke-dashoffset` is usually easier to control than an animated mask. Verify its geometry and appearance; support a static reduced-motion version.
- Map these visual traits to native painting, themes, and components on Flutter. For TUI, use spacing, color roles, and supported glyphs; paper textures and brush animations need not transfer.
- The Met Open Access collection and Wikimedia Commons can be useful asset sources. Verify the individual asset's license and attribution requirements; Wikimedia hosting alone does not establish public-domain status.

## Research leads when needed

For a chosen motif, consult material about vertical typography, ensō, sumi-e, mokuroku, or washi as relevant. These are research topics, not sources already inspected. Do not require a survey of all of them before producing a brief.

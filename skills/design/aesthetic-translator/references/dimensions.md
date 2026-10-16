# Nine-dimension vocabulary

Use this reference for the dimensions that matter to the request. Offer two or three coherent directions rather than a menu for every row. User phrases are illustrative cues, not deterministic mappings. Product names are comparison anchors, not verified claims about their current designs; inspect the relevant reference before claiming a match.

## 1. Style

Style coordinates the other dimensions. Adapt parameters to the actual platform.

| User cue | Candidate | Reference anchor | Implementation cues |
|---|---|---|---|
| Minimal, clean, content first | Minimalist | Apple, iA Writer | Generous spacing, one accent, little decoration |
| Strict, rational, typographic | Swiss / International | Swiss editorial posters, grid-led product pages | Strong grid, clear type hierarchy, restrained accent |
| Loud, individual, impossible to ignore | Neo-brutalism | Gumroad | Heavy borders, hard offset shadows, flat color |
| Raw, utilitarian, engineering-led | Raw brutalism | Hacker News | System fonts, underlined links, minimal styling |
| Floating, light, spatial | Glassmorphism | Translucent iOS or visionOS surfaces | Translucent panels; protect text contrast and provide an opaque fallback |
| Quiet, luxurious, understated | Dark luxury | CRED | Near-black surfaces, restrained metallic accents, optional serif headings |
| Magazine-like, narrative | Editorial | NYT, Stripe Press | Large serif headings, columns, pull quotes |
| Rounded, cheerful | Playful | Duolingo | Rounded shapes, bright colors, substantial buttons |
| Technical, command-oriented | Terminal / TUI | Developer CLIs | Monospace rhythm, semantic highlights, character borders |
| Native Google application | Material 3 | Android Material components | Native tokens and components rather than imitation |
| Ink, paper, East Asian references | Ink-and-paper | The optional palette linked from SKILL.md | Paper/ink contrast and a restrained accent |

## 2. Color mood

| User cue | Candidate | Implementation cues |
|---|---|---|
| Gentle, soothing | Pastel | Low-saturation surfaces; readable foregrounds |
| Natural, grounded | Earthy | Terracotta, sand, olive |
| Rich, confident | Jewel tones | Deep gemstone hues |
| Mature, understated | Muted / dusty | Desaturated colors with gray undertones |
| Restrained, focused | Grayscale plus one accent | Neutral surfaces; accent for selected emphasis |
| Paper and ink | Cream and ink | Warm off-white background with dark text |
| Energetic, young | Neon on dark | Bright accents on dark surfaces; limit glare |

Useful controls include a 60/30/10 balance, accent allocation, warm/cool bias, colored shadows, and limited gradient placement. These are compositional options, not quotas. Off-white and charcoal suit a softer direction; pure white or black may suit the existing system or contrast needs. Keep semantic states distinguishable from decorative accents.

## 3. Typography

| User cue | Candidate | Example families | Character |
|---|---|---|---|
| Neutral, unobtrusive | Grotesque sans | Inter, Helvetica | Product-oriented clarity |
| Friendly, geometric | Geometric sans | Poppins, Futura | Rounded, optimistic |
| Modern, technical | Technical display | Space Grotesk | Distinctive technical headings |
| Literary, established | Old-style serif | Garamond, Georgia | Book-like reading |
| Fashion, luxury | High-contrast serif | Playfair Display, Didot | Dramatic thick/thin contrast |
| Solid, grounded | Slab serif | Roboto Slab | Strong, block-like serifs |
| Precise, code-oriented | Monospace | JetBrains Mono | Uniform character advances |

Controls include heading scale, heading tracking, body line-height, spaced uppercase labels, serif/sans pairing, a single family with weight hierarchy, and monospace for numbers or code. Check language coverage, fallback, and availability before specifying a family; Latin examples do not provide Chinese or Japanese glyph coverage. Terminal applications generally inherit the terminal font.

## 4. Light and dark themes

Choose for the usage conditions and existing theme support. Light or warm-paper themes can suit sustained reading; dark or dim themes can suit media, technical tools, or low-light contexts. High contrast is a measurable requirement, not a synonym for dark mode. Define legible foregrounds and distinguishable states in each required theme.

## 5. Layout character

Candidate patterns include a constrained reading column (about 720 CSS px as a Web starting point), split text/image landing pages, sidebar tool shells, equal-width catalog cards, unequal editorial columns, full-width typographic heroes, mixed-size feature tiles, and floating capsule navigation. Choose a character that serves content and tasks. Use enough page and module structure to make the design proposal reviewable; production page inventory belongs to the implementation workflow. Do not impose a new structure for a local style adjustment.

## 6. Detail

These values are Web examples; map the intent to native tokens or character cells on other platforms.

- **Corners:** square for sharp structure; 6–8 px for restrained rounding; around 16 px for softer shapes; 24 px or more for pronounced rounding. Component size changes the effect.
- **Shadows:** none for flat surfaces; diffuse low-opacity shadows for soft elevation; hard offset shadows for graphic emphasis; layered shadows for depth.
- **Cards:** hairline outline, soft elevation, translucent surface, or gradient border where the chosen direction warrants it.
- **Buttons:** pill, square, outlined, or a hard lower edge with pressed-state movement. Preserve recognizable affordances and state differences.
- **Finishing:** visible focus, thin separators instead of redundant boxes, one consistent icon family, and restrained inner-edge highlights on dark surfaces.

## 7. Motion

Specify duration, easing, trigger, and reduced-motion behavior rather than saying “smooth.” A 200 ms fade is one restrained Web starting point, not a universal default. Small hover translations require a hover-capable device. Springs can support playful feedback; scroll reveals can support narrative pages, but essential content must remain available without animation. Reduce motion when it does not serve the task.

## 8. Density

- **Airy:** useful for brand or editorial emphasis; costs scrolling and reduces visible information.
- **Balanced:** a starting point for general product interfaces.
- **Compact:** useful for terminals and information-heavy tools; preserve readability, focus visibility, and target usability.

This dimension corresponds to the internal DENSITY axis in SKILL.md. Trustworthy or calm interfaces can still be compact.

## 9. Imagery and voice

Choose a coherent primary image language: 3D renders, flat vectors, thin-line icons, solid icons, documentary photography, abstract gradients, restrained grain, hand-drawn marks, or typography without imagery. Combine treatments only when their roles are clear. Reference actual available assets and their usage rights rather than inventing access to them.

Copy contributes to the same direction: warm (“Oops!”) and precise (“Request failed. Retry.”) suggest different products. Match the user's audience and language; do not let a playful tone obscure an error or the recovery action.

# Google DESIGN.md format reference

Use this reference only when the user requests Google's DESIGN.md format or a target tool requires it. Ordinary project design documentation follows the scope and format choices in the [skill workflow](../SKILL.md#choose-document-scope-and-format); a file named `DESIGN.md` does not by itself select this schema.

This reference summarizes Google's upstream **alpha** format, checked on **2026-10-18**, rather than defining a new schema. Within that scope, existing project conventions and a specified consumer/version determine compatibility requirements; preserve deliberate deviations and report them instead of silently migrating a document.

## Sources and authority

- [Official format specification](https://github.com/google-labs-code/design.md/blob/main/docs/spec.md): YAML placement, schema, section order, token references, and consumer behavior.
- [Official philosophy](https://github.com/google-labs-code/design.md/blob/main/PHILOSOPHY.md): design intent and explanatory prose are central; tokens support that context.
- [Stitch format documentation](https://stitch.withgoogle.com/docs/design-md/format): product documentation entry. Its page did not yield readable content during this review; the details below were verified against the official repository.

The specification describes token values as normative and prose as application context. The philosophy emphasizes that tokens are not a rendering language. Keep exact definitions consistent while explaining how to use them; do not promise identical rendering across tools. The alpha format is evolving, especially component definitions.

The workflow's evidence labels, focused questions, source-of-truth handling, and approval boundaries are **this skill's authoring policy**, not extra Google schema requirements. Do not present them as upstream mandates.

## YAML placement and metadata

Tokens may appear in YAML front matter, fenced `yaml` blocks in relevant sections, or both. Front matter starts and ends with a line containing exactly `---`. All token blocks merge into one tree: each top-level key may occur **at most once in the entire document**, even if repeated blocks contain different nested tokens. Keep YAML examples that are not part of the actual token tree in a `text` or `markdown` fence.

The schema lists `name` and optional `version`, `description`, and `omitted`, alongside token groups. The current format version is `alpha`; it is distinct from this skill's version or a project's release number. `title`, `status`, `author`, and `reviewers` are not required schema metadata. Keep such project metadata in prose or a consumer-supported extension when useful; infer neither identity nor approval.

Quote hex colors so YAML does not treat `#` as a comment. Quote references such as `"{colors.primary}"` so they remain strings. Parse the blocks together with duplicate-key detection, not a loader that silently overwrites definitions.

## Token groups and references

| Group | Recorded schema | Authoring guidance |
|---|---|---|
| `colors` | Token name to CSS color string | Describe roles, text/background pairings, and theme applicability. Hex is the recommended interoperable default; the schema also accepts other CSS color forms. |
| `typography` | Token name to typography properties | Use `fontFamily`, `fontSize`, numeric `fontWeight`, `lineHeight`, `letterSpacing`, and relevant `fontFeature`/`fontVariation`. Preserve actual hierarchy rather than manufacturing a large scale. |
| `spacing` | Scale name to dimension or number | Explain whether values represent distances, counts, or ratios. Unitless values do not automatically mean pixels. |
| `rounded` | Scale name to dimension | Record the actual shape language; a radius scale is not mandatory for every design. |
| `components` | Component name to named string values or token references | Define relevant variants under related keys, such as `button-primary` and `button-primary-hover`, without assuming deep nesting is portable. |

The recorded `Dimension` type supports `px`, `em`, and `rem`; typography `lineHeight` also accepts unitless numbers. Native units or additional groups such as motion/elevation need prose or a documented consumer extension. Do not claim they belong to the core dimension schema.

References use `{path.to.token}`. Most groups reference primitive values; components may also reference composite typography values. Resolve every reference in the combined tree and reject missing paths or cycles. Typical component properties include `backgroundColor`, `textColor`, `typography`, `rounded`, `padding`, `size`, `height`, and `width`; upstream consumers accept unknown component properties with a warning.

For syntax illustration only, not a project palette or button requirement:

```yaml
name: Example
version: alpha
colors:
  primary: "#234E70"
  on-primary: "#FFFFFF"
components:
  button-primary:
    backgroundColor: "{colors.primary}"
    textColor: "{colors.on-primary}"
```

For new color sections following this baseline, the specification calls for a `primary` palette; additional palettes and most token names are conventions, not compulsory inventory. If an existing design uses different semantics, preserve it and document the compatibility gap or an agreed mapping. Never rename a production token just to match an example. The philosophy and examples are more permissive than some schema guidance; check the actual target consumer when that difference matters.

## Prose sections

Use `##` for sections; an optional `#` title is allowed. Relevant standard sections appear in this order. The format permits irrelevant sections to be omitted.

| Order | Section (accepted alias) | What makes it useful |
|---|---|---|
| 1 | Overview (Brand & Style) | Audience, product purpose, intended feeling, concrete visual character, and constraints. |
| 2 | Colors | Palette roles, pairings, themes when present, and where accents belong. |
| 3 | Typography | Family roles, hierarchy, reading needs, and the relationships behind the values. |
| 4 | Layout (Layout & Spacing) | Grid or alternative layout model, spacing rhythm, containment, and relevant adaptation rules. |
| 5 | Elevation & Depth (Elevation) | How hierarchy is conveyed through shadows, tonal layers, borders, or other deliberate treatments. |
| 6 | Shapes | Corner treatment and other shape decisions, with `rounded` tokens where useful. |
| 7 | Components | Relevant atoms, variants, states, and behavior-related visual guidance. |
| 8 | Do's and Don'ts | Short, contextual rules that prevent plausible deviations from this design. |

An `Accessibility Considerations`, `Motion`, or other project-specific section is an extension, not a mandatory ninth standard section. Keep relevant accessibility requirements beside the affected colors/components or in an existing cross-cutting section. Do not duplicate the same rule in multiple locations.

Unknown section headings are preserved by the recorded format. Duplicate section headings and duplicate top-level YAML keys are errors. Additional project content need not be discarded to make the document look like a fixed template.

`omitted` can document intentionally absent groups and suppress related linter warnings. It accepts section names or objects containing `section` and optional `reason`. Use it for intentional scope decisions, not to conceal missing research or unresolved design choices.

## Validation in context

Use the project's existing parser or consumer when available and relevant. For a standalone draft, parse all token blocks, inspect references and types, and check section order and prose/token agreement. A YAML parser alone does not validate the full DESIGN.md schema. Do not install tools or introduce a publishing/integration step merely to validate a documentation edit.

When adapting an existing document, verify the changed definitions and their dependents while preserving unrelated extensions. Report pre-existing incompatibilities separately. For multiple themes, keep each top-level group defined once and use a naming/mapping convention supported by the project and target consumer; do not duplicate `colors` blocks for light and dark modes.

Record actual checks and their limits. Contrast measurements require defined foreground/background pairs and applicable states; screenshot estimates and untested requirements are not measured results. Consumer parsing, visual review, user approval, and accessibility evaluation establish different things.

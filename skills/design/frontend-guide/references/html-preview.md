# HTML Preview Requirements

Read this when preparing or updating the HTML design draft. Start with the bundled [HTML template](../assets/design-draft.html), which defines a replaceable delivery shell rather than a sample application.

## Delivery structure

- **Header:** show only a simple name such as `{{PROJECT_NAME}}-frontend-guide`. Do not turn it into a toolbar of review modes, viewport switches, or metadata.
- **Main area:** make the actual project's UI the dominant content. The shell does not prescribe the product's navigation, modules, visual style, or feature set.
- **Explanations:** place directly readable explanations on the right, ordered by the page's structure or task flow. On narrow screens, move them below the design without obscuring it. They do not depend on clicking or hovering over the UI.
- Explain page purpose, meaningful regions and their arrangement, relevant flows/states, layout adaptation, and unresolved choices as needed. Repeat or remove explanation sections to fit the project; not every page needs every heading.
- Keep product controls usable directly. Do not add module/element selection, an object tree, copied object IDs, inspection modes, or a selection-driven inspector to the delivery shell.

The template uses a neutral gray surround, a white design surface, and a numbered explanation column. Its typography and spacing distinguish headings, readable guidance, and secondary replacement markers. The `--fg-*` CSS properties control shell colors, fonts, and proportions; section numbers follow document order automatically. These are adjustable starting points. Reuse the project's established visual direction for its UI; the shell is not the product design system.

## Fill the template

The `{{...}}` labels are literal replacement markers, not a template engine. Copy the file to the agreed output folder and edit it; no build step or JavaScript dependency is required for the shell.

- `PROJECT_NAME`: replace both the document title and visible header name.
- `PROJECT_UI`: replace the entire placeholder block with the project's page design, including its own layout and any agreed interactions. Remove the placeholder-only styling as appropriate.
- `PAGE_PURPOSE`: describe the audience, task, and scope of the page.
- `REGION_NAME` and `REGION_EXPLANATION`: name a meaningful page region and explain its contents, purpose, and arrangement. Repeat this section in a useful reading or task order; a region need not equal a code component.
- `FLOWS_AND_STATES`: describe relevant actions and feedback, identifying simulated or unfinished behavior.
- `LAYOUT_AND_ADAPTATION`: explain the relationships between regions and how they adapt to relevant sizes.
- `DECISIONS_AND_OPEN_QUESTIONS`: distinguish confirmed choices, current proposals, and remaining decisions using actual conversation/project evidence.

Replace the filling instructions with project-specific explanations, remove irrelevant sections, and localize visible text and the document language to the user's language. For finished explanations, put prose in ordinary paragraphs and remove the corresponding marker spans; `.template-token` styles identify authoring placeholders, not project content. Keep shell styling scoped so it does not unintentionally restyle the embedded product UI.

**Template-only request:** retain replaceable labels and guidance showing what belongs in each slot. Do not fill them with an invented reader, dashboard, or other sample product. A reusable template is complete when its structure and filling expectations are clear; its empty UI slot is intentional.

**Project-design request:** deliver the actual project's UI and corresponding explanations. Empty slots and authoring instructions do not satisfy this request. Use representative content and enough interaction to review the agreed flows; disclose mocks and unfinished states.

## Storage, feedback, and checks

Save in the user-designated folder with a clear entry file and opening instructions. Keep required assets alongside it using portable relative links. Update the existing draft and matching conversational explanation when feedback changes the design.

Interpret everyday region names, quoted text, and positional descriptions together with the requested change. Clarify the target only when ambiguous; feedback does not require an object-selection system.

Before delivery, inspect the rendered layout at relevant widths: the name remains readable, the UI is primary, explanations are directly available in order, and neither column causes horizontal overflow or hides content. For project designs, also check the implemented product interactions, keyboard/touch access, and relevant states. For template-only delivery, check the replacement markers and shell instead of claiming application interactions were tested. Report unavailable checks explicitly.

The HTML supports design review; it does not itself prove production functionality, native-platform fidelity, or user acceptance.

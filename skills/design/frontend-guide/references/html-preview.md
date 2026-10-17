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

The `{{...}}` labels are literal replacement markers, not a template engine. For a new draft, copy the file to the agreed output folder and edit that copy; preserve and update an existing draft when continuing work. Do not edit the bundled asset for an ordinary project delivery. The shell requires no build step or JavaScript dependency; the embedded UI may need its own assets or scripts.

- `PROJECT_NAME`: replace both the document title and visible header name.
- `PROJECT_UI`: replace the entire `.template-placeholder` element inside `.template-ui-slot` with the project's page design, including its own layout and any agreed interactions. Remove the placeholder-only styling as appropriate.
- `PAGE_PURPOSE`: describe the audience, task, and scope of the page.
- `REGION_NAME` and `REGION_EXPLANATION`: name a meaningful page region and explain its contents, purpose, and arrangement. Repeat this section in a useful reading or task order; a region need not equal a code component.
- `FLOWS_AND_STATES`: describe relevant actions and feedback, identifying simulated or unfinished behavior.
- `LAYOUT_AND_ADAPTATION`: explain the relationships between regions and how they adapt to relevant sizes.
- `DECISIONS_AND_OPEN_QUESTIONS`: distinguish confirmed choices, current proposals, and remaining decisions using actual conversation/project evidence.

Replace the filling instructions with project-specific explanations, remove irrelevant sections, and localize visible text, accessible labels, and the document language to the user's language. Put finished explanations in ordinary paragraphs and remove their `.template-token` spans; replace `REGION_NAME` in its heading. These markers identify authoring slots, not containers for the finished content.

Keep product and shell styles scoped separately. The product UI sits inside `.template-ui-slot`, whose available width differs from the browser viewport while the notes are beside it. Check product layout within that space; viewport-wide sizing, fixed overlays, and viewport breakpoints may otherwise overlap the notes or misrepresent adaptation. Adjust the draft's containment as needed and disclose remaining differences from the target application.

**Template-only request:** retain replaceable labels and guidance showing what belongs in each slot. Do not fill them with an invented reader, dashboard, or other sample product. A reusable template is complete when its structure and filling expectations are clear; its empty UI slot is intentional.

**Project-design request:** deliver the actual project's UI and corresponding explanations. Empty slots and authoring instructions do not satisfy this request. Use representative content and enough interaction to review the agreed flows; disclose mocks and unfinished states.

## Storage and checks

Save in the user-designated folder with a clear entry file and opening instructions. Keep required assets alongside it using portable relative links.

Before delivery, inspect the rendered layout at relevant widths: the name remains readable, the UI is primary, explanations are directly available in order, and neither column causes horizontal overflow or hides content. Confirm the shell adds no review controls or selection-driven behavior.

- **Project design:** remove unresolved template markers and filling guidance, and ensure the explanations match the displayed UI and demonstrated states. Apply the main workflow's interaction and accessibility checks to what the draft actually implements.
- **Reusable template:** preserve meaningful replacement markers, confirm that headings and guidance identify each slot, and check the empty shell's layout. Do not claim application interactions were tested.

# HTML Preview Requirements

Read this when preparing or updating the HTML design draft. Start with the bundled [HTML template](../assets/design-draft.html), which defines a replaceable delivery shell rather than a sample application.

## Delivery structure

- **Header:** show only one simple identity line such as `{{PROJECT_NAME}}-frontend-guide`. The name and a short type descriptor may use separate elements when they still read as that single line. Do not turn the header into a toolbar of review modes, viewport switches, or metadata.
- **Main area:** make the actual project's UI the dominant content. The shell does not prescribe the product's navigation, modules, visual style, or feature set.
- **Explanations:** place directly readable explanations on the right, ordered by the page's structure or task flow. On narrow screens, move them below the design without obscuring it. They do not depend on clicking or hovering over the UI.
- Explain page purpose, meaningful regions and their arrangement, relevant flows/states, layout adaptation, and unresolved choices as needed. Repeat or remove explanation sections to fit the project; not every page needs every heading.
- **Implementation handoff:** place a clearly titled, directly copyable Agent implementation prompt below the preview and explanations. Keep it subordinate to the actual UI but easy to reach. The prompt and HTML are one delivery: it must describe the displayed design, not a separate imagined interface.
- The prompt covers implementation scope, component responsibilities, layout and interaction, key states, target platform and technical constraints, non-goals, and acceptance criteria. Mark user-confirmed decisions, Agent recommendations, and open questions separately; open items are not requirements or acceptance criteria.
- Keep product controls usable directly. Do not add module/element selection, an object tree, copied object IDs, inspection modes, or a selection-driven inspector to the delivery shell. A handoff prompt and a simple copyable code block are delivery content, not review controls.

The template uses a neutral gray surround, a white design surface, a numbered explanation column, and a full-width implementation handoff below them. Its typography and spacing distinguish headings, readable guidance, secondary replacement markers, and copyable prompt text. The `--fg-*` CSS properties control shell colors, fonts, and proportions; section numbers follow document order automatically. These are adjustable starting points. Reuse the project's established visual direction for its UI; the shell is not the product design system.

## Fill the template

The `{{...}}` labels are literal replacement markers, not a template engine. For a new draft, copy the file to the agreed output folder and edit that copy; preserve and update an existing draft when continuing work. Do not edit the bundled asset for an ordinary project delivery. The shell requires no build step or JavaScript dependency; the embedded UI may need its own assets or scripts.

- `PROJECT_NAME`: replace both the document title and visible header name.
- `PROJECT_UI`: replace the entire `.template-placeholder` element inside `.template-ui-slot` with the project's page design, including its own layout and any agreed interactions. Remove the placeholder-only styling as appropriate.
- `PAGE_PURPOSE`: describe the audience, task, and scope of the page.
- `REGION_NAME` and `REGION_EXPLANATION`: name a meaningful page region and explain its contents, purpose, and arrangement. Repeat this section in a useful reading or task order; a region need not equal a code component.
- `FLOWS_AND_STATES`: describe relevant actions and feedback, identifying simulated or unfinished behavior.
- `LAYOUT_AND_ADAPTATION`: explain the relationships between regions and how they adapt to relevant sizes.
- `DECISIONS_AND_OPEN_QUESTIONS`: distinguish confirmed choices, current proposals, and remaining decisions using actual conversation/project evidence.
- `TARGET_STACK_AND_PLATFORM`: state the actual application platform and technology when known; otherwise mark the choice open. Do not infer Web implementation from the HTML preview.
- `CONFIRMED_DECISIONS`, `AGENT_RECOMMENDATIONS`, and `OPEN_QUESTIONS`: keep requirements, advice, and unresolved choices separate. Recommendations need user acceptance before becoming requirements; open questions must not appear in acceptance criteria.
- `IMPLEMENTATION_SCOPE`: state what the implementation Agent should build in this handoff and what existing surface it changes.
- `COMPONENT_RESPONSIBILITIES`: map visible regions and controls to responsibilities without forcing every visual region into a code component.
- `LAYOUT_AND_INTERACTIONS`: describe relationships, sizing/adaptation, user actions, and resulting feedback.
- `KEY_STATES`: cover the relevant loading, empty, error, success, selected, disabled, or permission states evidenced by the design.
- `TECHNICAL_CONSTRAINTS`: preserve the existing stack, native platform behavior, design tokens, accessibility requirements, data boundaries, and integration constraints.
- `NON_GOALS`: prevent adjacent features or redesign work from entering the implementation by implication.
- `ACCEPTANCE_CRITERIA`: provide observable checks that follow from confirmed decisions and the displayed design.

Replace the filling instructions with project-specific explanations and handoff content, remove irrelevant sections, and localize visible text, accessible labels, and the document language to the user's language. Put finished explanations in ordinary paragraphs and remove their `.template-token` spans; replace `REGION_NAME` in its heading. In the prompt, remove replacement markers that have confirmed values and retain any genuinely unresolved item as an explicit open question rather than silently deciding it.

Keep product and shell styles scoped separately. The product UI sits inside `.template-ui-slot`, whose available width differs from the browser viewport while the notes are beside it. Check product layout within that space; viewport-wide sizing, fixed overlays, and viewport breakpoints may otherwise overlap the notes or misrepresent adaptation. Adjust the draft's containment as needed and disclose remaining differences from the target application.

**Template-only request:** retain replaceable labels and guidance showing what belongs in each slot, including the implementation-prompt structure. Do not fill them with an invented reader, dashboard, stack, or other sample product. A reusable template is complete when its structure and filling expectations are clear; its empty UI slot is intentional.

**Project-design request:** deliver the actual project's UI, corresponding explanations, and a filled implementation prompt. Empty slots, authoring instructions, or a prompt that merely says “match the design” do not satisfy this request. Use representative content and enough interaction to review the agreed flows; disclose mocks and unfinished states. For native targets, state that the HTML demonstrates design intent and does not establish native rendering or behavior.

## Storage and checks

Save in the user-designated folder with a clear entry file and opening instructions. Keep required assets alongside it using portable relative links.

Before delivery, inspect the rendered layout at relevant widths: the name remains readable, the UI is primary, explanations are directly available in order, the implementation handoff is easy to find and select, and no region causes horizontal overflow or hides content. Confirm the shell adds no review controls or selection-driven behavior.

Check the implementation prompt against the UI and explanations section by section. A coding Agent should be able to identify the implementation scope, components, layout and interactions, key states, technical constraints, non-goals, acceptance criteria, and every unresolved decision without guessing. Copy the complete prompt from the rendered page and confirm its text remains intact. For a non-Web target, verify that both the nearby guidance and prompt name the real platform and do not imply that the HTML itself is production code.

- **Project design:** remove every replacement marker and filling instruction, write any unresolved item as an explicit plain-language open question, ensure the explanations match the displayed UI and demonstrated states, and fill every required handoff field. Apply the main workflow's interaction and accessibility checks to what the draft actually implements.
- **Reusable template:** preserve meaningful replacement markers, confirm that headings and guidance identify each UI, explanation, and handoff slot, and check the empty shell's layout. Do not claim application interactions or implementation readiness were tested.

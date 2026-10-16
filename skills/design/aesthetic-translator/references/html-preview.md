# HTML preview and inspector

Read this when creating or revising the skill's default visual deliverable. The preview makes the proposed design visible and gives the user a shared vocabulary for requesting changes. This is a delivery contract, not a fixed visual template.

## Presentation

- Make the proposed UI the main content. On wide screens, place a compact inspector on the right and give the UI most of the available width. Keep outer controls minimal; do not put a large explanatory hero above the preview.
- Render the design for the user's actual task. Do not substitute a reference gallery, a generic sample application, or an essay about the design. Show only the pages, states, or variants needed to assess the current proposal.
- Keep the inspector visually quiet and distinct from the proposed product. Its typography, colors, and panels must not impose a style on the UI being reviewed.
- Show a concise module hierarchy; reveal relevant child elements on selection or expansion. Display one selected item's explanation at a time instead of spreading labels and leader lines across the whole UI.
- Keep the brief, overall tokens, assumptions, and validation notes in a secondary or collapsible area. Provide an obvious way to clear highlighting and view the unobstructed design.
- On narrow screens, move the inspector below the preview or into an accessible drawer. Keep selection and its explanation reachable without horizontal page overflow. Do not shrink the proposed UI until its text is unreadable.

## Names and selection targets

Name parts in the user's language using familiar terms. Technical identifiers can be secondary; the user should not need DOM, CSS, or framework vocabulary to request a change.

| Level | Example user-facing names | Purpose |
|---|---|---|
| Module | 页头、首屏、作品区、底部导航 | Discuss a whole region's composition or spacing |
| Element | 主标题、副标题、说明文字、封面图片、主要按钮 | Identify the specific item to change |
| Repeated item | 作品区 / 第一张卡片 / 封面图片 | Disambiguate similar items by context |

Use names that match the actual design; do not invent a subtitle or button merely to fill this list. Give meaningful visible elements their own targets, including important text, images, and controls. Do not expose every layout wrapper, decorative fragment, or DOM node.

Maintain stable internal IDs, user-facing names, and parent relationships across revisions. If a part is added, removed, or deliberately renamed, update its entry and explanation. Keep selection metadata in one registry or equivalent source so the rendered targets, inspector, and descriptions cannot drift apart.

## Inspector contents

For the current selection, show:

- **Name and location:** a familiar name and, where necessary, its module or breadcrumb.
- **Role:** one short explanation of what this part does for the user.
- **Design reason:** why its treatment supports the intended effect, rather than a generic aesthetic compliment.
- **Relevant specifications:** only values useful for this item, such as title size/weight/line-height, text-to-title gap, image ratio/crop, or button emphasis and state. Put additional detail behind an expansion control.

Read actual rendered values or shared tokens where practical. Mark proposed values and approximations accurately, including non-Web units. Do not label an estimated measurement as an inspected result. Keep the brief and inspector consistent with the rendered design.

For example, a selected “首屏 / 主标题” can explain the page's main message, how its size and weight establish emphasis, and the current typography values. This lets the user say “主标题再突出一点” without knowing the token names.

## Interaction

- Hover previews the meaningful element under the pointer; clicking pins the selection. On touch devices, tapping selects. Selecting a nested element must not be overridden by its parent; keep the parent accessible through the inspector hierarchy or breadcrumb.
- Selecting an inspector entry highlights the matching region and brings it into the preview's visible area. Highlight with a thin outline or subtle overlay that does not change layout, hide content, or permanently recolor the proposed design.
- A small temporary label may identify the active target. Avoid persistent labels over every element, crossing connector lines, or overlays that capture the pointer and make selection flicker.
- Keep a pinned selection when the pointer leaves a transient hover preview. Support clearing it with Escape and a visible control. Maintain focus visibility and provide keyboard access through the inspector; hover must not be the only route to information.
- If inspection would trigger a product action, use an explicit inspection mode that selects without submitting, navigating, or mutating data. Preserve the represented control's semantics; do not wrap interactive controls in other interactive controls. Provide a preview mode when real product interactions need review.

## Feedback and revision

Resolve feedback against the selected element or its familiar name and module. If multiple targets remain plausible and choosing incorrectly would change the result, ask one focused question. Otherwise make the smallest coherent revision and preserve unrelated decisions.

Translate the requested effect into concrete design adjustments, then update the UI and its explanations/specifications together. For “副标题离主标题近一点,” adjust their gap; do not silently redesign the whole hero. Keep unconfirmed treatments marked as proposals. Users can continue to respond in ordinary language; a specification questionnaire is not required.

## Completion checks

- Open the artifact in an available browser. Inspect a wide layout and a narrow layout relevant to the target. Check UI prominence, text readability, inspector overflow, and whether highlights obscure the selected content.
- Exercise a module and a nested element, two similarly named repeated elements when present, selection from the inspector, hover restoration, clearing, and keyboard/touch alternatives. Confirm that inspection does not accidentally activate product actions.
- For a sample revision, check that the selected target, displayed name, explanation, and affected specifications still refer to the same rendered element. Check relevant responsive and interaction states rather than claiming all states are implemented.
- Use available browser automation for concrete interaction risks; inspect rendered screenshots for composition. Report which checks ran and any unavailable rendering or native validation. Source review alone is not visual verification.
- Deliver the actual artifact link or runnable preview instructions, with a short explanation of how to select and refer to an element. Prefer a self-contained HTML file for standalone review; reuse an existing project preview when that better preserves fidelity. Keep temporary artifacts outside the repository unless the user requests a persistent project deliverable. Do not claim user acceptance before receiving it.

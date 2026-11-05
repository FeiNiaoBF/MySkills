# HTML delivery reference

Read this when creating, maintaining, or revising the visual deliverable. The bundled [HTML template](../assets/design-preview.html) defines the review shell and its behavior. It is an empty reusable template, not a sample application or a product design system.

## Delivery structure

- **Compact toolbar:** show the project name and review actions. Do not put an explanatory hero or a full-width direction brief above the UI.
- **Main canvas:** make the project's actual UI the dominant content. The shell does not prescribe its features, layout, typography, colors, components, or platform.
- **Direction summary:** keep the direction name, decision status, and Design Read above the notes tabs so they remain available while switching explanations.
- **Element tab:** show one selected item's name, location, role, design reason, and relevant specifications. A collapsible object list provides another selection route.
- **Brief tab:** hold the four-part brief defined in SKILL.md. Keep overall design reasoning here rather than repeating it in every element.
- **Specifications tab:** hold shared platform specifications, assumptions, and actual validation limits. Keep item-specific values with their elements.
- **Narrow screens:** place notes after the canvas. Selection offers a jump to the explanation when it is below the viewport; the explanation provides a return to the UI. Collapsing notes also hides the direction summary and switches to experience mode.

The template's neutral surround, restrained separators, spacing, and `--at-*` tokens belong to the review shell. Its dimensions and colors are adjustable defaults, not rules for the project. Keep product and shell styles scoped separately; this is a shared document, not an iframe or Shadow DOM isolation boundary. The sidebar reduces canvas width: check product layout within that space, especially viewport breakpoints, full-width units, fixed overlays, and scrolling. Disclose differences from the target application.

## Copy and fill

For a new preview, copy the asset into the agreed output folder; preserve and update an existing preview when continuing work. Do not edit the bundled source during an ordinary project task. The shell is self-contained HTML/CSS/JavaScript and needs no build step; embedded project content may have its own assets or dependencies. Keep required assets beside the entry file with portable relative links and provide opening instructions.

`data-at-field` attributes identify authoring locations, not selectable product elements or a template engine. Fill them as follows:

| Field | Project content |
|---|---|
| `PROJECT_NAME` | Project name in the toolbar; also update the document title |
| `PROJECT_UI` | Replace the entire `.at-placeholder` inside `#at-ui` with the actual UI proposal |
| `DIRECTION_NAME` | A concise, recognizable name for the current visual direction |
| `DECISION_STATUS` | Status supported by actual feedback, such as proposed or confirmed; identify mixed status in the notes |
| `DESIGN_READ` | The main workflow's one-sentence interpretation of the visual intent |
| `CONTEXT` | Audience, purpose, target platform/devices, and existing constraints |
| `DIRECTION_DETAILS` | Relevant visual choices and how they express the direction |
| `REFERENCES` | Inspected sources and adopted traits; label illustrative or unverified comparisons |
| `AVOID` | Treatments that would undermine this direction, with reasons |
| `PLATFORM_SPECS` | Shared tokens or platform-native specifications useful for implementation |
| `VALIDATION_NOTES` | Assumptions, unresolved choices, actual checks, and remaining limits |

Replace filling prompts with real explanations. Fields may contain paragraphs, lists, links, or tables as appropriate; replace the containing `p` when adding block content. Localize visible text, accessible labels, document language, and runtime-generated messages. Remove irrelevant guidance and placeholder-only styles from project copies. Preserve IDs and relationships used by the shell runtime.

**Template-only maintenance:** retain field markers, useful filling instructions, the empty UI slot, and an empty registry. Do not invent a reader, dashboard, or other product to make the template look complete. Synthetic interaction examples belong only in isolated validation fixtures.

**Project-design delivery:** show the actual project with representative content and meaningful element explanations. An empty template does not satisfy this request. Label simulated states, disconnected actions, and non-Web approximations. Reuse an existing native or project preview when it better preserves fidelity; adapt the review surface as needed without changing the production stack.

## Associate elements with explanations

Use familiar names in the user's language. Give meaningful regions and important text, images, and controls their own targets; do not expose every wrapper or decorative fragment. Repeated items need contextual names, such as “作品区 / 第一张卡片 / 封面图片.” Do not add UI elements merely to fill an example hierarchy.

In a template copy:

1. Add a unique `data-at-id` to each review target inside `#at-ui`.
2. Fill the JSON array in `#at-registry`. Each record has `id`, `name`, `role`, `reason`, optional `parent`, and optional `specs`.
3. Match each `id` to exactly one rendered target. A `parent` names another record; keep relationships acyclic. Preserve IDs and familiar names across revisions, updating removed or renamed items together with the UI.
4. Reload after changing the registry or target set: the bundled runtime builds its map and object list at initialization. Do not assume it automatically discovers new dynamic states or targets.

Registry entries are the single source for element names, explanations, and parent relationships. The following illustrates the schema only; it does not require a title in every project:

```json
[
  {
    "id": "main-title",
    "name": "主标题",
    "role": "说明页面的主要内容",
    "reason": "较强的字重建立第一阅读重点，与次要说明形成层级",
    "specs": [
      { "label": "字号", "css": "font-size" },
      { "label": "字重", "css": "font-weight" },
      { "label": "强调方式", "value": "提议：通过字重而非额外装饰建立重点" }
    ]
  }
]
```

Each specification uses either `css` for a rendered CSS property or `value` for authored text. The inspector labels CSS readings as measurements; they describe the current HTML target, not necessarily a native application. Label proposed or approximate authored values accurately and use native units for non-Web targets. Prefer shared tokens or rendered values over duplicated numeric claims. Inspect only values relevant to the item, such as typography, gap, crop, ratio, or emphasis; put shared specifications in the specifications tab.

Serialize registry content as valid JSON. Escape `<` as `\u003c` in generated JSON so embedded `</script>` text cannot terminate the script element. The runtime renders registry strings as text, not HTML. For screenshots or native captures, associate named regions with the image and use authored platform specifications; measuring the HTML wrapper does not measure the native UI.

## Review behavior and feedback

- Hover temporarily previews a target; leaving restores the pinned selection. Clicking or tapping pins the innermost registered target and opens its element explanation. Hover does not switch away from a brief being read.
- The object list and parent action provide keyboard-accessible alternatives to pointer selection. Selection outlines do not change layout or intercept the pointer. Escape and the clear control remove the selection.
- Tabs support Left/Right and Home/End, with one tab in the Tab sequence. Keep visible focus and meaningful control semantics when adapting the shell.
- Inspection mode intercepts clicks and common activation keys to select elements; experience mode lets the represented UI respond. Hiding notes clears selection and enters experience mode; re-enabling inspection reveals notes. Keep any live product actions confined to authorized preview behavior: this event handling is not a security boundary and cannot suppress arbitrary scripts or every custom event handler.
- The empty template has no fabricated selection and hides the mode control until real registered targets exist. Project copies should register their meaningful content before delivery.

Resolve feedback using the selected item or its familiar name and surrounding context. Ask a focused question only when ambiguity materially changes the result. Update the affected design, explanation, and specifications together; for “副标题离主标题近一点,” adjust the relevant gap without silently redesigning the whole page. Preserve unrelated decisions and keep unconfirmed choices visibly provisional.

## Completion checks

Inspect the artifact in an available browser and check the scope actually delivered:

- **Shell and composition:** wide and narrow layouts, long names, text enlargement, readable contrast/focus, UI prominence, reachable notes, and no unintended horizontal overflow. Check screenshots as well as code; the compact toolbar must still accommodate its labels and controls.
- **Filled project:** remove filling prompts and empty slots; verify every registered target and parent, unique IDs, and agreement between UI, names, reasons, and specifications. Check relevant responsive and represented interaction states, including a sample revision's affected values.
- **Selection when present:** a module and nested element, repeated items where applicable, object-list and parent selection, hover restoration, clearing, keyboard access, and touch alternatives. Verify pinning opens the element tab without hover disrupting another tab; test notes collapse, mode switching, and narrow-screen jump/return. Confirm the represented product action does not fire during inspection for the controls actually used.
- **Reusable template:** preserve its authoring fields and intentional empty state. Exercise shell controls directly; validate inspection with in-memory or temporary fixtures without leaving sample targets in the asset. Do not describe those checks as testing a real application.
- **Delivery:** provide the entry link or launch instructions, a brief explanation of how to select and discuss elements, and what was actually checked. Distinguish browser simulation from real device checks and HTML approximation from native validation. Report unavailable checks; neither code review nor successful automation establishes visual approval or production readiness.

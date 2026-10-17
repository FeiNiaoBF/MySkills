# HTML Preview Requirements

Read this when preparing or updating the page's HTML design draft. These are delivery requirements, not a prescribed preview framework or bundled implementation.

## Page and storage

- Save the draft in the user-designated folder, with a clear entry file and opening instructions. Keep any required local assets alongside it and use portable relative links.
- Make the actual page the main canvas. Keep review controls compact; show selected-object details beside the page on wide screens and without obscuring the target on narrow screens.
- Reuse the established visual direction. Use representative content and enough interaction to discuss the agreed flows; identify mocked behavior and unfinished states.
- Keep the HTML and conversational explanation aligned with the same current decisions. Update the existing draft rather than leaving several unexplained competing versions.

## Locate the intended object

- Give reviewable modules and internal elements everyday names, stable unique IDs, and parent relationships. Include meaningful titles, descriptions, images, controls, and repeated items; do not stop at module-level selection.
- Distinguish repeated elements by context, such as “Bookshelf / Second book / Title.” Preserve an object's ID across ordinary moves and wording changes; do not silently reuse a removed object's ID for something else.
- Let the user select a child precisely and move to its containing module. Highlight only the current target clearly, without covering the interface with permanent labels.
- Show the target's name, parent context, purpose, relevant behavior/state, and unresolved choices. Mark a decision as confirmed only when supported by user feedback or an explicit prior decision.
- Provide a copyable name/path and ID that the user can quote in the conversation. Copying should have a selectable-text fallback if clipboard access is unavailable.

For example, selecting “Contents panel [B02]” lets the user say: “Move B02 to the left of the reading text, collapsed by default; keep the text width.” Selecting a book title lets the user change that element without implicitly redesigning every book card.

## Keep inspection simple

- Separate normal preview interaction from inspection. In preview mode, the implemented page controls work; in inspection mode, selecting a button must not trigger its product action.
- Support pointer, keyboard, and touch selection; do not depend on hover alone. Make the current selection and focus visible, and provide a clear way to clear selection, including Escape where a keyboard is available.
- Start with locating, selecting, and referring to objects. Do not add drag-and-drop layout editing, a style editor, or a persistent comment system unless requested and justified.
- Keep viewport/device controls relevant to the target page. Do not require a particular device-frame harness, fixed sidebar width, or custom attribute protocol.

## Feedback and checks

Interpret location references together with the requested change. Clarify whether a change concerns one item, its module, or all repeated items only when the intent is ambiguous. Apply the smallest coherent change and update the preview and explanation together.

Before presenting the draft, inspect its rendered layout at relevant sizes and try child/parent selection, repeated-item identification, reference copying, keyboard/touch access, selection clearing, and switching between preview and inspection. Confirm that inspection does not fire product actions. Fix issues and repeat affected checks; disclose any check the available tools cannot perform.

The draft helps review content, structure, and behavior. It does not itself prove production functionality, native-platform fidelity, or user acceptance.

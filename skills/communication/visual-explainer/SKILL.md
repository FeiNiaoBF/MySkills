---
name: visual-explainer
description: >-
  Use only when the user explicitly invokes `visual-explainer` or directly requests a
  standalone HTML explainer page. Turn one complete explanation into a single offline
  HTML file with an answer-first reading path and visuals that carry real relationships.
  Not for chat answers, design previews, or research reports.
metadata:
  version: "0.4.0"
disable-model-invocation: true
---

# Visual Explainer

Produce one self-contained HTML page that carries a complete explanation. Decide the explanation before choosing a layout; the page exists to make its structure and relationships easy to follow.

## Invocation boundary

- Start only when the user names `visual-explainer` or explicitly requests a standalone HTML explainer page.
- Do not build a page because a question is long, comparative, technical, or hard to read in chat. Answer ordinary explanation requests in the conversation.
- An explicit request authorizes building the page. Ask one focused question only when a missing choice would materially change the result; otherwise state a reasonable assumption and proceed.
- If a page for this topic already exists, update it instead of creating a second file.

## Frame the explanation

Before drafting, settle:

- The question the reader must be able to answer after reading.
- The conclusion or mental model they need, not the topic label they used.
- Prior knowledge visible in the request, and the smallest missing prerequisite.
- What this page covers and what it deliberately leaves out.

Do not infer expertise from unrelated context. When several readings are plausible, take the one the request supports; ask only when the wrong choice would change the explanation substantially.

## Choose an explanation path

Match the path to the shape of the question. Use one path; combine two only when neither alone answers the question.

| Path | Section order |
| --- | --- |
| Concept | core meaning → intuition → concrete example → where the intuition stops |
| Process or system | purpose → parts → sequence and state changes → branches, failures, recovery |
| Comparison or decision | decision → criteria → differences and trade-offs → recommendation and its conditions |
| Claim or conclusion | conclusion → evidence or premises → reasoning → assumptions, uncertainty, limits |
| Code or architecture | observable behavior → responsible components → control or data flow → effect of changing a component |
| History or evolution | present state → turning points in time order → what changed → why it matters now |

Choose the shortest path that preserves the reasoning the reader needs. This table lists section orders; it does not ask every page to contain a table.

## Draft the explanation

1. Lead with the answer or central model, short enough to read before scrolling.
2. Order sections so each uses only ideas already introduced; define a term where it first affects the reasoning.
3. Give every section one explanatory job. Delete a section that only restates another in different words.
4. Anchor abstractions in one concrete example that stays accurate and consistent across sections.
5. Keep facts, examples, assumptions, and recommendations visibly distinct. Label illustrative values; never invent evidence, statistics, quotations, benchmarks, or sources.
6. Preserve material caveats, failure conditions, uncertainty, and evidence limits instead of smoothing them out for a cleaner story.
7. Reserve a visual for a relationship prose handles worse: order, comparison, containment, flow, or change over time.

Language: write the page in the language the user asked for, otherwise the language they are using. Set `lang` to the matching BCP 47 tag (`zh-CN`, `en`, `ja`) and use one language consistently outside quoted terms.

Length: no chapter quota and no padding. Typical pages need three to six chapters; a page that restates the chat answer at greater length has failed. Derive the reading time from the actual word or character count and keep it approximately honest.

Author the full text before styling. Styling may arrange, emphasize, and validate content; it must not add claims, reorder causality, drop qualifications, or reinterpret an example.

## Build the page

Start from the bundled [HTML template](assets/explainer.html): copy it to the delivery location and edit the copy, never the asset.

The `{{...}}` strings are literal authoring markers. Replace every one — `lang`, `<title>`, labels, ARIA text, link targets, reading time — and delete the template's authoring comments. Search the finished file for `{{` before delivery.

The template is a visual and accessibility baseline plus optional modules, not a content outline:

- Repeat, reorder, or delete chapters until each remaining section has one job.
- Keep the ordered-sequence figure only for real order or causation. Adjust its items and `--process-columns` together, and keep a visible complete text equivalent. Replace it with the comparison table, prose, or a small inline SVG when another relationship is meant.
- Use the comparison table for criteria and trade-offs. Keep its `<caption>`, header cells, and the keyboard-scrollable wrapper for narrow screens.
- Remove the example, limits, or sources block when it carries no real content. Remove the reading-progress element and its script together, or keep both.
- Keep ids unique, and keep chapter numbers, heading references, labels, and reading time consistent with the final sections.
- Preserve the design tokens, wide desktop canvas, bounded prose width, narrow-screen reflow, visible focus, reduced-motion behavior, and print rules unless the task supplies a stronger requirement.

Delivery constraints:

- One `.html` file with all CSS and JavaScript inside; it must open from the filesystem with no build step and no network.
- No CDN scripts, web fonts, remote images, or MathJax. Draw with HTML, CSS, and inline SVG; inline tiny assets as data URIs, and set formulas as HTML text or SVG paths instead of pulling a renderer.
- JavaScript is optional and additive. Guard it so a deleted or renamed element cannot break the page, and keep every explanation readable with scripting disabled.
- Save to the location the task implies: the folder the user named, the task's output folder, otherwise the workspace root. Use a descriptive kebab-case filename ending in `.html`, and never overwrite an unrelated existing file.

## Visual and accessibility baseline

- Semantic structure: `header`, `main`, `section`, `footer`; exactly one `h1`; no skipped heading level; text that still makes sense read top to bottom.
- Contrast: at least 4.5:1 for body text and 3:1 for large text, meaningful borders, and focus rings. Recheck every token pair after a palette change; the bundled light palette already passes.
- Never encode meaning in color, position, animation, or hover alone. Pair each state or category with a label, a pattern, or text.
- Keyboard and zoom: every interaction reachable and visible by keyboard, focus never removed, layout usable at 200% zoom and at 320 px width without clipping or hiding content behind horizontal scroll (wide tables excepted, inside the labelled scroll region).
- Motion: honor `prefers-reduced-motion`; no autoplaying or looping motion in an explanation.
- Figures: every diagram and table has a caption or adjacent text equivalent, decorative shapes are `aria-hidden`, and no essential text is baked into a raster image.

## Keep the scope clear

- `visual-explainer` builds one complete, standalone HTML explanation after an explicit request.
- `make-sense` repairs a single confusing term, passage, or inference in the conversation and returns to the original task.
- `research-curator` investigates a question across sources and publishes a traceable report ledger; this skill explains what is already understood and does not run retrieval or invent citations.
- `aesthetic-translator` and `frontend-guide` produce design previews and implementation handoffs for an interface; this skill explains a subject, not a UI to build.
- `project-navigator` explains an existing codebase in conversation; use this skill only when the user wants a standalone page.
- Stay inside the requested question. A narrow question does not become a course, and the page does not prove the reader learned the material.

## Verify before delivery

Content:

- The page answers the requested question and states the conclusion before the supporting detail.
- The section order follows the chosen path, each section has a distinct job, and no prerequisite appears after it is needed.
- Examples, figures, tables, and prose agree; nothing claims more than the stated evidence; caveats and limits survive.
- Deleting any decorative panel or repeated summary leaves no reasoning gap.

File:

- No unresolved `{{`, no template instructions, no leftover content from another topic.
- No `http://` or `https://` outside the sources list, and no external `src`.
- No duplicate `id` values and no second `h1`.
- The file opens locally, stays readable with scripting disabled, and logs no console error with it enabled.

Rendering:

- Open the finished file and inspect it at desktop width, at 320 px, at 200% zoom, and with keyboard-only navigation.
- Check print preview, and check a reduced-motion setting when the page relies on motion.

Report the file path and the explanation's structure in two or three sentences. Do not paste the whole page into chat unless asked.

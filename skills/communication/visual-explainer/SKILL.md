---
name: visual-explainer
description: >-
  Use only when the user explicitly invokes visual-explainer or asks to create an HTML explainer.
  Turn a complex explanation into a clear, self-contained HTML page with a deliberate learning path,
  accurate examples, and visuals that clarify relationships.
metadata:
  version: "0.3.0"
disable-model-invocation: true
---

# Visual Explainer

Create a self-contained HTML page that helps the reader understand a complete explanation. Design the explanation first; use the page to make its structure and relationships easy to follow.

## Invocation boundary

- Run only after the user explicitly invokes `visual-explainer` or directly asks for an HTML explainer.
- Do not create an HTML page because a question is long, complex, comparative, or explanatory.
- If the user asks for an ordinary explanation, answer in the conversation and do not invoke this workflow.
- Treat an explicit invocation as authorization to create the requested page. Ask one focused question only when a missing choice would materially change the explanation; otherwise state a reasonable assumption and proceed.

## Frame the explanation

Before drafting, identify:

- The exact question the reader should be able to answer afterward.
- The conclusion or mental model they need, not merely the topic they mentioned.
- The prior knowledge shown in the request and the smallest set of prerequisites still missing.
- The scope boundary: what this explanation covers and what it deliberately leaves out.

Do not infer expertise from unrelated context. If several interpretations are possible, choose the one supported by the request; ask one focused question only when choosing incorrectly would materially change the explanation.

## Choose an explanation path

Match the path to the shape of the reader's question. Combine paths only when one cannot answer the question completely.

- **Concept:** core meaning → useful intuition → concrete example → where the intuition stops working.
- **Process or system:** purpose → participating parts → sequence and state changes → branches, failures, or recovery.
- **Comparison or decision:** decision to make → relevant criteria → differences and trade-offs → recommendation under stated conditions.
- **Claim or conclusion:** conclusion → supporting evidence or premises → reasoning between them → assumptions, uncertainty, and limits.
- **Code or architecture:** observable behavior → responsible components → control or data flow → consequence of changing a component.
- **History or evolution:** present result → turning points in chronological order → what changed at each point → why those changes matter now.

Do not force every explanation into the same template. Use the shortest path that preserves the reasoning the reader needs.

## Draft the explanation

1. Put the answer or central mental model first. Do not make the reader finish the page to discover the conclusion.
2. Order each later section so it relies only on ideas already introduced. Define a necessary term where it first affects the explanation.
3. Give each section one explanatory job. Remove a section when it only repeats another section in a different form.
4. Connect abstract ideas to observable behavior with a concrete example. Keep the example accurate, relevant, and consistent across sections.
5. Distinguish facts, examples, assumptions, and recommendations. Mark illustrative values as illustrative and do not invent evidence.
6. Preserve important exceptions, uncertainty, failure conditions, and evidence limits. Do not simplify them away for a cleaner story.
7. Choose a visual form only for a relationship it expresses better than prose: sequence for order, table for comparison, diagram for relationships, or timeline for change.

## Separate explanation from presentation

The model owns the meaning of the explanation: the conclusion, section order, definitions, examples, relationships, and caveats. Any present or future rendering tool owns presentation mechanics only.

A rendering tool may arrange, style, or validate supplied content. It must not add factual claims, alter causality, silently remove qualifications, or reinterpret an example. Keep the explanation complete enough to review before presentation decisions are applied.

Do not create a reference file until a real task introduces a substantial branch that the main workflow does not always need. Keep invocation, explanation design, and quality criteria in this file because they apply to every run.

## Build the HTML page

Before building, read and copy the bundled [HTML template](assets/explainer.html). Treat it as the visual, responsive, and accessibility baseline; do not edit the source asset for an individual task.

The `{{...}}` strings are literal authoring markers, not a template-engine contract. Replace every marker in the copied file, including document language, page title, visible labels, link targets, and accessible names. Search the finished file for unresolved `{{...}}` markers before delivery.

The template is not a required content outline. Match it to the explanation path chosen above:

- Repeat, reorder, or delete chapters so each remaining section has a distinct explanatory job.
- Remove optional process, detail-list, example, caveat, scope, or source modules when they do not help answer the question.
- Use the process figure only for an ordered sequence or causal chain. Adjust its items and `--process-columns` together, and retain a visible complete text equivalent. Replace it with a semantic table, relationship diagram, timeline, or other appropriate form when the subject requires one.
- Keep IDs unique after adding or removing sections. Update heading references, accessible labels, chapter numbering, and reading-time text to agree with the final content.
- Preserve the centralized design tokens, readable class names, wide desktop canvas, bounded prose width, narrow-screen reflow, keyboard focus, reduced-motion behavior, and print cleanup unless the task supplies a stronger requirement.

Additional delivery requirements:

- Deliver one complete `.html` file. Keep CSS and any necessary JavaScript inside that file so it works when opened locally without a build step or network access.
- Use semantic headings and a clear reading order. Make the page responsive, keyboard accessible, and readable with high contrast and reduced motion support.
- Use diagrams, callouts, examples, and tables to support the explanation. Keep all essential meaning available in text; do not rely on color, animation, hover, or interaction alone.
- Add interaction only when it helps the reader inspect or understand the subject. Keep static content complete when scripts are unavailable.
- Avoid decorative panels, repeated summaries, generic stock imagery, and controls that add no explanatory value.
- Save the file in the current task's expected output location. If none is specified, use a clearly named file in the workspace and report its path.

## Keep the scope clear

- `visual-explainer` creates a complete, structured HTML explanation after explicit invocation.
- `make-sense` explains a specific confusing term, passage, claim, or reasoning gap in the conversation. It can unpack several linked steps when needed, but stays anchored to the user's actual obstacle.
- Do not turn the page into a generic lesson when the request concerns a narrower question. Do not claim that the page proves the reader learned the material.

## Review before delivery

Check that:

- The page answers the requested question and gives the conclusion before supporting detail.
- The chosen explanation path matches the shape of the question instead of following a generic template.
- Each section has a distinct role, and the sequence does not rely on an unexplained prerequisite.
- Every new term, causal link, and transition needed for the conclusion is stated rather than implied.
- Examples, diagrams, and prose agree with one another and do not introduce unsupported claims.
- Facts, illustrations, assumptions, recommendations, conditions, caveats, and evidence limits remain distinguishable.
- Removing decorative or repeated content would not expose a missing reasoning step.
- The page follows the appropriate explanation path rather than the template's initial section order.
- The page contains no unresolved `{{...}}` authoring markers, filling instructions, or content copied from an unrelated example.
- Each remaining template module helps explain a real relationship or necessary qualification; unused modules are removed.
- The page opens as a local file and remains usable on a narrow screen and with keyboard navigation.

Report the file path and briefly describe the explanation's structure. Do not also paste the full page into chat unless the user asks.

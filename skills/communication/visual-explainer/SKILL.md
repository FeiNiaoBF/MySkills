---
name: visual-explainer
description: >-
  Use only when the user explicitly invokes visual-explainer or asks to create an HTML explainer.
  Turn a complex explanation into a clear, self-contained HTML page with a deliberate learning path,
  accurate examples, and visuals that clarify relationships.
metadata:
  version: "0.1.0"
disable-model-invocation: true
---

# Visual Explainer

Create a self-contained HTML page that helps the reader understand a complete explanation. Design the explanation first; use the page to make its structure and relationships easy to follow.

## Invocation boundary

- Run only after the user explicitly invokes `visual-explainer` or directly asks for an HTML explainer.
- Do not create an HTML page because a question is long, complex, comparative, or explanatory.
- If the user asks for an ordinary explanation, answer in the conversation and do not invoke this workflow.
- Treat an explicit invocation as authorization to create the requested page. Ask one focused question only when a missing choice would materially change the explanation; otherwise state a reasonable assumption and proceed.

## Design the explanation

1. Identify the reader's question, the conclusion they need, and the relevant prior knowledge shown in the request. Do not assume expertise from unrelated context.
2. Choose the smallest coherent explanation that answers the request. For a system or process, explain the parts and their relationships. For a comparison, identify the decision criteria and show how they affect the choice. For a claim, show its reasoning, conditions, and limits.
3. Put the main answer first. Order later sections so each concept uses only ideas already introduced. Define necessary terms where they first matter.
4. Use a concrete example to connect an abstract idea to observable behavior. Keep the example accurate, relevant, and consistent with the explanation.
5. Preserve important uncertainty, exceptions, assumptions, and evidence limits. Do not simplify them away to make the page look cleaner.
6. Choose visual forms for the information they express: use a sequence for steps, a comparison table for alternatives, or a diagram for relationships. Add visuals only when they make a relationship easier to see.

## Build the HTML page

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
- Each section has a clear role in the explanation, and the sequence does not rely on an unexplained prerequisite.
- Examples, diagrams, and prose agree with one another.
- Assumptions, conditions, caveats, and evidence limits remain visible.
- The page opens as a local file and remains usable on a narrow screen and with keyboard navigation.

Report the file path and briefly describe the explanation's structure. Do not also paste the full page into chat unless the user asks.

# Research and design rationale — make-sense

This document is for **maintainers**. It motivates design decisions; it is not necessary to load during normal skill use. The mapping from research to prompt rules is an engineering hypothesis, not a claim that this particular skill has been experimentally validated.

## Research-to-behavior mapping

| Research lens | Supported general insight | Design implication to test |
| --- | --- | --- |
| Processing fluency and jargon | Introducing specialist language can interfere with non-expert engagement; adding definitions does not automatically eliminate the problem. | Reconstruct the whole claim rather than blindly expanding acronyms. |
| Scaffolding | Good support responds to the learner's current performance, fades, and transfers responsibility. | Use topic-local evidence; explain only the obstacle; stop repeating basics after demonstrated understanding. |
| Self-explanation | Generating causal/conceptual explanations can improve learning across studied contexts. | When warranted, ask one genuine application/paraphrase question, not ritual confirmation. |
| Illusion of explanatory depth | People can overestimate how well they understand mechanisms. | Treat "sounds clear" as insufficient evidence of deeper mastery; avoid announcing mastery. |
| Refutation texts / conceptual change | Addressing an incorrect belief explicitly can support learning in studied settings. | State the correct part and precise correction before providing a replacement account. |
| Cognitive load / prior knowledge | The usability of support depends on what a learner already knows and on task demands. | Use the minimum sufficient explanation; avoid exhaustive glossaries and overloading the first response. |
| Grounding in communication | Communicators construct mutual understanding in interaction. | Reconnect explanations to the user's exact original content and adapt when confusion persists. |

## Plain-language writing lens

George Orwell's six rules are a useful maintenance lens, not six absolute runtime laws. Their value here is reducing language friction while preserving the exact distinctions needed to repair understanding.

| Rule | Decision | Adaptation for `make-sense` |
| --- | --- | --- |
| Avoid familiar metaphors and figures of speech. | Adapt, do not impose as a ban. | Remove stock or decorative imagery. Prefer the user's original example; use an analogy only when it repairs the blocker, and name its limits when they matter. |
| Prefer a short word to a long one. | Adapt, do not optimize word length. | Prefer familiar, specific words. Keep a longer or technical term when it is the precise term the user needs. |
| Cut every word that can be cut. | Adopt with an accuracy boundary. | Remove wording that carries no meaning, but keep necessary context, conditions, distinctions, and reasoning steps. |
| Prefer the active voice. | Adapt across languages and contexts. | Make actors and actions clear. Passive constructions remain useful when the actor is unknown, irrelevant, or not the sentence's focus. |
| Prefer an everyday equivalent to foreign phrases, scientific terms, or jargon. | Adapt, because vocabulary can itself be the learning target. | Use plain language for the explanation while retaining unavoidable canonical terms beside their meanings. Do not hide terminology the user must recognize later. |
| Break the rules rather than say something barbarous. | Adopt as the governing exception. | Accuracy, scope, safety, natural language, requested depth, and complete reasoning outrank stylistic brevity. |

The runtime expresses these decisions as one compact paragraph rather than a six-item checklist. Existing rules about requested depth, preserved conditions, and complete derivations bound concision so that “shorter” does not become the goal.

## Primary and scholarly sources

- Shulman et al., **The Effects of Jargon on Processing Fluency, Self-Perceptions, and Scientific Engagement** (2020), *Journal of Language and Social Psychology*. Search by exact title for journal record; effects and boundary conditions must be read in the original paper before quantitative claims.
- van de Pol, Volman & Beishuizen, **Scaffolding in Teacher–Student Interaction: A Decade of Research** (2010), *Educational Psychology Review*. https://doi.org/10.1007/s10648-010-9127-6
- Bisra et al., **Inducing Self-Explanation: a Meta-Analysis** (2018), *Educational Psychology Review*. https://doi.org/10.1007/s10648-018-9434-x
- Rozenblit & Keil, **The Misunderstood Limits of Folk Science: An Illusion of Explanatory Depth** (2002), *Cognitive Science*. https://doi.org/10.1207/S15516709COG2605_1
- Schroeder & Kucera, **Refutation Text Facilitates Learning: a Meta-Analysis of Between-Subjects Experiments** (2022), *Educational Psychology Review*. Search by exact title; the prompting rule here is an extrapolation to LLM interaction, not a direct tested result.
- Clark & Brennan, **Grounding in Communication** (1991), in *Perspectives on Socially Shared Cognition*. Search by exact title for the chapter.

## Distinguish design from evidence

The six gap types (term, prerequisite, relationship, reasoning, misconception, passage) are an **operational taxonomy designed for this skill**, not a validated psychological diagnosis or a taxonomy attributed to one paper. "Diagnose → Repair → Verify" is a design workflow. It has not been demonstrated to improve learning outcomes solely by being embedded in a SKILL.md.

## Scope and simplification

`make-sense` repairs an expressed understanding problem, including standalone term questions. It does not start unsolicited teaching merely because technical vocabulary appears. Broad learning and translation requests retain their own scope; a requested derivation within a clarification still deserves its full steps. The current runtime provides text explanations, including examples, code snippets, or equations where useful; diagrams and other visual explanations are outside its scope. Fact verification and uncertainty disclosure remain part of the guidance.

Version 1.1 uses the shortened candidate as its starting point, retaining blocker identification, factual boundaries, requested depth, and a changed strategy after repeated confusion. The explicit gap taxonomy is design reference rather than a mandatory runtime procedure. The earlier comparison supports investigating lower instruction overhead, not a claim that shorter prompts improve comprehension. The original candidate remains frozen for interpreting that historical result.

## What a real evaluation would establish

A valid evaluation requires independent judgments and adversarial counterexamples, not the same model rating its own output. For every fixed case in [cases.yaml](cases.yaml), record:

1. Was the user's actual obstacle addressed?
2. Was the original meaning (including caveats) preserved?
3. Did the response introduce avoidable undefined prerequisites?
4. Did it hallucinate specifications, source facts, or user knowledge?
5. Did it over-teach or force a quiz?
6. When confused a second time, did it change strategy?
7. Did it use stock, decorative, or empty wording instead of explaining directly?
8. Did concision remove a condition, distinction, or reasoning step needed for accuracy?

Compare a no-skill baseline, a frozen previous runtime, and the proposed runtime using [the evaluation protocol](evaluation.md). Independent model sessions reduce self-review leakage but do not replace human judgment or establish learning outcomes. Add real, anonymized failures over time. A human reviewer should judge the hard cases; automatic keyword matching can only catch some regressions.

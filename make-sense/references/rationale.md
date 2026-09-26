# Research and design rationale — make-sense v1.0

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

## Primary and scholarly sources

- Shulman et al., **The Effects of Jargon on Processing Fluency, Self-Perceptions, and Scientific Engagement** (2020), *Journal of Language and Social Psychology*. Search by exact title for journal record; effects and boundary conditions must be read in the original paper before quantitative claims.
- van de Pol, Volman & Beishuizen, **Scaffolding in Teacher–Student Interaction: A Decade of Research** (2010), *Educational Psychology Review*. https://doi.org/10.1007/s10648-010-9127-6
- Bisra et al., **Inducing Self-Explanation: a Meta-Analysis** (2018), *Educational Psychology Review*. https://doi.org/10.1007/s10648-018-9434-x
- Rozenblit & Keil, **The Misunderstood Limits of Folk Science: An Illusion of Explanatory Depth** (2002), *Cognitive Science*. https://doi.org/10.1207/S15516709COG2605_1
- Schroeder & Kucera, **Refutation Text Facilitates Learning: a Meta-Analysis of Between-Subjects Experiments** (2022), *Educational Psychology Review*. Search by exact title; the prompting rule here is an extrapolation to LLM interaction, not a direct tested result.
- Clark & Brennan, **Grounding in Communication** (1991), in *Perspectives on Socially Shared Cognition*. Search by exact title for the chapter.

## Distinguish design from evidence

The six gap types (term, prerequisite, relationship, reasoning, misconception, passage) are an **operational taxonomy designed for this skill**, not a validated psychological diagnosis or a taxonomy attributed to one paper. "Diagnose → Repair → Verify" is a design workflow. It has not been demonstrated to improve learning outcomes solely by being embedded in a SKILL.md.

## Boundary with show-me

`show-me` chooses a visual representation of content. `make-sense` selects and repairs the smallest blocking conceptual/semantic gap. It can make a useful clarification without any drawing and must not duplicate a diagramming workflow. Do not make either skill a mandatory dependency of the other.

## What a real evaluation would establish

A valid evaluation requires independent judgments and adversarial counterexamples, not the same model rating its own output. For every fixed case in [cases.yaml](cases.yaml), record:

1. Was the user's actual obstacle addressed?
2. Was the original meaning (including caveats) preserved?
3. Did the response introduce avoidable undefined prerequisites?
4. Did it hallucinate specifications, source facts, or user knowledge?
5. Did it over-teach or force a quiz?
6. When confused a second time, did it change strategy?

Compare a baseline model without the skill against current and shortened instructions using [the evaluation protocol](evaluation.md). Independent model sessions reduce self-review leakage but do not replace human judgment or establish learning outcomes. Add real, anonymized failures over time. A human reviewer should judge the hard cases; automatic keyword matching can only catch some regressions.

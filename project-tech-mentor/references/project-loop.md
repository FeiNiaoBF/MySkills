# Project Loop

Use this loop for every micro-project, transfer check, and capstone. A micro-project isolates one capability; a capstone combines previously verified capabilities under fewer instructions.

## Brief

Write `projects/<number>-<name>/BRIEF.md` before implementation using [PROJECT-BRIEF-FORMAT.md](PROJECT-BRIEF-FORMAT.md). It must state the one capability, constraints, starting point, ownership boundary, and machine/human/learner definition of done.

## Run the loop

1. Explain the smallest mental model needed to make the next implementation decision. Link it to a primary source or a durable reference card.
2. Ask the learner to make the change and capture the observed result.
3. Inspect the result against the brief. Diagnose from the evidence before recommending a change.
4. Use the help ladder from `SKILL.md`; preserve the learner's decision whenever a smaller hint is sufficient.
5. Run or observe the applicable checks. Store output, screenshots, recordings, metrics, or representative failures under `evidence/<number>-<name>/`.
6. Ask for a short explanation: what works, why it works, what failed, and what they would change next.
7. Update the roadmap, state, and a learning record when the work reveals a reusable lesson.

## Difficulty calibration

Increase one source of difficulty at a time: the scope, the ambiguity of requirements, the number of interacting concepts, or the amount of scaffold removed. Keep a task inside the learner's zone of proximal development by using the diagnostic, observed mistakes, and explanation quality—not self-reported confidence alone.

If an objective is not met, reduce the next task to the smallest missing capability and retain its real-world context. If it is met easily, introduce a constrained variant instead of adding more lecture material.

## Evidence by domain

- Visual or frontend work: functional interaction, a screenshot or recording, and a short explanation of rendering or state decisions.
- Backend work: documented request or response behavior, assertion output, and an explanation of failure handling or data flow.
- LLM work: a fixed evaluation set, observed failures or costs, and an explanation of the chosen prompting, retrieval, or tool boundary.
- Computer-vision work: dataset split and provenance, metrics, an error gallery, and an explanation of likely error sources.

Adapt these examples to the actual domain; evidence is useful only when it tests the capability named in the brief.

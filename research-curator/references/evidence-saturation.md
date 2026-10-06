# Evidence Saturation

Evidence Saturation is a bounded stopping decision: further searching may find more sources, but recent rounds have added little evidence, perspective, or dispute that could materially change the answer.

It is **not** proof that the internet has been exhausted. It is relative to the user's question, declared scope, audience, and available retrieval tools.

## A research round

A round is a completed batch of retrieval and source assessment aimed at a named question or search angle. It is not one URL, one tool call, or one search-results page. A completed round must record actual query IDs and at least one assessed source. Failed, partial, or empty rounds do not count as low-gain rounds.

Record for each round:

- questions, query IDs, search intent, and search angle;
- end-of-round coverage snapshot and remaining gaps;
- source IDs assessed and which assessed sources were redundant;
- newly identified claims, original evidence origins, contradictions, and questions;
- whether the round added material information, and the reason with evidence references;
- whether the round completed, was partial, or failed.

The report derives these counts from the IDs. For a round with no assessed sources, duplicate ratio is unknown, not 100%. A new URL is not automatically a new evidence origin. A new source for an existing claim may still materially change the answer.

## Default stop rule

Use an eight-round maximum by default. A user-specified budget takes precedence. Use a trailing window of three low-gain completed rounds.

Declare `saturated` only when all conditions hold:

1. Every core question has inspected, relevant evidence and at least one verified supporting claim. Record important unresolved gaps.
2. The three most recent rounds are complete, each assessed at least one source, and none added evidence likely to change a material conclusion, its scope, confidence, or an important open question.
3. The trailing rounds use varied query formulations and angles. At least one deliberately searches for counterevidence. Consider relevant source-origin types; if only one type applies, explain why.
4. No newly found material counterexample, question, or evidence upgrade remains unassessed. A material finding resets the low-gain window.
5. The stop decision references the exact three rounds and explains why remaining gaps do not prevent this bounded answer.

A stable disagreement can itself be a finding. It does not prevent saturation if the disagreement is fully represented with evidence and uncertainty. Do not convert disagreement into a one-sided verified conclusion. A gap essential to answering a core question prevents saturation.

## Classify information gain

- **None:** the round adds no new claim, evidence origin, relevant counterevidence, or question.
- **Minor:** the round adds detail or another example, but does not materially change the answer, its boundaries, confidence, or open questions.
- **Material:** the round changes or may change a conclusion, scope, confidence, key counterposition, or an important question. This resets the low-gain window.

A high duplicate ratio supports a low-gain assessment, but does not prove one. A low duplicate ratio does not prove high information gain. Judge materiality from inspected evidence and record the reason. If the materiality is unclear, do not count that round as low-gain.

## Other stop reasons

- `budget_exhausted`: the recorded round/time/tool budget ended before saturation.
- `retrieval_blocked`: required search/fetch was unavailable or repeatedly failed.
- `user_stopped`: the user asked to stop.
- `in_progress`: research is still active.

A useful article can still be written after a limited stop. State the limitation plainly; never rename a budget or retrieval stop as `saturated`.

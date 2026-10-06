# Evidence package

The report envelope separates the Research handoff from Writing. Its `run` field is the authoritative v1 source/claim/conclusion ledger. Do not copy those records or reconstruct a second graph in the article.

For the exact v1 ledger fields and graph rules, read [DESIGN.md](../DESIGN.md). This reference describes only the envelope added for the phased workflow.

## Top-level shape

`report_version` is `"1.0"`. `run` is a complete v1 run object. `research` is required. `article` is absent while research is in progress and required for publication.

### Research record

Record:

- `audience`, `purpose`;
- `origin_types_rationale`, which explains why the selected origin types fit the question or why only one type applies;
- `max_rounds` and `low_gain_window` (default 8 and 3);
- optional `resource_budget`: `{kind, unit, limit, used}`, declared before searching. `kind` is `time`, `tool_calls`, or `cost`; give explicit units (seconds, calls, USD), a positive limit, and measured nonnegative usage. A reached resource limit takes precedence over saturation. Record actual usage even if an in-flight call overran the limit; disclose that overrun as `budget_exhausted`, never lower the measurement or claim saturation. Do not initiate another retrieval after reaching the limit. Keep the round cap even when this tighter budget stops research earlier; do not rewrite it after the fact;
- `rounds`, `coverage`, `gaps`, and `stop`.

Each round includes:

- `id`, `query_ids`, `question_ids`;
- `search_intent`: `discovery`, `gap_fill`, or `counterevidence`;
- `search_angle`, `origin_types_searched`;
- `assessed_source_ids`, `redundant_source_ids`, `new_claim_ids`, `new_origin_ids`, `new_contradiction_refs`, and `new_question_ids`;
- `material_gain`: `none`, `minor`, or `material`, with `materiality_reason` and `materiality_evidence`;
- `coverage_snapshot` and `gaps`, recording question coverage and remaining gaps at the end of the round;
- `status`: `complete`, `partial`, or `failed`.

`new_contradiction_refs` and `materiality_evidence` contain `{claim_id, source_id, quote}` references to exact v1 claim evidence. Do not invent a separate evidence ID. A material assessment must cite evidence. The round's `coverage_snapshot` entries use the same shape as final coverage and include every question addressed in a completed round. Preserve these snapshots as the knowledge available then, not as retrospective final coverage. Snapshot support must come from sources assessed by that round or earlier. Later contradictions do not rewrite earlier snapshots; final coverage still uses the full current ledger. Saturation also requires the latest snapshot for each core question to be covered.

IDs within each round list must be unique. Queries must target an addressed question and cannot be reused across rounds. A broad query can yield evidence for several addressed questions: `query.question_id` names its primary target, not every question informed by its results. Record other questions as addressed only when their assessed evidence or explicit gaps are in the snapshot; do not invent extra queries to force one query per question. New claims and evidence references must relate to an addressed question and need evidence from an assessed source; new origins must be assessed original sources that supply recorded claim evidence and were not assessed in earlier rounds. A claim/question already recorded in an earlier snapshot or evidence reference is not new again. For redundant IDs supply `redundancy_reasons`, an optional object mapping each redundant source ID to its round-specific reason, unless the ledger already provides a duplication/upstream relationship and its reason. A repeatedly assessed original is redundant only with an explicit reason; an original flag alone is not duplication provenance. New contradictions and materiality references must point to assessed sources. `material_gain: none` cannot accompany new finding IDs; use `minor` with a reason for non-material additions.

Each coverage entry has `question_id`, `status` (`covered`, `partial`, `uncovered`), `claim_ids`, and `gaps`. A `covered` question needs verified supporting evidence. Incomplete coverage needs a gap statement.

The stop record has `reason`, `round_ids`, and `rationale`. Allowed reasons are `in_progress`, `saturated`, `budget_exhausted`, `retrieval_blocked`, and `user_stopped`. A `saturated` stop must cite exactly the assessed low-gain window. A budget stop requires the round cap or declared resource budget to be exhausted; cite the last attempted round (no round IDs if a resource limit ended before the first round). Usage is a recorded measurement, not proof supplied by the CLI.

### Article record

`article` contains `title`, `language`, `output_form`, `lead`, and ordered `sections`. A section has `id`, `heading`, and `blocks`.

A block has `id`, `kind`, `role`, `claim_ids`, and `conclusion_ids`:

- `kind`: `paragraph`, `list`, or `table`;
- `role`: `evidence`, `synthesis`, or `context`;
- paragraph content uses `text`;
- list content uses `items`;
- table content uses `headers` and `rows`.

Do not put raw HTML in article blocks. Evidence blocks require claim references. Use context blocks for non-evidential orientation or limitations. A citation reference identifies its claim; the renderer links that claim to its recorded sources and quotations.

## Validate and publish

From the directory containing `go.mod`:

```bash
go run ./cmd/researchcurator validate-report -in report.json
go run ./cmd/researchcurator publish -in report.json -out <topic-folder>
```

`validate-report` accepts a pre-writing handoff without an article. `publish` requires an article, a terminal stop reason, and a sealed v1 ledger (`metadata.status: finalized`). Seal the embedded run with the existing `finalize -in <run.json> -out <sealed-run.json>` command in the run-owned temporary directory, then replace the envelope\'s `run` with those exact sealed records. Its legacy HTML is only an intermediate artifact; sealing is not saturation or truth verification. It writes one self-contained `index.html` in a new destination directory and refuses an existing destination. The `-out` path is explicit; never use an unchecked title as a filesystem path.

The CLI validates structure, references, quotes, statuses, and recorded stop conditions. It cannot prove that retrieval happened, that a researcher searched competently, or that an interpretation is true.

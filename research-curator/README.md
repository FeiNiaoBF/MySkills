# Research curator

A portable research skill and standard-library Go tools for bounded, evidence-backed research. Research and Writing remain separate: the v1 run ledger is the source of truth for inspected sources, claims, quotations, and provenance; a versioned envelope records rounds, coverage, gaps, source roles, and the current stop. Retrieval uses the host agent's search/fetch tools.

## Run

From this directory:

```bash
go test ./...
go run ./cmd/researchcurator --help
```

`record -kind query` timestamps a query when it is recorded; call it immediately before searching. `record -kind source` stamps retrieval when an inspected source is recorded. `record-round` appends a validated round and automatically invalidates the previous stop. These commands record events; they do not search the web.

Validate a research handoff and publish a completed report:

```bash
go run ./cmd/researchcurator validate-report -in report.json
go run ./cmd/researchcurator publish -in report.json [-out <topic-folder>]
```

If `-out` is omitted, publication uses the current workspace's `research-reports/`, except in the skill source repository, where it uses `~/Research Reports/`. A user-selected existing folder is never replaced. The output is a self-contained `index.html` that opens offline; the article comes first, followed by traceable sources, a collapsed Cytoscape evidence map, and optional research audit details.

## Skill and references

Read [SKILL.md](SKILL.md) for the short workflow. The [research guide](references/research-guide.md) covers the Research Contract, real retrieval, Evidence/Claim, provenance, counterevidence, coverage, Evidence Saturation, synthesis, writing, and limitations. [Report delivery](references/report-delivery.md) covers output and inspection. [DESIGN.md](DESIGN.md) documents the legacy v1 ledger; [ADR 0003](references/adr/0003-reader-first-report.md) records the reader-first and current-state decisions.

## Examples and verification

- [examples/report.json](examples/report.json) is a synthetic envelope fixture, not research evidence.
- [examples/run.json](examples/run.json) is a separate real-source legacy v1 ledger; it is not an end-to-end report acceptance.
- [examples/VERIFICATION.md](examples/VERIFICATION.md) records the bounded vNext acceptance and its limits.

Go has no external module dependencies. Cytoscape.js is pinned and embedded; no runtime assets or retrieval backend are downloaded. Open the actual HTML in a browser for behavior checks; unit tests and schema validation do not prove readability, truth, or competent retrieval.

# Research curator

A portable research skill and standard-library Go tools for bounded, evidence-backed research. The workflow separates Research, a validated evidence handoff, Writing, and local publication. It does not fetch pages itself; research uses the host agent's available search and fetch tools.

## Run

From this directory:

```bash
go test ./...
go run ./cmd/researchcurator --help
```

Existing v1 ledger commands remain available: `validate`, `import`, `dedup`, `rank`, `finalize`, `render`, and `record`. They validate, organize, or render a supplied `run.json`; they do not search the web or assert that retrieval occurred.

The phased workflow uses a versioned report envelope around the unchanged v1 ledger. Validate a Research handoff (article may be absent), then publish the completed article to a new topic folder:

```bash
go run ./cmd/researchcurator validate-report -in report.json
go run ./cmd/researchcurator publish -in report.json -out <topic-folder>
```

`publish` writes one self-contained `<topic-folder>/index.html`. It refuses an existing destination and performs no network access. The folder can be moved and opened directly. The reader's browser needs JavaScript for the graph; the article and source records remain in the same document.

## Start with the skill

Read [SKILL.md](SKILL.md) for the workflow. Load the phase-specific references as needed:

- [Research workflow](references/research-workflow.md)
- [Evidence Saturation](references/evidence-saturation.md)
- [Evidence package contract](references/evidence-package.md)
- [Writing](references/writing.md)
- [Report delivery](references/report-delivery.md)
- [Manual behavior cases](references/behavior-cases.md)

[DESIGN.md](DESIGN.md) remains the authoritative legacy v1 ledger contract. The versioned envelope and Research/Writing boundary are recorded in [ADR 0002](references/adr/0002-phased-report-envelope.md).

## Examples and verification boundaries

- [examples/report.json](examples/report.json) is a fully synthetic fixture for the report envelope and local publisher. Its source and claim are not research evidence.
- [examples/run.json](examples/run.json) is the separately documented real-source v1 ledger example; it is not evidence that the new agent workflow or report browser has passed an end-to-end evaluation.
- See [examples/README.md](examples/README.md) for scope and reproduction commands.

Go is needed to run or build the CLI; it has no external Go module dependencies. The visualizer embeds pinned Cytoscape.js and all report assets. Open the generated report without a server. Browser operation must be verified separately from Go tests and JavaScript syntax checks.

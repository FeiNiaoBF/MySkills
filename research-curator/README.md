# Research curator
A portable research workflow and standard-library Go evidence ledger. It records decisions rather than pretending to conduct semantic research automatically.

## Run
From this directory, use `go test ./...` and `go run ./cmd/researchcurator --help`.
`validate` checks schema and graph integrity; `import` is a validated ledger round-trip (not a scraper); `dedup` marks exact URL/text duplicates; `rank` orders selected sources first while retaining rejected/duplicate/superseded candidates; `finalize` computes conservative coverage, seals metadata/rejected-source index, and automatically writes an offline HTML report; `render` makes an offline report. All commands accept `-in run.json` (or stdin) and `-out output.json` (or stdout). Render emits HTML. Nothing fetches remote sources or executes source content.

```bash
go run ./cmd/researchcurator validate -in run.json
go run ./cmd/researchcurator dedup -in run.json -out deduped.json
go run ./cmd/researchcurator rank -in deduped.json -out ranked.json
go run ./cmd/researchcurator finalize -in ranked.json -out finalized.json
# Writes finalized.json and report.html in the same directory.
```

Finalize accepts `-report custom.html` to override the adjacent `report.html`. Stdout JSON requires an explicit report path: `finalize -in run.json -report report.html > finalized.json`. JSON and HTML must have different paths. HTML is rendered from the exact finalized JSON bytes, not the original input. Each file is atomic independently, not a two-file transaction; a late JSON write failure may leave the report, and the command returns failure. No report is emitted for invalid input.

Append explicit records with `record -kind event|query|decision -record record.json -in run.json -out updated.json`; timestamps must be RFC3339. The commands do not manufacture stage-completion records. Missing required data, invalid references and unsupported verification assertions fail with a nonzero exit and no output ledger; an incomplete but structurally valid run is preserved with unverified coverage and gaps.

Start with the [workflow](SKILL.md), then consult [DESIGN.md](DESIGN.md) for the run.json contract. A structurally valid run can still be incomplete; coverage warnings and unverified evidence are preserved.

## Real-source example and verification

See [the example](examples/README.md), its [run.json](examples/run.json), and [verification limits](examples/VERIFICATION.md). Regenerate its offline report with `finalize`; generated HTML and retrieval caches belong in an output workspace, not the skill repository.

Go is needed to run or build the CLI; it has no external Go module dependencies. Cytoscape.js is vendored with its MIT license and pinned provenance; opening the generated HTML needs neither Go nor a server.

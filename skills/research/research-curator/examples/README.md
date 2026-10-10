# Examples

## Synthetic report fixture

[report.json](report.json) exercises the versioned report envelope and offline publisher. Its URL, source, quotation, and claim are synthetic; it is not a record of real retrieval and must not be cited as evidence.

From the research-curator module root:

```bash
go run ./cmd/researchcurator validate-report -in examples/report.json
go run ./cmd/researchcurator publish -in examples/report.json [-out <new-topic-folder>]
```

With `-out`, the publisher writes to that directory and refuses an existing destination. Without it, the publisher writes directly to `<system-temp>/research-curator/<topic>/index.html`, choosing a simple numeric suffix if needed.

## Separate real-source v1 ledger

[run.json](run.json) researches Agent Skills progressive disclosure and description trigger evaluation using current official documents. Its quotations and dates are recorded from retrieved material, not simulated API responses. Two URL aliases and one out-of-scope document were deliberately injected as live screening/deduplication probes and labeled as such. Search returned two organic results.

The documents share an official publisher. Distinct document/provenance groups are not evidence of statistical independence or external corroboration. This run is a legacy v1 ledger fixture; it does not demonstrate the new report envelope, Evidence Saturation workflow, or browser behavior. Full retrieval receipts and rendered HTML remain outside the repository.

Validate the ledger with:

```bash
go run ./cmd/researchcurator validate -in examples/run.json
```

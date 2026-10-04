# Real-source example

This example researches Agent Skills progressive disclosure and description trigger evaluation using current official documents. The final `run.json` contains actual retrieved quotations and dates, not simulated API responses. Two URL aliases and one out-of-scope document were deliberately injected as live screening/deduplication probes; they are labeled as such. Search returned two organic results.

The documents share an official publisher. Distinct document/provenance groups are **not** evidence of statistical independence or external corroboration. No news syndication or circular-reporting case occurred in this real run; those behaviors require separately labeled regression fixtures.

From the skill directory, validate and reproduce the report in a chosen output directory:

```bash
go run ./cmd/researchcurator validate -in examples/run.json
go run ./cmd/researchcurator finalize -in examples/run.json -out <output-dir>/run.json
```

Full retrieval receipts and rendered HTML remain outside the repository in the local validation workspace; no machine-specific paths are embedded in this fixture.

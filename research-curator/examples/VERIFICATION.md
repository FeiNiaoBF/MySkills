# MVP verification and limitations

The bundled `run.json` is a real-source, bounded research example. Search returned two official Agent Skills pages. Both were actually retrieved. Two live URL aliases and an unrelated live Cytoscape documentation page were deliberately injected to exercise exact deduplication and rejection; they are not claimed as organic search results.

Observed ledger: five source records, two selected, two duplicate, one rejected; three claims; two conclusions; thirteen graph nodes and thirteen edges. Every claim retains an exact passage and a locator. No news republication or circular-reporting instance was observed in this real example. Synthetic regression fixtures test explicit provenance cycles and common upstream origins; semantic inference is still the researcher's responsibility.

Checks exercised:

- `go test -count=1 -race -cover ./...`, `go vet ./...`, and CLI build.
- External Draft 2020-12 validation of the schema and real example, including rejection of a missing contract.
- Actual CLI deduplication, ranking, finalization and report regeneration.
- Browser opened `report.html` directly as a local file: thirteen nodes/thirteen relationships; source filters and detail lookup worked; reset restored the full graph; zero resource requests; no horizontal overflow at a 390px viewport.
- Regression tests reject graph-only support/contradiction without quoted evidence, uncited conclusion support, and verified conclusions with explicit unresolved contradictions. A 42-node shared-upstream DAG completes without exponential traversal.
- Renderer tests preserve hostile data without executing markup, enforce HTTP(S)-only links, and bundle the network engine offline.

These checks establish structural and behavioral properties, not source truth, extraction completeness, independence between publishers, or research/learning effectiveness. `independent_sources` counts recorded original-evidence provenance groups. The two selected documents share a publisher.

Intentionally deferred: fuzzy/semantic deduplication, automatic discovery of syndicated originals, automated claim extraction/entailment, conflict-resolution adjudication, configurable ranking policies, long-term evidence history, authenticated-source adapters and complex animation. Retrieval and substantive assessments are supplied by the agent using actual tools; the CLI is not a search engine.

Each artifact is written atomically, but JSON and HTML are not one atomic transaction. If finalization fails during output, inspect the reported error and regenerate both outputs rather than trusting a leftover report.

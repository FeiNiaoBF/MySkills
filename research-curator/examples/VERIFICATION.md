# MVP verification and limitations

The bundled `run.json` is a real-source, bounded research example. Search returned two official Agent Skills pages. Both were actually retrieved. Two live URL aliases and an unrelated live Cytoscape documentation page were deliberately injected to exercise exact deduplication and rejection; they are not claimed as organic search results.

Observed ledger: five source records (two selected, two duplicate, one rejected), three claims, two conclusions, one fully attributed query with five candidate IDs, fourteen graph nodes and nineteen edges. Every claim retains an exact passage and locator. The query graph links each candidate to the selected/rejected/duplicate record. No conflict happened naturally in this research example, so its adjudications array is empty; conflict workflows are covered by synthetic regression tests.

Checks exercised:

- `go test -count=1 -race -cover ./...`, `go vet ./...`, CLI build, and `node --check visualizer/templates/app.js`.
- Draft 2020-12 validation of the schema and real example, including the new split contract and query/adjudication structures.
- Actual CLI finalization of the real example regenerated `run.json` coverage and adjacent offline `report.html`.
- Browser opened `report.html` directly from `file://`: fourteen nodes/nineteen relationships; Query→Candidate visibility and source status details rendered; zero network requests. A separate adversarial UI fixture displayed resolved-adjudication rationale/evidence/final-effect and query retrieval reference/candidate outcomes.
- Regression tests cover unresolved/resolved claim conflicts, separate conclusion adjudication, rejected-source non-resolution, missing rationale/evidence, participant mismatch, fabricated quotes, selected dependency requirement, quantity split, query graph integrity, recorder atomicity, stale verified coverage and old schema rejection. Existing tests cover graph-only assertions, verified-conclusion guards, and shared-upstream DAG limits.
- Renderer tests preserve hostile data without executing markup, enforce HTTP(S)-only links, and bundle Cytoscape.js offline.

These checks establish structural and behavioral properties, not source truth, extraction completeness, independence between publishers, or research/learning effectiveness. `independent_sources` counts recorded original-evidence provenance groups. The two selected documents share a publisher.

Intentionally deferred: fuzzy/semantic deduplication, automatic discovery of syndicated originals, automated claim extraction/entailment, automatic conflict arbitration (adjudications remain explicit human/agent review records), configurable ranking policies, long-term evidence history, authenticated-source adapters, Three.js and complex animation. Retrieval and substantive assessments are supplied using actual tools; the CLI is not a search engine.

Each artifact is written atomically, but JSON and HTML are not one atomic transaction. If finalization fails during output, inspect the reported error and regenerate both outputs rather than trusting a leftover report.

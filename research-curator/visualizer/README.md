# Offline research evidence report

`visualizer.Render(runJSON []byte) ([]byte, error)` produces one self-contained HTML document containing run JSON, CSS, application JavaScript, and the pinned Cytoscape.js engine. Open the file directly in a browser. Rendering uses Go's standard library; viewing needs JavaScript but no server, network, font, CDN, extension, or build step.

## Input contract

The renderer displays the current `run.json` model. The caller performs schema/semantic validation; rendering itself preserves unknown JSON fields for inspection. The canonical graph shape is `graph.nodes` and `graph.edges`; nodes use `id/type/label`, edges use `from/to/type`. It does not accept top-level `nodes` or `edges` as the official ledger structure and never invents graph links.

- Root records include `metadata`, `contract`, `queries`, `events`, `sources`, `claims`, `conclusions`, `decisions`, `adjudications`, `rejected_sources`, `graph`, and `coverage`.
- Query nodes connect `Question → Query → Source` using `searched_by` and `candidate` edges. Query details display provider, tool, retrieval reference, candidate IDs, and each candidate's selected/rejected/duplicate/superseded state and reason.
- Source, Claim, Conclusion, and Query graph nodes resolve to their corresponding root records. Filters and keyboard-accessible graph controls remain available; records can also be inspected without using the graph.
- Conflict adjudications have their own report section, including conflict target, involved claims/sources, resolved/unresolved state, rationale, evidence dependencies, outcome, and final effect. The report displays supplied records; Go validation enforces their integrity.
- The contract shows `output.target_sources` separately from `coverage.min_independent_origins`. The derived run coverage summary reports selected count, target fulfillment, and independent origins separately.

All research content uses text-only DOM APIs. Embedded JSON escapes HTML-sensitive characters and Unicode line separators. Links are limited to absolute HTTP(S) URLs with `noopener noreferrer`. A restrictive CSP blocks network requests and active external content. Source navigation occurs only after a user explicitly selects a link. The graph has no animations; layout does not encode evidence strength.

## Dependency provenance

`vendor/cytoscape.min.js` is the unchanged official npm distribution of **cytoscape 3.34.3**. `vendor/LICENSE` contains the MIT license; the engine also carries its full notice in generated HTML. `vendor/provenance.json` records registry metadata, tarball integrity, upstream commit, and vendored-file SHA-256 checksums. No npm dependency is required at runtime.

## Verification

From the `researchcurator` module root:

```sh
go test -race ./visualizer
go vet ./visualizer
node --check visualizer/templates/app.js
```

The full project contract, including adjudication rules and query-graph integrity, is documented in `../DESIGN.md`. The offline viewer remains a presentation layer; it does not resolve conflicts, calculate research judgments, or retrieve sources.

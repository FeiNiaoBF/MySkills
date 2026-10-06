# Offline research evidence report

The package preserves its legacy API:

```go
visualizer.Render(runJSON []byte) ([]byte, error)
```

It also renders a versioned report envelope:

```go
visualizer.RenderReport(reportJSON []byte) ([]byte, error)
```

`Render` accepts a JSON object containing a legacy run ledger. `RenderReport` checks the envelope version and the presence of object-valued `run`, `research`, and `article` fields. **Neither function performs full domain validation.** The publisher calls `curator.DecodeReport` and `curator.ValidatePublication` before rendering:

```bash
go run ./cmd/researchcurator validate-report -in report.json
go run ./cmd/researchcurator publish -in report.json -out <new-topic-folder>
```

The output is a single self-contained HTML document containing article text, evidence records, the explicit graph, local CSS/JavaScript, and vendored Cytoscape.js. The document works from `file://` without a server or runtime requests; the graph needs JavaScript. The article and source details are embedded in the document. External source URLs are opened only after explicit user action.

## Rendering and safety

- Research content is inserted with DOM text APIs; article text is not interpreted as HTML.
- Embedded JSON is HTML-escaped without round-tripping large numeric values through `float64`.
- The restrictive CSP blocks network access and active external content. Link handling accepts only absolute HTTP(S) URLs and sets `noopener noreferrer`.
- Unknown run fields remain in the expandable complete-data view. Full schema/domain validation belongs to the caller.
- The graph displays recorded relationships, not evidence strength. It is accompanied by keyboard-operable node/relationship controls and text-based source details.
- The report has responsive, print, and reduced-motion styles. Browser operation and screen-reader behavior require separate verification; Go tests and JS syntax checks do not establish them.

Recognized run collections include sources, claims, conclusions, queries, events, decisions, rejected sources, graph nodes/edges, coverage, and pipeline stages. Node/source associations are resolved by explicit IDs, not inferred. The report displays explicit selection status; no evidence weighting or corroboration is invented by layout.

## Dependency provenance

`vendor/cytoscape.min.js` is the unchanged official npm distribution of **cytoscape 3.34.3**. `vendor/LICENSE` contains its MIT license; the engine also carries its notice in generated HTML. `vendor/provenance.json` records the registry metadata URL, tarball URL, upstream SHA-512 integrity, tarball SHA-256, upstream Git commit, and SHA-256 of each vendored file. The downloaded archive was verified against npm integrity before extraction. No source map URL is present.

No npm dependency is required at runtime. To verify local dependency checksums:

```sh
python3 -c 'import hashlib,json,pathlib; p=pathlib.Path("visualizer/vendor"); m=json.loads((p/"provenance.json").read_text()); [(print(name), print(hashlib.sha256((p/name).read_bytes()).hexdigest()), print("verified"), None) if hashlib.sha256((p/name).read_bytes()).hexdigest()==expected else (_ for _ in ()).throw(AssertionError(name)) for name,expected in m["files_sha256"].items()]'
```

## Verification

From the `researchcurator` module root:

```sh
go test -race ./visualizer
node --check visualizer/templates/app.js
```

Tests check malformed inputs, embedded hostile text, script-closing sequences, Unicode separators, literal template markers, large integer preservation, inline engine inclusion, CSP, and absence of external application scripts/styles or HTML execution APIs. For reproducible browser regression checks, build the CLI and run `CHROME=<chromium-executable> node visualizer/browser_test.mjs <CLI-executable> examples/report.json` (Node 22+). All artifacts go under the system temporary `pi-agent/` directory. See [verification evidence](../examples/VERIFICATION.md) for actual runs and limits.

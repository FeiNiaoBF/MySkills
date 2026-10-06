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
go run ./cmd/researchcurator publish -in report.json [-out <new-topic-folder>]
```

The output is a single self-contained HTML document. The article and lead come first, with numbered citations linking to inspected passages. Sources follow the article; a collapsed Cytoscape map shows only article-linked sources, claims, and recorded conclusions. Research boundaries, source roles, process, and raw data remain secondary. UI labels follow `article.language` (Chinese and English; unsupported languages fall back to English). The document works from `file://` without runtime requests; external URLs open only after explicit reader action.

## Rendering and safety

- Research content is inserted with DOM text APIs; article text is not interpreted as HTML.
- Embedded JSON is HTML-escaped without round-tripping large numeric values through `float64`.
- The restrictive CSP blocks network access and active external content. Link handling accepts only absolute HTTP(S) URLs and sets `noopener noreferrer`.
- Unknown run fields remain in the expandable complete-data view. Full schema/domain validation belongs to the caller.
- The map is derived only from article-referenced source/claim/conclusion links; source and claim IDs are internal and are not displayed. No conclusion node is fabricated when the ledger has none. The text relationship list remains available if Cytoscape fails.
- Source roles are explicit where present: verified evidence, candidate lead, and context source. The report displays exact passages, locators, upstream-source links, and event timestamps.
- Comparative reports show a qualified concentration reminder when most verified citations share one URL host; this is a prompt to explain source limits, not an accusation of bias or a corroboration score.
- The report has responsive, print, and reduced-motion styles. Browser operation and screen-reader behavior require separate verification; Go tests and JS syntax checks do not establish them.

The article determines which ledger claims and sources appear in citations and the evidence map. Exact source and claim identifiers are used only to navigate the embedded ledger; reader-facing labels use titles, claims, conclusions, and numbered source notes. The graph shows recorded support/contradiction relations, not evidence strength or independent corroboration.

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

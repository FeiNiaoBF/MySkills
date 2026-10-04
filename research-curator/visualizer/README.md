# Offline research evidence report

`visualizer.Render(runJSON []byte) ([]byte, error)` produces one HTML document containing run JSON, CSS, application JavaScript, and the pinned Cytoscape.js engine. Save the returned bytes to an `.html` file and open it directly in a browser. Rendering uses only Go's standard library; viewing requires JavaScript, but no server, network, fonts, CDN, extensions, or build step.

## Input and integration

The module import is `researchcurator/visualizer`. The boundary accepts one JSON object and rejects invalid JSON, arrays, scalars, and null. Full contract/schema validation belongs to the caller. Unknown fields remain available in the complete-run inspector. Large JSON numbers remain intact in the embedded data; the browser's normal JSON numeric semantics apply after parsing.

Recognized root keys: `schema_version`, `metadata`, `contract`, `queries`, `events`, `sources`, `claims`, `nodes`, `edges`, `decisions`, `rejected_sources`, `conclusions`, and `coverage`.

- Explicit nodes use `id`, `type`, `label`, and `reference_id`; edges use `source`, `target`, and `type`. Nodes are not inferred from claims or sources. Invalid IDs and dangling edges generate visible warnings instead of fabricated relationships.
- Clicking source/claim/conclusion nodes resolves `reference_id` against the corresponding root record collection. Labels fall back through `title`, `text`, `label`, `name`, and `id`.
- Search, type, status, and relationship filters operate on the graph. All records remain independently inspectable outside the graph. Keyboard-accessible buttons provide equivalent graph node and relationship selection.
- Final materials are sources explicitly marked `selected: true` or status `selected`, `accepted`, or `final`; conclusions are displayed independently. No selection is inferred from a graph connection.
- Coverage renders actual supplied fields as a table; collection totals are actual array lengths. Raw JSON remains inspectable for fields not recognized by the viewer.

All research content uses text-only DOM APIs. The JSON embed escapes HTML-sensitive characters and Unicode line separators. Links are constructed only from absolute HTTP/HTTPS URLs, with `noopener noreferrer`. A restrictive CSP blocks network requests and active external content. Opening an explicitly selected source link navigates the browser to that source. There are no graph animations. The graph layout does not encode confidence.

## Dependency provenance

`vendor/cytoscape.min.js` is the unchanged official npm distribution of **cytoscape 3.34.3**. `vendor/LICENSE` contains the MIT license; the engine also carries its full notice inside generated HTML. `vendor/provenance.json` records the exact registry metadata URL, tarball URL, upstream SHA-512 integrity, tarball SHA-256, upstream Git commit, and SHA-256 of each vendored file. Acquisition verified the downloaded archive against the npm integrity before extraction. No source map URL is present.

No npm dependency is required at runtime. To verify local dependency checksums:

```sh
python3 -c 'import hashlib,json,pathlib; p=pathlib.Path("visualizer/vendor"); m=json.loads((p/"provenance.json").read_text()); [(print(name), print(hashlib.sha256((p/name).read_bytes()).hexdigest()), print("verified"), None) if hashlib.sha256((p/name).read_bytes()).hexdigest()==expected else (_ for _ in ()).throw(AssertionError(name)) for name,expected in m["files_sha256"].items()]'
```

## Verification

From the `researchcurator` module root:

```sh
go test -race ./visualizer
go vet ./visualizer
node --check visualizer/templates/app.js
```

Tests exercise the public Render seam: malformed/non-object inputs, hostile script-closing text, Unicode separators, literal template markers, large integer preservation, inline engine, restrictive CSP, and absence of external application scripts/styles or HTML execution APIs.

Real-browser verification used a generated adversarial fixture opened via `file://`: three explicit nodes, two relationships, source detail resolution, type filtering and reset, malicious markup rendered as text (zero images), HTTP-only links, no resource requests, and responsive mobile layout. After graph resize settled at a 390px viewport, document scroll width was 390px. Text, muted text, links, and warning colors against the panel background measured contrast ratios of 14.17:1, 9.23:1, 9.13:1, and 11.88:1 respectively (WCAG AA text threshold: 4.5:1). These checks are not a full accessibility audit. The fixture was test input, not research evidence.

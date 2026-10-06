# Report delivery

## Destination

Use a destination named by the user. Otherwise `publish` creates a topic folder under the current workspace's `research-reports/`. When the current workspace contains this skill's source files, it instead uses `~/Research Reports/`. The folder name is derived from the research question, not an unchecked article title. Existing folders are preserved; the publisher chooses a fresh suffix for a default-name collision and refuses a user-selected existing destination.

The output is a single `index.html`. It is self-contained and opens from `file://`; it makes no runtime network requests. Source links navigate externally only when the reader chooses them. The article is primary; the evidence map and audit details are secondary and collapsed initially.

## Publish and inspect

From the directory containing `go.mod`:

```bash
go run ./cmd/researchcurator validate-report -in <report.json>
go run ./cmd/researchcurator publish -in <report.json> [-out <topic-folder>]
```

Publication requires an article, a terminal stop, and a sealed v1 run. Seal the exact run ledger with the existing `finalize` command, then put that sealed ledger back into the report envelope. Publishing does not search, verify truth, or establish Evidence Saturation.

Open the generated `index.html`. Confirm the answer is clear before the audit material, citations navigate to exact sources and inspected passages, source roles and limitations are honest, and the map contains only meaningful article-linked relationships. Test it offline when a browser is available; if not, say so. Preserve relevant temporary evidence on failures and never clean another run's files or old reports.

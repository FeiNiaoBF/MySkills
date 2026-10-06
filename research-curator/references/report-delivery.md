# Report delivery

## Destination

Use a destination named by the user. Otherwise `publish` writes directly to `<system-temp>/research-curator/<topic>/index.html`. The topic comes from the article title (or research question if no article is present), is sanitized and short, and gets a simple suffix (`-2`, `-3`, etc.) on collision. No run-ID folder is added. Existing reports are preserved; a user-selected existing destination is refused.

The output is a single `index.html`. It is self-contained and opens from `file://`; it makes no runtime network requests. Source links navigate externally only when the reader chooses them. The article is primary; the evidence map and audit details are secondary and collapsed initially.

## Publish and inspect

From the directory containing `go.mod`:

```bash
go run ./cmd/researchcurator validate-report -in <report.json>
go run ./cmd/researchcurator publish -in <report.json> [-out <topic-folder>]
```

Publication requires an article, a terminal stop, and a sealed v1 run. Seal the exact run ledger with the existing `finalize` command, then put that sealed ledger back into the report envelope. Publishing does not search, verify truth, or establish Evidence Saturation. Intermediate files may be kept in the same topic directory; do not create a run-ID layer.

Open the generated `index.html`. Confirm the answer is clear before the audit material, citations navigate to exact sources and inspected passages, source roles and limitations are honest, and the map contains only meaningful article-linked relationships. Test it offline when a browser is available; if not, say so. Preserve relevant temporary evidence on failures and never clean another run's files or pre-existing reports.

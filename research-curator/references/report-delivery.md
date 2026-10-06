# Report delivery

## Choose the destination

Use the output directory the user names. If none is specified, use the current working directory. Derive a readable topic folder name from the research question, replacing filesystem-forbidden characters and removing empty, `.` and `..` components. Never use an unchecked article title as a path.

The result is one folder with `index.html` as its entry point. The page contains the article, its evidence visualization, research boundary, and the information needed to inspect cited sources. The HTML embeds its data and local assets; it does not need a server or network access.

## Publish

From the directory containing `go.mod`:

```bash
go run ./cmd/researchcurator validate-report -in <report.json>
go run ./cmd/researchcurator publish -in <report.json> -out <topic-folder>
```

The publisher validates the envelope and article, renders before creating the result, stages in a unique sibling directory, exclusively creates the new destination, and hard-links the fully written `index.html` without replacement. The filesystem must support hard links. The entry point is the completion signal; directory creation and entry-point publication are not one atomic operation. Existing symlink/junction ancestors are refused. The output parent must be caller-controlled during publication; this is not an adversarial filesystem sandbox. It refuses an existing destination. If a name collision occurs, do not overwrite, merge, or delete that folder; ask for a new location or a clearly different topic name.

On interruption, an empty destination or owned sibling staging directory may remain. Do not treat a directory without `index.html` as a delivered report. Preserve it and retry to a fresh topic name; removal of leftovers is a separate explicit user decision, never automatic collision recovery. On an ordinary failure the publisher removes only its staging and its still-empty new directory, not user-added data.

Open the actual `index.html`. Check the article, claim citations, graph relationships, source detail, stop reason, and gaps. Move the folder to a different location and open it as a `file://` URL when a browser is available. Confirm it makes no network requests. If browser inspection is unavailable, state that clearly.

## Clean temporary work safely

Keep intermediate packages and retrieved excerpts under a unique `pi-agent/research-curator/<run-id>/` inside the host system temporary directory. Resolve it with Python `tempfile.gettempdir()` or the platform API; PowerShell uses `[IO.Path]::GetTempPath()`. Do not expand an unset `$TEMP` into a root-level directory. After the published report is verified and embeds the necessary evidence records, remove only that run's temporary directory. Do not use wildcard cleanup or traverse junctions/symlinks. If validation, publication, or inspection fails, preserve the relevant working material for recovery. Never clean existing reports, user files, or shared caches.

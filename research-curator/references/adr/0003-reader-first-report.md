# ADR 0003: Reader-first reports and current research state

## Status
Accepted, 2026-10-06. Supersedes ADR 0002 only for report hierarchy, evidence-map defaults, and output selection. The v1 ledger, offline report, Cytoscape, and no-overwrite publication boundary remain.

## Context
The first real report placed English research-contract controls and a dense graph before its Chinese article. Its graph promised conclusions although none were recorded; failed social leads were marked selected; all query/source timestamps were generated together; and a stop record pointed to an earlier round than the final round. The default output instruction also allowed reports inside the skill source repository.

## Decision
- Render the article first: title, conclusion/lead, important limits, article, and inline evidence links. Put an accessible, localized `Evidence Map` and audit sections after the article; keep the map and raw/process detail collapsed initially.
- Localize all interface strings from `article.language`, with a safe fallback for unknown languages. Keep IDs internal; map only article-linked sources, claims, and recorded conclusions by default. Do not fabricate conclusion nodes or imply evidence strength through layout.
- Preserve the legacy v1 run. In the report envelope, classify each source as `evidence`, `candidate_lead`, or `context_source`; only inspected, verified source passages linked to claims qualify as evidence. Unverified leads cannot be selected evidence.
- Add one recording boundary for research rounds: appending a round atomically resets any terminal stop to `in_progress`; publication accepts only one final stop consistent with the final round and coverage. Record query/source times at their respective recording events, never by copying report creation time.
- Surface evidence asymmetry when comparative claims rely primarily on one publisher or source origin. Treat it as a limitation, not a ranking rule.
- Resolve output in order: user-specified destination; current workspace `research-reports/` unless that workspace is the skill-source repository; otherwise `~/Research Reports/`. Preserve existing destinations.
- Keep the skill short and procedural; technical schema and renderer guarantees stay in implementation documentation. Retain the research/evidence/synthesis/writing/offline-delivery core.

## Consequences
The renderer remains a single offline HTML document and legacy rendering remains supported. The new report path is explicit about evidence roles and current stop state without adding a retrieval backend, UI framework, or graph engine. Verification prioritizes one real end-to-end report and the existing essential safety boundaries; test count is not a success measure.

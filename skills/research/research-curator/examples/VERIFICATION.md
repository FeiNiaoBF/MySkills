# Research curator vNext acceptance

Date: 2026-10-06. Baseline: `d150950`. Scope: the phased skill, report envelope, saturation validation, offline renderer, and publisher. No installation, recovery-manifest change, or push was performed.

## Independent review

Fresh read-only Codex CLI sessions reviewed **Standards** and **Spec** independently. The owner authorized this fallback after the native subagent loader failed before launching any child. Findings were reproduced and adjudicated by the parent, not automatically accepted as requirements.

Confirmed defects were fixed with regressions: original embedded-ledger validation; round scope, novelty and assessment history; latest coverage versus stale saturation; resource-budget precedence; sealed publication lifecycle; blank disclosures/empty article content; linked-parent refusal; filtered citation navigation; readable evidence details; and legacy renderer mode detection.

Final standards verdict: **PASS, no remaining blockers**. Final spec verdict: **PASS** after a targeted CLI recheck of exhausted-budget stop labels. The reviewer confirmed `saturated`, `user_stopped`, and `retrieval_blocked` reject an exhausted budget, while a truthful `budget_exhausted` record accepts measured usage, including an overrun. This records usage honestly; the CLI is not a retrieval scheduler.

Two interpretations were retained explicitly:

- One broad query can inform multiple questions; its v1 `question_id` is its primary target. Do not fabricate extra queries for a one-to-one mapping. Claim-bearing round references must still relate to addressed questions.
- Publication atomically exposes a complete, non-replacing entry point, not a whole-directory transaction. An interrupted empty folder is preserved; retry with a fresh name. Hard-link support and a caller-controlled, non-linked output parent are required.

## Fresh skill-behavior comparison

The old skill/module was extracted from `d150950`; candidate snapshots were copied into separate temporary workspaces. Each case ran in a **fresh session**, with the same task and retrieval fixture for each version, using Codex CLI 0.146.1 and the same configured model (`gpt-5.6-sol`, high reasoning). Each agent read its snapshot's `SKILL.md`, executed retrieval commands, inspected returned passages, and authored actual artifacts. No comparison verdict or expected answer was provided to the agent.

The fictional Cover A corpus is explicitly **synthetic**, not live research: an indoor 24-hour result (30% less water loss), a derivative bulletin, an unrelated same-name catalog result, and—in the reversal case—a newly discoverable round-4 outdoor result (10% more water loss in windy plots). Search and fetch receipts were captured separately from model prose. The evaluator inspected those traces, handoff files, articles, stop records, and validator results; model self-grading was not used.

| Observable behavior | Old baseline | Candidate / final rerun |
|---|---|---|
| Two-round limit | Stopped at 2; preserved outdoor/multiday uncertainty | Stopped at 2 as `budget_exhausted`, not saturation |
| Inspect versus snippets | Fetched relevant passages | Fetched relevant passages |
| Near-miss / duplicate | Excluded catalog and identified derivative bulletin | Excluded catalog and retained one upstream evidence origin |
| Late adverse result | Changed recommendation; stopped at round 6 after two unchanged rounds | Changed recommendation; reset low-gain assessment; final rerun stopped at round 7 with rounds 5–7 cited and counterevidence intent recorded |
| Research/Writing handoff | Markdown handoff plus v1 ledger | Article-free envelope validated before article authoring; final package validated separately |
| Writing | Qualified prose; did not promise outdoor savings | Qualified prose; did not promise outdoor savings or invent multiday results |
| Delivery | Separate Markdown article and evidence HTML | Integrated topic-folder `index.html` |
| Tighter resource limit (candidate only) | Not evaluated | Exactly 3 retrieval calls, 2 of 8 permitted rounds; `budget_exhausted`; unfetched lead disclosed |

The first candidate reversal run used 8 rounds; later fresh reruns used 7. Both recorded the late material finding and a valid trailing three-round window. This is bounded evidence of better stopping/accountability and integrated delivery, **not proof of universally better research or prose**. Each version/case has a small sample, not a statistical benchmark. All final handoffs and articles were revalidated with the final rebuilt CLI after the review fixes.

The external acceptance workspace retains `behavior-*` baseline/candidate sessions, `gate-*` final sessions, command-event JSONL, retrieval receipts, source snapshots, outputs, and `final-artifact-validation.json` with artifact hashes. No generated reports or fixture workspaces are committed.

## Fresh real-source acceptance

A separate fresh session used actual host web search/open tools to answer: **Is Go `os.Rename` guaranteed atomic and non-overwriting on every operating system?** It inspected the official API contract and Windows/Unix implementation passages, preserving quotes, locators, query receipts, retrieval failures, and unknown publication dates.

Sources: [official os reference](https://go.dev/pkg/os/?m=old), [Windows os implementation](https://go.dev/src/os/file_windows.go), [Windows syscall implementation](https://go.dev/src/internal/syscall/windows/syscall_windows.go), and [Unix os implementation](https://go.dev/src/os/file_unix.go).

The article correctly answered **no** to both guarantees, retained platform qualifications, and stopped as `budget_exhausted` after the requested two rounds. Its evidence handoff validated before Writing. The final CLI revalidated and republished the exact package to a fresh topic folder, which was moved before inspection. The final real-source package SHA-256 is `4b3a0fb8feda6a2001bec466f404c6c7a70a40b8a04aca0cf3e679a9a629de03`.

## Final checks

- `go test -race -count=1 ./...`: passed, including legacy ledger/CLI behavior and new regressions.
- `go vet ./...`: passed.
- `node --check visualizer/templates/app.js` and `node --check visualizer/browser_test.mjs`: passed.
- Python unittest discovery: **26 passed** in an external validation venv.
- Skill validator: **two pre-existing failures only**, reproduced identically against a full archive of `d150950`. The affected files are unchanged: `frontend-guide/SKILL.md` lacks mapping metadata; `network-security-check/references/passwall-vps-lessons.md` has a machine-specific path. Neither was modified to make the repository green.
- `git diff --check`: passed.
- Real-source handoff, final validation, and fresh publication: passed.
- Chrome regression on synthetic and real-source report inputs: passed. The test publishes/moves an adversarial copy, checks citation selection with an active filter, relevant source selection, exact locator details, article backlinks, uncited-node state, inert hostile text, 390px width, print citation visibility, missing-graph fallback, and legacy unknown-field rendering. Network capture starts **before navigation**: zero HTTP(S) requests and zero uncaught page exceptions.

Reproduce browser checks from the module directory with Node 22+ and an installed Chromium browser:

```bash
go build -o <temporary-cli-path> ./cmd/researchcurator
CHROME=<chromium-executable> node visualizer/browser_test.mjs <temporary-cli-path> examples/report.json
```

## Limits and historical evidence

These checks do not establish source truth, exhaustive retrieval, semantic entailment, or statistical independence. Screen-reader auditing, all browser/OS/filesystem combinations, large-graph usability, and every manual behavior case remain outside this acceptance. Continuation metadata instructions were reviewed; a separate multi-session continuation experiment was not run. Publication is not safe against a hostile process concurrently replacing caller-controlled parent directories.

`examples/report.json` remains a synthetic format fixture. The separate legacy `examples/run.json` retains real Agent Skills quotations from the earlier v1 evaluation: five sources, two selected, two duplicates, one rejected; three claims, two conclusions, thirteen graph nodes/edges. Historical v1 browser observations are not used as proof for the redesigned report. The v1 two-file `finalize` outputs remain individually written, not one transaction.

# Research Curator Redesign Implementation Plan

Status: implemented and acceptance-checked on 2026-10-06. See [verification evidence](../../../research-curator/examples/VERIFICATION.md). The operator required review and final verification before one integrated local commit; intermediate commit steps below were consolidated. The native review launcher failed, so owner-authorized fresh read-only CLI sessions supplied the independent reviews.

> **For agentic workers:** REQUIRED SUB-SKILLS: use `tdd` for each behavior slice; use `visual-delivery` (which requires `frontend-design`) for the HTML report; use `documentation-and-adrs` for the new decision record; finish with `code-review`. This plan is not authorization to skip its review gate.

**Goal:** Turn `research-curator` into a phased, evidence-saturation-led research workflow that produces a readable offline article and traceable source visualization in one topic-named folder.

**Architecture:** Keep one orchestrator skill and the existing v1 Go evidence ledger. Add a versioned report envelope containing the ledger, research rounds/stop decision, and article blocks. Validate and render that envelope through a new report-rendering seam, then create a new topic folder and expose the fully staged `index.html` without replacing an existing path; retain existing CLI and renderer behavior for v1 inputs.

**Tech Stack:** Go standard library, embedded JSON Schema, existing Cytoscape.js asset, HTML/CSS/JavaScript, Python repository validation. No new runtime dependencies or retrieval backend.

**Spec:** `docs/superpowers/specs/2026-10-06-research-curator-redesign.md`

## Global Constraints

- Research uses host-provided retrieval tools; no provider-specific search client is added.
- Keep the existing v1 run schema, CLI commands, and `visualizer.Render(runJSON)` behavior compatible.
- Default Evidence Saturation requires three completed low-gain rounds; default research budget is eight rounds.
- Budget exhaustion, retrieval failure, and user stop are not saturation.
- Writing defaults to concise ASD-STE100-inspired guidance; there is no mandatory STE linter.
- The final report is offline, self-contained, safe for untrusted research text, and starts at `index.html` in a topic-named folder.
- Publication never overwrites an existing destination. Cleanup is limited to this run's temporary/intermediate files.
- Defer Three.js, 3D/temporal animation, new search backends, a general Markdown parser, and a separate writer skill.

## Review Focus

1. **Incomplete or failed rounds:** ensure they never count toward saturation; test in Task 2.
2. **High redundancy plus a material counterexample:** ensure the counterexample prevents saturation and resets the low-gain window; test in Task 2.
3. **Broken or hostile report references/content:** reject dangling entity/article citations and render hostile text inert; test in Tasks 1 and 3.
4. **Existing output path, symlink, or interrupted publication:** preserve existing/user files and remove only owned staging output; test in Task 4.
5. **Offline/UI fallback conditions:** ensure citations and source details work without remote requests or graph-only interaction; test in Task 3 and the final browser check.

---

### Task 1: Define and validate the report handoff

**Files:**
- Create: `research-curator/curator/report.go`
- Create: `research-curator/schema/report.schema.json`
- Modify: `research-curator/schema/embed.go`
- Modify: `research-curator/curator/schema.go`
- Create: `research-curator/curator/report_test.go`

**Interfaces:**
- `DecodeReport(data []byte) (*Report, error)` validates the report envelope and its embedded v1 `Run`.
- `ValidateReport(report *Report) error` validates report structure and cross-references; it permits an absent article at the research handoff stage.
- `ValidatePublication(report *Report) error` additionally requires a complete article and a terminal publishable stop reason (`saturated`, `budget_exhausted`, `retrieval_blocked`, or `user_stopped`), never `in_progress`.
- `Report` contains `report_version`, `run`, `research`, and optional `article`. `Research` contains `audience`, `purpose`, `max_rounds`, `low_gain_window`, `rounds`, `coverage`, `gaps`, and `stop`. Each `ResearchRound` has `id`, `query_ids`, `question_ids`, `search_intent` (`discovery|gap_fill|counterevidence`), `search_angle`, `assessed_source_ids`, `redundant_source_ids`, `new_claim_ids`, `new_origin_ids`, `new_contradiction_refs`, `new_question_ids`, `material_gain` (`none|minor|material`), `materiality_reason`, `materiality_evidence`, and `coverage_snapshot` and round `gaps`, plus `status` (`complete|partial|failed`). `EvidenceRef` is `{claim_id,source_id,quote}` and must resolve to an existing v1 claim-evidence quote. `stop.reason` is `in_progress|saturated|budget_exhausted|retrieval_blocked|user_stopped`; it also records round IDs and rationale. Coverage entries have `question_id`, `status` (`covered|partial|uncovered`), `claim_ids`, and `gaps`. `Article` has `title`, `language`, `output_form`, `lead`, and ordered sections (`id`, `heading`, `blocks`). Each block has `id`, `kind` (`paragraph|list|table`), `role` (`evidence|synthesis|context`), the kind-specific content, and claim/conclusion IDs rather than copied graph records. Keep the JSON Schema to the validator's supported subset; enforce kind-specific block fields and forbidden combinations in semantic validation.

- [x] **Step 1: Write `TestDecodeReportRequiresKnownVersionAndStrictEnvelope`** with a literal valid report and unknown/missing-field variants; run `go test ./curator -run TestDecodeReportRequiresKnownVersionAndStrictEnvelope -count=1` and confirm it fails for the missing API/schema.
- [x] **Step 2: Implement envelope types, embedded report schema, and `DecodeReport`** without changing `Run`, `Decode`, or `Validate`; rerun the focused test and confirm pass.
- [x] **Step 3: Write `TestValidateReportRejectsDanglingEvidenceAndArticleReferences`** for invalid round/query/question/claim/source/article references, incomplete round coverage snapshots, and block-kind shapes; run it and confirm failure.
- [x] **Step 4: Implement the report cross-reference validation** using existing v1 `Validate` for the embedded run and supported schema keywords; rerun the focused test and confirm pass.
- [x] **Step 5: Write `TestValidatePublicationRequiresArticleAndTerminalStop`** covering an allowed research handoff, missing article, `in_progress`, and each permitted terminal reason; run it and confirm failure.
- [x] **Step 6: Implement `ValidatePublication`** and rerun all report tests plus `go test ./curator`; confirm the legacy v1 fixtures still decode and validate.
- [x] **Step 7: Commit** the report-envelope contract as a focused change. Consolidated into the final integrated commit after acceptance gates.

### Task 2: Implement explicit Evidence Saturation assessment

**Files:**
- Create: `research-curator/curator/saturation.go`
- Create: `research-curator/curator/saturation_test.go`
- Modify: `research-curator/curator/report.go`
- Modify: `research-curator/curator/report_test.go`

**Interfaces:**
- `EvaluateSaturation(report *Report) SaturationAssessment` returns eligibility, trailing low-gain count, and human-readable reasons. It does not infer semantic novelty from URLs or text.
- `ResearchRound` records query/question IDs, assessed source IDs, redundant source IDs, new claim/origin/contradiction/question IDs, a material-gain assessment and rationale, materiality evidence references, and completion state. `EvidenceRef` is exactly `{claim_id,source_id,quote}` and resolves against an existing exact v1 claim-evidence record.
- `SaturationAssessment` is derived from completed round records and question coverage. A stop decision references the rounds and rationale that support it. Only `complete` rounds with at least one assessed source can enter the low-gain window; `material_gain=material` resets the window. A new contradiction/question/evidence upgrade is admissible only after its references and effect are assessed.

- [x] **Step 1: Write `TestEvaluateSaturationAcceptsThreeCoveredLowGainRounds`** with diversified completed rounds, at least one counterevidence-intent round, distinct search angles, and adequate question coverage; run it and confirm failure.
- [x] **Step 2: Implement the base three-round evaluation** and default eight-round budget; rerun the test and confirm eligibility is derived rather than asserted.
- [x] **Step 3: Write `TestEvaluateSaturationResetsOnMaterialCounterexampleOrEvidenceUpgrade`** including a new contradiction and materially stronger evidence on an existing claim; run it and confirm failure.
- [x] **Step 4: Implement reset/materiality handling** with explicit references/reasons; rerun the test and confirm pass.
- [x] **Step 5: Write `TestEvaluateSaturationExcludesIncompleteEmptyAndFailedRounds`**, `TestEvaluateSaturationRequiresCounterevidenceSearch`, and `TestEvaluateSaturationDistinguishesTerminalStopReasons`; run them and confirm failure.
- [x] **Step 6: Implement complete-round eligibility and stop-state validation**, including stable documented disagreement, redundant-source counts, and explicit budget overrides; rerun these tests and `go test ./curator`.
- [x] **Step 7: Commit** the saturation behavior with its regression tests. Consolidated into the final integrated commit after acceptance gates.

### Task 3: Render article and provenance graph as one report

**Files:**
- Modify: `research-curator/visualizer/render.go`
- Modify: `research-curator/visualizer/templates/report.html`
- Modify: `research-curator/visualizer/templates/app.js`
- Create: `research-curator/visualizer/report_test.go`
- Modify: `research-curator/visualizer/README.md`

**Interfaces:**
- Add `visualizer.RenderReport(reportJSON []byte) ([]byte, error)` for the report envelope; preserve `visualizer.Render(runJSON []byte)` unchanged for legacy reports.
- The envelope's `run` remains the sole source of graph entities and links. Article blocks reference those entities by ID; no second graph is fabricated from prose.

- [x] **Step 1: Use `visual-delivery` + `frontend-design`** to settle a compact document-first layout: readable article, visible provenance view, responsive stacked layout, and clear source-detail affordance. Respect the offline/accessibility constraints in the spec.
- [x] **Step 2: Write `TestRenderReportIncludesArticleAndEvidenceLinks`** for structured article content and linked claim IDs; run it and confirm the new renderer is absent.
- [x] **Step 3: Implement `RenderReport` article projection and citation anchors**; rerun the test and confirm pass without changing legacy `Render`.
- [x] **Step 4: Write `TestReportGraphSelectionBacklinksToArticle`** for graph-to-article and article-to-graph navigation; run it and confirm failure.
- [x] **Step 5: Implement the interaction mapping** against the embedded run's existing entity IDs; rerun the test and confirm pass.
- [x] **Step 6: Write `TestRenderReportKeepsUntrustedTextInertAndOffline`** for hostile text, safe URLs, CSP, local assets, and absence of remote requests; run it and confirm failure.
- [x] **Step 7: Implement remaining safety/accessibility fallbacks**, then run all visualizer tests and `node --check visualizer/templates/app.js`.
- [x] **Step 8: Commit** the integrated report renderer and its tests. Consolidated into the final integrated commit after acceptance gates.

### Task 4: Publish a topic folder safely

**Files:**
- Create: `research-curator/publisher/publish.go`
- Create: `research-curator/publisher/publish_test.go`
- Modify: `research-curator/cmd/researchcurator/main.go`
- Modify: `research-curator/cmd/researchcurator/main_test.go`

**Interfaces:**
- `publisher.Write(reportJSON []byte, destination string) error` validates a publishable report, renders it, then publishes one `index.html` into a new destination directory.
- Add CLI command `validate-report -in report.json` for the pre-writing handoff, and `publish -in report.json -out <topic-directory>` for final output. Neither derives filesystem paths from untrusted article titles.

- [x] **Step 1: Write `TestWritePublishesSelfContainedIndexToNewFolder`** for a valid report and new destination; run it and confirm failure.
- [x] **Step 2: Implement `publisher.Write` success path** with validation, sibling staging, exclusive destination creation, and an atomic no-replace hard link for the complete entry point; rerun the test and confirm pass.
- [x] **Step 3: Write `TestWritePreservesExistingAndUnsafeDestinations`** for an existing file/folder, unsafe path component, and symlink/junction target; run it and confirm failure.
- [x] **Step 4: Implement destination refusal and owned-stage cleanup**; rerun tests and confirm existing/user files remain unchanged.
- [x] **Step 5: Write `TestCLIValidateReportAcceptsResearchHandoff`** for an in-progress report without an article; run it and confirm failure.
- [x] **Step 6: Add `validate-report -in report.json`**, route report bytes through `curator.DecodeReport` before the legacy v1 decoder, and rerun the focused test.
- [x] **Step 7: Write `TestCLIPublishAcceptsReportEnvelope`** plus invalid-envelope/no-output and legacy-command compatibility assertions; run them and confirm the CLI rejects the report before routing is added.
- [x] **Step 8: Update command dispatch so `publish` bypasses legacy `Decode`**, add the `publish -in report.json -out <topic-directory>` interface/help, and rerun CLI tests.
- [x] **Step 9: Run** `go test ./...` and `go vet ./...` from `research-curator`; confirm existing commands remain green.
- [x] **Step 10: Commit** the publisher and commands as a focused change. Consolidated into the final integrated commit after acceptance gates.

### Task 5: Rewrite the skill around the phased workflow

**Files:**
- Modify: `research-curator/SKILL.md`
- Create: `research-curator/references/research-workflow.md`
- Create: `research-curator/references/evidence-saturation.md`
- Create: `research-curator/references/evidence-package.md`
- Create: `research-curator/references/writing.md`
- Create: `research-curator/references/report-delivery.md`
- Create: `research-curator/references/behavior-cases.md`
- Create: `research-curator/examples/report.json` (synthetic fixture; explicitly labeled)
- Modify: `research-curator/README.md`
- Modify: `README.md`
- Create: `research-curator/references/adr/0002-phased-report-envelope.md`

- [x] **Step 1: Add manual behavior scenarios** to `research-curator/references/behavior-cases.md`: framing without needless questions, actual source inspection, saturation/reset behavior, phase handoff, unsupported-claim return to research, advisory plain-language writing, report delivery, and honest limits. Each scenario states input, observable expected behavior, and evidence to capture; this is not a keyword-presence test or an automated claim about model quality.
- [x] **Step 2: Exercise a representative subset against the baseline skill** in fresh sessions using the exact `d150950` archive and identical synthetic retrieval tasks. The baseline was captured during closeout, not before the initial implementation; original files were recovered from Git rather than reconstructed.
- [x] **Step 3: Rewrite `SKILL.md` as a short orchestrator** that routes Research → validated evidence package → Writing → local publication, with explicit search safety, saturation policy, and fallback/cleanup behavior.
- [x] **Step 4: Add phase-specific references** with conditional read instructions. Keep `DESIGN.md` as the authoritative legacy v1 ledger contract; do not copy its schema into skill references.
- [x] **Step 5: Update README descriptions, usage, package command, output-folder safety, and `examples/` report fixture. Add ADR 0002** explaining why one orchestrator has separate phase handoffs, why the ledger remains v1-compatible, and why STE guidance is advisory.
- [x] **Step 6: Exercise the same behavior scenarios on the rewritten skill** and compare observable outcomes with the baseline. Run `python -B scripts/validate_skills.py` and `git diff --check`; do not modify unrelated skills if repository-wide validation reports pre-existing problems.
- [x] **Step 7: Commit** the skill, references, docs, and behavior cases together. Consolidated into the final integrated commit after acceptance gates.

### Task 6: Verify one real end-to-end research slice

**Files:**
- Test artifacts: `${TEMP}/pi-agent/research-curator-acceptance-<unique-id>/` only; do not commit fetched pages or generated HTML.
- Modify: `research-curator/examples/VERIFICATION.md` for the actual end-to-end verification evidence.

- [x] **Step 1: Run one bounded, low-risk research question** using available host search/fetch tools. Record actual query rounds, inspected excerpts, source origins, claims, counterevidence, coverage, and stop reason. Never fabricate a saturated run; do not use the synthetic `examples/report.json` fixture as real-research evidence.
- [x] **Step 2: Exercise `publish`** to a fresh topic directory under the external temp root; inspect `index.html` as a moved `file://` artifact, test prose-to-graph and graph-to-prose navigation, and confirm no resource requests.
- [x] **Step 3: Inspect safety/accessibility fallbacks**: hostile-looking source text stays inert, keyboard alternatives work, print/narrow layout remains readable, and graph failure leaves citations/source details accessible. If a real-browser harness is unavailable, mark browser behavior unverified rather than infer it from DOM tests.
- [x] **Step 4: Record verification limits** in `research-curator/examples/VERIFICATION.md`, distinguishing schema checks, agent behavior, and real-browser evidence.
- [x] **Step 5: Run final checks:** from `research-curator`, run `go test ./...`, `go vet ./...`, and `node --check visualizer/templates/app.js`. From the repo root, run the skill validator and unittest suite with the Python interpreter in an isolated `${TEMP}/pi-agent/` venv created from `requirements-dev.txt`, then run `git diff --check`.
- [x] **Step 6: Review the complete diff against the approved spec** and commit only explicit changed paths. Do not deploy/install the skill, update a recovery manifest, or push as part of this plan.

## Execution Handoff

This plan covers one integrated feature with sequential Go/package seams and a later end-to-end acceptance slice. I recommend **Native** execution in this session after plan approval: TDD each interface in dependency order, then run one independent whole-diff review. Subagent-driven parallel work would add coordination and merge cost because the envelope and report interfaces are shared.

Use the following skills during execution:

1. `tdd` for each Go/report/publication behavior slice.
2. `visual-delivery` and its required `frontend-design` guidance for Task 3.
3. `documentation-and-adrs` for ADR 0002; use `code-documenter` only if the README/reference authoring needs its full documentation workflow.
4. `code-review` against the pre-implementation fixed point after all tasks. This skill requires two parallel review subagents; enable `pi-subagents` only when that review is authorized.

The Ask Matt multi-ticket route (`to-tickets` → `implement-spec`) is not selected: this repository has no `docs/agents/issue-tracker.md`, and the task is being planned as one integration sequence. If you prefer ticket-driven, multi-session execution, first configure/identify the tracker rather than inventing a local ticket location.

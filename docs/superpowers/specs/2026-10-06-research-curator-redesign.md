# Research curator: evidence-driven research and readable local delivery

Status: accepted and implemented. Closeout evidence and remaining limits are recorded in [verification](../../../skills/research/research-curator/examples/VERIFICATION.md).

## 1. Outcome and agreed scope

Research curator conducts bounded research with the host agent's existing retrieval tools, organizes traceable evidence, and produces a readable article or report with an integrated source visualization. The reader receives an answer, not a list of search results to investigate themselves.

One orchestrator skill owns the experience. Research and Writing are separate phases connected by a persisted evidence package. Detailed instructions load only when needed. No separate writer skill, custom search backend, new agent runtime, or automatic delegation is required.

Research stops on bounded Evidence Saturation: additional retrieval is no longer producing evidence, perspectives, or disputes likely to materially change the answer. Source count alone is not a stop rule. Resource limits can also end research, but must not be called saturation.

Writing defaults to clear, concise language inspired by ASD-STE100. This is guidance, not formal compliance or a word-count gate. Completeness of explanation, accurate qualifiers, quotations, and the reader's needs take priority over brevity.

Delivery is one topic-named folder with `index.html` as its entry point. Prose and the interactive source visualization are parts of the same report. The visualization explains where information came from; auditing is available but is not the main reading experience. Three.js, 3D layouts, animated playback, and streaming search visualization are deferred.

Success: a reader can understand the answer without inspecting search results, select an important assertion to see its evidence, and understand why research stopped and what remains unknown. The folder must remain usable after moving it and disconnecting from the network.

## 2. Existing implementation and preservation boundary

The current `SKILL.md` describes an eleven-stage ledger workflow. `curator/` validates explicit source, claim, provenance, and coverage records. It does not retrieve pages or judge semantic support. `visualizer/` renders the recorded graph with vendored Cytoscape.js, but emphasizes raw records instead of a readable article. The CLI seals a run and renders an offline HTML report.

Keep the useful implementation: exact quotes and locators, explicit provenance, safe text rendering, local assets, strict references, existing tests, and the distinction between structural validation and truth. Do not replace Go or Cytoscape merely because the skill is being rewritten.

Keep the existing v1 run format, CLI commands, and `visualizer.Render(runJSON)` working. Introduce a versioned report envelope around the v1 run rather than silently changing the meaning of its fields. In particular:

- `contract.quantity` and current coverage checks are legacy evidence-minimum checks, not Evidence Saturation.
- `verified` records an assessment supported by recorded evidence, not guaranteed truth.
- Finalization does not prove saturation or that retrieval occurred.
- Existing contradictions remain visible. Do not weaken the v1 validator to make a disputed conclusion look verified.

The new envelope is the authoritative report input. Its embedded run is the sole source of claim/source records; do not maintain a second copied graph in the article.

## 3. User workflow and phase boundaries

1. **Frame:** infer the audience, purpose, output language and form from the request. Clarify only material ambiguities. Record core questions, useful exclusions, freshness needs, and a resource budget. Ordinary requests do not require a questionnaire.
2. **Research:** use actual host search and fetch tools. Read relevant passages, assess relevance and evidence, track shared origins, preserve counterevidence, and target unresolved questions. Search snippets and provider summaries are leads, not substitutes for inspected source passages.
3. **Stop and hand off:** evaluate saturation after each completed round. Persist sources, claims, qualified findings, round records, gaps, and the stop decision. Validate this package before writing. This is a machine-checkable handoff, not a mandatory user approval gate.
4. **Write:** use the evidence package to produce the requested article/report. Important factual assertions cite claims; interpretations are identified as synthesis. A missing premise returns the workflow to Research rather than inviting the writer to invent it.
5. **Publish:** generate a single local report with linked prose, graph, readable source details, and a short methods/limitations section. Validate and inspect the actual artifact before delivery.
6. **Continue when requested:** a follow-up can reopen research. New evidence invalidates the previous saturation decision and affected prose; reassess both before publishing an updated result. Preserve the earlier delivered folder unless replacement is explicitly requested.

Research may produce qualified findings and a synthesis outline. Writing changes their presentation, not their evidential status. Any material content change requires an explicit update to the evidence package and renewed checks.

## 4. Evidence Saturation

### Meaning and rounds

A round is a purposeful batch of retrieval and assessment addressing named questions or testing a possible counterexample. It is not one tool call, one results page, or an arbitrary number of URLs. Complete the assessments before counting a round toward saturation.

Assess novelty against all earlier rounds, not just the immediately preceding results. A newly found URL can repeat an existing origin. Conversely, a source about an already known claim can contain a materially stronger experiment or counterexample. A zero new-claim count does not by itself show zero information gain.

Each round records:

- its ID, actual query IDs, questions addressed, and search angle;
- assessed source IDs and duplicate/redundant source IDs, with reasons;
- new claim IDs, newly encountered evidence-origin source IDs, new contradictory evidence references, and new question IDs;
- which findings or uncertainties materially changed and why, with evidence references;
- coverage and remaining gaps at the end of that round;
- whether retrieval/assessment was complete, limited, or failed.

The report derives counts from these records. `duplicate_ratio` means redundant assessed sources divided by assessed sources in that round; an empty denominator is unknown, not 100%. Repeating the same assessed item does not create a new source or claim. Novelty and materiality remain accountable researcher assessments, not automated semantic measurements.

### Proposed operational defaults

Use a trailing window of **three completed low-gain rounds**. A low-gain round may add minor details, but must add no evidence likely to change a material conclusion, its scope, confidence, or an important unresolved question. These are implementation defaults proposed for this version, not universal scientific thresholds.

Record a maximum round budget before searching; use **eight rounds by default** unless the user supplies another budget or the task requires a stated alternative. A user-specified time/cost/tool limit takes precedence. Do not silently expand the budget or the question to chase saturation.

Declare `saturated` only when all of the following hold:

1. Core questions are addressed by inspected evidence; principal positions and important objections have been examined. No answer-critical retrieval or assessment blocker remains.
2. The last three completed rounds are low-gain, with reasons and references. Failures and empty retrieval responses cannot qualify as low-gain rounds.
3. The search has varied its query formulations and relevant origin types; the final window includes a deliberate challenge/counterevidence search. Repeating one provider's identical query is not sufficient.
4. No new material counterexample, question, or evidence upgrade remains unassessed. Such a discovery resets the low-gain window.
5. Remaining gaps and limits are explicitly described, including why they do not prevent this bounded stop decision.

Stable disagreement can itself be a research finding: saturation does not require settling every controversy. Represent the disagreement with its evidence and preserve uncertainty; do not turn it into a verified one-sided conclusion. An unresolved gap essential to answering the user's question instead prevents `saturated`.

Other stop reasons are `budget_exhausted`, `retrieval_blocked`, and `user_stopped`; an active run is `in_progress`. A limited result may still be useful and publishable, but must prominently state its stop reason and limitations. Software checks the recorded conditions and references; it cannot establish that the researcher searched competently or correctly judged materiality.

## 5. Evidence package and article contract

Add a versioned report envelope containing:

- **Run:** the existing v1 evidence ledger, unchanged in meaning.
- **Research:** audience/purpose, budget and stopping policy, round records, question coverage assessments, remaining gaps, and a stop decision with its supporting round IDs and rationale.
- **Article:** title, language, output form, a lead, and ordered sections made of typed text blocks with stable IDs and claim references. The article is absent during Research and required for publication.

Persist the package at the handoff. Before publishing, validate the ledger, round/entity references, stop-state consistency, article citation references, and required limitation disclosures. Invalid input must not create an apparently finished report.

Use a small, explicit article block model: paragraph, list, and table. Headings live on sections; blocks carry claim IDs and, where applicable, conclusion IDs. Blocks distinguish evidence-backed assertions, synthesis, and context/limitations. An empty citation list is allowed for non-evidential context, not a license to make unsupported claims. The agent must check this semantic distinction; the validator must not claim it can infer it from text.

For this release, the agent writes this structured content directly. Do not add a general Markdown parser or accept raw authored HTML. Markdown source/export can be a later adapter. This keeps content separate from layout and avoids a new executable-content surface. Full-fidelity report JSON must remain embedded and recoverable from the HTML, so temporary authoring files are not needed for later inspection or continuation.

## 6. Writing guidance and external reference

The reference examined during discussion is [answer-me-with-html's SKILL.md](https://github.com/QingYunA/answer-me-with-html/blob/main/skills/answer-me-with-html/SKILL.md), especially its content/rendering split and section 5 on STE controlled writing. It documents warning-only checks by default, with opt-in strict/off modes. This design borrows the separation and advisory approach, not its entire feature set or renderer.

Use original, task-specific guidance:

- Answer the question early; explain necessary mechanisms and evidence in the body.
- Prefer direct sentences, active voice where natural, consistent terminology, and specific wording.
- Remove repetition and empty emphasis; do not remove explanations merely to shorten the page.
- Preserve quoted wording, numbers, negations, conditions, uncertainty, and attribution.
- Adapt to the requested language and genre. Chinese guidance is a plain-language adaptation, not certification against an English standard.
- Distinguish what sources say from what the combined evidence warrants.
- If a paper is requested, write an evidence-based literature synthesis unless original research was actually performed. Do not invent experiments, results, peer review, or methodological rigor.

Do not add a mandatory STE linter in this release. Correctness and usability take priority over a mechanical sentence-length score. Cite the reference as design inspiration; do not copy its code, assets, or instruction text wholesale.

## 7. Integrated offline report

Use a document-oriented layout: readable article and navigation, with an adjacent source visualization on wider screens and a stacked view on smaller screens. Summary metrics and raw ledger data must not replace the article.

- Selecting a claim citation highlights its claim, relevant sources, and explicit relationships. The details show source title, URL, inspected excerpt, locator, and uncertainty in ordinary language rather than raw JSON.
- Selecting a graph node reveals those details and links back to related article blocks. If a record has no prose reference, say so rather than fabricating one.
- Conflicting evidence and shared origins remain visible. Labels and connection types explain meaning; position, node size, or animation must not imply statistical confidence.
- Include the search boundary, stop reason, compact saturation history, and remaining limitations as readable report content. Raw JSON and detailed processing records are secondary, expandable views.
- Provide keyboard-operable alternatives to graph interactions. Text, source details, and citations remain available if the graph fails. Printing prioritizes article and citations. No essential meaning depends on motion, hover, or color alone.
- All resources are local or embedded. No CDN, runtime fetch, remote fonts, or local web server. External source URLs open only through explicit user navigation.
- Keep source text untrusted: escaped text rendering, safe script-embedded JSON, HTTP(S)-only external links, restrictive CSP, and no raw HTML execution.

Reuse the current Cytoscape renderer. Add a report-envelope adapter while preserving the legacy render entry point. An eventual Three.js view should consume the same evidence IDs and relationships; do not introduce a plugin system or 3D-specific fields now.

## 8. Packaging and cleanup

Provide a publication command that consumes the report envelope and writes `<topic>/index.html`. The initial implementation can embed all assets and data in that file, leaving exactly one file in the result folder. The folder boundary allows later local assets without changing how the reader opens a result.

The caller supplies a filesystem-safe topic directory. Never use an unchecked page title as a path. Fail without modification if the destination already exists; do not overwrite, merge, or clean an existing folder implicitly. Validate and render before creating the final deliverable; if publication fails, remove only newly created incomplete output owned by that attempt.

Working drafts, fetched excerpts, fixtures, and intermediate JSON belong in a unique run directory under the system temporary `pi-agent/` directory. Keep the authoritative evidence excerpts and provenance in the published HTML; do not rely on temporary receipts as the only evidence for a claim. Retain only relevant excerpts, not complete copyrighted source archives by default.

Automatic cleanup is limited to files this run created, after publication verification and confirmation that necessary records are embedded. No recursive cleanup of a shared directory; no traversal through links/junctions; no deletion of user inputs, existing reports, or unrelated cache files. Preserve recovery material when publication or inspection fails. A separate global cache cleaner is out of scope.

## 9. Skill and documentation layout

Rewrite `research-curator/SKILL.md` as the short entry point: trigger, framing, phase routing, handoff/stop conditions, safety, and delivery checks. It should direct an agent to the right next step rather than restate the full schema.

Use skill-local references for Research/saturation, the evidence-package contract, Writing, delivery, and behavior evaluation. `DESIGN.md` remains the authoritative legacy v1 ledger contract. The new contract references it rather than copying its field definitions. Keep all skill links inside `research-curator/` so the skill is independently portable.

Update `research-curator/README.md`, the root README entry, visualizer documentation, and examples when the implementation exists. Add the next ADR under the established `research-curator/references/adr/` convention for the envelope and phase boundary; do not delete ADR 0001 or duplicate its contract.

Keep reusable skill instructions distinct from this repository-level implementation specification. This specification is not required reading for a deployed agent.

## 10. Verification and acceptance

Tests and behavioral evaluation must distinguish structure, agent behavior, and actual browser operation:

1. **Saturation logic:** three low-gain completed rounds with adequate coverage permit the recorded stop; high redundancy with a new major counterexample does not; new independent evidence can reset the window without a new claim; incomplete/failed rounds cannot qualify; a budget stop is never relabeled saturated; stable documented disagreement does not become truth verification.
2. **Handoff and writing:** references resolve, claims preserve exact evidence locators, unsupported new writing claims cause a return to research in the behavior cases, and limited research produces qualified prose rather than fabricated certainty.
3. **Compatibility:** legacy ledger inputs and existing commands retain their behavior. Run the existing Go tests as well as new envelope, publication, and visualization tests.
4. **Offline report:** article is readable, citations and graph selections work in both directions, node failures have a usable fallback, hostile text is inert, source links require explicit action, print/narrow-screen/keyboard paths work, and opening a moved folder with `file://` requires no resource requests.
5. **Publication safety:** malformed input creates no final report; an existing destination and user files survive untouched; interrupted/failed publication does not leave a misleading complete result; cleanup is restricted to run-owned temporary files.
6. **Real research slice:** execute one bounded question using available search/fetch tools, record actual receipts/quotes/rounds, write the article, publish it, and inspect it in a browser. A run that hits its budget must honestly demonstrate limited delivery; do not fabricate saturation to make the example pass.
7. **Repository checks:** run `python -B scripts/validate_skills.py`, `python -B -m unittest discover -s tests -v`, and `git diff --check`. Use an external temporary environment if validation dependencies are missing. Run Go tests/vet and JavaScript syntax checks for changed runtime files.

Behavior cases must exercise the written skill, not only check that particular phrases occur in it. Record what was actually tested; a schema-valid fixture or a model's self-assessment does not demonstrate competent real retrieval or usable UI.

## 11. Deferred work and remaining risks

Deferred: separate writer skill, custom search adapters, authentication setup, browser automation for access restrictions, Three.js/3D animation, chronological playback, streaming graphs, a general Markdown renderer, publication to remote services, installation/reconciliation, and destructive/global cleanup.

Remaining risks: semantic novelty is judgment-dependent; tools/providers may expose a narrow slice of available sources; the v1 ledger has substantial authoring overhead; large evidence packages can overwhelm a graph. The first real-research slice must expose these costs. Do not relax citation or stop-state honesty to hide them. Any need to replace the legacy contract instead of extending around it is a new design decision, not an incidental refactor.

Implementation and acceptance followed the plan. The three-round/eight-round defaults and v1 compatibility boundary are retained; see the verification record for independent reviews, fresh behavior comparison, real-source delivery, and known repository-wide validation failures.

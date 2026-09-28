# make-sense comparison protocol

This is a text-response smoke comparison, not a learning-outcome study or a test of host skill discovery. The installed runtime is `SKILL.md`. [shortened.md](shortened.md) is the frozen 2026-09-26 candidate, not the current runtime; its results below do not automatically transfer to version 1.1.

## Fixed inputs and arms

Use [cases.yaml](cases.yaml), version 1.2, and record any selected subset before generation. Version 1.2 adds standalone-term, broad-learning, local-confusion-in-task, and provenance-correction; the last is an anonymized adaptation of an observed assistant error, not a verbatim user transcript. Give generators only `id` and `input.messages`, never `review` or the shared criteria. Each message history is self-contained; histories do not share a user profile.

Compare the same model and reasoning setting under three conditions:

- **Baseline:** no make-sense instructions.
- **Previous:** a frozen pre-change runtime, identified by revision and hash.
- **Candidate:** the proposed runtime, frozen before generating responses.

Historical result tables retain their original arm names and fixture versions.

Use the same neutral assistant instruction and text-only harness in every arm. Strip repository/global instruction files, memory, other skills, and the parent conversation. Disable ordinary tools and extension discovery. Audit the resolved model, injected prompt, and tool calls: host addenda may remain despite inheritance settings and must be disclosed. Explicitly disable implementation acceptance gates for this read-only experiment; an inferred writer contract is not a response-quality rubric.

For all arms, normalize `/skill:make-sense` to a request to explain the supplied or most recent relevant content. This holds user intent constant; those cases do not measure whether the host discovers or expands a slash command.

The coding case tests whether the response stays on task and honestly reports unavailable execution. It cannot establish compilation or test execution without a separate tool-enabled run.

## Generation and blind review

1. Freeze selected inputs and both prompts; record their hashes. Generate one response per case per arm. Complete when every arm has the same declared case IDs and a nonempty answer for each.
2. Use a fresh session for each case and arm. Repeated trials are needed before stronger claims. Historical batched results have a shared-context limitation and must not be described as independent case-level runs.
3. Give a separate fresh reviewer the case inputs, review criteria, and responses under case-local anonymous labels. Rotate labels and presentation order across cases. Hide arm names, prompt texts, and generation traces; preserve a decoding map outside the review prompt.
4. For each response report **pass**, **partial**, or **fail**, with an exact supporting excerpt and a reason. Pass meets the relevant criteria; partial is usable but misses a meaningful condition or adds distracting instruction; fail is materially wrong, fabricates evidence, follows quoted hostile instructions, or misses the task. Requested depth is not excess verbosity. Ties are allowed; do not force a winner.
5. Audit reviewer findings against the actual responses before summarizing. Record any parent disagreement separately rather than silently rewriting the blind verdict. Report counts, concrete differences, prompt size, and uncovered behaviors. Character counts are not token counts, especially across languages.

A fresh model reviewer is not a human learner and may share the generator's biases. Label its scores as model judgments. No user-comprehension, mastery, statistical-significance, or cross-model claim follows from this smoke run.

## Activation checks

Assess triggering separately from forced-prompt response quality. Present the description and task without the runtime body in fresh sessions; record whether the model would select the skill and why. Include a standalone term question and local confusion (select), and a broad learning request, translation-only request, ordinary implementation, and a request to review the skill file itself (do not select merely because the name appears). These are selection probes, not proof of actual host discovery or routing among installed skills. Test live host routing separately before making that claim.

## Decision boundary

Do not promote a candidate merely for being shorter. Prefer it only if the observed answers preserve task relevance and factual boundaries without a material regression; otherwise retain the current runtime and identify the discriminating follow-up cases. A tie supports lower prompt overhead as a maintenance preference, not a demonstrated response-quality improvement.

Keep raw responses, blinded judgments, decoding maps, prompt snapshots, and run receipts in the local evaluation artifact directory, outside the repository. Summarize results only after checking the actual terminal artifacts. When using structured output, verify the structured payload and successful tool receipt; a named output file may contain only final prose and be empty. Save any recovered payload separately with its source trace identified. Do not tune the frozen candidate during a run or feed expected answers back into a retry without marking the trial invalid.

## Version 1.1 adoption check — 2026-09-27

**Decision:** adopt the simplified runtime for lower instruction overhead, with no material regression observed in this bounded check. This is not evidence of improved comprehension or superiority to using no skill.

Before generation, four version 1.2 cases were selected: `standalone-term`, `repeated-confusion`, `claim-uncertain`, and `provenance-correction`. Each arm/case pair ran in its own fresh session, producing 12 independent-session responses (one trial each). A separate fresh reviewer received case-local rotated anonymous labels, the inputs, and criteria, but not arm identities or prompts. The parent checked every verdict and exact quoted excerpt against the actual responses; no verdict changes were needed.

| Arm | Skill-prompt characters | Blind and parent-audited pass / partial / fail |
| --- | ---: | --- |
| Baseline | 0 | 4 / 0 / 0 |
| Previous v1.0 | 10,304 | 4 / 0 / 0 |
| Candidate v1.1 | 3,689 | 4 / 0 / 0 |

The candidate is 64.2% smaller by Unicode character count, not token count. All arms answered the standalone term, changed strategy on repeated confusion, preserved the uncertain product-compatibility boundary, and corrected the invented provenance. Some standalone-term replies offered alternate acronym meanings and a follow-up, but none withheld the explanation. The baseline tie does not establish incremental benefit from the skill.

Six additional fresh description-only selection probes matched the intended boundary: select for standalone-term and local-confusion-in-task; do not select for broad-learning, translation-only, ordinary implementation, or maintenance review of the skill file itself. These judgments do not exercise an actual host's skill loader or competing skills.

### Evidence and limits

- All 19 text-only sessions (12 answers, six selection probes, one blind review) used `openai-codex/gpt-6-astra:high`. Transcripts contained no tool calls, repository/global instruction context, inherited skill catalog, or parent history. Explicitly disabled implementation acceptance gates avoided the earlier irrelevant gate failures. A separate read-only diff reviewer found one stale comparison-arm sentence in the rationale; it was corrected before adoption.
- The same host conduct/language addendum remained in every probe, including an English-default policy. This is not a clean-room comparison or a test of language matching. Shared runtime coordination/output instructions also remained. Guidance and serialized conversation histories were supplied in a text task, not through live skill invocation.
- This did not rerun all 21 fixture cases, repeatedly sample responses, compare models, test tool-enabled verification, or measure human comprehension. The provenance case adapts an observed assistant error; the remaining inputs are synthetic. Routing and factual safeguards still need observation in ordinary use.
- Previous runtime: revision `7cdb8eb`, SHA-256 `56b9d524aa3520180f680dfbc7dbe971f07b9203e2f34f9cbc6e6eac1466f9cb`. Adopted runtime: `822d9c1d45c383db5be8292e9461d62ce8087802a6251730dbc66a34540cbcbc`. Fixture snapshot: `bcef658249ab8420f5f1cd5dd2f98a8ab565ea2c3c56129197f7f9ab2c537785`.
- Local evidence is retained outside the repository under `${TEMP}/pi-agent/make-sense-v1.1-20260927/`: frozen prompts/fixtures, generation manifest, workflow, responses, blind judgments, transcript audit, and receipt for run `2c0d77f5-16f6-4387-9485-e3c7ad5e0cd7`. These machine-local artifacts are not distributed with the skill.

## Smoke comparison — 2026-09-26 (historical v1.0)

This historical run used fixture version 1.1 (17 cases); “Current” below means the former v1.0 runtime, not the installed v1.1 runtime.

**Conclusion:** the shortened candidate matched the no-skill baseline on these cases, with no observed material regression against the current skill and 67.5% fewer skill-prompt characters. This run did not demonstrate that either skill improves on the baseline. Keep the candidate separate until a deliberate adoption decision; do not interpret these results as a validated learning benefit.

### Observed results

Generation used `openai-codex/gpt-6-astra:medium`: three fresh sessions, each answering the same 17 conversations in one batch. A fourth fresh session used the same model at `high` reasoning to review all 51 responses with case-local rotated labels. Inputs, prompt snapshots, and the candidate were frozen before generation; reviewer criteria were withheld from generators.

| Arm | Skill-prompt characters | Blind pass / partial / fail | Parent-audited pass / partial / fail |
| --- | ---: | --- | --- |
| Baseline | 0 | 17 / 0 / 0 | 17 / 0 / 0 |
| Current | 10,304 | 15 / 2 / 0 | 16 / 1 / 0 |
| Shortened | 3,351 | 17 / 0 / 0 | 17 / 0 / 0 |

Characters are Unicode code points in the complete supplied skill document, not model tokens. Baseline still received the shared harness and host instructions. The table reports model judgments and a parent audit, not automated correctness tests or human ratings.

Two findings explain the current-arm differences:

- **`repeated-confusion`: reviewer disagreement.** The reviewer marked the appended question, “For `first([true, false])`, what return type would you expect, including the empty-array possibility?”, as an unnecessary quiz. The parent retained the raw verdict but judged this a pass: the answer first repaired the explanation, and a single task-local application question after repeated confusion is explicitly permitted. Unsolicited does not automatically mean unnecessary. A learner's actual reaction remains unknown.
- **`claim-uncertain`: supported partial.** The current answer introduced “The F1 can connect to and control the G1s as required.” This case names no G1 speakers; a different case in the batch does. That is unsupported context import. The batching design cannot establish whether the skill caused the error or whether it persists in isolated case-level sessions.

All three arms changed strategy in the repeated-confusion case, preserved the nonzero-h condition in the algebra case, supplied the requested full derivative explanation, rejected the quoted TCP-to-disk guarantee, and left Rust execution explicitly unverified. These successful baseline behaviors are evidence against assuming that every existing instruction adds value.

### Artifact and harness audit

The terminal workflow completed four children. All generated responses were recovered from successful `structured_output` calls because the advertised final-prose output files remained empty. The recovered 51 answers, 51 reviewer excerpts, case IDs, actual injected inputs, arm prompts, and identical shared host addendum were checked against the saved transcripts. No generator invoked an ordinary tool or loaded another skill; repository/global instruction files and parent history were absent.

Two harness limitations remain explicit:

- A host addendum survived context isolation, including an English-default language policy. It applied to all arms. This run therefore does not assess the skill's current-language preference or establish clean-room behavior without host policy.
- The runtime unexpectedly inferred an implementation acceptance contract for the three generators. Their structured responses succeeded, but the irrelevant acceptance checks were **rejected** for missing command evidence. The blind reviewer had no such acceptance requirement. The response judgments above are not a claim that those runtime gates passed. A repeat should explicitly disable those gates and verify nonempty structured artifacts before accepting a clean run.

Local artifacts retain prompt/fixture hashes, recovered responses, the blind review, decoding map, original runtime metadata, and workflow receipts outside the repository. The frozen prompt SHA-256 values are `56b9d524aa3520180f680dfbc7dbe971f07b9203e2f34f9cbc6e6eac1466f9cb` (current) and `a96aa3d72d691b3e3ea74ed3ffe549074746ed5c21fa2a1883d19cf2d817ff64` (shortened).

### What this supports

If retaining make-sense as explicit reusable guidance, the shortened candidate is the lower-overhead choice supported by this smoke run. It is not a proven quality improvement over using no skill. Before stronger claims, repeat difficult cases in fresh case-level sessions with the harness issues removed, include real anonymized failures and held-out cases, and obtain human review. Skill discovery, competing-skill selection, tool-enabled verification, cross-model behavior, and learner comprehension remain untested.

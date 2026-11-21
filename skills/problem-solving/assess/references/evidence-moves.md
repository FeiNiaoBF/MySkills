# Evidence moves by target type

Use these as prompts, not a checklist. Stop when a check cannot change the verdict.

## Code or function

**Look for:** the contract and its callers; error, empty, boundary, and partial-failure paths; state and concurrency; resource cleanup; trust boundaries; performance only where it matters.

**Verify with:** read callers and tests; run focused tests; reproduce a suspected failure; inspect configuration, versions, and generated artifacts when they affect behavior.

**Limit claims to:** what the exercised paths establish. A focused test does not prove general correctness; an unmeasured path does not establish performance.

## Document or specification

**Look for:** the audience and intended action; material claims; internal consistency; stale versus missing content; requirements versus evidence of compliance.

**Verify with:** check material claims against project facts or authoritative sources; compare the document with current code, configuration, or behavior; sample the sections that carry the decision.

**Limit claims to:** what the sources establish. An uncited claim is unverified, not automatically false; an existing document is not automatically current; a specification is not proof of implementation.

## UI

**Look for:** the rendered state; loading, empty, error, disabled, and permission states; labels and contrast; keyboard and focus behavior; responsive and zoom behavior.

**Verify with:** open the actual surface; interact with it; inspect the DOM or accessibility tree; compare screenshots with observed behavior.

**Limit claims to:** states actually exercised. A screenshot does not establish unseen interaction; attributes alone do not establish accessibility; a happy path does not establish error handling.

## Plan, idea, or proposal

**Look for:** the problem it solves; critical assumptions; constraints; cost and reversibility; alternatives including the status quo; who bears the risk.

**Verify with:** check critical assumptions against evidence; compare with a simpler alternative; identify the cheapest check that would discriminate between the live possibilities.

**Limit claims to:** what the evidence supports. Plausibility is not feasibility; value without a baseline is incomplete; stakeholder agreement must be checked, not assumed.

## Data or model

**Look for:** provenance; fitness for the intended use; coverage and missingness; leakage; uncertainty; the comparison baseline.

**Verify with:** inspect the schema and source; reproduce a metric; compare against a baseline; check assumptions and validation scope.

**Limit claims to:** the tested use and sample. Correlation is not causation; fit on one population does not establish fit on another; a metric without a baseline does not establish improvement.

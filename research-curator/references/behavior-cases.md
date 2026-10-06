# Manual behavior cases

These cases check observable agent behavior. They are not keyword checks and do not prove general research competence. Run them in a fresh session with the skill available; capture the tool calls, retrieved passages, package, output page, and limitations. Never fabricate expected artifacts when a required host tool is unavailable.

| Case | Input | Observable pass condition |
|---|---|---|
| Frame without friction | Ask for a short, well-scoped explanation for a named audience. | Starts with a reasonable scope; does not ask the user to fill a generic multi-field contract. |
| Search, then inspect | Ask a current factual question with two core subquestions. | Performs real searches using distinct formulations, fetches relevant pages, and cites inspected passages rather than relying on snippets. |
| Near-miss source | Include a high-ranking but off-topic candidate. | Records a specific fit reason and excludes it without pretending it supports the answer. |
| Duplicate origin | Search results include aliases or syndicated copies. | Deduplicates exact matches and identifies shared upstream origins instead of counting copies as independent evidence. |
| Saturation reached | Provide evidence whose first rounds add major claims and whose next three completed rounds add little, including a counterevidence search. | Stops as `saturated` only after core coverage, references the exact three low-gain rounds, and reports remaining gaps. |
| Material counterexample | A late result changes a conclusion or its scope. | Marks material gain, resets the low-gain window, and searches/assesses the new issue before any saturation stop. |
| Budget exhausted | Use a small explicit budget while important questions remain unresolved. | Stops as `budget_exhausted`, writes a qualified useful result if possible, and does not label it saturated. |
| Research-to-writing boundary | Ask for an article, then inspect the evidence package before Writing. | Package validates; article claims reference existing evidence; unsupported premises return to Research. |
| Plain-language default | Ask for an English or Chinese reader-facing explanation containing necessary technical qualifications. | Uses direct wording and consistent terms, but preserves exact caveats and does not force artificial sentence limits. |
| Provenance visualization | Open a published report, select a body citation, then select its graph claim/source. | Navigation works both ways; source excerpt, locator, uncertainty, and related article block are inspectable. |
| Offline and safe | Open the report from a moved local folder using adversarial source text. | No runtime network requests; hostile text is inert; keyboard alternatives work; a graph failure does not remove the article or citations. |
| Collision and cleanup | Publish where the folder already exists; inspect temporary cleanup boundaries. | Existing contents remain unchanged; the run reports collision; cleanup is restricted to its owned temporary directory. |

The 2026-10-06 [acceptance record](../examples/VERIFICATION.md) reports the fresh old/new comparison, final reruns, actual retrieval/browser checks, and untested cases. This is not a blanket pass for every scenario above.

Record each case as `pass`, `fail`, or `unverified`, with a short evidence reference. Separate structural validation, observed agent behavior, and actual browser behavior.

# HTML Lesson Format

Lessons live in `./lessons/` and are optional. Use one when a visual model, interactive simulator, debugging flow, or print-friendly reference materially improves recall. The runnable project remains the primary learning artifact.

## File contract

Name a lesson `0001-<dash-case-topic>.html`. It should be self-contained, link to the relevant project and reference files, and contain one tightly scoped decision or mental model.

Each lesson includes:

- A visible statement of the project capability it supports.
- The smallest necessary explanation, with links to the primary source in `RESOURCES.md`.
- A retrieval prompt, interactive check, or small observable task.
- A link to the next project brief and any related glossary term.
- A reminder that the learner can ask follow-up questions.

Use shared styles and interactive components from `./assets/`. Add a reusable asset only when a second lesson will use it; otherwise keep the lesson self-contained. Open or render the HTML when possible and check that text, controls, and links are usable before treating it as complete.

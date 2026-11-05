# TRACK.md Format

`TRACK.md` describes the domain-specific capabilities, verification, and likely misconceptions for one technology. The ordered project ladder and its status live in `ROADMAP.md`; refer to it rather than maintaining a second copy here.

## Template

```md
# {Technology} Track

## Scope
- Learner: {diagnostic result and relevant prior ability}
- North-star artifact: {what they want to make}
- Environment: {runtime, hardware, budget, and access constraints}
- Version boundary: {technology and important dependency versions}

## Capability map
| Capability | Observable evidence | Prerequisites | Primary source |
| --- | --- | --- | --- |
| {capability} | {proof} | {known skill} | {direct link} |

## Domain verification
{Commands, visual checks, metrics, or evaluation protocol for this technology.}

## Likely misconceptions
| Misconception | Observation that reveals it | Corrective project action |
| --- | --- | --- |
| {belief} | {failure or explanation} | {smallest corrective task} |
```

## Rules

- Derive the capability map from the learner's goal, current ability, and the domain's actual dependencies. Keep the ordered projects in `ROADMAP.md`; do not prewrite projects that the evidence cannot support.
- Record source URLs and versions for APIs, models, datasets, standards, security, and evaluation choices.
- Keep domain-specific verification here; keep the help ladder and state contract in the shared references.

## Source rules

Use official documentation, original papers, standards, or authoritative maintainers for material that determines APIs, behavior, security, compatibility, evaluation, or version choices. Record the URL and the version or publication date when it can affect the route. Community examples may inspire project scenarios, but they do not replace the primary source behind a technical claim.

## Project selection rules

Use the diagnostic and the learner's desired artifact to select a first project. Add later projects when they introduce a distinct capability or useful transfer constraint; adjust their order when new evidence changes prerequisites or difficulty. A later project earns its place only when its expected evidence differs from prior work.

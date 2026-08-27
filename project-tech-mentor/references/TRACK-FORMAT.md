# TRACK.md Format

`TRACK.md` describes one technology route. It is the bridge between the common project workflow and domain-specific capabilities.

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

## Project ladder
| Order | Project | Capability added | Constraint | Evidence | Transfer check |
| --- | --- | --- | --- | --- | --- |
| 1 | {name} | {one capability} | {scope boundary} | {proof} | {variant} |

## Domain verification
{Commands, visual checks, metrics, or evaluation protocol for this technology.}

## Likely misconceptions
| Misconception | Observation that reveals it | Corrective project action |
| --- | --- | --- |
| {belief} | {failure or explanation} | {smallest corrective task} |
```

## Rules

- Use a tested default ladder, then adapt scenario and constraints to the learner's mission.
- Preserve capability order unless diagnostic evidence gives a concrete reason to change it.
- Record source URLs and versions for APIs, models, datasets, standards, security, and evaluation choices.
- Keep domain-specific verification here; keep the help ladder and state contract in the shared references.

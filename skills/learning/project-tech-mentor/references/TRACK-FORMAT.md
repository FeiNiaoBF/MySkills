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

## Source rules

Use official documentation, original papers, standards, or authoritative maintainers for material that determines APIs, behavior, security, compatibility, evaluation, or version choices. Record the URL and the version or publication date when it can affect the route. Community examples may inspire project scenarios, but they do not replace the primary source behind a technical claim.

## Project selection rules

Start with a tested default ladder, then adjust its scenario and constraints to the learner's north-star artifact. Preserve the capability order unless diagnostic evidence gives a concrete reason to change it. A route should normally contain four to eight projects; a project earns its place only when its evidence differs from the prior project's evidence.

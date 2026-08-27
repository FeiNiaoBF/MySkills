# Domain Track Card

This is the legacy name for the domain-route guidance. The canonical `TRACK.md` file shape is [TRACK-FORMAT.md](TRACK-FORMAT.md); use this page for the source and project-selection rules when revising a route.

```markdown
# <Technology> Track

## Scope
- Learner: <relevant current ability and diagnostic result>
- North-star artifact: <what they want to make>
- Environment: <runtime, hardware, budget, and access constraints>
- Version boundary: <technology and important dependency versions>

## Capability map
| Capability | Observable evidence | Prerequisites | Primary source |
| --- | --- | --- | --- |
| <capability> | <proof> | <known skill> | <direct link> |

## Project ladder
| Order | Project | Capability added | Constraint | Evidence | Transfer check |
| --- | --- | --- | --- | --- | --- |
| 1 | <name> | <one capability> | <scope boundary> | <proof> | <variant> |

## Domain verification
<Commands, visual checks, metrics, or evaluation protocol appropriate to this technology.>

## Likely misconceptions
| Misconception | Observation that reveals it | Corrective project action |
| --- | --- | --- |
| <belief> | <failure or explanation> | <smallest corrective task> |
```

## Source rules

Use official documentation, original papers, standards, or authoritative maintainers for material that determines APIs, behavior, security, compatibility, evaluation, or version choices. Record the URL and the version or publication date when it can affect the route. Community examples may inspire project scenarios, but they do not replace the primary source behind a technical claim.

## Project selection rules

Start with a tested default ladder, then adjust its scenario and constraints to the learner's north-star artifact. Preserve the capability order unless diagnostic evidence gives a concrete reason to change it. A route should normally contain four to eight projects; a project earns its place only when its evidence differs from the prior project's evidence.

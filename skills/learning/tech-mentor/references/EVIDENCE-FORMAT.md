# Evidence Format

Store evidence in `evidence/<number>-<project-name>/`. Evidence is the proof that a project works and that the learner can explain the capability, not a dump of every terminal log.

## Recommended contents

```text
evidence/0001-render-loop/
  README.md              # index and verdict
  machine.txt            # command, version, result, and relevant output
  human.png              # screenshot, or a recording path
  examples/              # representative inputs, outputs, or failures
  metrics.json           # optional structured evaluation
```

## `README.md` template

```md
# Evidence: {Project}

## Machine
- Command or protocol: `{exact command}`
- Result: {pass, metric, or known limitation}

## Human
- Artifact: {path or link}
- Observable behavior: {what was inspected}

## Learner explanation
{Why the key design works, what failed, and what would change next.}

## Assistance and reproducibility
{Material hints, examples, or code supplied; setup/version facts needed to reproduce the result.}
```

Keep only evidence that supports the brief. For Three.js or frontend work, favor interaction capture; for backend work, request assertions and logs; for LLM work, fixed evaluation outputs and failures; for computer vision, metrics and an error gallery.

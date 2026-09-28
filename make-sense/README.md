# make-sense

A context-aware understanding-repair Agent Skill. Use it when a term, passage, or reasoning step does not make sense, or when you want to check an interpretation. It repairs the blocking gap and returns to your task without imposing a lesson or quiz. Standalone term questions qualify; broad topic teaching, translation-only requests, and ordinary implementation retain their own scope.

## Files

- `SKILL.md` — the installable runtime Skill; this is the only file required by the Agent Skills format.
- `references/rationale.md` — maintainer-only research/design notes, not loaded in ordinary use.
- `references/cases.yaml` — self-contained conversation inputs and separate reviewer criteria; not executable tests.
- `references/shortened.md` — frozen candidate from the 2026-09-26 comparison, retained as historical evidence; not the installed runtime skill.
- `references/evaluation.md` — comparison protocol, smoke results, and evidence limits.

## Install

Copy the **whole `make-sense` directory** into your agent's skill directory. The parent directory name must remain `make-sense`.

For a shared Pi/Codex user-level location on a compatible recent installation:

```sh
mkdir -p ~/.agents/skills
cp -R make-sense ~/.agents/skills/
```

For Pi only, a typical alternate destination is `~/.pi/agent/skills/make-sense/`. Pi skill commands use `/skill:make-sense`; user-provided arguments can follow it. Do not assume `/make-sense` is an alias. Reload skills or restart the host if it does not discover newly installed skills; check the host's active skill settings if it is disabled.

The metadata description enables contextual selection **when the host supports model-invoked skills**; it does not guarantee automatic invocation on every eligible message.

## Review checklist

Compare a no-skill baseline, a frozen previous runtime, and the proposed runtime using [the evaluation protocol](references/evaluation.md) and [the fixed cases](references/cases.yaml). Give response generators only each case's `input`; keep `review` and the shared criteria hidden until scoring. Human review should judge meaning preservation, targeting, excess jargon, invented specifics, unnecessary tutorials, and whether repeated confusion results in a genuinely different repair strategy.

## Scope

This Skill does not perform web search, generate diagrams, keep permanent user knowledge profiles, or override the user's main task. It does not require scripts or additional installed packages.

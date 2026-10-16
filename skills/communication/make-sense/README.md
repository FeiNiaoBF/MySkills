# make-sense

A context-aware understanding-repair Agent Skill. Use it when a term, passage, or reasoning step does not make sense, or when you want to check an interpretation. It uses precise text explanations to repair the blocking gap and return to your task without imposing a lesson or quiz. Standalone term questions qualify; broad topic teaching, translation-only requests, and ordinary implementation retain their own scope.

## Files

- `SKILL.md` — the installable runtime Skill; this is the only file required by the Agent Skills format.
- `references/rationale.md` — maintainer-only research/design notes, not loaded in ordinary use.
- `references/cases.yaml` — self-contained conversation inputs and separate reviewer criteria; not executable tests.
- `references/shortened.md` — frozen candidate from the 2026-09-26 comparison, retained as historical evidence; not the installed runtime skill.
- `references/evaluation.md` — comparison protocol, smoke results, and evidence limits.

## Install

Install the skill from the MySkills GitHub repository with Vercel's `skills` CLI:

```sh
npx skills add FeiNiaoBF/MySkills --skill make-sense -g -a pi
```

Replace `pi` with another supported agent name as needed. To list all skills in the repository, run `npx skills add FeiNiaoBF/MySkills --list`. For updates, follow the repository’s shared [installation and update guidance](https://github.com/FeiNiaoBF/MySkills#安装与更新).

Pi skill commands use `/skill:make-sense`; user-provided arguments can follow it. Do not assume `/make-sense` is an alias. Reload skills or restart the host if it does not discover newly installed skills; check the host's active skill settings if it is disabled.

The metadata description enables contextual selection **when the host supports model-invoked skills**; it does not guarantee automatic invocation on every eligible message.

## Review checklist

Compare a no-skill baseline, a frozen previous runtime, and the proposed runtime using [the evaluation protocol](references/evaluation.md) and [the fixed cases](references/cases.yaml). Give response generators only each case's `input`; keep `review` and the shared criteria hidden until scoring. Human review should judge meaning preservation, targeting, excess jargon, invented specifics, unnecessary tutorials, and whether repeated confusion results in a genuinely different repair strategy.

## Scope

This Skill provides text explanations, including examples, code snippets, or equations where useful. It does not generate diagrams or other visual explanations, keep permanent user knowledge profiles, or override the user's main task. It may use available verification tools for consequential facts and states what remains uncertain when verification is unavailable. It does not require scripts or additional installed packages.

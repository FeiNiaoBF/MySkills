#!/usr/bin/env python3
"""Report per-turn and on-demand context size; fail on over-budget descriptions."""
from __future__ import annotations

import argparse
import re
from pathlib import Path
import sys

import yaml

ROOT = Path(__file__).resolve().parents[1]
AUTO_BUDGET = 140
MANUAL_BUDGET = 100


def skill_frontmatter(path: Path) -> dict:
    text = path.read_text(encoding="utf-8")
    if not text.startswith("---\n"):
        raise ValueError(f"{path}: frontmatter must start on the first line")
    end = text.find("\n---", 4)
    if end == -1:
        raise ValueError(f"{path}: unterminated frontmatter")
    values = yaml.safe_load(text[4:end])
    if not isinstance(values, dict):
        raise ValueError(f"{path}: frontmatter must be a mapping")
    return values


def description_size(values: dict) -> int:
    description = values.get("description", "")
    if not isinstance(description, str) or not description.strip():
        raise ValueError("description must be a nonempty string")
    return len(re.sub(r"\s+", " ", description).strip())


def skill_budget(values: dict) -> int:
    return MANUAL_BUDGET if values.get("disable-model-invocation") is True else AUTO_BUDGET


def text_size(path: Path) -> int:
    return len(path.read_text(encoding="utf-8"))


def collect(root: Path) -> tuple[list[tuple[Path, int, int]], dict[str, int]]:
    rows: list[tuple[Path, int, int]] = []
    sizes = {"agents": 0, "descriptions": 0, "bodies": 0, "references": 0, "repo_docs": 0}
    agents = root / "AGENTS.md"
    if agents.exists():
        sizes["agents"] = text_size(agents)
    for skill in sorted(root.glob("skills/*/*/SKILL.md")):
        values = skill_frontmatter(skill)
        size = description_size(values)
        rows.append((skill, size, skill_budget(values)))
        sizes["descriptions"] += size
        sizes["bodies"] += text_size(skill)
        sizes["references"] += sum(
            text_size(path)
            for path in skill.parent.joinpath("references").rglob("*")
            if path.is_file()
        )
    for name in ("README.md", "CONTRIBUTING.md"):
        path = root / name
        if path.exists():
            sizes["repo_docs"] += text_size(path)
    docs = root / "docs"
    if docs.exists():
        sizes["repo_docs"] += sum(
            text_size(path) for path in docs.rglob("*.md") if path.is_file()
        )
    return rows, sizes


def report(rows: list[tuple[Path, int, int]], sizes: dict[str, int], root: Path) -> list[str]:
    print("per-turn")
    print(f"  AGENTS.md          {sizes['agents']:>7}")
    print(f"  descriptions       {sizes['descriptions']:>7}  ({len(rows)} skills)")
    print("on-demand")
    print(f"  skill bodies       {sizes['bodies']:>7}")
    print(f"  references         {sizes['references']:>7}")
    print(f"  repo docs          {sizes['repo_docs']:>7}")

    failures = []
    for path, size, budget in rows:
        if size > budget:
            rel = path.relative_to(root).as_posix()
            failures.append(f"{rel}: description {size} > {budget}")
    return failures


def main() -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--root", type=Path, default=ROOT, help="repository root")
    args = parser.parse_args()

    root = args.root.resolve()
    try:
        rows, sizes = collect(root)
    except (OSError, ValueError, yaml.YAMLError) as exc:
        print(f"FAIL: {exc}")
        return 1

    failures = report(rows, sizes, root)
    if failures:
        for failure in failures:
            print(failure)
        print("FAIL")
        return 1
    print("PASS")
    return 0


if __name__ == "__main__":
    sys.exit(main())

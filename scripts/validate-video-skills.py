from __future__ import annotations

import re
import sys
from pathlib import Path

ROOT = Path(__file__).resolve().parents[1]
REQUIRED = {
    "math-video-workflow",
    "video-smart-cut",
    "math-video-review",
    "math-manim-insertion",
    "manim-video-insertion",
    "manim-recap-video",
}
ABSOLUTE_PATH = re.compile(r"(?:[A-Za-z]:[\\/]|/home/[^/]+/)")
LINK = re.compile(r"\[[^]]+\]\(([^)]+)\)")


def frontmatter(text: str) -> dict[str, str]:
    if not text.startswith("---\n"):
        raise ValueError("frontmatter must start at byte 0")
    end = text.find("\n---\n", 4)
    if end < 0:
        raise ValueError("frontmatter closing marker missing")
    values: dict[str, str] = {}
    for line in text[4:end].splitlines():
        match = re.match(r"^([a-z_]+):\s*(.+)$", line)
        if match:
            values[match.group(1)] = match.group(2).strip().strip('"')
    if not text[end + 5 :].strip():
        raise ValueError("body is empty")
    return values


def main() -> int:
    errors: list[str] = []
    for name in sorted(REQUIRED):
        path = ROOT / name / "SKILL.md"
        if not path.exists():
            errors.append(f"{name}: missing SKILL.md")
            continue
        text = path.read_text(encoding="utf-8")
        try:
            fm = frontmatter(text)
        except ValueError as exc:
            errors.append(f"{name}: {exc}")
            continue
        if fm.get("name") != name:
            errors.append(f"{name}: frontmatter name={fm.get('name')!r}")
        description = fm.get("description", "")
        if not description:
            errors.append(f"{name}: missing description")
        if len(description) > 57:
            errors.append(f"{name}: description exceeds 57 characters")
        if "metadata:" not in text[: text.find("\n---\n", 4)]:
            errors.append(f"{name}: missing metadata mapping")
        for match in ABSOLUTE_PATH.finditer(text):
            errors.append(f"{name}: machine-local path: {match.group(0)}")
        for target in LINK.findall(text):
            if "://" in target or target.startswith("#"):
                continue
            if not (path.parent / target).exists():
                errors.append(f"{name}: broken link: {target}")
    if errors:
        print("FAIL")
        print("\n".join(f"- {error}" for error in errors))
        return 1
    print(f"PASS: validated {len(REQUIRED)} video skills")
    return 0


if __name__ == "__main__":
    sys.exit(main())

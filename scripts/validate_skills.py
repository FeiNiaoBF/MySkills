"""Validate repository packaging; semantic quality and secrets still need review."""
from __future__ import annotations

import argparse
from pathlib import Path
import re
import sys
from urllib.parse import unquote, urlsplit

try:
    from markdown_it import MarkdownIt
    import yaml
except ImportError as exc:
    raise SystemExit(
        "Missing validation dependencies. Run: python -m pip install -r requirements-dev.txt"
    ) from exc

ROOT = Path(__file__).resolve().parents[1]
INFRASTRUCTURE_DIRS = {"docs", "scripts", "tests"}
FRONTMATTER_FIELDS = {
    "name",
    "description",
    "metadata",
    "license",
    "compatibility",
    "allowed-tools",
    # Claude Code extension for skills that require explicit user invocation.
    "disable-model-invocation",
}
TEXT_SUFFIXES = {".md", ".yaml", ".yml", ".json"}
MACHINE_PATH = re.compile(
    r"(?<![A-Za-z0-9])(?:[A-Za-z]:[\\/]|/(?:home|Users)/[^/\s<>]+/|/root/)"
)
PRIVATE_KEY = re.compile(r"-----BEGIN (?:[A-Z0-9]+ )?PRIVATE KEY-----")
MARKDOWN = MarkdownIt("commonmark")


class UniqueSafeLoader(yaml.SafeLoader):
    """Reject silent overwrites while retaining SafeLoader's tag restrictions."""

    def construct_mapping(self, node, deep=False):
        self.flatten_mapping(node)
        mapping = {}
        for key_node, value_node in node.value:
            key = self.construct_object(key_node, deep=deep)
            if key in mapping:
                raise yaml.constructor.ConstructorError(
                    None, None, "duplicate mapping key", key_node.start_mark
                )
            mapping[key] = self.construct_object(value_node, deep=deep)
        return mapping


def frontmatter(text: str) -> dict:
    lines = text.splitlines()
    if not lines or lines[0] != "---":
        raise ValueError("frontmatter must start on the first line")
    try:
        end = lines.index("---", 1)
    except ValueError as exc:
        raise ValueError("frontmatter closing marker missing") from exc
    try:
        values = yaml.load("\n".join(lines[1:end]), Loader=UniqueSafeLoader)
    except (yaml.YAMLError, TypeError, ValueError) as exc:
        raise ValueError("invalid YAML frontmatter (including duplicate keys or unsafe tags)") from exc
    if not isinstance(values, dict):
        raise ValueError("frontmatter must be a mapping")
    if not "\n".join(lines[end + 1:]).strip():
        raise ValueError("skill body is empty")
    return values


def markdown_targets(text: str):
    # Mask YAML instead of parsing it as Markdown; preserve source line offsets.
    lines = text.splitlines(keepends=True)
    if lines and lines[0].strip() == "---":
        for end in range(1, len(lines)):
            if lines[end].strip() == "---":
                lines[:end + 1] = ["\n"] * (end + 1)
                break

    def visit(tokens, line=1):
        for token in tokens:
            token_line = token.map[0] + 1 if token.map else line
            attribute = {"link_open": "href", "image": "src"}.get(token.type)
            if attribute:
                yield token_line, token.attrGet(attribute)
            if token.children:
                yield from visit(token.children, token_line)

    yield from visit(MARKDOWN.parse("".join(lines)))


def validate_repository(root: Path) -> tuple[int, list[str]]:
    root = root.resolve()
    if not root.is_dir():
        return 0, ["repository root is not a directory"]
    issues: list[str] = []
    skills = []
    # Every visible top-level directory except repository infrastructure is a skill.
    for directory in sorted(root.iterdir()):
        if not directory.is_dir() or directory.name.startswith(".") or directory.name in INFRASTRUCTURE_DIRS:
            continue
        if not (directory / "SKILL.md").is_file():
            issues.append(f"{directory.name}/SKILL.md: missing skill entrypoint")
        else:
            skills.append(directory)
    if not skills:
        issues.append("no skills found")

    def report(path: Path, message: str, line: int = 1):
        issues.append(f"{path.relative_to(root).as_posix()}:{line}: {message}")

    def read_text(path: Path, boundary: Path):
        try:
            if not path.resolve().is_relative_to(boundary):
                report(path, "file resolves outside allowed directory")
                return None
            return path.read_text(encoding="utf-8")
        except (OSError, UnicodeError, RuntimeError):
            report(path, "cannot read UTF-8 file or resolve its path")
            return None

    # Each file is checked once, with either skill-local or repository-local links.
    documents = {path: root for path in root.glob("*.md") if path.is_file()}
    docs = root / "docs"
    if docs.is_dir():
        documents.update({path: root for path in docs.rglob("*.md") if path.is_file()})
    for skill in skills:
        if not skill.resolve().is_relative_to(root):
            report(skill, "skill resolves outside allowed directory")
            continue
        text = read_text(skill / "SKILL.md", skill.resolve())
        if text is not None:
            try:
                values = frontmatter(text)
            except ValueError as exc:
                report(skill / "SKILL.md", str(exc))
            else:
                if values.get("name") != skill.name:
                    report(skill / "SKILL.md", "name must match the skill directory")
                description = values.get("description")
                if not isinstance(description, str) or not description.strip():
                    report(skill / "SKILL.md", "description must be a nonempty string")
                if not isinstance(values.get("metadata"), dict):
                    report(skill / "SKILL.md", "metadata must be a mapping")
                if set(values) - FRONTMATTER_FIELDS:
                    report(skill / "SKILL.md", "custom frontmatter fields belong under metadata")
        documents.update({
            path: skill.resolve()
            for path in skill.rglob("*")
            if path.is_file() and path.suffix.lower() in TEXT_SUFFIXES
        })

    for path, boundary in sorted(documents.items()):
        text = read_text(path, boundary)
        if text is None:
            continue
        for number, line in enumerate(text.splitlines(), 1):
            if MACHINE_PATH.search(line):
                report(path, "machine-specific absolute path; use a placeholder", number)
            if PRIVATE_KEY.search(line):
                report(path, "private-key marker; remove private key material", number)
        if path.suffix.lower() != ".md":
            continue
        for line, target in markdown_targets(text):
            try:
                parts = urlsplit(target)
                if parts.scheme or parts.netloc or not parts.path:
                    continue
                destination = (path.parent / unquote(parts.path)).resolve()
                if not destination.is_relative_to(boundary):
                    report(path, f"link resolves outside allowed directory: {target}", line)
                elif not destination.exists():
                    report(path, f"missing local target: {target}", line)
            except (ValueError, OSError, RuntimeError):
                report(path, "invalid local link target", line)
    return len(skills), sorted(set(issues))


def main(argv: list[str] | None = None) -> int:
    parser = argparse.ArgumentParser(description=__doc__)
    parser.add_argument("--root", type=Path, default=ROOT, help="repository root (default: this checkout)")
    args = parser.parse_args(argv)
    count, issues = validate_repository(args.root)
    if issues:
        print(f"FAIL: checked {count} skill(s); {len(issues)} issue(s)")
        print("\n".join(f"- {issue}" for issue in issues))
        return 1
    print(f"PASS: checked {count} skill(s), their text resources, and repository Markdown")
    print("Manual review still required: trigger in first 57 characters, sensitive values, and behavior.")
    return 0


if __name__ == "__main__":
    sys.exit(main())

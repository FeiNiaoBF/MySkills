from __future__ import annotations

import contextlib
import io
import os
from pathlib import Path
import subprocess
import sys
import tempfile
import unittest

from scripts.validate_skills import frontmatter, main, validate_repository

REPO = Path(__file__).resolve().parents[1]
TEMP_ROOT = Path(tempfile.gettempdir()) / "pi-agent"
TEMP_ROOT.mkdir(parents=True, exist_ok=True)


def skill_text(name: str = "example", body: str = "Explain the requested task.") -> str:
    return (
        f"---\nname: {name}\ndescription: Use when explaining a task.\n"
        f"metadata: {{version: '1.0'}}\n---\n\n{body}\n"
    )


class FrontmatterTests(unittest.TestCase):
    def test_folded_long_description_is_not_a_57_character_cap(self):
        text = skill_text().replace(
            "description: Use when explaining a task.",
            "description: >-\n  Use when explaining a task.\n"
            "  Retain important conditions while clarifying the user's immediate question.",
        )
        parsed = frontmatter(text)
        self.assertGreater(len(parsed["description"]), 57)
        self.assertIn("Retain important", parsed["description"])

    def test_rejects_malformed_or_unsafe_yaml(self):
        for value in ("[", "!!python/object/apply:os.system ['echo unsafe']"):
            with self.subTest(value=value), self.assertRaises(ValueError):
                frontmatter(skill_text().replace("metadata: {version: '1.0'}", f"metadata: {value}"))

    def test_rejects_duplicate_keys(self):
        for field in ("name: example\nname: another", "name: example\nmetadata: {}"):
            with self.subTest(field=field), self.assertRaises(ValueError):
                frontmatter(skill_text().replace("name: example", field))

    def test_requires_frontmatter_mapping_and_body(self):
        for text in ("Body", "---\nname: x", "---\n- a\n---\nBody", "---\n{}\n---\n"):
            with self.subTest(text=text), self.assertRaises(ValueError):
                frontmatter(text)


class RepositoryTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix="skills-validator-", dir=TEMP_ROOT)
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.write("example/SKILL.md", skill_text())

    def write(self, relative: str, text: str):
        path = self.root / relative
        path.parent.mkdir(parents=True, exist_ok=True)
        path.write_text(text, encoding="utf-8")
        return path

    def errors(self):
        return validate_repository(self.root)[1]

    def test_discovers_nonvideo_skills_and_ignores_infrastructure(self):
        self.write("new-skill/SKILL.md", skill_text("new-skill"))
        for name in ("scripts", "tests", "docs", ".git", ".cache"):
            (self.root / name).mkdir()
        self.assertEqual(validate_repository(self.root), (2, []))

    def test_directory_missing_skill_is_reported(self):
        (self.root / "unfinished").mkdir()
        self.assertTrue(any("unfinished/SKILL.md" in e for e in self.errors()))

    def test_empty_and_nonexistent_roots_do_not_pass(self):
        empty = self.root / "empty"
        empty.mkdir()
        self.assertTrue(validate_repository(empty)[1])
        self.assertTrue(validate_repository(self.root / "absent")[1])

    def test_required_fields_and_custom_metadata(self):
        variants = {
            "name must match": skill_text("wrong"),
            "description must be a nonempty string": skill_text().replace(
                "description: Use when explaining a task.", "description: []"
            ),
            "metadata must be a mapping": skill_text().replace("metadata: {version: '1.0'}\n", ""),
            "custom frontmatter fields": skill_text().replace("name: example", "name: example\nauthor: example"),
        }
        for expected, text in variants.items():
            with self.subTest(expected=expected):
                self.write("example/SKILL.md", text)
                self.assertTrue(any(expected in e for e in self.errors()), self.errors())

    def test_standard_optional_frontmatter_fields(self):
        self.write("example/SKILL.md", skill_text().replace(
            "name: example", "name: example\nlicense: MIT\ncompatibility: Python\nallowed-tools: Read"
        ))
        self.assertEqual(self.errors(), [])

    def test_all_skill_markdown_not_only_entrypoint_is_checked(self):
        self.write("example/references/note.md", "[Missing](absent.md)\n")
        self.assertTrue(any("example/references/note.md:1" in e and "missing local target" in e for e in self.errors()))

    def test_root_and_shared_docs_links_are_checked(self):
        self.write("README.md", "[Skill](example/SKILL.md)\n[Missing](missing.md)\n")
        self.write("docs/guide.md", "[README](../README.md)\n[Missing](missing.md)\n")
        errors = self.errors()
        self.assertEqual(len(errors), 2, errors)
        self.assertTrue(any(e.startswith("README.md:") for e in errors))
        self.assertTrue(any(e.startswith("docs/guide.md:") for e in errors))

    def test_reference_links_images_and_encoded_balanced_paths(self):
        self.write("example/guide (draft).md", "# Heading\n")
        self.write("example/SKILL.md", skill_text(body=(
            "[Guide](guide%20%28draft%29.md#heading)\n\n"
            "[Reference][guide]\n\n[guide]: <guide (draft).md>\n\n"
            "![Missing image](missing.png)\n"
        )))
        errors = self.errors()
        self.assertEqual(len(errors), 1, errors)
        self.assertIn("missing.png", errors[0])

    def test_reference_link_missing_target(self):
        self.write("example/README.md", "[Guide][ref]\n\n[ref]: absent.md\n")
        self.assertTrue(any("missing local target: absent.md" in e for e in self.errors()))

    def test_markdown_examples_are_not_live_links(self):
        self.write("example/README.md", "`[Example](missing.md)`\n\n```md\n[Example](missing.md)\n```\n")
        self.assertEqual(self.errors(), [])

    def test_external_links_fragments_and_doi_are_not_machine_paths(self):
        self.write("example/README.md", (
            "[DOI](https://doi.org/10.1007/example)\n\n"
            "[Email](mailto:example@example.invalid)\n\n[Section](#section)\n"
        ))
        self.assertEqual(self.errors(), [])

    def test_skill_link_cannot_escape_even_if_target_exists(self):
        self.write("README.md", "# Root\n")
        for target in ("../README.md", "%2E%2E/README.md"):
            with self.subTest(target=target):
                self.write("example/README.md", f"[Root]({target})\n")
                self.assertTrue(any("outside allowed directory" in e for e in self.errors()))

    def test_symlink_target_cannot_escape_skill(self):
        outside = self.write("outside.md", "# Outside\n")
        link = self.root / "example" / "linked.md"
        try:
            link.symlink_to(outside)
        except OSError as exc:
            self.skipTest(f"Symlinks unavailable: {exc}")
        self.write("example/README.md", "[Outside](linked.md)\n")
        self.assertTrue(any("outside allowed directory" in e for e in self.errors()))

    def test_invalid_encoding_reports_error_instead_of_crashing(self):
        (self.root / "example" / "SKILL.md").write_bytes(b"\xff")
        self.assertTrue(any("cannot read UTF-8" in e for e in self.errors()))

    def test_machine_paths_in_examples_and_yaml_are_reported_without_values(self):
        # Assemble synthetic private-looking values rather than publishing real paths.
        windows = "C:" + "/Users/" + "fixture-person/private"
        unix = "/home/" + "fixture-person/private"
        self.write("example/README.md", f"```sh\n{windows}\n```\n")
        self.write("example/references/data.yaml", f"path: {unix}\n")
        errors = self.errors()
        self.assertEqual(len(errors), 2, errors)
        self.assertTrue(all("machine-specific absolute path" in e for e in errors))
        self.assertTrue(all("fixture-person" not in e for e in errors))

    def test_private_key_markers_are_reported(self):
        marker = "-----BEGIN " + "OPENSSH PRIVATE KEY-----"
        self.write("example/README.md", marker)
        self.assertTrue(any("private-key marker" in e for e in self.errors()))

    def test_portable_placeholders_are_allowed(self):
        self.write("example/README.md", "`<VPS_IP>` `<SSH_HOST>` `/home/<user>/project` `${TEMP}/pi-agent/`\n")
        self.assertEqual(self.errors(), [])

    def test_cli_exit_code_and_counts(self):
        for valid in (True, False):
            if not valid:
                self.write("example/README.md", "[Broken](missing.md)\n")
            with contextlib.redirect_stdout(io.StringIO()) as output:
                code = main(["--root", str(self.root)])
            self.assertEqual(code, 0 if valid else 1)
            self.assertIn("1 skill", output.getvalue())

    def test_legacy_entrypoint_delegates_to_generic_validator(self):
        env = {**os.environ, "PYTHONDONTWRITEBYTECODE": "1"}
        for script in ("validate_skills.py", "validate-video-skills.py"):
            result = subprocess.run(
                [sys.executable, str(REPO / "scripts" / script), "--root", str(self.root)],
                capture_output=True, text=True, env=env,
            )
            self.assertEqual(result.returncode, 0, result.stderr)
            self.assertIn("1 skill", result.stdout)
        self.write("example/README.md", "[Broken](missing.md)\n")
        result = subprocess.run(
            [sys.executable, str(REPO / "scripts" / "validate-video-skills.py"), "--root", str(self.root)],
            capture_output=True, text=True, env=env,
        )
        self.assertEqual(result.returncode, 1, result.stderr)


if __name__ == "__main__":
    unittest.main()

from __future__ import annotations

import os
from pathlib import Path
import shutil
import subprocess
import tempfile
import unittest

REPO = Path(__file__).resolve().parents[1]
POWERSHELL = shutil.which("pwsh") or shutil.which("powershell")
TEMP_ROOT = Path(tempfile.gettempdir()) / "pi-agent"
TEMP_ROOT.mkdir(parents=True, exist_ok=True)


@unittest.skipUnless(os.name == "nt" and POWERSHELL, "Windows PowerShell junction tests")
class WindowsLinkTests(unittest.TestCase):
    def setUp(self):
        self.temp = tempfile.TemporaryDirectory(prefix="skills-links-", dir=TEMP_ROOT)
        self.addCleanup(self.temp.cleanup)
        self.root = Path(self.temp.name)
        self.repo = self.root / "repo"
        self.home = self.root / "home"
        self.script = self.repo / "scripts" / "link-skills.ps1"
        self.script.parent.mkdir(parents=True)
        self.home.mkdir()
        shutil.copyfile(REPO / "scripts" / "link-skills.ps1", self.script)
        self.env = {**os.environ, "USERPROFILE": str(self.home)}
        for name in ("example", "other"):
            directory = self.repo / name
            directory.mkdir()
            (directory / "SKILL.md").write_text(f"# {name}\n", encoding="utf-8")
        self.base = self.home / ".agents" / "skills"
        self.link = self.base / "myskills"

    def run_script(self, *args, expected_code=0):
        result = subprocess.run(
            [POWERSHELL, "-NoLogo", "-NoProfile", "-NonInteractive", "-File", str(self.script), *args],
            env=self.env, capture_output=True, text=True,
        )
        self.assertEqual(result.returncode, expected_code, result.stderr)
        return result

    def make_junction(self, link: Path, target: Path):
        link.parent.mkdir(parents=True, exist_ok=True)
        result = subprocess.run(
            [POWERSHELL, "-NoLogo", "-NoProfile", "-NonInteractive", "-Command",
             '$ErrorActionPreference = "Stop"; New-Item -ItemType Junction -Path $env:TEST_LINK -Target $env:TEST_TARGET | Out-Null'],
            env={**self.env, "TEST_LINK": str(link), "TEST_TARGET": str(target)},
            capture_output=True, text=True,
        )
        self.assertEqual(result.returncode, 0, result.stderr)

    def test_collection_mount_is_live_idempotent_and_removable(self):
        self.run_script()
        self.assertEqual(self.link.resolve(), self.repo.resolve())
        self.assertEqual((self.link / "example" / "SKILL.md").read_text(), "# example\n")

        added = self.repo / "added"
        added.mkdir()
        (added / "SKILL.md").write_text("# added\n", encoding="utf-8")
        self.assertTrue((self.link / "added" / "SKILL.md").is_file())

        self.assertIn("ok", self.run_script().stdout)
        self.run_script("-Remove")
        self.assertFalse(self.link.exists())
        self.assertEqual((added / "SKILL.md").read_text(), "# added\n")
        self.assertFalse((self.home / ".pi" / "agent" / "skills").exists())

    def test_foreign_collection_path_is_never_replaced_or_removed(self):
        self.link.mkdir(parents=True)
        sentinel = self.link / "local.txt"
        sentinel.write_text("keep", encoding="utf-8")
        self.run_script(expected_code=1)
        self.run_script("-Remove")
        self.assertEqual(sentinel.read_text(), "keep")

    def test_owned_legacy_links_are_migrated_but_foreign_entries_are_preserved(self):
        self.make_junction(self.base / "example", self.repo / "example")
        foreign = self.base / "other"
        foreign.mkdir(parents=True)
        sentinel = foreign / "local.txt"
        sentinel.write_text("keep", encoding="utf-8")

        self.run_script()

        self.assertFalse((self.base / "example").exists())
        self.assertEqual(sentinel.read_text(), "keep")
        self.assertTrue((self.link / "example" / "SKILL.md").is_file())


if __name__ == "__main__":
    unittest.main()

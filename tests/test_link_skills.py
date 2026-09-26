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
        self.link = self.home / ".agents" / "skills" / "example"

    def run_script(self, *args):
        result = subprocess.run(
            [POWERSHELL, "-NoLogo", "-NoProfile", "-NonInteractive", "-File", str(self.script), *args],
            env=self.env, capture_output=True, text=True,
        )
        self.assertEqual(result.returncode, 0, result.stderr)
        return result

    def test_selected_install_is_idempotent_and_owned_removal_preserves_source(self):
        self.run_script("-Only", "example")
        self.assertEqual(self.link.resolve(), (self.repo / "example").resolve())
        self.assertFalse((self.link.parent / "other").exists())
        self.assertFalse((self.home / ".pi" / "agent" / "skills").exists())
        self.assertIn("ok", self.run_script("-Only", "example").stdout)
        self.run_script("-Remove", "-Only", "example")
        self.assertFalse(self.link.exists())
        self.assertEqual((self.repo / "example" / "SKILL.md").read_text(), "# example\n")

    def test_existing_real_directory_is_never_replaced_or_removed(self):
        self.link.mkdir(parents=True)
        sentinel = self.link / "local.txt"
        sentinel.write_text("keep", encoding="utf-8")
        self.run_script("-Only", "example")
        self.run_script("-Remove", "-Only", "example")
        self.assertEqual(sentinel.read_text(), "keep")

    def test_foreign_junction_is_never_replaced_or_removed(self):
        foreign = self.root / "foreign"
        foreign.mkdir()
        (foreign / "local.txt").write_text("keep", encoding="utf-8")
        self.link.parent.mkdir(parents=True)
        result = subprocess.run(
            [POWERSHELL, "-NoLogo", "-NoProfile", "-NonInteractive", "-Command",
             '$ErrorActionPreference = "Stop"; New-Item -ItemType Junction -Path $env:TEST_LINK -Target $env:TEST_TARGET | Out-Null'],
            env={**self.env, "TEST_LINK": str(self.link), "TEST_TARGET": str(foreign)},
            capture_output=True, text=True,
        )
        self.assertEqual(result.returncode, 0, result.stderr)
        self.run_script("-Only", "example")
        self.assertEqual(self.link.resolve(), foreign.resolve())
        self.run_script("-Remove", "-Only", "example")
        self.assertTrue(self.link.exists(), "Must preserve a junction owned by another source")
        self.assertEqual((self.link / "local.txt").read_text(), "keep")


if __name__ == "__main__":
    unittest.main()

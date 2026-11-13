import tempfile
import unittest
from pathlib import Path

from scripts.context_budget import AUTO_BUDGET, MANUAL_BUDGET, collect


def write_skill(root: Path, name: str, description: str, manual: bool = False) -> None:
    skill = root / "skills" / "demo" / name
    skill.mkdir(parents=True)
    disabled = "disable-model-invocation: true\n" if manual else ""
    (skill / "SKILL.md").write_text(
        f"---\nname: {name}\ndescription: {description}\n{disabled}license: MIT\n---\n\n# Demo\n",
        encoding="utf-8",
    )


class ContextBudgetTests(unittest.TestCase):
    def test_auto_description_within_budget_passes(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            write_skill(root, "demo", "x" * AUTO_BUDGET)
            rows, _ = collect(root)
            self.assertEqual(rows[0][1], AUTO_BUDGET)

    def test_auto_description_over_budget_is_reported(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            write_skill(root, "demo", "x" * (AUTO_BUDGET + 1))
            rows, _ = collect(root)
            self.assertGreater(rows[0][1], rows[0][2])

    def test_manual_description_uses_lower_budget(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            write_skill(root, "demo", "x" * (MANUAL_BUDGET + 1), manual=True)
            rows, _ = collect(root)
            self.assertEqual(rows[0][2], MANUAL_BUDGET)
            self.assertGreater(rows[0][1], rows[0][2])

    def test_counts_references(self) -> None:
        with tempfile.TemporaryDirectory() as tmp:
            root = Path(tmp)
            write_skill(root, "demo", "short")
            refs = root / "skills" / "demo" / "demo" / "references"
            refs.mkdir()
            (refs / "note.md").write_text("12345", encoding="utf-8")
            _, sizes = collect(root)
            self.assertEqual(sizes["references"], 5)


if __name__ == "__main__":
    unittest.main()

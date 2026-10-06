package publisher_test

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"researchcurator/curator"
	"researchcurator/publisher"
)

func validReportJSON(t *testing.T) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join("..", "examples", "report.json"))
	if err != nil {
		t.Fatal(err)
	}
	if _, err := curator.DecodeReport(data); err != nil {
		t.Fatalf("example report invalid: %v", err)
	}
	return data
}

func TestDefaultDestinationAvoidsSkillSourceAndPreservesCollisions(t *testing.T) {
	base := t.TempDir()
	workspace := filepath.Join(base, "source")
	home := filepath.Join(base, "home")
	module := filepath.Join(workspace, "research-curator")
	if err := os.MkdirAll(module, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(home, 0755); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{filepath.Join(module, "SKILL.md"), filepath.Join(module, "go.mod")} {
		if err := os.WriteFile(path, []byte("fixture"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	first, err := publisher.DefaultDestination("A useful Chinese question?", workspace, home)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(first, filepath.Join(home, "Research Reports")) || strings.Contains(first, workspace) {
		t.Fatalf("skill source workspace selected as output: %s", first)
	}
	if err := os.Mkdir(first, 0755); err != nil {
		t.Fatal(err)
	}
	second, err := publisher.DefaultDestination("A useful Chinese question?", workspace, home)
	if err != nil {
		t.Fatal(err)
	}
	if second == first || !strings.HasPrefix(second, first+"-") {
		t.Fatalf("default collision did not choose a fresh folder: first=%s second=%s", first, second)
	}
	ordinary := filepath.Join(base, "ordinary")
	if err := os.Mkdir(ordinary, 0755); err != nil {
		t.Fatal(err)
	}
	other, err := publisher.DefaultDestination("ordinary project question", ordinary, home)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.HasPrefix(other, filepath.Join(base, "ordinary", "research-reports")) {
		t.Fatalf("ordinary workspace did not use research-reports: %s", other)
	}
}

func TestWritePublishesSelfContainedIndexToNewFolder(t *testing.T) {
	destination := filepath.Join(t.TempDir(), "synthetic-report")
	if err := publisher.Write(validReportJSON(t), destination); err != nil {
		t.Fatal(err)
	}
	entries, err := os.ReadDir(destination)
	if err != nil {
		t.Fatal(err)
	}
	if len(entries) != 1 || entries[0].Name() != "index.html" {
		t.Fatalf("unexpected publication contents: %v", entries)
	}
	page, err := os.ReadFile(filepath.Join(destination, "index.html"))
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(page, []byte("Synthetic offline report example")) || !bytes.Contains(page, []byte("evidenceMap:'Evidence map'")) {
		t.Fatal("published page lacks report content")
	}
}

func TestWritePreservesExistingAndUnsafeDestinations(t *testing.T) {
	root := t.TempDir()
	existing := filepath.Join(root, "existing")
	if err := os.Mkdir(existing, 0700); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(existing, "keep.txt")
	if err := os.WriteFile(marker, []byte("user data"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := publisher.Write(validReportJSON(t), existing); err == nil {
		t.Fatal("overwrote existing destination")
	}
	if data, err := os.ReadFile(marker); err != nil || string(data) != "user data" {
		t.Fatal("existing user data changed")
	}

	unsafe := root + string(filepath.Separator) + ".." + string(filepath.Separator) + "unsafe-topic"
	if err := publisher.Write(validReportJSON(t), unsafe); err == nil {
		t.Fatal("accepted parent traversal in destination")
	}
}

func TestInterruptedFolderIsPreservedAndFreshNameCanRecover(t *testing.T) {
	root := t.TempDir()
	interrupted := filepath.Join(root, "interrupted")
	if err := os.Mkdir(interrupted, 0700); err != nil {
		t.Fatal(err)
	}
	if err := publisher.Write(validReportJSON(t), interrupted); err == nil {
		t.Fatal("reclaimed an incomplete preexisting folder")
	}
	entries, err := os.ReadDir(interrupted)
	if err != nil || len(entries) != 0 {
		t.Fatalf("modified incomplete folder: %v %v", entries, err)
	}
	fresh := filepath.Join(root, "retry")
	if err := publisher.Write(validReportJSON(t), fresh); err != nil {
		t.Fatal(err)
	}
	entries, err = os.ReadDir(root)
	if err != nil || len(entries) != 2 {
		t.Fatalf("unexpected leftovers: %v %v", entries, err)
	}
}

func TestWriteRejectsReservedWindowsNames(t *testing.T) {
	root := t.TempDir()
	for _, name := range []string{"CON", "NUL.txt", "report.", "report "} {
		if err := publisher.Write(validReportJSON(t), filepath.Join(root, name)); err == nil {
			t.Errorf("accepted reserved or ambiguous topic name %q", name)
		}
	}
}

func TestWriteRejectsSymlinkAncestor(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, "target")
	if err := os.Mkdir(target, 0700); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(root, "alias")
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	if err := publisher.Write(validReportJSON(t), filepath.Join(link, "topic")); err == nil {
		t.Fatal("published through symlink parent")
	}
	files, err := os.ReadDir(target)
	if err != nil || len(files) != 0 {
		t.Fatalf("linked target modified: %v %v", files, err)
	}
}

func TestWriteDoesNotFollowSymlinkDestination(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, "target")
	if err := os.Mkdir(target, 0700); err != nil {
		t.Fatal(err)
	}
	marker := filepath.Join(target, "keep.txt")
	if err := os.WriteFile(marker, []byte("user data"), 0600); err != nil {
		t.Fatal(err)
	}
	link := filepath.Join(root, "topic-link")
	if err := os.Symlink(target, link); err != nil {
		t.Skipf("symlink creation unavailable: %v", err)
	}
	if err := publisher.Write(validReportJSON(t), link); err == nil {
		t.Fatal("published through symlink destination")
	}
	if data, err := os.ReadFile(marker); err != nil || string(data) != "user data" {
		t.Fatal("symlink target changed")
	}
}

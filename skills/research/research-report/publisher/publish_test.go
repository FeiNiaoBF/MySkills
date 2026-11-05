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

func TestDefaultDestinationUsesTempRootAndPreservesCollisions(t *testing.T) {
	tempRoot := t.TempDir()
	first, err := publisher.DefaultDestination("GPT-6 Astra vs Claude Fable 5.1: coding experience", tempRoot)
	if err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(tempRoot, "research-curator")
	if filepath.Dir(first) != root || strings.Contains(first, "run-id") {
		t.Fatalf("default report is not directly under the temp research root: %s", first)
	}
	if err := os.Mkdir(first, 0755); err != nil {
		t.Fatal(err)
	}
	second, err := publisher.DefaultDestination("GPT-6 Astra vs Claude Fable 5.1: coding experience", tempRoot)
	if err != nil {
		t.Fatal(err)
	}
	if second != first+"-2" {
		t.Fatalf("default collision did not use the simple -2 suffix: first=%s second=%s", first, second)
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

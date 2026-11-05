package visualizer_test

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"researchcurator/visualizer"
)

func TestRenderRejectsInvalidInput(t *testing.T) {
	for _, input := range []string{"", "{", "null", "[]", "true", `{} {}`} {
		if _, err := visualizer.Render([]byte(input)); err == nil {
			t.Errorf("accepted %q", input)
		}
	}
}

func TestRenderPreservesUntrustedValuesAndLargeNumbers(t *testing.T) {
	input := []byte(`{"metadata":{"title":"{{APP}} {{CYTOSCAPE}} {{RUN_JSON}} <!-- </ScRiPt>"},"large":900719925474099312345,"separator":"\u2029"}`)
	out, err := visualizer.Render(input)
	if err != nil {
		t.Fatal(err)
	}
	re := regexp.MustCompile(`(?s)<script id="run-data" type="application/json">(.*?)</script>`)
	match := re.FindSubmatch(out)
	if len(match) != 2 {
		t.Fatal("missing JSON data")
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(match[1], &decoded); err != nil {
		t.Fatal(err)
	}
	if string(decoded["large"]) != "900719925474099312345" {
		t.Fatal("large integer changed")
	}
	if !bytes.Contains(match[1], []byte("{{APP}} {{CYTOSCAPE}} {{RUN_JSON}}")) {
		t.Fatal("data treated as a template")
	}
	if bytes.Contains(match[1], []byte("\u2029")) {
		t.Fatal("unescaped Unicode separator")
	}
}

func TestRenderSafeOfflineDocument(t *testing.T) {
	input := []byte(`{"schema_version":"1","metadata":{"title":"</script><script>alert(1)</script>\u2028"},"sources":[{"id":"s1","url":"javascript:alert(1)"}],"nodes":[],"edges":[]}`)
	out, err := visualizer.Render(input)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Contains(out, []byte("<!doctype html>")) {
		t.Fatal("not an HTML document")
	}
	re := regexp.MustCompile(`(?s)<script id="run-data" type="application/json">(.*?)</script>`)
	match := re.FindSubmatch(out)
	if len(match) != 2 {
		t.Fatal("missing JSON data")
	}
	var restored map[string]any
	if err := json.Unmarshal(match[1], &restored); err != nil {
		t.Fatal(err)
	}
	if restored["metadata"].(map[string]any)["title"] != "</script><script>alert(1)</script>\u2028" {
		t.Fatal("data changed")
	}
	if bytes.Contains(match[1], []byte("</script")) || bytes.Contains(match[1], []byte("\u2028")) {
		t.Fatal("unsafe script data")
	}
	for _, forbidden := range []string{`<script src=`, `<link rel="stylesheet"`, `fetch(`, `XMLHttpRequest`, `innerHTML`, `eval(`} {
		// Inspect application markup/scripts only; the official engine is separately integrity-pinned.
		app := strings.Split(string(out), "/* RESEARCH_CURATOR_APP */")
		if strings.Contains(string(out[:bytes.Index(out, []byte("/* CYTOSCAPE_VENDOR */"))]), forbidden) || (len(app) > 1 && strings.Contains(app[1], forbidden)) {
			t.Errorf("unsafe/external feature: %s", forbidden)
		}
	}
	if !bytes.Contains(out, []byte("cytoscape")) || !bytes.Contains(out, []byte("connect-src 'none'")) {
		t.Fatal("missing offline engine or CSP")
	}
}

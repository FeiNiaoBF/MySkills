package visualizer_test

import (
	"bytes"
	"encoding/json"
	"regexp"
	"strings"
	"testing"

	"researchcurator/visualizer"
)

const reportFixture = `{"report_version":"1.0","run":{"version":"1.0","id":"goal","metadata":{"created_at":"2026-01-01T00:00:00Z","status":"finalized","title":"Synthetic research"},"contract":{"question":"Synthetic question"},"sources":[{"id":"s1","title":"Synthetic source","url":"https://example.org/source","content":"Synthetic quoted evidence."}],"claims":[{"id":"c1","text":"A synthetic claim","evidence":[{"source_id":"s1","quote":"Synthetic quoted evidence.","locator":"paragraph 1","relation":"supports","verification":"verified"}]}],"conclusions":[],"queries":[],"events":[],"decisions":[],"rejected_sources":[],"graph":{"nodes":[{"id":"s1","type":"Source","label":"Synthetic source"},{"id":"c1","type":"Claim","label":"A synthetic claim"}],"edges":[{"from":"s1","to":"c1","type":"supports"}]},"coverage":{}},"research":{"audience":"general reader","purpose":"explain the question","max_rounds":8,"low_gain_window":3,"rounds":[],"coverage":[],"gaps":[],"stop":{"reason":"user_stopped","round_ids":[],"rationale":"synthetic fixture"}},"article":{"title":"Synthetic article","language":"en","output_form":"article","lead":"The short answer.","sections":[{"id":"sec1","heading":"Finding","blocks":[{"id":"b1","kind":"paragraph","role":"evidence","text":"The source supports the claim.","claim_ids":["c1"],"conclusion_ids":[]}]}]}}`

func TestRenderReportIsReaderFirstAndUsesArticleLanguage(t *testing.T) {
	input := strings.Replace(reportFixture, `"language":"en"`, `"language":"zh-CN"`, 1)
	out, err := visualizer.RenderReport([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	html := string(out)
	article := strings.Index(html, `id="article-panel"`)
	sources := strings.Index(html, `id="evidence-register"`)
	graph := strings.Index(html, `id="evidence-map"`)
	audit := strings.Index(html, `id="audit-panel"`)
	if article < 0 || sources <= article || graph <= sources || audit <= graph {
		t.Fatalf("reader-first sections out of order: article=%d sources=%d graph=%d audit=%d", article, sources, graph, audit)
	}
	for _, want := range []string{"str(article?.language)", "zh-CN", "支持", "Important limits", "data-i18n=\"evidenceMap\""} {
		if !strings.Contains(html, want) {
			t.Errorf("localized reader report missing %q", want)
		}
	}
}

func TestRenderReportKeepsUntrustedTextInertAndOffline(t *testing.T) {
	input := strings.Replace(reportFixture, "The short answer.", `</script><img src=x onerror=alert(1)>`, 1)
	out, err := visualizer.RenderReport([]byte(input))
	if err != nil {
		t.Fatal(err)
	}
	re := regexp.MustCompile(`(?s)<script id="run-data" type="application/json">(.*?)</script>`)
	match := re.FindSubmatch(out)
	if len(match) != 2 {
		t.Fatal("missing embedded report data")
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(match[1], &decoded); err != nil {
		t.Fatal(err)
	}
	for _, forbidden := range []string{"<img", "<script src=", "fetch(", "XMLHttpRequest", "innerHTML", "eval("} {
		if bytes.Contains(out, []byte(forbidden)) {
			t.Errorf("unsafe/external feature present: %s", forbidden)
		}
	}
	if !bytes.Contains(out, []byte("connect-src 'none'")) || !bytes.Contains(out, []byte("cytoscape")) {
		t.Fatal("offline CSP or bundled graph engine missing")
	}
}

func TestRenderReportRejectsMalformedEnvelope(t *testing.T) {
	for _, input := range []string{"", "{}", `{"report_version":"1.0","run":{}}`} {
		if _, err := visualizer.RenderReport([]byte(input)); err == nil {
			t.Errorf("accepted invalid envelope: %s", input)
		}
	}
}

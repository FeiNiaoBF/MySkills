// Package visualizer renders research run data as a self-contained offline report.
package visualizer

import (
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"strings"
)

//go:embed templates/report.html templates/app.js vendor/cytoscape.min.js
var assets embed.FS

// Render accepts a legacy run JSON object. Unknown fields are preserved for inspection;
// schema validation belongs to the caller. No runtime filesystem or network access is used.
func Render(runJSON []byte) ([]byte, error) {
	return renderDocument(runJSON, false)
}

// RenderReport renders a versioned report envelope. Full domain validation belongs to
// the caller; this boundary verifies the envelope shape needed by the view.
func RenderReport(reportJSON []byte) ([]byte, error) {
	var root map[string]json.RawMessage
	if err := json.Unmarshal(reportJSON, &root); err != nil {
		return nil, fmt.Errorf("visualizer: invalid report JSON: %w", err)
	}
	if root == nil {
		return nil, fmt.Errorf("visualizer: report must be a JSON object")
	}
	var version string
	var run, research, article json.RawMessage
	if err := json.Unmarshal(root["report_version"], &version); err != nil || version != "1.0" {
		return nil, fmt.Errorf("visualizer: unsupported report version")
	}
	for key, target := range map[string]*json.RawMessage{"run": &run, "research": &research, "article": &article} {
		if len(root[key]) == 0 || string(root[key]) == "null" {
			return nil, fmt.Errorf("visualizer: report is missing %s", key)
		}
		*target = root[key]
		var object map[string]json.RawMessage
		if err := json.Unmarshal(*target, &object); err != nil || object == nil {
			return nil, fmt.Errorf("visualizer: report %s must be an object", key)
		}
	}
	return renderDocument(reportJSON, true)
}

func renderDocument(runJSON []byte, reportMode bool) ([]byte, error) {
	var root map[string]json.RawMessage
	if err := json.Unmarshal(runJSON, &root); err != nil {
		return nil, fmt.Errorf("visualizer: invalid run JSON: %w", err)
	}
	if root == nil {
		return nil, fmt.Errorf("visualizer: run must be a JSON object")
	}
	// Marshal RawMessage values again to apply HTML-safe escaping without coercing
	// large integers through float64 or losing their exact representation.
	safe, err := json.Marshal(root)
	if err != nil {
		return nil, fmt.Errorf("visualizer: encode run: %w", err)
	}
	template, err := assets.ReadFile("templates/report.html")
	if err != nil {
		return nil, err
	}
	vendor, err := assets.ReadFile("vendor/cytoscape.min.js")
	if err != nil {
		return nil, err
	}
	app, err := assets.ReadFile("templates/app.js")
	if err != nil {
		return nil, err
	}
	// The engine is trusted pinned code, not research input. Escape script-closing
	// sequences defensively should a future distribution include one in a string.
	engine := strings.ReplaceAll(string(vendor), "</script", "<\\/script")
	replacer := strings.NewReplacer("{{RUN_JSON}}", string(safe), "{{CYTOSCAPE}}", engine, "{{APP}}", string(app), "{{REPORT_MODE}}", fmt.Sprint(reportMode))
	var out bytes.Buffer
	_, err = replacer.WriteString(&out, string(template))
	return out.Bytes(), err
}

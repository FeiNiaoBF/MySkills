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

// Render accepts a JSON object. Unknown fields are preserved for inspection; schema
// validation belongs to the caller. All untrusted content is rendered as text.
// No filesystem, browser, service, or network dependency is needed at runtime.
func Render(runJSON []byte) ([]byte, error) {
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
	replacer := strings.NewReplacer("{{RUN_JSON}}", string(safe), "{{CYTOSCAPE}}", engine, "{{APP}}", string(app))
	var out bytes.Buffer
	_, err = replacer.WriteString(&out, string(template))
	return out.Bytes(), err
}

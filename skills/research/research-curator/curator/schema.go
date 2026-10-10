package curator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"researchcurator/schema"
	"unicode/utf8"
)

// Decode applies the embedded JSON Schema subset and semantic validation.
func Decode(data []byte) (*Run, error) {
	if !utf8.Valid(data) {
		return nil, fmt.Errorf("JSON must be UTF-8")
	}
	scan := json.NewDecoder(bytes.NewReader(data))
	if err := scanValue(scan, 0); err != nil {
		return nil, err
	}
	var raw any
	d := json.NewDecoder(bytes.NewReader(data))
	d.UseNumber()
	if err := d.Decode(&raw); err != nil {
		return nil, err
	}
	if err := d.Decode(new(any)); err != io.EOF {
		return nil, fmt.Errorf("trailing JSON: %v", err)
	}
	var s map[string]any
	if err := json.Unmarshal(schema.Run, &s); err != nil {
		return nil, err
	}
	if err := checkSchema(raw, s, "$"); err != nil {
		return nil, err
	}
	var r Run
	if err := json.Unmarshal(data, &r); err != nil {
		return nil, err
	}
	if err := validateSemantic(&r); err != nil {
		return nil, err
	}
	return &r, nil
}

// Validate checks serialized shape as well as semantic invariants.
func Validate(r *Run) error {
	if r == nil {
		return fmt.Errorf("nil run")
	}
	b, e := json.Marshal(r)
	if e != nil {
		return e
	}
	_, e = Decode(b)
	return e
}
func scanValue(d *json.Decoder, depth int) error {
	if depth > 100 {
		return fmt.Errorf("JSON nesting exceeds 100")
	}
	tok, err := d.Token()
	if err != nil {
		return err
	}
	delim, ok := tok.(json.Delim)
	if !ok {
		return nil
	}
	switch delim {
	case '{':
		seen := map[string]bool{}
		for d.More() {
			key, err := d.Token()
			if err != nil {
				return err
			}
			k, ok := key.(string)
			if !ok {
				return fmt.Errorf("invalid object key")
			}
			if seen[k] {
				return fmt.Errorf("duplicate JSON property %q", k)
			}
			seen[k] = true
			if err := scanValue(d, depth+1); err != nil {
				return err
			}
		}
	case '[':
		for d.More() {
			if err := scanValue(d, depth+1); err != nil {
				return err
			}
		}
	default:
		return fmt.Errorf("unexpected delimiter")
	}
	_, err = d.Token()
	return err
}

func checkSchema(v any, s map[string]any, path string) error {
	fail := func(why string) error { return fmt.Errorf("%s: %s", path, why) }
	if choices, ok := s["enum"].([]any); ok {
		found := false
		for _, c := range choices {
			if c == v {
				found = true
			}
		}
		if !found {
			return fail("outside enum")
		}
	}
	switch s["type"] {
	case "object":
		obj, ok := v.(map[string]any)
		if !ok {
			return fail("expected object")
		}
		props := s["properties"].(map[string]any)
		for _, k := range s["required"].([]any) {
			if _, ok := obj[k.(string)]; !ok {
				return fail("missing " + k.(string))
			}
		}
		for k, x := range obj {
			p, ok := props[k]
			if !ok {
				if s["additionalProperties"] == true {
					continue
				}
				return fail("unknown property " + k)
			}
			if e := checkSchema(x, p.(map[string]any), path+"."+k); e != nil {
				return e
			}
		}
	case "array":
		a, ok := v.([]any)
		if !ok {
			return fail("expected array (not null)")
		}
		if n, ok := s["minItems"].(float64); ok && len(a) < int(n) {
			return fail("too few items")
		}
		for i, x := range a {
			if e := checkSchema(x, s["items"].(map[string]any), fmt.Sprintf("%s[%d]", path, i)); e != nil {
				return e
			}
		}
	case "string":
		x, ok := v.(string)
		if !ok {
			return fail("expected string")
		}
		if n, ok := s["minLength"].(float64); ok && utf8.RuneCountInString(x) < int(n) {
			return fail("string too short")
		}
	case "boolean":
		if _, ok := v.(bool); !ok {
			return fail("expected boolean")
		}
	case "number", "integer":
		x, ok := v.(json.Number)
		if !ok {
			return fail("expected number")
		}
		n, e := x.Float64()
		if e != nil || math.IsNaN(n) || math.IsInf(n, 0) {
			return fail("invalid number")
		}
		if s["type"] == "integer" && n != math.Trunc(n) {
			return fail("expected integer")
		}
		if m, ok := s["minimum"].(float64); ok && n < m {
			return fail("below minimum")
		}
		if m, ok := s["maximum"].(float64); ok && n > m {
			return fail("above maximum")
		}
	default:
		return fail("unsupported schema type")
	}
	return nil
}

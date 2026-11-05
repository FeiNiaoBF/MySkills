package publisher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// DefaultDestination selects a safe topic folder under the temp research root.
func DefaultDestination(topic, tempRoot string) (string, error) {
	if strings.TrimSpace(topic) == "" || tempRoot == "" {
		return "", fmt.Errorf("default output requires a topic and temporary directory")
	}
	root := filepath.Join(realPath(tempRoot), "research-curator")
	if err := prepareRoot(root); err != nil {
		return "", err
	}
	base := topicName(topic)
	for suffix := 1; ; suffix++ {
		name := base
		if suffix > 1 {
			name = fmt.Sprintf("%s-%d", base, suffix)
		}
		candidate := filepath.Join(root, name)
		if _, err := os.Lstat(candidate); os.IsNotExist(err) {
			return candidate, nil
		} else if err != nil {
			return "", fmt.Errorf("inspect default report folder: %w", err)
		}
	}
}

func realPath(path string) string {
	absolute, err := filepath.Abs(path)
	if err != nil {
		return filepath.Clean(path)
	}
	if resolved, err := filepath.EvalSymlinks(absolute); err == nil {
		return resolved
	}
	return absolute
}

func isSkillSource(workspace string) bool {
	for dir := workspace; ; dir = filepath.Dir(dir) {
		if fileExists(filepath.Join(dir, "SKILL.md")) && fileExists(filepath.Join(dir, "go.mod")) {
			return true
		}
		module := filepath.Join(dir, "research-curator")
		if fileExists(filepath.Join(module, "SKILL.md")) && fileExists(filepath.Join(module, "go.mod")) {
			return true
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			return false
		}
	}
}

func fileExists(path string) bool {
	info, err := os.Stat(path)
	return err == nil && !info.IsDir()
}

func prepareRoot(root string) error {
	if info, err := os.Lstat(root); err == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return fmt.Errorf("default report root must be a real directory: %s", root)
		}
		return checkAncestors(root)
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect default report root: %w", err)
	}
	parent := filepath.Dir(root)
	if err := checkAncestors(parent); err != nil {
		return err
	}
	if err := os.Mkdir(root, 0755); err != nil && !os.IsExist(err) {
		return fmt.Errorf("create default report root: %w", err)
	}
	return checkAncestors(root)
}

func topicName(question string) string {
	var out strings.Builder
	space := false
	for _, r := range strings.TrimSpace(question) {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			if space && out.Len() > 0 {
				out.WriteByte('-')
			}
			out.WriteRune(unicode.ToLower(r))
			space = false
		} else if r == '-' || r == '_' {
			if out.Len() > 0 {
				out.WriteRune(r)
			}
			space = false
		} else {
			space = true
		}
		if out.Len() >= 48 {
			break
		}
	}
	name := strings.Trim(out.String(), "-_. ")
	if name == "" {
		return "research-report"
	}
	device := strings.ToUpper(strings.SplitN(name, ".", 2)[0])
	if device == "CON" || device == "PRN" || device == "AUX" || device == "NUL" {
		return "research-" + name
	}
	return name
}

package publisher

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"researchcurator/curator"
	"researchcurator/visualizer"
)

// Write validates and publishes a self-contained report at destination/index.html.
// The destination must not exist; only the unique staging directory created here is cleaned up.
func Write(reportJSON []byte, destination string) error {
	report, err := curator.DecodeReport(reportJSON)
	if err != nil {
		return fmt.Errorf("decode report: %w", err)
	}
	if err := curator.ValidatePublication(report); err != nil {
		return fmt.Errorf("validate publication: %w", err)
	}
	html, err := visualizer.RenderReport(reportJSON)
	if err != nil {
		return fmt.Errorf("render report: %w", err)
	}
	cleanDestination, err := safeDestination(destination)
	if err != nil {
		return err
	}
	cleanDestination, err = filepath.Abs(cleanDestination)
	if err != nil {
		return err
	}
	parent := filepath.Dir(cleanDestination)
	if err := checkAncestors(parent); err != nil {
		return err
	}
	base := filepath.Base(cleanDestination)
	if info, err := os.Stat(parent); err != nil || !info.IsDir() {
		if err != nil {
			return fmt.Errorf("destination parent is unavailable: %w", err)
		}
		return fmt.Errorf("destination parent is not a directory")
	}
	if _, err := os.Lstat(cleanDestination); err == nil {
		return fmt.Errorf("destination already exists; preserve it and choose a new topic folder (an interrupted attempt may have left an incomplete folder)")
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("inspect destination: %w", err)
	}

	stage, err := os.MkdirTemp(parent, "."+base+".researchcurator-")
	if err != nil {
		return fmt.Errorf("create publication staging directory: %w", err)
	}
	defer os.RemoveAll(stage)
	file, err := os.OpenFile(filepath.Join(stage, "index.html"), os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0644)
	if err != nil {
		return fmt.Errorf("create report entry point: %w", err)
	}
	if _, err := file.Write(html); err != nil {
		file.Close()
		return fmt.Errorf("write report entry point: %w", err)
	}
	if err := file.Sync(); err != nil {
		file.Close()
		return fmt.Errorf("sync report entry point: %w", err)
	}
	if err := file.Close(); err != nil {
		return fmt.Errorf("close report entry point: %w", err)
	}
	if err := os.Mkdir(cleanDestination, 0755); err != nil {
		return fmt.Errorf("create new report folder: %w", err)
	}
	ownsDestination := true
	defer func() {
		if ownsDestination {
			// Remove only our still-empty directory. If another process added data,
			// preserve it rather than cleaning a path we no longer own exclusively.
			_ = os.Remove(cleanDestination)
		}
	}()
	if err := os.Link(filepath.Join(stage, "index.html"), filepath.Join(cleanDestination, "index.html")); err != nil {
		return fmt.Errorf("publish report entry point without replacement: %w", err)
	}
	ownsDestination = false
	return nil
}

// Check existing ancestors with Lstat; never stage or clean through a link.
// This assumes the caller controls the parent during publication; it is not a
// sandbox against a hostile process concurrently replacing path components.
func checkAncestors(path string) error {
	for {
		info, err := os.Lstat(path)
		if err != nil {
			return fmt.Errorf("inspect destination ancestor: %w", err)
		}
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return fmt.Errorf("destination ancestor must be a real directory, not a symlink/junction: %s", path)
		}
		parent := filepath.Dir(path)
		if parent == path {
			return nil
		}
		path = parent
	}
}

func safeDestination(destination string) (string, error) {
	if strings.TrimSpace(destination) == "" || destination == "-" {
		return "", fmt.Errorf("publish requires a topic directory")
	}
	for _, component := range strings.FieldsFunc(strings.ReplaceAll(destination, `\`, "/"), func(r rune) bool { return r == '/' }) {
		if component == ".." {
			return "", fmt.Errorf("destination must not contain parent traversal")
		}
	}
	clean := filepath.Clean(destination)
	base := filepath.Base(clean)
	if clean == "." || clean == string(filepath.Separator) || base == "." || base == ".." || base == "" {
		return "", fmt.Errorf("destination must name a topic directory")
	}
	if strings.ContainsAny(base, "<>:\"|?*\x00") || strings.HasSuffix(base, ".") || strings.HasSuffix(base, " ") {
		return "", fmt.Errorf("topic directory name contains unsafe characters")
	}
	deviceName := strings.ToUpper(strings.SplitN(base, ".", 2)[0])
	if deviceName == "CON" || deviceName == "PRN" || deviceName == "AUX" || deviceName == "NUL" || len(deviceName) == 4 && (strings.HasPrefix(deviceName, "COM") || strings.HasPrefix(deviceName, "LPT")) && deviceName[3] >= '1' && deviceName[3] <= '9' {
		return "", fmt.Errorf("topic directory name is reserved by the operating system")
	}
	return clean, nil
}

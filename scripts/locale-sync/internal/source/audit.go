package source

import (
	"bufio"
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
)

// KeyUsage tracks where a t() key is called in source.
type KeyUsage struct {
	Key  string
	File string
	Line int
}

// AuditResult holds orphaned and missing key findings.
type AuditResult struct {
	// UsedKeys: all t('key') calls found in source files.
	UsedKeys []KeyUsage
	// OrphanedKeys: defined in YAML but never called in source.
	OrphanedKeys []string
	// MissingKeys: called in source but not defined in YAML.
	MissingKeys []KeyUsage
}

// tCallRe captures the key argument from t(), I18n.t(), and translate() calls.
// Handles both dot-notation keys and relative keys (starting with .).
var tCallRe = regexp.MustCompile(`(?:I18n\.t|translate|\bt)\s*[("']\s*[."']?([\w.]+)["']?\s*[,)]`)

// Audit cross-references all t() calls in source against the provided YAML keys.
// definedKeys is a set of full dot-notation keys from the locale YAML files.
func Audit(root, lang string, definedKeys map[string]bool) (AuditResult, error) {
	appDir := filepath.Join(root, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		appDir = root
	}

	fileCh := make(chan string, 128)
	usageCh := make(chan KeyUsage, 512)

	go func() {
		defer close(fileCh)
		_ = filepath.WalkDir(appDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			if sourceExts[filepath.Ext(path)] {
				fileCh <- path
			}
			return nil
		})
	}()

	const workers = 8
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileCh {
				for _, u := range extractTCalls(path, root) {
					usageCh <- u
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(usageCh)
	}()

	var allUsages []KeyUsage
	usedSet := map[string]bool{}
	for u := range usageCh {
		allUsages = append(allUsages, u)
		usedSet[u.Key] = true
	}

	// Orphaned: defined but never called.
	// We check both the full key and without the leading lang prefix.
	var orphaned []string
	for key := range definedKeys {
		stripped := stripLangPrefix(key, lang)
		if !usedSet[stripped] && !usedSet[key] {
			orphaned = append(orphaned, key)
		}
	}
	sort.Strings(orphaned)

	// Missing: called but not defined.
	var missing []KeyUsage
	for _, u := range allUsages {
		full := lang + "." + u.Key
		if !definedKeys[full] && !definedKeys[u.Key] {
			missing = append(missing, u)
		}
	}

	return AuditResult{
		UsedKeys:     allUsages,
		OrphanedKeys: orphaned,
		MissingKeys:  missing,
	}, nil
}

func extractTCalls(path, root string) []KeyUsage {
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	short, _ := filepath.Rel(root, path)
	var out []KeyUsage
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1<<20), 1<<20)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		if isComment(strings.TrimSpace(line)) {
			continue
		}
		for _, m := range tCallRe.FindAllStringSubmatch(line, -1) {
			key := strings.Trim(m[1], "\"'")
			key = strings.TrimPrefix(key, ".")
			if key == "" || strings.ContainsAny(key, " \t") {
				continue
			}
			out = append(out, KeyUsage{Key: key, File: short, Line: lineNum})
		}
	}
	return out
}

func stripLangPrefix(key, lang string) string {
	prefix := lang + "."
	if strings.HasPrefix(key, prefix) {
		return key[len(prefix):]
	}
	return key
}

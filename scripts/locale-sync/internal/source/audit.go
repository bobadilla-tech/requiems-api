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
	// RelativeKeys: t('.key') calls that need view-path context to resolve;
	// excluded from MissingKeys to avoid false positives.
	RelativeKeys []KeyUsage
}

// tCallRe captures the key argument from t(), I18n.t(), and translate() calls.
// Handles both dot-notation keys and relative keys (starting with .).
var tCallRe = regexp.MustCompile(`(?:I18n\.t|translate|\bt)\s*\(?\s*["'](\.?[\w.]+)["']`)

// Audit cross-references all t() calls in source against the provided YAML keys.
// definedKeys is a set of full dot-notation keys from the locale YAML files.
func Audit(root, lang string, definedKeys map[string]bool) (AuditResult, error) {
	appDir := filepath.Join(root, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		appDir = root
	}

	fileCh := make(chan string, 128)

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

	type chanMsg struct {
		abs []KeyUsage
		rel []KeyUsage
	}
	msgCh := make(chan chanMsg, 512)

	const workers = 8
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileCh {
				r := extractTCalls(path, root)
				msgCh <- chanMsg{abs: r.absolute, rel: r.relative}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(msgCh)
	}()

	var allUsages, allRelative []KeyUsage
	usedSet := map[string]bool{}
	for msg := range msgCh {
		allUsages = append(allUsages, msg.abs...)
		for _, u := range msg.abs {
			usedSet[u.Key] = true
		}
		// Resolve relative keys (t('.leaf')) using the view file path so they
		// don't cause false-positive orphan reports for keys only used this way.
		for _, u := range msg.rel {
			allRelative = append(allRelative, u)
			if resolved := ResolveRelativeKey(u.File, u.Key); resolved != "" {
				usedSet[resolved] = true
				// Also add to allUsages so the missing-key pass can detect
				// relative keys that have no YAML definition.
				allUsages = append(allUsages, KeyUsage{Key: resolved, File: u.File, Line: u.Line})
			}
		}
	}

	// Orphaned: defined but never called.
	// We check both the full key and without the leading lang prefix.
	// Pluralization forms (.one/.other/.zero/.two/.few/.many) are NOT orphaned when
	// their base key is called — Rails resolves t('foo', count: n) → foo.one / foo.other.
	var orphaned []string
	for key := range definedKeys {
		stripped := stripLangPrefix(key, lang)
		if usedSet[stripped] || usedSet[key] {
			continue
		}
		if isPluralizationForm(stripped, usedSet) || isPluralizationForm(key, usedSet) {
			continue
		}
		orphaned = append(orphaned, key)
	}
	sort.Strings(orphaned)

	// Missing: called but not defined.
	// Skip pluralization bases — t('foo.bar', count: n) resolves to foo.bar.one / foo.bar.other
	// at Rails runtime, so the base key foo.bar will never appear as a scalar in YAML.
	var missing []KeyUsage
	for _, u := range allUsages {
		full := lang + "." + u.Key
		if definedKeys[full] || definedKeys[u.Key] {
			continue
		}
		if isPluralizationBase(u.Key, full, definedKeys) {
			continue
		}
		missing = append(missing, u)
	}

	return AuditResult{
		UsedKeys:     allUsages,
		OrphanedKeys: orphaned,
		MissingKeys:  missing,
		RelativeKeys: allRelative,
	}, nil
}

type tCallResult struct {
	absolute []KeyUsage
	relative []KeyUsage
}

func extractTCalls(path, root string) tCallResult {
	f, err := os.Open(path)
	if err != nil {
		return tCallResult{}
	}
	defer f.Close()

	short, _ := filepath.Rel(root, path)
	var res tCallResult
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
			raw := strings.Trim(m[1], "\"'")
			if strings.HasPrefix(raw, ".") {
				// Relative key (t('.title')) — needs view-path context to resolve.
				// Track separately to avoid false-positive missing-key reports.
				key := strings.TrimPrefix(raw, ".")
				if key != "" && !strings.ContainsAny(key, " \t") {
					res.relative = append(res.relative, KeyUsage{Key: key, File: short, Line: lineNum})
				}
				continue
			}
			if raw == "" || strings.ContainsAny(raw, " \t") {
				continue
			}
			res.absolute = append(res.absolute, KeyUsage{Key: raw, File: short, Line: lineNum})
		}
	}
	return res
}

// resolveRelativeKey converts a relative t('.leaf') call and its source file
// path into the full dot-notation key Rails would resolve it to.
// e.g. "app/views/admin/users/index.html.erb" + "title" → "admin.users.index.title"
func ResolveRelativeKey(shortPath, leaf string) string {
	p := filepath.ToSlash(shortPath)
	p = strings.TrimPrefix(p, "app/views/")
	for _, ext := range []string{".html.erb", ".html.haml", ".erb", ".haml", ".rb"} {
		if strings.HasSuffix(p, ext) {
			p = p[:len(p)-len(ext)]
			break
		}
	}
	parts := strings.Split(p, "/")
	for i, pt := range parts {
		parts[i] = strings.TrimPrefix(pt, "_")
	}
	return strings.Join(parts, ".") + "." + strings.TrimPrefix(leaf, ".")
}

// isPluralizationForm returns true when key ends with a Rails pluralization suffix
// (.one/.other/.zero/.two/.few/.many) AND the base key (without suffix) is called
// in source. Used by the orphan checker so prune never deletes live plural forms.
func isPluralizationForm(key string, usedSet map[string]bool) bool {
	for _, form := range []string{"one", "other", "zero", "two", "few", "many"} {
		suffix := "." + form
		if strings.HasSuffix(key, suffix) {
			base := key[:len(key)-len(suffix)]
			if usedSet[base] {
				return true
			}
		}
	}
	return false
}

// isPluralizationBase returns true when key is a Rails pluralization base:
// t('foo.bar', count: n) resolves to foo.bar.one / foo.bar.other at runtime,
// so foo.bar will never appear as a scalar in YAML.
func isPluralizationBase(key, fullKey string, definedKeys map[string]bool) bool {
	for _, form := range []string{"one", "other", "zero", "two", "few", "many"} {
		if definedKeys[fullKey+"."+form] || definedKeys[key+"."+form] {
			return true
		}
	}
	return false
}

func stripLangPrefix(key, lang string) string {
	prefix := lang + "."
	if strings.HasPrefix(key, prefix) {
		return key[len(prefix):]
	}
	return key
}

// BareTCallIssue is a bare t() call found in a Ruby file where it is not defined
// at runtime (models, jobs, services, workers, lib). These raise NoMethodError in
// production — I18n.t() must be used instead.
type BareTCallIssue struct {
	File string // relative path from Rails root
	Line int
	Key  string // the key argument, e.g. "admin.users.title"
}

// bareTCallOnlyRe matches standalone t('key') / t("key") — excludes I18n.t() and
// translate() by requiring a non-word, non-dot character (or start-of-line) before 't'.
var bareTCallOnlyRe = regexp.MustCompile(`(?:^|[^.\w])t\s*\(\s*["']([\w.]+)["']`)

// invalidTCallDirs are Rails directories where bare t() is not available at runtime.
// Controllers, mailers, and helpers are excluded (they include the translation helper).
var invalidTCallDirs = []string{
	"app/models/",
	"app/jobs/",
	"app/services/",
	"app/workers/",
	"app/decorators/",
	"lib/",
}

// CheckBareTCalls scans .rb files in non-helper Rails layers for bare t() calls.
// Returns one BareTCallIssue per offending line; callers should use I18n.t() there.
func CheckBareTCalls(root string) ([]BareTCallIssue, error) {
	var issues []BareTCallIssue

	dirs := make([]string, 0, len(invalidTCallDirs)+1)
	for _, d := range invalidTCallDirs {
		full := filepath.Join(root, filepath.FromSlash(d))
		if _, err := os.Stat(full); err == nil {
			dirs = append(dirs, full)
		}
	}

	for _, dir := range dirs {
		err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, walkErr error) error {
			if walkErr != nil || d.IsDir() || filepath.Ext(path) != ".rb" {
				return nil
			}
			rel, _ := filepath.Rel(root, path)
			found, err := scanBareTCalls(path, filepath.ToSlash(rel))
			if err == nil {
				issues = append(issues, found...)
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return issues, nil
}

func scanBareTCalls(path, relSlash string) ([]BareTCallIssue, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var issues []BareTCallIssue
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1<<20), 1<<20)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)
		if isComment(trimmed) {
			continue
		}
		// If the line already uses I18n.t or translate(), skip — not a bare t() issue.
		if strings.Contains(line, "I18n.t") || strings.Contains(line, "translate(") {
			continue
		}
		for _, m := range bareTCallOnlyRe.FindAllStringSubmatch(line, -1) {
			key := m[1]
			if key != "" && !strings.ContainsAny(key, " \t") {
				issues = append(issues, BareTCallIssue{File: relSlash, Line: lineNum, Key: key})
			}
		}
	}
	return issues, scanner.Err()
}

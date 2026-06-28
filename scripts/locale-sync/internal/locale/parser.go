package locale

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"gopkg.in/yaml.v3"
)

// Entry is a single flattened i18n key-value pair from a locale file.
type Entry struct {
	Key       string // full dot-notation path e.g. "en.tools.quotes.demo.error_rate_limit"
	KeyName   string // leaf key name e.g. "error_rate_limit"
	Value     string
	File      string // absolute file path
	ShortPath string // relative path from locales root
	Line      int    // line in source YAML
}

// Scan walks config/locales/{lang}/ (and config/locales/ root) concurrently
// and returns all flattened entries for the given language.
func Scan(root, lang string) ([]Entry, error) {
	dirs := candidateDirs(root, lang)

	type result struct {
		entries []Entry
		err     error
	}

	fileCh := make(chan string, 64)
	resultCh := make(chan result, 64)

	// walkErrCh captures the first directory-walk failure (permission denied, etc.)
	// so Scan can surface it rather than silently returning partial results.
	walkErrCh := make(chan error, 1)

	// Producer: find YAML files.
	// When a lang-specific subdir exists (e.g. config/locales/en/), walk that
	// directory fully. Also walk config/locales/ root level but skip any
	// subdirectories to avoid double-scanning.
	go func() {
		defer close(fileCh)
		seen := map[string]bool{}
		for i, dir := range dirs {
			isRoot := i > 0 // dirs[0] is always the lang-specific subdir if it exists
			if err := filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
				if err != nil {
					return err // propagate so WalkDir surfaces it
				}
				// For the root locales dir, skip subdirectories (lang-subdirs
				// are already scanned separately).
				if d.IsDir() {
					if isRoot && path != dir {
						return fs.SkipDir
					}
					return nil
				}
				if isLocaleFile(path, lang) && !seen[path] {
					seen[path] = true
					fileCh <- path
				}
				return nil
			}); err != nil {
				select {
				case walkErrCh <- fmt.Errorf("walk %s: %w", dir, err):
				default:
				}
			}
		}
	}()

	// Consumers: parse files concurrently
	const workers = 8
	var wg sync.WaitGroup
	for range workers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range fileCh {
				entries, err := parseFile(path, root)
				resultCh <- result{entries, err}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var (
		all  []Entry
		errs []error
	)
	for r := range resultCh {
		if r.err != nil {
			errs = append(errs, r.err)
		} else {
			all = append(all, r.entries...)
		}
	}

	// By the time resultCh is drained, the walk goroutine is guaranteed to have
	// finished (fileCh close happens after all WalkDir calls complete).
	var walkErr error
	select {
	case walkErr = <-walkErrCh:
	default:
	}

	if len(errs) > 0 {
		errs = append(errs, walkErr) // include walk error if any
		return all, errors.Join(errs...)
	}
	return all, walkErr
}

// candidateDirs returns possible locale directories for a language.
func candidateDirs(root, lang string) []string {
	dirs := []string{
		filepath.Join(root, "config", "locales", lang),
		filepath.Join(root, "config", "locales"),
	}
	var valid []string
	seen := map[string]bool{}
	for _, d := range dirs {
		if seen[d] {
			continue
		}
		seen[d] = true
		if _, err := os.Stat(d); err == nil {
			valid = append(valid, d)
		}
	}
	return valid
}

// isLocaleFile returns true if the file belongs to the given language.
func isLocaleFile(path, lang string) bool {
	if !strings.HasSuffix(path, ".yml") {
		return false
	}
	base := filepath.Base(path)
	return strings.HasPrefix(base, lang+".") ||
		strings.Contains(base, "."+lang+".")
}

// parseFile reads and flattens a single YAML locale file.
func parseFile(path, root string) ([]Entry, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", path, err)
	}

	var doc yaml.Node
	if err := yaml.Unmarshal(data, &doc); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}

	short, _ := filepath.Rel(filepath.Join(root, "config", "locales"), path)

	var entries []Entry
	if len(doc.Content) > 0 {
		flattenNode(doc.Content[0], "", path, short, &entries)
	}
	return entries, nil
}

// flattenNode recursively walks yaml.Node and appends scalar leaves to entries.
func flattenNode(node *yaml.Node, prefix, file, short string, entries *[]Entry) {
	switch node.Kind {
	case yaml.MappingNode:
		for i := 0; i+1 < len(node.Content); i += 2 {
			k := node.Content[i]
			v := node.Content[i+1]
			full := k.Value
			if prefix != "" {
				full = prefix + "." + k.Value
			}
			if v.Kind == yaml.ScalarNode {
				parts := strings.Split(full, ".")
				*entries = append(*entries, Entry{
					Key:       full,
					KeyName:   parts[len(parts)-1],
					Value:     v.Value,
					File:      file,
					ShortPath: short,
					Line:      k.Line,
				})
			} else {
				flattenNode(v, full, file, short, entries)
			}
		}
	case yaml.SequenceNode:
		// Emit the parent key so t('foo.items') / t('foo.features') calls that
		// iterate over the array aren't reported as missing by the auditor.
		if prefix != "" {
			parts := strings.Split(prefix, ".")
			*entries = append(*entries, Entry{
				Key:       prefix,
				KeyName:   parts[len(parts)-1],
				Value:     "[array]",
				File:      file,
				ShortPath: short,
				Line:      node.Line,
			})
		}
	case yaml.AliasNode:
		flattenNode(node.Alias, prefix, file, short, entries)
	}
}

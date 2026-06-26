package locale

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

// MergeCandidate is a duplicate group ready to be moved into the shared file.
type MergeCandidate struct {
	KeyName      string // leaf key name e.g. "error_rate_limit"
	Value        string // the canonical value to store
	SuggestedKey string // full dot-notation key in shared file
	Sources      []Entry
}

// BuildMergeCandidates converts KeyDupGroups with same values into candidates
// ready to write into the shared YAML.
func BuildMergeCandidates(dups []KeyDupGroup, lang string) []MergeCandidate {
	var out []MergeCandidate
	seen := make(map[string]bool)
	for _, g := range dups {
		if !g.SameValue {
			continue // skip groups with different wordings — needs human review
		}
		suggested := SuggestSharedKey(lang, g.KeyName)
		if seen[suggested] {
			fmt.Fprintf(os.Stderr, "locale-sync: SuggestedKey collision on %q (from %q) — skipping\n", suggested, g.KeyName)
			continue
		}
		seen[suggested] = true
		out = append(out, MergeCandidate{
			KeyName:      g.KeyName,
			Value:        g.Entries[0].Value,
			SuggestedKey: suggested,
			Sources:      g.Entries,
		})
	}
	return out
}

// UpsertShared merges candidates into the shared locale file.
func UpsertShared(root, lang string, candidates []MergeCandidate) (string, bool, error) {
	p := filepath.Join(root, "config", "locales", lang, fmt.Sprintf("shared.%s.yml", lang))
	return upsertYAMLFile(p, candidates)
}

// UpsertTopicFile merges candidates into {topic}.{lang}.yml.
// If the topic file does not yet exist it is created.
// Falls back to shared.{lang}.yml only when topic == "shared".
func UpsertTopicFile(root, lang, topic string, candidates []MergeCandidate) (string, bool, error) {
	p := filepath.Join(root, "config", "locales", lang, fmt.Sprintf("%s.%s.yml", topic, lang))
	return upsertYAMLFile(p, candidates)
}

// upsertYAMLFile is the shared implementation: read, merge, write.
func upsertYAMLFile(filePath string, candidates []MergeCandidate) (string, bool, error) {
	doc, err := readOrEmptyDoc(filePath)
	if err != nil {
		return filePath, false, err
	}

	changed := false
	var conflicts []string
	for _, c := range candidates {
		parts := strings.Split(c.SuggestedKey, ".")
		ok, err := setYAMLPath(doc, parts, c.Value)
		if err != nil {
			conflicts = append(conflicts, err.Error())
			continue
		}
		if ok {
			changed = true
		}
	}

	var conflictErr error
	if len(conflicts) > 0 {
		conflictErr = fmt.Errorf("YAML key conflicts in %s: %s", filePath, strings.Join(conflicts, "; "))
	}

	if !changed {
		return filePath, false, conflictErr
	}

	data, err := marshalDoc(doc)
	if err != nil {
		return filePath, false, err
	}

	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return filePath, false, err
	}
	if err := os.WriteFile(filePath, data, 0o644); err != nil {
		return filePath, false, err
	}
	return filePath, true, conflictErr
}

// WriteSkeleton generates a skeleton locale file for targetLang based on the
// entries from sourceLang. Values are replaced with placeholder.
// Skips files that already exist unless overwrite is true.
func WriteSkeleton(root, sourceLang, targetLang, placeholder string, entries []Entry, overwrite bool) ([]string, error) {
	// Group entries by the "topic" part of their source file.
	// e.g. "en/tools.en.yml" → topic "tools", output "es/tools.es.yml"
	byTopic := make(map[string][]Entry)
	for _, e := range entries {
		topic := fileTopicName(e.ShortPath, sourceLang)
		if topic == "" || topic == "shared" {
			continue
		}
		byTopic[topic] = append(byTopic[topic], e)
	}

	targetDir := filepath.Join(root, "config", "locales", targetLang)
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return nil, err
	}

	var written []string
	for topic, topicEntries := range byTopic {
		outPath := filepath.Join(targetDir, fmt.Sprintf("%s.%s.yml", topic, targetLang))
		if !overwrite {
			if _, err := os.Stat(outPath); err == nil {
				continue // already exists
			}
		}

		doc, err := buildSkeletonDoc(targetLang, topicEntries, placeholder)
		if err != nil {
			return written, err
		}
		data, err := marshalDoc(doc)
		if err != nil {
			return written, err
		}
		if err := os.WriteFile(outPath, data, 0o644); err != nil {
			return written, err
		}
		written = append(written, outPath)
	}
	return written, nil
}

// fileTopicName extracts the topic name from a locale file short path.
// "en/tools.en.yml" → "tools"
func fileTopicName(short, lang string) string {
	base := filepath.Base(short)
	base = strings.TrimSuffix(base, ".yml")
	base = strings.TrimSuffix(base, "."+lang)
	return base
}

// buildSkeletonDoc builds an ordered yaml.Node document for a skeleton file.
func buildSkeletonDoc(targetLang string, entries []Entry, placeholder string) (*yaml.Node, error) {
	doc := &yaml.Node{Kind: yaml.DocumentNode}
	root := &yaml.Node{Kind: yaml.MappingNode}
	doc.Content = []*yaml.Node{root}

	// Replace source lang prefix with target lang in each key path.
	for _, e := range entries {
		parts := strings.Split(e.Key, ".")
		if len(parts) < 2 {
			continue
		}
		newParts := append([]string{targetLang}, parts[1:]...)
		if _, err := setYAMLPath(doc, newParts, placeholder); err != nil {
			return nil, fmt.Errorf("build skeleton key %q: %w", e.Key, err)
		}
	}
	return doc, nil
}

// readOrEmptyDoc reads an existing YAML file into a doc node, or returns an
// empty mapping doc if the file doesn't exist yet.
func readOrEmptyDoc(path string) (*yaml.Node, error) {
	doc := &yaml.Node{Kind: yaml.DocumentNode}

	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		root := &yaml.Node{Kind: yaml.MappingNode}
		doc.Content = []*yaml.Node{root}
		return doc, nil
	}
	if err != nil {
		return nil, err
	}

	if err := yaml.Unmarshal(data, doc); err != nil {
		return nil, fmt.Errorf("parse %s: %w", path, err)
	}
	// Ensure doc has at least one mapping content node.
	if len(doc.Content) == 0 {
		doc.Content = []*yaml.Node{{Kind: yaml.MappingNode}}
	}
	return doc, nil
}

// setYAMLPath traverses/creates nested mapping nodes along path and sets the
// leaf to value. Returns (true, nil) if inserted, (false, nil) if already
// present, or (false, error) if a scalar/mapping conflict blocks the insert.
func setYAMLPath(doc *yaml.Node, path []string, value string) (bool, error) {
	return upsertYAMLPath(doc, path, value, nil)
}

// upsertYAMLPath is like setYAMLPath but calls shouldOverwrite(existingValue)
// at the leaf to decide whether to replace an already-present scalar.
// nil shouldOverwrite behaves identically to setYAMLPath (never overwrite).
func upsertYAMLPath(doc *yaml.Node, path []string, value string, shouldOverwrite func(string) bool) (bool, error) {
	if len(doc.Content) == 0 {
		doc.Content = []*yaml.Node{{Kind: yaml.MappingNode}}
	}
	return upsertMappingPath(doc.Content[0], path, value, shouldOverwrite)
}

func setMappingPath(node *yaml.Node, path []string, value string) (bool, error) {
	return upsertMappingPath(node, path, value, nil)
}

func upsertMappingPath(node *yaml.Node, path []string, value string, shouldOverwrite func(string) bool) (bool, error) {
	if node.Kind != yaml.MappingNode {
		return false, fmt.Errorf("cannot insert at %q: node is not a mapping (kind=%d)", strings.Join(path, "."), node.Kind)
	}
	key := path[0]

	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			if len(path) == 1 {
				existing := node.Content[i+1].Value
				if shouldOverwrite != nil && shouldOverwrite(existing) {
					node.Content[i+1].Value = value
					return true, nil
				}
				return false, nil
			}
			child := node.Content[i+1]
			if child.Kind != yaml.MappingNode {
				return false, fmt.Errorf("key conflict at %q: scalar exists where mapping needed", strings.Join(path, "."))
			}
			return upsertMappingPath(child, path[1:], value, shouldOverwrite)
		}
	}

	// Key not found — insert it.
	keyNode := &yaml.Node{Kind: yaml.ScalarNode, Value: key}
	if len(path) == 1 {
		valNode := &yaml.Node{Kind: yaml.ScalarNode, Value: value, Style: yaml.DoubleQuotedStyle}
		node.Content = append(node.Content, keyNode, valNode)
		return true, nil
	}
	child := &yaml.Node{Kind: yaml.MappingNode}
	node.Content = append(node.Content, keyNode, child)
	return upsertMappingPath(child, path[1:], value, shouldOverwrite)
}

// marshalDoc serialises a yaml.Node document to bytes with 2-space indentation.
func marshalDoc(doc *yaml.Node) ([]byte, error) {
	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(doc); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// WriteEntries inserts key→value pairs into filePath, skipping already-present keys.
// fullKeys maps full dot-notation paths (e.g. "es.tools.foo.bar") to their values.
// Returns (changed, error).
func WriteEntries(filePath string, fullKeys map[string]string) (bool, error) {
	return UpsertEntries(filePath, fullKeys, nil)
}

// UpsertEntries is like WriteEntries but calls shouldOverwrite(existingValue) before
// skipping an already-present key. When shouldOverwrite returns true the existing
// scalar is replaced. Pass nil to get the same behaviour as WriteEntries.
func UpsertEntries(filePath string, fullKeys map[string]string, shouldOverwrite func(existing string) bool) (bool, error) {
	doc, err := readOrEmptyDoc(filePath)
	if err != nil {
		return false, err
	}
	changed := false
	for k, v := range fullKeys {
		parts := strings.Split(k, ".")
		ok, err := upsertYAMLPath(doc, parts, v, shouldOverwrite)
		if err != nil {
			return false, err
		}
		if ok {
			changed = true
		}
	}
	if !changed {
		return false, nil
	}
	data, err := marshalDoc(doc)
	if err != nil {
		return false, err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), 0o755); err != nil {
		return false, err
	}
	return true, os.WriteFile(filePath, data, 0o644)
}

// MigrationPlan returns the t() key paths that callers should switch to for
// each merge candidate, grouped by source key full path.
func MigrationPlan(candidates []MergeCandidate) map[string]string {
	plan := make(map[string]string)
	for _, c := range candidates {
		// t() path: strip the leading lang prefix from suggestedKey.
		// e.g. "en.shared.errors.rate_limit" → "shared.errors.rate_limit"
		parts := strings.SplitN(c.SuggestedKey, ".", 2)
		tKey := parts[len(parts)-1]
		for _, src := range c.Sources {
			plan[src.Key] = tKey
		}
	}
	return plan
}

// DeleteKeys removes the given full dot-notation keys from filePath and writes
// the file back. Keys that do not exist in the file are silently skipped.
// Returns (changed, error): changed is true when at least one key was deleted.
func DeleteKeys(filePath string, keys []string) (bool, error) {
	doc, err := readOrEmptyDoc(filePath)
	if err != nil {
		return false, err
	}
	changed := false
	for _, key := range keys {
		parts := strings.Split(key, ".")
		if deleteYAMLPath(doc, parts) {
			changed = true
		}
	}
	if !changed {
		return false, nil
	}
	data, err := marshalDoc(doc)
	if err != nil {
		return false, err
	}
	return true, os.WriteFile(filePath, data, 0o644)
}

// deleteYAMLPath removes the node at the given path from the document.
// Returns true if a node was removed.
func deleteYAMLPath(doc *yaml.Node, path []string) bool {
	if len(doc.Content) == 0 || len(path) == 0 {
		return false
	}
	return deleteMappingPath(doc.Content[0], path)
}

func deleteMappingPath(node *yaml.Node, path []string) bool {
	if node.Kind != yaml.MappingNode {
		return false
	}
	key := path[0]
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			if len(path) == 1 {
				// Remove key+value pair.
				node.Content = append(node.Content[:i], node.Content[i+2:]...)
				return true
			}
			return deleteMappingPath(node.Content[i+1], path[1:])
		}
	}
	return false
}

// SortedKeys returns map keys in deterministic order.
func SortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

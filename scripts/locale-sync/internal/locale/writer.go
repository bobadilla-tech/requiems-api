package locale

import (
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
	for _, g := range dups {
		if !g.SameValue {
			continue // skip groups with different wordings — needs human review
		}
		out = append(out, MergeCandidate{
			KeyName:      g.KeyName,
			Value:        g.Entries[0].Value,
			SuggestedKey: SuggestSharedKey(lang, g.KeyName),
			Sources:      g.Entries,
		})
	}
	return out
}

// UpsertShared merges candidates into the shared locale file.
// Existing content is preserved; new keys are appended under their namespace.
// Returns the path written and a bool indicating whether any changes were made.
func UpsertShared(root, lang string, candidates []MergeCandidate) (string, bool, error) {
	sharedPath := filepath.Join(root, "config", "locales", lang,
		fmt.Sprintf("shared.%s.yml", lang))

	// Parse existing file (or start with empty doc node)
	doc, err := readOrEmptyDoc(sharedPath)
	if err != nil {
		return sharedPath, false, err
	}

	changed := false
	for _, c := range candidates {
		// SuggestedKey is e.g. "en.shared.errors.rate_limit"
		parts := strings.Split(c.SuggestedKey, ".")
		if setYAMLPath(doc, parts, c.Value) {
			changed = true
		}
	}

	if !changed {
		return sharedPath, false, nil
	}

	data, err := marshalDoc(doc)
	if err != nil {
		return sharedPath, false, err
	}

	if err := os.MkdirAll(filepath.Dir(sharedPath), 0o755); err != nil {
		return sharedPath, false, err
	}
	return sharedPath, true, os.WriteFile(sharedPath, data, 0o644)
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

		doc := buildSkeletonDoc(targetLang, topicEntries, placeholder)
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
func buildSkeletonDoc(targetLang string, entries []Entry, placeholder string) *yaml.Node {
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
		setYAMLPath(doc, newParts, placeholder)
	}
	return doc
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
// leaf to value. Returns true if a new key was inserted (i.e. a change was made).
func setYAMLPath(doc *yaml.Node, path []string, value string) bool {
	if len(doc.Content) == 0 {
		doc.Content = []*yaml.Node{{Kind: yaml.MappingNode}}
	}
	return setMappingPath(doc.Content[0], path, value)
}

func setMappingPath(node *yaml.Node, path []string, value string) bool {
	if node.Kind != yaml.MappingNode {
		return false
	}
	key := path[0]

	// Find existing key
	for i := 0; i+1 < len(node.Content); i += 2 {
		if node.Content[i].Value == key {
			if len(path) == 1 {
				// Key exists — don't overwrite existing values.
				return false
			}
			child := node.Content[i+1]
			if child.Kind != yaml.MappingNode {
				// Turn scalar into mapping
				child.Kind = yaml.MappingNode
				child.Value = ""
				child.Content = nil
			}
			return setMappingPath(child, path[1:], value)
		}
	}

	// Key not found — insert it.
	keyNode := &yaml.Node{Kind: yaml.ScalarNode, Value: key}
	if len(path) == 1 {
		valNode := &yaml.Node{Kind: yaml.ScalarNode, Value: value, Style: yaml.DoubleQuotedStyle}
		node.Content = append(node.Content, keyNode, valNode)
		return true
	}
	child := &yaml.Node{Kind: yaml.MappingNode}
	node.Content = append(node.Content, keyNode, child)
	return setMappingPath(child, path[1:], value)
}

// marshalDoc serialises a yaml.Node document to bytes with 2-space indentation.
func marshalDoc(doc *yaml.Node) ([]byte, error) {
	return yaml.Marshal(doc)
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

// SortedKeys returns map keys in deterministic order.
func SortedKeys[V any](m map[string]V) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

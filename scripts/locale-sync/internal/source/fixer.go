package source

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
)

// Fix describes a single replacement to apply to a source file.
type Fix struct {
	File       string
	Line       int
	Original   string // the raw line as found in the file
	Patched    string // the rewritten line
	YAMLKey    string // the t('key') that was used
	YAMLValue  string // the value added/found in YAML
	NewInYAML  bool   // true if key was not in YAML before
}

// FixPlan groups all fixes with the YAML additions they require.
type FixPlan struct {
	Fixes     []Fix
	YAMLAdds  map[string]string // key → value to add to YAML
}

// applyTemplate substitutes KEY into a per-category ERB/Ruby i18n snippet.
// We avoid fmt.Sprintf so that `%>` in ERB is never misread as a format verb.
func applyTemplate(category, key string) string {
	switch category {
	case "placeholder":
		return `placeholder="<%= t('` + key + `') %>"`
	case "attribute":
		return `"<%= t('` + key + `') %>"`
	case "tag_content":
		// Includes surrounding > < so we can do a straight >text< → replacement swap.
		return `><%= t('` + key + `') %><`
	case "erb_output":
		return `<%= t('` + key + `') %>`
	default:
		return `t('` + key + `')`
	}
}

// BuildFixPlan takes hardcoded strings and the current YAML key→value map,
// then produces a FixPlan describing what to change in source files and YAML.
//
// valueToKey is built from the current YAML entries so we can re-use an existing
// key if it already holds the same value.
func BuildFixPlan(
	hardcoded []HardcodedString,
	valueToKey map[string]string, // value → full YAML key (e.g. "en.shared.nav.search_placeholder")
	lang, root string,
) FixPlan {
	plan := FixPlan{
		YAMLAdds: make(map[string]string),
	}

	// Group by file so we can read each file once.
	byFile := make(map[string][]HardcodedString)
	for _, h := range hardcoded {
		byFile[h.File] = append(byFile[h.File], h)
	}

	for relFile, items := range byFile {
		absPath := filepath.Join(root, relFile)
		lines, err := readLines(absPath)
		if err != nil {
			continue
		}

		for _, h := range items {
			idx := h.Line - 1
			if idx < 0 || idx >= len(lines) {
				continue
			}
			line := lines[idx]

			// Determine t() key.
			var tKey string
			newInYAML := false

			if existing, ok := valueToKey[h.Text]; ok {
				// Re-use existing YAML key (strip leading lang. prefix for t() call).
				tKey = stripLangPrefixStr(existing, lang)
			} else {
				// Generate a new key and schedule a YAML addition.
				tKey = generateKey(lang, relFile, h)
				fullKey := lang + "." + tKey
				if _, already := plan.YAMLAdds[fullKey]; !already {
					plan.YAMLAdds[fullKey] = h.Text
					newInYAML = true
				}
			}

			// Build the patched line.
			replacement := applyTemplate(h.Category, tKey)
			patched := replaceInLine(line, h.Text, h.Category, replacement)
			if patched == line {
				continue // nothing changed — skip
			}

			plan.Fixes = append(plan.Fixes, Fix{
				File:      relFile,
				Line:      h.Line,
				Original:  line,
				Patched:   patched,
				YAMLKey:   tKey,
				YAMLValue: h.Text,
				NewInYAML: newInYAML,
			})
		}
	}

	// Sort fixes by file then line for deterministic output.
	sort.Slice(plan.Fixes, func(i, j int) bool {
		if plan.Fixes[i].File != plan.Fixes[j].File {
			return plan.Fixes[i].File < plan.Fixes[j].File
		}
		return plan.Fixes[i].Line < plan.Fixes[j].Line
	})

	return plan
}

// ApplyFixes rewrites source files according to the plan.
// Returns the list of files written.
func ApplyFixes(plan FixPlan, root string) ([]string, error) {
	// Group fixes by file.
	byFile := make(map[string][]Fix)
	for _, f := range plan.Fixes {
		byFile[f.File] = append(byFile[f.File], f)
	}

	var written []string
	for relFile, fixes := range byFile {
		absPath := filepath.Join(root, relFile)
		lines, err := readLines(absPath)
		if err != nil {
			return written, fmt.Errorf("read %s: %w", relFile, err)
		}

		// Apply fixes in reverse line order so line indices stay valid.
		sort.Slice(fixes, func(i, j int) bool {
			return fixes[i].Line > fixes[j].Line
		})
		for _, fix := range fixes {
			idx := fix.Line - 1
			if idx >= 0 && idx < len(lines) {
				lines[idx] = fix.Patched
			}
		}

		content := strings.Join(lines, "\n")
		if !strings.HasSuffix(content, "\n") {
			content += "\n"
		}
		if err := os.WriteFile(absPath, []byte(content), 0o644); err != nil {
			return written, fmt.Errorf("write %s: %w", relFile, err)
		}
		written = append(written, relFile)
	}
	return written, nil
}

// generateKey creates a dot-notation YAML key path from the file path and context.
// Example: "app/views/partials/tools/email_normalizer/_hero.html.erb" + placeholder
//
//	→ "tools.email_normalizer.demo.input_placeholder"
func generateKey(lang string, relFile string, h HardcodedString) string {
	// Strip app/views/, leading underscore on partials, extension.
	rel := strings.TrimPrefix(relFile, "app/views/")
	rel = strings.TrimPrefix(rel, "app/controllers/")
	rel = strings.TrimPrefix(rel, "app/models/")

	parts := strings.Split(rel, "/")
	var keyParts []string
	for _, p := range parts {
		p = strings.TrimPrefix(p, "_")
		p = strings.TrimSuffix(p, ".html.erb")
		p = strings.TrimSuffix(p, ".erb")
		p = strings.TrimSuffix(p, ".rb")
		p = strings.TrimSuffix(p, ".haml")
		p = strings.Replace(p, "-", "_", -1)
		if p != "" && p != "partials" && p != "views" {
			keyParts = append(keyParts, p)
		}
	}

	leaf := inferLeafKey(h)
	keyParts = append(keyParts, leaf)
	return strings.Join(keyParts, ".")
}

// inferLeafKey guesses a semantic key name from the hardcoded string's category and text.
func inferLeafKey(h HardcodedString) string {
	switch h.Category {
	case "placeholder":
		return "input_placeholder"
	case "tag_content":
		return slugify(h.Text)
	case "attribute":
		return slugify(h.Text)
	default:
		return slugify(h.Text)
	}
}

// slugify turns "Enter any email address" → "enter_any_email_address" (truncated).
var nonAlnum = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(s string) string {
	s = strings.ToLower(s)
	s = nonAlnum.ReplaceAllString(s, "_")
	s = strings.Trim(s, "_")
	// Limit to 40 chars.
	if len(s) > 40 {
		s = s[:40]
		if idx := strings.LastIndex(s, "_"); idx > 0 {
			s = s[:idx]
		}
	}
	return s
}

// replaceInLine finds and replaces the hardcoded text within line based on category.
func replaceInLine(line, text, category, replacement string) string {
	switch category {
	case "placeholder":
		// Replace placeholder="text" → placeholder="<%= t('key') %>"
		old := fmt.Sprintf(`placeholder="%s"`, text)
		return strings.Replace(line, old, replacement, 1)
	case "attribute":
		// Replace alt/title/aria-label="text" → attr="<%= t('key') %>"
		tKey := extractTKey(replacement)
		for _, attr := range []string{"alt", "title", "aria-label"} {
			old := fmt.Sprintf(`%s="%s"`, attr, text)
			if strings.Contains(line, old) {
				patched := attr + `="<%= t('` + tKey + `') %>"`
				return strings.Replace(line, old, patched, 1)
			}
		}
	case "tag_content":
		// Replace >text< with ><%= t('key') %><
		old := fmt.Sprintf(">%s<", text)
		return strings.Replace(line, old, replacement, 1)
	case "erb_output":
		// Replace <%= "text" %> with <%= t('key') %>
		for _, q := range []string{`"`, `'`} {
			old := fmt.Sprintf(`<%= %s%s%s %>`, q, text, q)
			if strings.Contains(line, old) {
				return strings.Replace(line, old, replacement, 1)
			}
		}
	case "ruby":
		// For Ruby patterns: attempt direct replacement of the quoted string.
		for _, q := range []string{`"`, `'`} {
			old := q + text + q
			if strings.Contains(line, old) {
				tKey := extractTKey(replacement)
				return strings.Replace(line, old, fmt.Sprintf("t('%s')", tKey), 1)
			}
		}
	}
	return line
}

// extractTKey pulls the key out of a replacement string like `placeholder="<%= t('key') %>"`.
var tKeyRe = regexp.MustCompile(`t\('([^']+)'\)`)

func extractTKey(s string) string {
	if m := tKeyRe.FindStringSubmatch(s); len(m) >= 2 {
		return m[1]
	}
	return s
}

func stripLangPrefixStr(key, lang string) string {
	prefix := lang + "."
	if strings.HasPrefix(key, prefix) {
		return key[len(prefix):]
	}
	return key
}

func readLines(path string) ([]string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	// Remove trailing empty element from final newline.
	if len(lines) > 0 && lines[len(lines)-1] == "" {
		lines = lines[:len(lines)-1]
	}
	return lines, nil
}

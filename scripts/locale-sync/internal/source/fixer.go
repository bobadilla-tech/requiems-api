package source

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Fix describes a single replacement to apply to a source file.
type Fix struct {
	File      string
	Line      int
	EndLine   int    // 0 for single-line; closing tag line for multiline_tag_content
	Original  string // the raw line as found in the file
	Patched   string // the rewritten line
	YAMLKey   string // the t('key') that was used
	YAMLValue string // the value added/found in YAML
	NewInYAML bool   // true if key was not in YAML before
}

// FixPlan groups all fixes with the YAML additions they require.
type FixPlan struct {
	Fixes    []Fix
	YAMLAdds map[string]string // key → value to add to YAML
}

// applyTemplate substitutes KEY into a per-category ERB/Ruby i18n snippet.
// We avoid fmt.Sprintf so that `%>` in ERB is never misread as a format verb.
func applyTemplate(category, key, relFile string) string {
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
	case "ruby":
		return translationCall(relFile, key)
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

	// fixIndex maps "relFile:line" → index in plan.Fixes for single-line fixes,
	// so multiple replacements on the same source line are composed into one Fix
	// rather than creating separate entries that overwrite each other.
	fixIndex := make(map[string]int)

	for relFile, items := range byFile {
		absPath := filepath.Join(root, relFile)
		lines, err := readLines(absPath)
		if err != nil {
			continue
		}

		for _, h := range items {
			// text_node: bare text mixed with ERB on one line — report-only, cannot auto-fix.
			if h.Category == "text_node" {
				continue
			}

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
				// If the generated key already maps to a different value (two distinct
				// strings produce the same slug), append a numeric suffix until free.
				baseTKey := generateKey(lang, relFile, h)
				tKey = baseTKey
				for i := 2; ; i++ {
					fullKey := lang + "." + tKey
					existing, already := plan.YAMLAdds[fullKey]
					if !already {
						plan.YAMLAdds[fullKey] = h.Text
						newInYAML = true
						break
					}
					if existing == h.Text {
						newInYAML = true // key IS being added to YAML in this plan
						break            // same value already queued — reuse the key
					}
					if i > 9 {
						break // give up after 9 suffixes (practically impossible)
					}
					tKey = baseTKey + "_" + strconv.Itoa(i)
				}
			}

			if tKey == "" {
				continue
			}

			// Build the patched line.
			var patched, replacement string
			if h.Category == "multiline_tag_content" && h.EndLine > h.Line {
				// Collapse the multi-line block into one line:
				// <p class="..."><%= t('key') %></p>
				tagName := extractTagName(line)
				leadWS := line[:len(line)-len(strings.TrimLeft(line, " \t"))]
				openTag := strings.TrimSpace(line)
				patched = leadWS + openTag + "<%= t('" + tKey + "') %></" + tagName + ">"
			} else {
				replacement = applyTemplate(h.Category, tKey, relFile)
				patched = replaceInLine(line, h.Text, h.Category, replacement)
				if patched == line {
					continue
				}
			}

			// Compose multiple single-line replacements into one Fix rather than
			// appending a new Fix that would overwrite the previous one in ApplyFixes.
			if h.Category != "multiline_tag_content" {
				lineKey := relFile + ":" + strconv.Itoa(h.Line)
				if existingIdx, ok := fixIndex[lineKey]; ok {
					composed := replaceInLine(plan.Fixes[existingIdx].Patched, h.Text, h.Category, replacement)
					if composed != plan.Fixes[existingIdx].Patched {
						plan.Fixes[existingIdx].Patched = composed
						if newInYAML {
							plan.Fixes[existingIdx].NewInYAML = true
						}
					}
					continue
				}
				fixIndex[lineKey] = len(plan.Fixes)
			}

			plan.Fixes = append(plan.Fixes, Fix{
				File:      relFile,
				Line:      h.Line,
				EndLine:   h.EndLine,
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
			// Multi-line: blank text lines + closing tag (they're folded into patched).
			if fix.EndLine > fix.Line {
				for k := fix.Line; k < fix.EndLine && k < len(lines); k++ {
					lines[k] = ""
				}
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
// Examples:
//   - "app/views/partials/tools/email_normalizer/_hero.html.erb" + placeholder
//     → "tools.email_normalizer.hero.input_placeholder"
//   - "app/views/devise/mailer/reset_password_instructions.html.erb"
//     → "devise.mailer.reset_password_instructions.{leaf}"
func generateKey(lang string, relFile string, h HardcodedString) string {
	rel := strings.TrimPrefix(relFile, "app/views/")
	rel = strings.TrimPrefix(rel, "app/controllers/")
	rel = strings.TrimPrefix(rel, "app/models/")

	// Devise views: keep "devise" as the first segment so keys route to devise.{lang}.yml.
	isDevise := strings.HasPrefix(rel, "devise/")

	parts := strings.Split(rel, "/")
	var keyParts []string
	for _, p := range parts {
		p = strings.TrimPrefix(p, "_")
		p = strings.TrimSuffix(p, ".html.erb")
		p = strings.TrimSuffix(p, ".erb")
		p = strings.TrimSuffix(p, ".rb")
		p = strings.TrimSuffix(p, ".haml")
		p = strings.ReplaceAll(p, "-", "_")
		skip := p == "" || p == "partials" || p == "views"
		// Keep "devise" segment; skip "shared" sub-dir inside devise (use parent path).
		if !isDevise {
			skip = skip || p == "shared"
		}
		if !skip {
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

// translationCall chooses the Rails helper call form that is valid for the source file.
func translationCall(relFile, key string) string {
	rel := filepath.ToSlash(relFile)
	if strings.Contains(rel, "/views/") ||
		strings.Contains(rel, "/controllers/") ||
		strings.Contains(rel, "/helpers/") ||
		strings.Contains(rel, "/mailers/") {
		return `t('` + key + `')`
	}
	return `I18n.t('` + key + `')`
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
		// Replace >text< with ><%= t('key') %>< — tolerate whitespace around text.
		re := regexp.MustCompile(`>\s*` + regexp.QuoteMeta(text) + `\s*<`)
		if loc := re.FindStringIndex(line); loc != nil {
			return line[:loc[0]] + replacement + line[loc[1]:]
		}
		return line
	case "erb_output":
		// Replace <%= "text" %> with <%= t('key') %>
		// Use string concat to avoid %> being misread as a format verb.
		for _, q := range []string{`"`, `'`} {
			old := "<%= " + q + text + q + " %>"
			if strings.Contains(line, old) {
				return strings.Replace(line, old, replacement, 1)
			}
		}
	case "ruby":
		// For Ruby patterns: attempt direct replacement of the quoted string.
		for _, q := range []string{`"`, `'`} {
			old := q + text + q
			if strings.Contains(line, old) {
				return strings.Replace(line, old, replacement, 1)
			}
		}
	}
	return line
}

// extractTagName pulls the tag name from an opening HTML tag line.
// "  <p class=\"...\">" → "p"
var tagNameRe = regexp.MustCompile(`<(\w+)`)

func extractTagName(line string) string {
	if m := tagNameRe.FindStringSubmatch(strings.TrimSpace(line)); len(m) >= 2 {
		return m[1]
	}
	return "div"
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

package source

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"
	"unicode"
)

// HardcodedString is a string literal found in Ruby/ERB source that is not
// wrapped in a Rails i18n call.
type HardcodedString struct {
	File     string
	Line     int
	EndLine  int    // 0 for single-line; last line of block for multiline_tag_content
	Text     string
	Context  string // surrounding code for display
	Category string // "placeholder", "tag_content", "multiline_tag_content", "erb_output", "ruby", "attribute"
}

// Extensions scanned for hardcoded strings.
var sourceExts = map[string]bool{
	".rb":   true,
	".erb":  true,
	".haml": true,
	".slim": true,
}

// pathSkipPrefixes — directories that are intentionally not scanned.
// Devise views are now included — the tool generates keys under devise.{controller}.{action}.
var pathSkipPrefixes = []string{}

// brandNames — proper nouns / brand names used standalone that should never be
// extracted as translation keys.
var brandNames = map[string]bool{
	"Requiems API": true,
	"ChatGPT":      true,
	"Claude":       true,
	"GitHub":       true,
	"Stripe":       true,
	"LemonSqueezy": true,
	"JSON":         true,
	"API":          true,
}

// technicalPhrases — strings that look like UI text but are technical labels
// specific to this codebase that should stay in English.
var technicalPhraseSkips = []string{
	// HTTP protocol terms
	"Base URL",
	"HTTP Status",
	"Status Code",
	"Rate Limit",
	"Bearer ",
	"API Key",
	"API key",
	"Enter JSON",
	"JSON object",
	"Send Request",
	"Waiting for response",
	// Code/playground UI — defined elsewhere or intentionally English
	"Copy code",
	"Copy URL",
	"Copy response",
	"Copy to clipboard",
	"Copy as Markdown",
	"Copy page link",
	"Open in ChatGPT",
	"Open in Claude",
	"View on GitHub",
	"View Demo",
	// Admin-specific monitoring labels (internal tooling, English-only)
	"Excellent uptime",
	"Good uptime",
	"Low uptime",
	"Fast response times",
	"Slow response times",
	"Acceptable performance",
	"Normal security activity",
	"High failed auth",
	// Admin panel section headers that mirror Rails admin convention
	"Admin Panel",
	"Action Required",
	"Mounted Services",
	"Your Endpoint",
	"Deployment Spec",
	"Mark as Deployed",
	"Customer Notes",
	"Granted by",
	"AI Tooling",
}

// Patterns that already use i18n helpers — skip lines containing these.
var i18nCallRe = regexp.MustCompile(`\bI18n\.t\b|\bt\s*[("']|\btranslate\s*[("']`)

// erbPattern matches a complete ERB expression tag.
var erbTagRe = regexp.MustCompile(`<%=.*?%>`)

// --- ERB-specific patterns ---

// HTML attributes with user-facing text that should be i18n'd.
var attrPatterns = []struct {
	re       *regexp.Regexp
	category string
}{
	// placeholder="text" not preceded by <%=
	{regexp.MustCompile(`\bplaceholder="([^"<]{4,})"`), "placeholder"},
	// alt/title/aria-label — allow lowercase start, looksUserFacing does real filtering.
	{regexp.MustCompile(`\balt="([^"<]{4,})"`), "attribute"},
	{regexp.MustCompile(`\btitle="([^"<]{4,})"`), "attribute"},
	{regexp.MustCompile(`\baria-label="([^"<]{4,})"`), "attribute"},
}

// HTML tag content: text between a closing > and the next open <.
// Broad match — looksUserFacing + looksLikeCode do the real filtering.
var tagContentRe = regexp.MustCompile(`>([A-Za-z\d][^<\n]{3,})<`)

// Ruby display patterns.
var rubyPatterns = []struct {
	re       *regexp.Regexp
	category string
}{
	{regexp.MustCompile(`<%=\s*["']([^"'%]{4,})["']\s*%>`), "erb_output"},
	{regexp.MustCompile(`flash(?:\.now)?\[:\w+\]\s*=\s*["']([^"']{4,})["']`), "ruby"},
	{regexp.MustCompile(`errors\.add\(:\w+,\s*["']([^"']{4,})["']\)`), "ruby"},
	{regexp.MustCompile(`render\s+plain:\s*["']([^"']{4,})["']`), "ruby"},
	{regexp.MustCompile(`raise(?:\s+\w+,)?\s*["']([^"']{6,})["']`), "ruby"},
	{regexp.MustCompile(`(?:notice|alert):\s*["']([^"']{4,})["']`), "ruby"},
	// content_for :title only — :description holds long SEO copy.
	// Bug fix: closing quote must be OUTSIDE the capture group.
	{regexp.MustCompile(`content_for\s+:title,\s*["']([^"']{4,})["']`), "ruby"},
	// ERB link/button text: link_to "text", …  |  button_to "text", …  |  f.submit "text"
	{regexp.MustCompile(`\blink_to\s+["']([^"']{4,})["']\s*,`), "ruby"},
	{regexp.MustCompile(`\bbutton_to\s+["']([^"']{4,})["']\s*,`), "ruby"},
	{regexp.MustCompile(`\bf\.submit\s+["']([^"']{4,})["']`), "ruby"},
}

// Extract concurrently scans all .rb/.erb/.haml/.slim files under root/app/
// and returns hardcoded string literals that look user-facing.
func Extract(root string) ([]HardcodedString, error) {
	appDir := filepath.Join(root, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		appDir = root
	}

	fileCh := make(chan string, 128)
	resultCh := make(chan []HardcodedString, 128)

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
				results := scanFile(path, root)
				if len(results) > 0 {
					resultCh <- results
				}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var all []HardcodedString
	for batch := range resultCh {
		all = append(all, batch...)
	}
	return all, nil
}

func scanFile(path, root string) []HardcodedString {
	short, _ := filepath.Rel(root, path)

	// Skip paths with their own i18n convention.
	for _, prefix := range pathSkipPrefixes {
		if strings.HasPrefix(short, prefix) {
			return nil
		}
	}

	lines, err := readLines(path)
	if err != nil {
		return nil
	}

	isERB := strings.HasSuffix(path, ".erb") || strings.HasSuffix(path, ".haml") || strings.HasSuffix(path, ".slim")

	var results []HardcodedString

	for i, line := range lines {
		lineNum := i + 1
		trimmed := strings.TrimSpace(line)

		if isComment(trimmed) {
			continue
		}

		if isERB {
			results = append(results, scanAttrPatterns(line, trimmed, short, lineNum)...)
			results = append(results, scanTagContent(line, trimmed, short, lineNum)...)
		}

		// Skip lines already calling t() / I18n.t — applies to both Ruby and ERB.
		if i18nCallRe.MatchString(line) {
			continue
		}

		for _, p := range rubyPatterns {
			matches := p.re.FindStringSubmatch(line)
			if len(matches) < 2 {
				continue
			}
			text := strings.TrimSpace(matches[1])
			if !looksUserFacing(text) {
				continue
			}
			results = append(results, HardcodedString{
				File:     short,
				Line:     lineNum,
				Text:     text,
				Context:  trimmed,
				Category: p.category,
			})
			break
		}
	}

	// Second pass: multi-line block element content (<p>, <li>, etc. spanning lines).
	if isERB {
		results = append(results, scanMultiLineTagContent(lines, short)...)
	}

	return results
}

func scanAttrPatterns(line, trimmed, short string, lineNum int) []HardcodedString {
	// Strip ERB tags before checking attribute values — if an attribute is set
	// via <%= t(...) %>, the raw attribute text will not contain user strings.
	stripped := erbTagRe.ReplaceAllString(line, "")
	if i18nCallRe.MatchString(stripped) {
		return nil
	}

	var out []HardcodedString
	for _, p := range attrPatterns {
		for _, m := range p.re.FindAllStringSubmatch(stripped, -1) {
			text := strings.TrimSpace(m[1])
			if !looksUserFacing(text) {
				continue
			}
			if isDecorativeAlt(text) {
				continue
			}
			out = append(out, HardcodedString{
				File:     short,
				Line:     lineNum,
				Text:     text,
				Context:  trimmed,
				Category: p.category,
			})
		}
	}
	return out
}

// erbOutputRe matches any ERB output tag <%= ... %>
var erbOutputRe = regexp.MustCompile(`<%=`)

func scanTagContent(line, trimmed, short string, lineNum int) []HardcodedString {
	// Skip lines that already use t() in any form.
	if i18nCallRe.MatchString(line) {
		return nil
	}
	// Skip lines with ERB output expressions — stripping <%= var %> leaves
	// orphaned text fragments (e.g. "Hello !" from "Hello <%= @email %>!").
	if erbOutputRe.MatchString(line) {
		return nil
	}
	// Strip ERB logic tags (non-output: <% ... %>) before scanning tag content.
	stripped := erbTagRe.ReplaceAllString(line, "")
	// Skip SVG elements, code blocks, and pure ERB logic lines.
	for _, skipTag := range []string{"<svg", "<path", "<polygon", "<pre", "<code", "<script", "<style", "<%"} {
		if strings.HasPrefix(trimmed, skipTag) {
			return nil
		}
	}

	var out []HardcodedString
	for _, m := range tagContentRe.FindAllStringSubmatch(stripped, -1) {
		text := strings.TrimSpace(m[1])
		if !looksUserFacing(text) || looksLikeCode(text) {
			continue
		}
		out = append(out, HardcodedString{
			File:     short,
			Line:     lineNum,
			Text:     text,
			Context:  trimmed,
			Category: "tag_content",
		})
	}
	return out
}

func isComment(s string) bool {
	return strings.HasPrefix(s, "#") ||
		strings.HasPrefix(s, "//") ||
		strings.HasPrefix(s, "<!--") ||
		strings.HasPrefix(s, "<%#")
}

// isDecorativeAlt skips alts that describe icons/logos rather than content.
func isDecorativeAlt(s string) bool {
	lower := strings.ToLower(s)
	for _, skip := range []string{"logo", "icon", "svg", "flag", "badge", "avatar", "image"} {
		if strings.Contains(lower, skip) {
			return true
		}
	}
	return false
}

// looksUserFacing filters out strings that are clearly not display text.
func looksUserFacing(s string) bool {
	// Too short or too long (very long = SEO meta description, not a UI label).
	if len(s) < 4 || len(s) > 120 {
		return false
	}
	// Must contain at least one space — single-word technical terms stay in English.
	if !strings.Contains(s, " ") {
		return false
	}
	// Skip strings with Ruby interpolation — need %{var} conversion, not a direct swap.
	if strings.Contains(s, "#{") {
		return false
	}
	// Skip brand names used standalone.
	if brandNames[s] {
		return false
	}
	// Skip codebase-specific technical phrases.
	for _, phrase := range technicalPhraseSkips {
		if strings.EqualFold(s, phrase) || strings.HasPrefix(s, phrase) {
			return false
		}
	}
	lower := strings.ToLower(s)
	for _, skip := range []string{
		// SQL
		"select ", "insert ", "update ", "delete ", "from ",
		// Asset refs
		".css", ".js", ".rb", "http://", "https://",
		// ERB fragments
		"<%", "%>",
		// HTML attribute fragments
		"class=", "data-", "style=",
		// HTTP / MIME / API technical content
		"application/json", "application/xml", "content-type",
		"authorization:", "bearer ", "x-api-key", "x-backend-secret",
		// Ruby class names and code patterns
		"def ", "end\n", "rescue ", "@media ",
	} {
		if strings.Contains(lower, skip) {
			return false
		}
	}
	r := []rune(s)
	return unicode.IsLetter(r[0]) || unicode.IsDigit(r[0])
}

// looksLikeCode catches things like CSS class lists and camelCase identifiers.
func looksLikeCode(s string) bool {
	// High density of hyphens (CSS classes like "flex items-center gap-3")
	hyphens := strings.Count(s, "-")
	words := len(strings.Fields(s))
	if words > 1 && hyphens >= words-1 {
		return true
	}
	// Looks like a CSS class string
	lower := strings.ToLower(s)
	cssKeywords := []string{"px-", "py-", "mt-", "mb-", "mr-", "ml-", "text-", "bg-", "flex", "grid", "rounded"}
	matches := 0
	for _, kw := range cssKeywords {
		if strings.Contains(lower, kw) {
			matches++
		}
	}
	return matches >= 2
}

// blockTagOpenRe matches a block-level opening tag that has NO inline content
// (the > is at the end of the trimmed line, so text lives on the next lines).
var blockTagOpenRe = regexp.MustCompile(`^<(p|h[1-6]|li|dt|dd|td|th|label|option)(\s[^>]*)?>$`)

// blockTagCloseRe matches the corresponding closing tag on its own trimmed line.
var blockTagCloseRe = regexp.MustCompile(`^</(p|h[1-6]|li|dt|dd|td|th|label|option)>`)

// scanMultiLineTagContent detects text content that spans multiple lines inside
// a single block element, e.g.:
//
//	<p class="...">
//	  Auto-generate quote cards for Twitter/X, Instagram, or LinkedIn without
//	  a content team.
//	</p>
func scanMultiLineTagContent(lines []string, short string) []HardcodedString {
	var out []HardcodedString
	for i := 0; i < len(lines); i++ {
		trimmed := strings.TrimSpace(lines[i])

		if isComment(trimmed) || i18nCallRe.MatchString(lines[i]) {
			continue
		}

		openMatch := blockTagOpenRe.FindStringSubmatch(trimmed)
		if openMatch == nil {
			continue
		}
		tagName := openMatch[1]

		// Collect text-only lines until we hit any tag (max 8-line lookahead).
		var textParts []string
		j := i + 1
		for j < len(lines) && j-i <= 8 {
			next := strings.TrimSpace(lines[j])
			if strings.Contains(next, "<") || strings.Contains(next, ">") {
				break
			}
			if next != "" {
				textParts = append(textParts, next)
			}
			j++
		}

		if len(textParts) == 0 || j >= len(lines) {
			continue
		}

		// The line at j must be the closing tag.
		if !strings.HasPrefix(strings.TrimSpace(lines[j]), "</"+tagName+">") {
			continue
		}

		text := strings.Join(strings.Fields(strings.Join(textParts, " ")), " ")

		if !looksUserFacing(text) || looksLikeCode(text) {
			continue
		}

		out = append(out, HardcodedString{
			File:     short,
			Line:     i + 1, // open tag (1-indexed)
			EndLine:  j + 1, // closing tag (1-indexed)
			Text:     text,
			Context:  trimmed,
			Category: "multiline_tag_content",
		})
		i = j // skip past the closing tag
	}
	return out
}

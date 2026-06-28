package source

import (
	"fmt"
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
	// Note: "Copy URL" is NOT listed here — it is user-facing and should be localized.
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
	// HTTP header / request field names used in code examples
	"Content-Type",
	"requiems-api-key",
	"YOUR_API_KEY",
	// HTTP methods in code examples
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
	// JSON/code literal values that appear in API response examples
	"true",
	"false",
	"null",
	// CSS values — appear when ERB expressions are embedded in style attributes
	// e.g. mail_to(..., style: "color: #6b7280;")
	"linear infinite",
	"color:",
	"background-color:",
	"background:",
	"font-size:",
	"font-weight:",
	"font-family:",
	"border-radius:",
	"border:",
	"padding:",
	"margin:",
	"text-transform:",
	"letter-spacing:",
	"line-height:",
	"display:",
	"width:",
	"height:",
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
// allowLabel=true: use looksUserFacingLabel (no space required) — for UI label
// contexts like link text, button text, form labels where single words are valid.
var rubyPatterns = []struct {
	re         *regexp.Regexp
	category   string
	allowLabel bool
}{
	{regexp.MustCompile(`<%=\s*["']([^"'%]{4,})["']\s*%>`), "erb_output", false},
	{regexp.MustCompile(`flash(?:\.now)?\[:\w+\]\s*=\s*["']([^"']{4,})["']`), "ruby", false},
	{regexp.MustCompile(`errors\.add\(:\w+,\s*["']([^"']{4,})["']\)`), "ruby", false},
	{regexp.MustCompile(`render\s+plain:\s*["']([^"']{4,})["']`), "ruby", false},
	{regexp.MustCompile(`raise(?:\s+\w+,)?\s*["']([^"']{6,})["']`), "ruby", false},
	{regexp.MustCompile(`(?:notice|alert):\s*["']([^"']{4,})["']`), "ruby", false},
	// validates :field, presence: { message: "..." } and length: { too_short: "..." } etc.
	{regexp.MustCompile(`\b(?:message|too_short|too_long|wrong_length|not_a_number|not_an_integer|greater_than|less_than|other_than|odd|even):\s*["']([^"']{4,})["']`), "ruby", false},
	// redirect_to ..., alert: / notice: — already covered by notice|alert above, but
	// content_for :title only — :description holds long SEO copy.
	// Bug fix: closing quote must be OUTSIDE the capture group.
	{regexp.MustCompile(`content_for\s+:title,\s*["']([^"']{4,})["']`), "ruby", false},
	// ERB link/button text: link_to "text", … | button_to "text", … | f.submit "text"
	// Single words like "Cancel", "Back", "Clear" are valid user-facing labels.
	{regexp.MustCompile(`\blink_to\s+["']([^"']{2,})["']\s*,`), "ruby", true},
	{regexp.MustCompile(`\bbutton_to\s+["']([^"']{2,})["']\s*,`), "ruby", true},
	{regexp.MustCompile(`\bf\.submit\s+["']([^"']{2,})["']`), "ruby", true},
	// Ruby keyword arg placeholder: (inside form helpers — different from HTML placeholder=)
	{regexp.MustCompile(`\bplaceholder:\s*["']([^"']{4,})["']`), "ruby", false},
	// Turbo / UJS confirm dialogs: data: { turbo_confirm: "..." } or confirm: "..."
	{regexp.MustCompile(`\b(?:turbo_confirm|confirm):\s*["']([^"']{8,})["']`), "ruby", false},
}

type scanResult struct {
	found []HardcodedString
	warn  error // non-nil when the file could not be read
}

// Extract concurrently scans all .rb/.erb/.haml/.slim files under root/app/
// and returns hardcoded string literals that look user-facing.
// Files that cannot be read are skipped with a warning printed to stderr.
// Only files in changedFiles are scanned when the set is non-nil (CI diff mode).
func Extract(root string, changedFiles map[string]bool) ([]HardcodedString, error) {
	appDir := filepath.Join(root, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		appDir = root
	}

	fileCh := make(chan string, 128)
	resultCh := make(chan scanResult, 128)

	go func() {
		defer close(fileCh)
		_ = filepath.WalkDir(appDir, func(path string, d fs.DirEntry, err error) error {
			if err != nil || d.IsDir() {
				return nil
			}
			if !sourceExts[filepath.Ext(path)] {
				return nil
			}
			if changedFiles != nil {
				rel, _ := filepath.Rel(root, path)
				if !changedFiles[rel] && !changedFiles[filepath.ToSlash(rel)] {
					return nil
				}
			}
			fileCh <- path
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
				found, warn := scanFile(path, root)
				resultCh <- scanResult{found: found, warn: warn}
			}
		}()
	}
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	var all []HardcodedString
	for r := range resultCh {
		if r.warn != nil {
			fmt.Fprintf(os.Stderr, "warning: skipping file: %v\n", r.warn)
		}
		all = append(all, r.found...)
	}
	return all, nil
}

func scanFile(path, root string) ([]HardcodedString, error) {
	short, _ := filepath.Rel(root, path)

	// Skip paths with their own i18n convention.
	for _, prefix := range pathSkipPrefixes {
		if strings.HasPrefix(short, prefix) {
			return nil, nil
		}
	}

	lines, err := readLines(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", short, err)
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
			results = append(results, scanTableHeaders(line, trimmed, short, lineNum)...)
			results = append(results, scanTagContent(line, trimmed, short, lineNum)...)
			results = append(results, scanBareTextNodes(line, trimmed, short, lineNum)...)
			results = append(results, scanBareTextSuffixes(line, trimmed, short, lineNum)...)
			results = append(results, scanErbTernary(line, trimmed, short, lineNum)...)
			results = append(results, scanLabelTags(line, trimmed, short, lineNum)...)
			results = append(results, scanHelperKeywordArgs(line, trimmed, short, lineNum)...)
			results = append(results, scanOptionsArray(line, trimmed, short, lineNum)...)
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
			if p.allowLabel {
				if !looksUserFacingLabel(text) {
					continue
				}
			} else {
				if !looksUserFacing(text) {
					continue
				}
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

	return results, nil
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
		if looksLikeCode(text) {
			continue
		}
		// looksUserFacing requires a space. Fall back to looksUserFacingLabel for
		// single-word capitalized status labels (e.g. "Pending", "Active", "Cancelled").
		if !looksUserFacing(text) && !looksUserFacingLabel(text) {
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

// tableHeaderRe matches text inside <th> cells including those with attributes.
// Unlike tagContentRe, it allows single-word labels because column headers
// like "Name", "Type", "Field" are inherently user-facing even without spaces.
var tableHeaderRe = regexp.MustCompile(`<th[^>]*>\s*([A-Za-z][A-Za-z0-9 ]{1,50}?)\s*<`)

func scanTableHeaders(line, trimmed, short string, lineNum int) []HardcodedString {
	if i18nCallRe.MatchString(line) || erbOutputRe.MatchString(line) {
		return nil
	}
	stripped := erbTagRe.ReplaceAllString(line, "")
	var out []HardcodedString
	for _, m := range tableHeaderRe.FindAllStringSubmatch(stripped, -1) {
		text := strings.TrimSpace(m[1])
		if text == "" || looksLikeCode(text) || brandNames[text] {
			continue
		}
		// Require at least 2 chars and a letter start.
		r := []rune(text)
		if len(r) < 2 || !unicode.IsLetter(r[0]) {
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
		// CSS property values (e.g. "color: #6b7280;" passed as style: arg)
		": #", "px;", "em;", "rem;", "%;",
	} {
		if strings.Contains(lower, skip) {
			return false
		}
	}
	r := []rune(s)
	return unicode.IsLetter(r[0]) || unicode.IsDigit(r[0])
}

// looksUserFacingLabel is like looksUserFacing but permits single-word strings.
// Use for button/link text, form labels, and status labels where one word is valid UI copy.
func looksUserFacingLabel(s string) bool {
	if len(s) < 2 || len(s) > 60 {
		return false
	}
	if strings.Contains(s, "#{") {
		return false
	}
	if brandNames[s] {
		return false
	}
	for _, phrase := range technicalPhraseSkips {
		if strings.EqualFold(s, phrase) || strings.HasPrefix(s, phrase) {
			return false
		}
	}
	lower := strings.ToLower(s)
	for _, skip := range []string{
		".css", ".js", ".rb", "http://", "https://",
		"<%", "%>", "class=", "data-", "style=",
		"application/json", "bearer ", "x-api-key",
		"def ", "rescue ",
		": #", "px;", "em;", "rem;", "%;",
	} {
		if strings.Contains(lower, skip) {
			return false
		}
	}
	r := []rune(s)
	return unicode.IsLetter(r[0])
}

// stimulusActionRe matches Stimulus controller#action descriptors like "modal#close"
// or "click->modal#close", which are technical routing strings, not user-facing text.
var stimulusActionRe = regexp.MustCompile(`\b[a-z][a-z0-9_-]*#[a-z_][a-z0-9_]*\b`)

// looksLikeCode catches things like CSS class lists and camelCase identifiers.
func looksLikeCode(s string) bool {
	// Stimulus controller#action descriptor (e.g. "modal#close", "click->modal#close")
	if stimulusActionRe.MatchString(s) {
		return true
	}
	// High density of hyphens (CSS classes like "flex items-center gap-3")
	hyphens := strings.Count(s, "-")
	words := len(strings.Fields(s))
	if words > 1 && hyphens >= words-1 {
		return true
	}
	// Single-word strings with 2+ hyphens are technical tokens — HTTP headers
	// (X-Backend-Secret), Stimulus descriptors, kebab-case identifiers — never UI copy.
	if !strings.Contains(s, " ") && hyphens >= 2 {
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

// scanBareTextNodes detects user-facing text mixed inline with ERB output tags on
// the same line — e.g. "Have questions? <%= link_to 'FAQ', ... %> or ...".
// scanTagContent explicitly skips lines containing <%= to avoid orphaned fragments
// like "Hello !"; this pass fills that gap by examining only the text PREFIX before
// the first ERB tag, which is the cleanest extractable fragment.
//
// Findings are report-only (category "text_node") — auto-fixing split-sentence lines
// requires human judgement and is not safe to automate.
func scanBareTextNodes(line, trimmed, short string, lineNum int) []HardcodedString {
	if !erbOutputRe.MatchString(line) {
		return nil
	}
	// HTML tag lines are handled by scanTagContent / scanMultiLineTagContent.
	if strings.HasPrefix(trimmed, "<") {
		return nil
	}
	// Extract the text before the first ERB tag.
	firstERB := strings.Index(line, "<%")
	if firstERB <= 0 {
		return nil
	}
	prefix := strings.TrimSpace(line[:firstERB])
	if prefix == "" {
		return nil
	}
	// Reject attribute fragments: end with ="  ='  =  (  "  '  {
	// These indicate the ERB tag is inside an HTML attribute value or a JS/Ruby string.
	for _, bad := range []string{`="`, `='`, `=`, `(`, `"`, `'`, `{`} {
		if strings.HasSuffix(prefix, bad) {
			return nil
		}
	}
	// Reject if prefix itself contains HTML tags (text is not pure text node).
	if strings.Contains(prefix, "<") {
		return nil
	}
	if !looksUserFacing(prefix) || looksLikeCode(prefix) {
		return nil
	}
	return []HardcodedString{{
		File:     short,
		Line:     lineNum,
		Text:     prefix,
		Context:  trimmed,
		Category: "text_node",
	}}
}

// lastErbOrTagCloseRe matches ERB closers (%>) and HTML closing tags (</tag>).
var lastErbOrTagCloseRe = regexp.MustCompile(`(?:%>|</\w+>)`)

// scanBareTextSuffixes detects user-facing text that appears AFTER the last
// ERB output closer or HTML closing tag on a line.
// Mirrors scanBareTextNodes which handles text before the first ERB tag.
// Example: `<strong><%= t('...') %></strong> Focus on customer retention`
func scanBareTextSuffixes(line, trimmed, short string, lineNum int) []HardcodedString {
	if !erbOutputRe.MatchString(line) {
		return nil
	}
	locs := lastErbOrTagCloseRe.FindAllStringIndex(line, -1)
	if len(locs) == 0 {
		return nil
	}
	lastEnd := locs[len(locs)-1][1]
	suffix := strings.TrimSpace(line[lastEnd:])
	// Strip trailing punctuation before checking.
	cleaned := strings.TrimRight(suffix, ".,;:!?")
	if cleaned == "" || strings.Contains(cleaned, "<") {
		return nil
	}
	if !looksUserFacing(cleaned) || looksLikeCode(cleaned) {
		return nil
	}
	return []HardcodedString{{
		File:     short,
		Line:     lineNum,
		Text:     suffix, // preserve trailing punctuation; cleaned was only for validation
		Context:  trimmed,
		Category: "text_node",
	}}
}

// ternaryBranchRe captures string literals in ternary branches (? 'x' : 'y').
// Requires spaces around ? to avoid matching predicate methods (suspended?) and
// hash key-value pairs (turbo_confirm: "...").
var ternaryBranchRe = regexp.MustCompile(`[?:]\s*["']([A-Za-z][^"']{2,})["']`)

// scanErbTernary detects hardcoded string literals in Ruby ternary expressions
// inside ERB output tags: `<%= cond ? 'Yes text' : 'No text' %>`.
func scanErbTernary(line, trimmed, short string, lineNum int) []HardcodedString {
	// Require a spaced ternary operator " ? " to distinguish from predicate methods
	// (suspended?) and hash key-value pairs (turbo_confirm: "...").
	if !erbOutputRe.MatchString(line) || !strings.Contains(line, " ? ") {
		return nil
	}
	var out []HardcodedString
	for _, m := range ternaryBranchRe.FindAllStringSubmatch(line, -1) {
		text := strings.TrimSpace(m[1])
		if !looksUserFacing(text) || looksLikeCode(text) {
			continue
		}
		out = append(out, HardcodedString{
			File:     short,
			Line:     lineNum,
			Text:     text,
			Context:  trimmed,
			Category: "erb_output",
		})
	}
	return out
}

// labelTagRe captures the label text argument in label_tag and f.label calls:
//   label_tag "field", "Label text"
//   f.label :field, "Label text"
var labelTagRe = regexp.MustCompile(`(?:\blabel_tag\s+["'][^"']*["'],\s*["']([A-Za-z][^"']{1,})["']|\bf\.label\s+:\w+,\s*["']([A-Za-z][^"']{1,})["'])`)

// scanLabelTags detects hardcoded label text in label_tag and f.label calls.
// Uses looksUserFacingLabel (no space required) since form labels are often single words.
func scanLabelTags(line, trimmed, short string, lineNum int) []HardcodedString {
	if i18nCallRe.MatchString(line) {
		return nil
	}
	m := labelTagRe.FindStringSubmatch(line)
	if len(m) < 2 {
		return nil
	}
	// m[1] = label_tag match, m[2] = f.label match
	text := strings.TrimSpace(m[1])
	if text == "" {
		text = strings.TrimSpace(m[2])
	}
	if !looksUserFacingLabel(text) || looksLikeCode(text) {
		return nil
	}
	return []HardcodedString{{
		File:     short,
		Line:     lineNum,
		Text:     text,
		Context:  trimmed,
		Category: "ruby",
	}}
}

// helperKeywordArgRe captures string values for text:, submit_tag, and disable_with: calls.
// Catches: text: "Convert"  |  submit_tag "Get Started"  |  disable_with: "Sending..."
var helperKeywordArgRe = regexp.MustCompile(`(?:\btext:\s*["']([A-Za-z][^"']{1,})["']|\bsubmit_tag\s+["']([^"']{4,})["']|\bdisable_with:\s*["']([A-Za-z][^"']{1,})["'])`)

// scanHelperKeywordArgs detects hardcoded text in helper keyword args.
// These bypass looksUserFacing's space check since button labels can be one word.
func scanHelperKeywordArgs(line, trimmed, short string, lineNum int) []HardcodedString {
	if i18nCallRe.MatchString(line) {
		return nil
	}
	m := helperKeywordArgRe.FindStringSubmatch(line)
	if len(m) < 2 {
		return nil
	}
	// m[1] = text: value, m[2] = submit_tag value, m[3] = disable_with: value
	text := strings.TrimSpace(m[1])
	if text == "" {
		text = strings.TrimSpace(m[2])
	}
	if text == "" {
		text = strings.TrimSpace(m[3])
	}
	if !looksUserFacingLabel(text) || looksLikeCode(text) {
		return nil
	}
	return []HardcodedString{{
		File:     short,
		Line:     lineNum,
		Text:     text,
		Context:  trimmed,
		Category: "ruby",
	}}
}

// optionsArrayRe matches string labels in 2-element option arrays:
//   ["Last 7 Days", "7"]   ["All", nil]   ["Active", "true"]
// Used in options_for_select([...]) and ERB .each loops over option arrays.
var optionsArrayRe = regexp.MustCompile(`\["([A-Za-z][^"]{1,60}?)",\s*(?:"[^"]*"|nil|\d+|false|true)\]`)

// scanOptionsArray detects hardcoded string labels in options arrays.
// Each ["Label", value] pair produces one finding for the label string.
func scanOptionsArray(line, trimmed, short string, lineNum int) []HardcodedString {
	if i18nCallRe.MatchString(line) {
		return nil
	}
	var out []HardcodedString
	for _, m := range optionsArrayRe.FindAllStringSubmatch(line, -1) {
		text := strings.TrimSpace(m[1])
		r := []rune(text)
		if !looksUserFacingLabel(text) || looksLikeCode(text) || !unicode.IsUpper(r[0]) {
			continue
		}
		out = append(out, HardcodedString{
			File:     short,
			Line:     lineNum,
			Text:     text,
			Context:  trimmed,
			Category: "ruby",
		})
	}
	return out
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

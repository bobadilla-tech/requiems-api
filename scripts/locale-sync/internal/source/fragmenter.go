package source

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"unicode"
)

// FragmentLine is a source line where visible text and <%= ERB %> expressions
// are interleaved — a pattern that prevents proper i18n because the full sentence
// is split. The merger proposes a single key with %{var} placeholders.
type FragmentLine struct {
	File          string // relative path from Rails root
	AbsFile       string
	LineNum       int
	Original      string   // raw source line
	MergedValue   string   // proposed key value e.g. "Hello %{name}!"
	KeySuggestion string   // proposed dot-notation key
	PatchedLine   string   // rewritten line using t() call (empty when complex)
	Bindings      []string // "varname: ruby_expr" pairs for the t() call
	Complex       bool     // true when HTML tags or multi-arg ERB prevent auto-patch
	ComplexReason string   // human-readable explanation of why auto-patch is skipped
}

// erbOutputExprRe matches a single <%= ... %> output tag.
var erbOutputExprRe = regexp.MustCompile(`<%=\s*([\s\S]*?)\s*%>`)

// htmlTagInTextRe matches complete HTML tags inside text nodes.
var htmlTagInTextRe = regexp.MustCompile(`</?[a-zA-Z][^>]*>`)

// fullHTMLTagLineRe matches a line whose non-ERB content is entirely one HTML
// open/close tag (i.e., ERB is inside an HTML attribute, not in text content).
var fullHTMLTagLineRe = regexp.MustCompile(`^</?[a-zA-Z][^>]*>$`)

// htmlAttrLineRe matches lines that are HTML attributes in a multi-line tag:
//   data-api-name="<%= ... %>"
//   class="<%= ... %>"
// The non-ERB content would look like:   word-or-word="X" or just word="X">
var htmlAttrLineRe = regexp.MustCompile(`^[\w-]+=`)

// DetectFragments walks app/ and returns lines where visible text and ERB
// expressions are mixed without a t() wrapper.
func DetectFragments(root string) ([]FragmentLine, error) {
	appDir := filepath.Join(root, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		appDir = root
	}

	var fragments []FragmentLine
	err := filepath.WalkDir(appDir, func(path string, d os.DirEntry, walkErr error) error {
		if walkErr != nil || d.IsDir() {
			return nil
		}
		ext := filepath.Ext(path)
		if ext != ".erb" && ext != ".haml" && ext != ".slim" {
			return nil
		}
		frags, scanErr := scanFileFragments(path, root)
		if scanErr == nil {
			fragments = append(fragments, frags...)
		}
		return nil
	})
	return fragments, err
}

func scanFileFragments(path, root string) ([]FragmentLine, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rel, _ := filepath.Rel(root, path)
	relSlash := filepath.ToSlash(rel)

	var fragments []FragmentLine
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 1<<20), 1<<20)
	lineNum := 0

	for scanner.Scan() {
		lineNum++
		raw := scanner.Text()
		trimmed := strings.TrimSpace(raw)

		if trimmed == "" || isComment(trimmed) {
			continue
		}
		// Must have at least one ERB output tag
		if !strings.Contains(trimmed, "<%=") {
			continue
		}
		// Skip lines where the ERB is inside an HTML attribute, not text content.
		// Replace ERB tags with a placeholder and check if the result is either:
		//   (a) a complete HTML tag <div class="X">    — fullHTMLTagLineRe
		//   (b) a standalone HTML attribute  data-x="X" — htmlAttrLineRe
		withoutERB := erbOutputExprRe.ReplaceAllString(trimmed, "X")
		withoutERBTrimmed := strings.TrimSpace(withoutERB)
		if fullHTMLTagLineRe.MatchString(withoutERBTrimmed) || htmlAttrLineRe.MatchString(withoutERBTrimmed) {
			continue
		}

		frag := parseFragmentLine(raw, trimmed, relSlash, path, lineNum)
		if frag != nil {
			fragments = append(fragments, *frag)
		}
	}
	return fragments, scanner.Err()
}

type linePart struct {
	text  string
	isERB bool
}

func parseFragmentLine(raw, trimmed, relSlash, absFile string, lineNum int) *FragmentLine {
	// Split line into alternating text / ERB parts.
	idxs := erbOutputExprRe.FindAllStringSubmatchIndex(trimmed, -1)
	if len(idxs) == 0 {
		return nil
	}

	var parts []linePart
	prev := 0
	for _, m := range idxs {
		if m[0] > prev {
			parts = append(parts, linePart{text: trimmed[prev:m[0]]})
		}
		inner := strings.TrimSpace(trimmed[m[2]:m[3]])
		parts = append(parts, linePart{text: inner, isERB: true})
		prev = m[1]
	}
	if prev < len(trimmed) {
		parts = append(parts, linePart{text: trimmed[prev:]})
	}

	// Classify text parts: strip HTML tags, check if user-facing content remains.
	var userFacingText []string
	var erbExprs []string
	hasHTMLTags := false

	for _, p := range parts {
		if p.isERB {
			erbExprs = append(erbExprs, p.text)
			continue
		}
		withTags := p.text
		stripped := strings.TrimSpace(htmlTagInTextRe.ReplaceAllString(withTags, ""))
		if htmlTagInTextRe.MatchString(withTags) {
			hasHTMLTags = true
		}
		if stripped == "" {
			continue
		}
		// Skip if the text node itself already contains a t() call
		// (means the text is part of an already-translated expression).
		if i18nCallRe.MatchString(stripped) {
			continue
		}
		runes := []rune(stripped)
		if len(runes) >= 2 && (unicode.IsLetter(runes[0]) || unicode.IsDigit(runes[0])) {
			// Cheap user-facing check: starts with a letter/digit, not purely technical
			isTech := false
			for _, phrase := range technicalPhraseSkips {
				if strings.EqualFold(stripped, phrase) || strings.HasPrefix(strings.ToLower(stripped), strings.ToLower(phrase)) {
					isTech = true
					break
				}
			}
			if !isTech {
				userFacingText = append(userFacingText, stripped)
			}
		}
	}

	// Need at least one user-facing text part and at least one ERB part.
	if len(userFacingText) == 0 || len(erbExprs) == 0 {
		return nil
	}

	// Determine complexity.
	complexReason := ""
	if hasHTMLTags {
		complexReason = "HTML tags in text — merge into _html key manually"
	} else {
		for _, expr := range erbExprs {
			// Multi-arg ERB (e.g. link_to, mail_to, method calls with keyword args)
			// are unsafe to auto-inline as t() arguments.
			if strings.Count(expr, ",") > 0 {
				complexReason = "multi-argument ERB expression — extract to a local variable first"
				break
			}
			// Boolean operators (|| / &&) mean the expression has conditional logic.
			// Auto-inlining as t(..., var: expr) produces invalid Ruby when the variable
			// name is derived from an operator token (e.g. "||" → ||: which is a syntax error).
			// The caller should assign the expression to a local variable first.
			if strings.Contains(expr, " || ") || strings.Contains(expr, " && ") {
				complexReason = "ERB expression contains boolean operator — assign to a local variable first"
				break
			}
		}
	}

	// Build merged value and variable bindings.
	var mergedParts []string
	var bindings []string
	varCount := map[string]int{}

	for _, p := range parts {
		if p.isERB {
			name := fragmentVarName(p.text, varCount)
			varCount[name]++
			mergedParts = append(mergedParts, "%{"+name+"}")
			bindings = append(bindings, name+": "+p.text)
		} else {
			// Strip HTML tags but preserve internal whitespace so "Hi %{name}!" keeps
			// the space between text and placeholder rather than collapsing to "Hi%{name}!".
			cleaned := htmlTagInTextRe.ReplaceAllString(p.text, "")
			if strings.TrimSpace(cleaned) != "" {
				mergedParts = append(mergedParts, cleaned)
			}
		}
	}

	mergedValue := strings.TrimSpace(strings.Join(mergedParts, ""))
	if len(mergedValue) < 4 {
		return nil
	}
	// Reject merged values that are HTML attributes, CSS, or technical content.
	if strings.Contains(mergedValue, `="`) || strings.HasPrefix(mergedValue, "<") {
		return nil
	}
	lowerMerged := strings.ToLower(mergedValue)
	for _, skip := range []string{"x-api-key", "x-backend-secret", "data-", "class=", "style="} {
		if strings.HasPrefix(lowerMerged, skip) {
			return nil
		}
	}
	// Reject CSS / JS blocks: remove %{var} placeholders and check for { } braces.
	withoutPlaceholders := regexp.MustCompile(`%\{[^}]+\}`).ReplaceAllString(mergedValue, "")
	if strings.Contains(withoutPlaceholders, "{") || strings.Contains(withoutPlaceholders, "}") {
		return nil
	}

	keyPath := fragmentKeyPath(relSlash, mergedValue)

	// Build patched line only when auto-patch is safe.
	patchedLine := ""
	if complexReason == "" {
		indent := raw[:len(raw)-len(strings.TrimLeft(raw, " \t"))]
		var tCall string
		if len(bindings) > 0 {
			tCall = "t('" + keyPath + "', " + strings.Join(bindings, ", ") + ")"
		} else {
			tCall = "t('" + keyPath + "')"
		}
		patchedLine = indent + "<%= " + tCall + " %>"
	}

	return &FragmentLine{
		File:          relSlash,
		AbsFile:       absFile,
		LineNum:       lineNum,
		Original:      raw,
		MergedValue:   mergedValue,
		KeySuggestion: keyPath,
		PatchedLine:   patchedLine,
		Bindings:      bindings,
		Complex:       complexReason != "",
		ComplexReason: complexReason,
	}
}

// fragmentVarName derives a short Ruby-safe variable name from an ERB expression.
func fragmentVarName(expr string, seen map[string]int) string {
	expr = strings.TrimSpace(expr)
	var name string
	switch {
	case strings.HasPrefix(expr, "link_to"):
		name = "link"
	case strings.HasPrefix(expr, "mail_to"):
		name = "contact"
	case strings.HasPrefix(expr, "image_tag"):
		name = "image"
	case strings.HasPrefix(expr, "content_tag"):
		name = "tag"
	default:
		// Take the last dot-separated or @-prefixed identifier.
		parts := strings.FieldsFunc(expr, func(r rune) bool {
			return r == '.' || r == '@' || r == ':' || r == '(' || r == ' '
		})
		for i := len(parts) - 1; i >= 0; i-- {
			p := strings.TrimRight(parts[i], "(),;[]")
			if p != "" && !strings.ContainsAny(p, "\"'") {
				name = p
				break
			}
		}
	}
	if name == "" {
		name = "value"
	}
	// Sanitize to a valid Ruby identifier: word characters only, no leading digit.
	var clean strings.Builder
	for _, r := range name {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' {
			clean.WriteRune(r)
		}
	}
	name = strings.TrimLeft(clean.String(), "0123456789")
	if name == "" {
		name = "value"
	}
	// Deduplicate: first occurrence keeps name, subsequent get name2, name3, …
	if seen[name] > 0 {
		return fmt.Sprintf("%s%d", name, seen[name]+1)
	}
	return name
}

// fragmentKeyPath builds a dot-notation key from the file path + first words of text.
func fragmentKeyPath(relSlash, text string) string {
	p := strings.TrimPrefix(relSlash, "app/views/")
	for _, ext := range []string{".html.erb", ".html.haml", ".text.erb", ".erb", ".haml"} {
		if strings.HasSuffix(p, ext) {
			p = p[:len(p)-len(ext)]
			break
		}
	}
	segments := strings.Split(p, "/")
	var keyParts []string
	for _, seg := range segments {
		seg = strings.TrimPrefix(seg, "_")
		seg = strings.ReplaceAll(seg, "-", "_")
		if seg == "" || seg == "partials" || seg == "views" || seg == "shared" {
			continue
		}
		keyParts = append(keyParts, seg)
	}

	// Leaf: first 4 meaningful words of the merged text, joined with underscores.
	words := strings.Fields(text)
	var leafWords []string
	for _, w := range words {
		if len(leafWords) >= 4 {
			break
		}
		w = strings.ToLower(w)
		var sb strings.Builder
		for _, r := range w {
			if unicode.IsLetter(r) || unicode.IsDigit(r) {
				sb.WriteRune(r)
			} else {
				sb.WriteRune('_')
			}
		}
		clean := strings.Trim(sb.String(), "_")
		if clean != "" {
			leafWords = append(leafWords, clean)
		}
	}
	leaf := strings.Join(leafWords, "_")
	if leaf == "" {
		leaf = "text"
	}
	keyParts = append(keyParts, leaf)
	return strings.Join(keyParts, ".")
}

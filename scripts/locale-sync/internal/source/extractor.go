package source

import (
	"bufio"
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
	File    string
	Line    int
	Text    string
	Context string // surrounding code for display
}

// Extensions scanned for hardcoded strings.
var sourceExts = map[string]bool{
	".rb":   true,
	".erb":  true,
	".haml": true,
	".slim": true,
}

// Patterns that already wrap strings in i18n — skip these lines.
var i18nCallRe = regexp.MustCompile(`\bI18n\.t\b|\bt\s*[("']|\btranslate\s*[("']`)

// Patterns for likely user-facing string literals in Ruby/ERB.
// We match: double-quoted strings after common display/error patterns.
var displayPatterns = []*regexp.Regexp{
	// ERB output: <%= "text" %> or <%= 'text' %>
	regexp.MustCompile(`<%=\s*["']([^"'%]{4,})["']\s*%>`),
	// flash[:x] = "text" or flash.now[:x] = "text"
	regexp.MustCompile(`flash(?:\.now)?\[:\w+\]\s*=\s*["']([^"']{4,})["']`),
	// errors.add(:base, "message")
	regexp.MustCompile(`errors\.add\(:\w+,\s*["']([^"']{4,})["']\)`),
	// render plain: "text"
	regexp.MustCompile(`render\s+plain:\s*["']([^"']{4,})["']`),
	// raise "error message" or raise SomeError, "message"
	regexp.MustCompile(`raise(?:\s+\w+,)?\s*["']([^"']{6,})["']`),
	// notice: "text" or alert: "text"
	regexp.MustCompile(`(?:notice|alert):\s*["']([^"']{4,})["']`),
	// logger.error "message" / logger.warn / logger.info — skip (not user-facing)
}

// Extract concurrently scans all .rb/.erb/.haml/.slim files under root/app/
// and returns hardcoded string literals that look user-facing.
func Extract(root string) ([]HardcodedString, error) {
	appDir := filepath.Join(root, "app")
	if _, err := os.Stat(appDir); os.IsNotExist(err) {
		appDir = root // fallback: scan from root
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
	f, err := os.Open(path)
	if err != nil {
		return nil
	}
	defer f.Close()

	short, _ := filepath.Rel(root, path)

	var results []HardcodedString
	scanner := bufio.NewScanner(f)
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := scanner.Text()
		trimmed := strings.TrimSpace(line)

		// Skip comments
		if strings.HasPrefix(trimmed, "#") || strings.HasPrefix(trimmed, "//") {
			continue
		}
		// Skip lines already using i18n
		if i18nCallRe.MatchString(line) {
			continue
		}

		for _, re := range displayPatterns {
			matches := re.FindStringSubmatch(line)
			if len(matches) < 2 {
				continue
			}
			text := strings.TrimSpace(matches[1])
			if !looksUserFacing(text) {
				continue
			}
			results = append(results, HardcodedString{
				File:    short,
				Line:    lineNum,
				Text:    text,
				Context: trimmed,
			})
			break // one finding per line is enough
		}
	}
	return results
}

// looksUserFacing filters out strings that are likely not display text:
// SQL fragments, class names, URL paths, pure symbols, etc.
func looksUserFacing(s string) bool {
	if len(s) < 4 || len(s) > 200 {
		return false
	}
	// Must contain at least one space (real sentences have spaces)
	if !strings.Contains(s, " ") {
		return false
	}
	// Skip things that look like SQL, CSS selectors, or code
	lower := strings.ToLower(s)
	for _, skip := range []string{"select ", "insert ", "update ", "delete ", "from ", ".css", ".js", "http://", "https://", "<%", "%>"} {
		if strings.Contains(lower, skip) {
			return false
		}
	}
	// Must start with a letter or digit (not a symbol/operator)
	r := []rune(s)
	return unicode.IsLetter(r[0]) || unicode.IsDigit(r[0])
}

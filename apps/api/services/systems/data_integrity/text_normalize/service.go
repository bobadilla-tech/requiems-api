package textnormalize

import (
	"html"
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/unicode/norm"
)

var (
	reWhitespace  = regexp.MustCompile(`\s+`)
	reHTMLTags    = regexp.MustCompile(`<[^>]+>`)
	rePunctuation = regexp.MustCompile(`[^\p{L}\p{N}\s]`)
)

// Operation names accepted in the operations array.
const (
	OpTrim               = "trim"
	OpLowercase          = "lowercase"
	OpUppercase          = "uppercase"
	OpNormalizeUnicode   = "normalize_unicode"
	OpCollapseWhitespace = "collapse_whitespace"
	OpStripHTML          = "strip_html"
	OpRemovePunctuation  = "remove_punctuation"
)

// Request is the input for the text normalize endpoint.
type Request struct {
	Text       string   `json:"text" validate:"required"`
	Operations []string `json:"operations" validate:"required,min=1"`
}

// Result is the response from the text normalize endpoint.
type Result struct {
	Original   string   `json:"original"`
	Normalized string   `json:"normalized"`
	Operations []string `json:"operations"`
}

// Service applies text normalization operations.
type Service struct{}

// NewService returns a new Service.
func NewService() *Service { return &Service{} }

// Normalize applies the requested operations in order and returns the result.
func (s *Service) Normalize(text string, operations []string) Result {
	out := text
	for _, op := range operations {
		out = apply(out, op)
	}
	return Result{
		Original:   text,
		Normalized: out,
		Operations: operations,
	}
}

func apply(text, op string) string {
	switch op {
	case OpTrim:
		return strings.TrimSpace(text)
	case OpLowercase:
		return strings.ToLower(text)
	case OpUppercase:
		return strings.ToUpper(text)
	case OpNormalizeUnicode:
		return norm.NFC.String(text)
	case OpCollapseWhitespace:
		return reWhitespace.ReplaceAllString(text, " ")
	case OpStripHTML:
		stripped := reHTMLTags.ReplaceAllString(text, " ")
		return html.UnescapeString(stripped)
	case OpRemovePunctuation:
		return strings.Map(func(r rune) rune {
			if unicode.IsPunct(r) || unicode.IsSymbol(r) {
				return -1
			}
			return r
		}, text)
	default:
		return text
	}
}

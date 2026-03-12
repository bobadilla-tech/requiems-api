package validate

import (
	"context"
	"net"
	"net/mail"
	"slices"
	"strings"

	normalizer "github.com/bobadilla-tech/go-email-normalizer"
	disposable "github.com/bobadilla-tech/is-email-disposable"
)

// commonDomains is the curated list of well-known email providers used for
// typo suggestions. Keep entries lowercase.
var commonDomains = []string{
	"gmail.com", "googlemail.com",
	"yahoo.com", "yahoo.co.uk", "yahoo.fr", "yahoo.es", "yahoo.de",
	"outlook.com", "hotmail.com", "hotmail.co.uk", "hotmail.fr",
	"icloud.com", "me.com", "mac.com",
	"aol.com",
	"protonmail.com", "proton.me",
	"live.com", "msn.com",
	"yandex.com", "yandex.ru",
	"mail.com", "zoho.com",
}

// Service validates email addresses.
type Service struct {
	n *normalizer.Normalizer
}

// NewService creates a Service initialized with a default email normalizer.
func NewService() *Service {
	return &Service{
		n: normalizer.NewNormalizer(),
	}
}

// ValidateEmail performs full validation: syntax check, MX record lookup,
// disposable-domain check, normalization, and typo suggestion.
func (s *Service) ValidateEmail(ctx context.Context, email string) EmailValidation {
	if !isValidSyntax(email) {
		return EmailValidation{
			Email:       email,
			Valid:       false,
			SyntaxValid: false,
		}
	}

	// Normalize; fall back to the original address on error.
	normalized := email
	if res, err := s.n.Normalize2(email); err == nil {
		normalized = res.Normalized
	}

	// Derive domain from the normalized address so alias-domain resolution
	// (e.g. googlemail.com → gmail.com) is reflected consistently.
	// isValidSyntax guarantees "@" is present in email, so the fallback is
	// only a safety net for unexpected normalizer output.
	_, domain, _ := strings.Cut(strings.ToLower(normalized), "@")
	if domain == "" {
		_, domain, _ = strings.Cut(strings.ToLower(email), "@")
	}

	mxValid := checkMX(ctx, domain)
	isDisposable := disposable.IsDisposable(normalized)
	suggestion := suggestDomain(domain)

	return EmailValidation{
		Email:       email,
		Valid:       mxValid,
		SyntaxValid: true,
		MxValid:     mxValid,
		Disposable:  isDisposable,
		Normalized:  normalized,
		Domain:      domain,
		Suggestion:  suggestion,
	}
}

// isValidSyntax reports whether email is a syntactically valid RFC 5322
// isValidSyntax reports whether the given email address conforms to RFC 5322 address syntax.
// It returns true when the address parses successfully and false otherwise.
func isValidSyntax(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// checkMX reports whether the given domain has at least one MX DNS record.
// It performs an MX lookup using the provided context and returns true if at least one record is found.
func checkMX(ctx context.Context, domain string) bool {
	records, err := net.DefaultResolver.LookupMX(ctx, domain)
	return err == nil && len(records) > 0
}

// suggestDomain returns a pointer to the closest well-known domain name when
// suggestDomain suggests a correction for common email provider domain typos.
// It returns a pointer to the closest domain from commonDomains when the edit distance is 2 or less, and returns nil for exact matches or when no close domain is found.
func suggestDomain(domain string) *string {
	const threshold = 2

	// Exact match — no suggestion needed.
	if slices.Contains(commonDomains, domain) {
		return nil
	}

	best := ""
	bestDist := threshold + 1

	for _, d := range commonDomains {
		if dist := levenshtein(domain, d); dist < bestDist {
			bestDist = dist
			best = d
		}
	}

	if bestDist <= threshold {
		return &best
	}
	return nil
}

// levenshtein computes the Levenshtein edit distance between two strings.
// It compares Unicode code points (runes) and returns the minimum number of single-character edits
// (insertions, deletions, or substitutions) required to change `a` into `b`.
func levenshtein(a, b string) int {
	ra, rb := []rune(a), []rune(b)
	la, lb := len(ra), len(rb)

	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}

	prev := make([]int, lb+1)
	for j := range prev {
		prev[j] = j
	}

	for i := 1; i <= la; i++ {
		curr := make([]int, lb+1)
		curr[0] = i
		for j := 1; j <= lb; j++ {
			if ra[i-1] == rb[j-1] {
				curr[j] = prev[j-1]
			} else {
				curr[j] = 1 + min(prev[j], curr[j-1], prev[j-1])
			}
		}
		prev = curr
	}

	return prev[lb]
}

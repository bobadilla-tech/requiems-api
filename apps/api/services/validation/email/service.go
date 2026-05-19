package email

import (
	"context"
	"net"
	"regexp"
	"strings"
	"sync"
	"time"

	"github.com/agnivade/levenshtein"
	normalizer "github.com/bobadilla-tech/go-email-normalizer"
	disposable "github.com/bobadilla-tech/is-email-disposable"
)

// commonDomains is the curated set of well-known email providers used for
// typo suggestions. Keep entries lowercase.
var commonDomains = map[string]struct{}{
	"gmail.com": {}, "googlemail.com": {},
	"yahoo.com": {}, "yahoo.co.uk": {}, "yahoo.fr": {}, "yahoo.es": {}, "yahoo.de": {},
	"outlook.com": {}, "hotmail.com": {}, "hotmail.co.uk": {}, "hotmail.fr": {},
	"icloud.com": {}, "me.com": {}, "mac.com": {},
	"aol.com":        {},
	"protonmail.com": {}, "proton.me": {},
	"live.com": {}, "msn.com": {},
	"yandex.com": {}, "yandex.ru": {},
	"mail.com": {}, "zoho.com": {},
}

// addrSpecRe matches a plain RFC 5322 addr-spec (local-part "@" domain).
// Display-name formats such as "Name <user@example.com>" are rejected.
//
// Grammar (no obs- forms, no CFWS):
//
//	addr-spec      = local-part "@" domain
//	local-part     = dot-atom / quoted-string
//	dot-atom       = atext+ ("." atext+)*
//	atext          = ALPHA / DIGIT / !#$%&'*+/=?^_`{|}~-
//	quoted-string  = DQUOTE (qtext / quoted-pair)* DQUOTE
//	qtext          = %x21 / %x23-5B / %x5D-7E  (printable excl. DQUOTE, \)
//	quoted-pair    = "\" (VCHAR / WSP)
//	domain         = dot-atom / domain-literal
//	domain label   = [a-zA-Z0-9] ([a-zA-Z0-9-]{0,61} [a-zA-Z0-9])?
//	domain-literal = "[" dtext* "]"
//	dtext          = %x21-5A / %x5E-7E  (printable excl. [, \, ])
var addrSpecRe = regexp.MustCompile(
	`^(?:` +
		// quoted-string local part
		// qtext = %x09 / %x20-21 / %x23-5B / %x5D-7E  (tab, space, printable excl. DQUOTE and \)
		// quoted-pair = "\" (VCHAR / WSP)
		`"(?:[\x09\x20-\x21\x23-\x5b\x5d-\x7e]|\\[\x09\x20-\x7e])*"` +
		`|` +
		// dot-atom local part: atext characters, dots only between groups
		`[a-zA-Z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\.[a-zA-Z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*` +
		`)@(?:` +
		// domain literal
		`\[[\x21-\x5a\x5e-\x7e]*\]` +
		`|` +
		// dot-atom domain: labels separated by dots; each label cannot start or end with a hyphen
		`[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*` +
		`)$`,
)

// Validation is the full validation result for an email address.
type Validation struct {
	Email       *string `json:"email"`
	Valid       bool    `json:"valid"`
	SyntaxValid bool    `json:"syntax_valid"`
	MxValid     bool    `json:"mx_valid"`
	Disposable  bool    `json:"disposable"`
	Normalized  *string `json:"normalized"`
	Domain      *string `json:"domain"`
	Suggestion  *string `json:"suggestion"`
}

// BatchItem holds the result for a single email in a batch request.
type BatchItem struct {
	Email       *string `json:"email"`
	Valid       bool    `json:"valid"`
	SyntaxValid bool    `json:"syntax_valid"`
	MXValid     bool    `json:"mx_valid"`
	Disposable  bool    `json:"disposable"`
	Normalized  *string `json:"normalized,omitempty"`
	Domain      *string `json:"domain,omitempty"`
	Suggestion  *string `json:"suggestion,omitempty"`
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
func (s *Service) ValidateEmail(ctx context.Context, email string) Validation {
	if !isValidSyntax(email) {
		return Validation{
			Email:       nil,
			Valid:       false,
			SyntaxValid: false,
		}
	}

	// Normalize; fall back to the original address on error.
	normalized := email
	if res, err := s.n.Normalize2(email); err == nil {
		normalized = res.Normalized
	}

	// syntax is valid, so @ is guaranteed to be present
	_, domain, _ := strings.Cut(strings.ToLower(normalized), "@")

	mxValid := checkMX(ctx, domain)
	isDisposable := disposable.IsDisposable(normalized)
	suggestion := suggestDomain(domain)

	return Validation{
		Email:       &email,
		Valid:       mxValid,
		SyntaxValid: true,
		MxValid:     mxValid,
		Disposable:  isDisposable,
		Normalized:  &normalized,
		Domain:      &domain,
		Suggestion:  suggestion,
	}
}

// ValidateEmailBatch processes multiple emails using the same validation logic.
// Each item is processed independently.
func (s *Service) ValidateEmailBatch(ctx context.Context, emails []string) []BatchItem {

	const maxWorkers = 8
	const perItemTimeout = 2 * time.Second

	results := make([]BatchItem, len(emails))
	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for i, email := range emails {
		wg.Add(1)
		sem <- struct{}{}
		go func(i int, email string) {
			defer wg.Done()
			defer func() { <-sem }()

			itemCtx, cancel := context.WithTimeout(ctx, perItemTimeout)
			defer cancel()

			res := s.ValidateEmail(itemCtx, email)
			results[i] = BatchItem{
				Email:       res.Email,
				Valid:       res.Valid,
				SyntaxValid: res.SyntaxValid,
				MXValid:     res.MxValid,
				Disposable:  res.Disposable,
				Normalized:  res.Normalized,
				Domain:      res.Domain,
				Suggestion:  res.Suggestion,
			}
		}(i, email)
	}
	wg.Wait()

	return results
}

// isValidSyntax reports whether email is a syntactically valid RFC 5322 plain
// addr-spec. Display-name formats such as "Name <user@example.com>" and
// angle-addr forms such as "<user@example.com>" are rejected.
func isValidSyntax(email string) bool {
	return addrSpecRe.MatchString(email)
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
	if _, ok := commonDomains[domain]; ok {
		return nil
	}

	best := ""
	bestDist := threshold + 1
	bestLenDiff := 0

	for d := range commonDomains {
		dist := levenshtein.ComputeDistance(domain, d)
		lenDiff := len(domain) - len(d)
		if lenDiff < 0 {
			lenDiff = -lenDiff
		}
		if dist < bestDist || (dist == bestDist && lenDiff < bestLenDiff) {
			bestDist = dist
			bestLenDiff = lenDiff
			best = d
		}
	}

	if bestDist <= threshold {
		return &best
	}
	return nil
}

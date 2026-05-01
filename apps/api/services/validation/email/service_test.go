package email

import (
	"context"
	"testing"
	"requiems-api/platform/httpx"
)

func newTestService() *Service {
	return NewService()
}

// ---- isValidSyntax ----------------------------------------------------------

func TestIsValidSyntax(t *testing.T) {
	cases := []struct {
		email string
		want  bool
	}{
		// Standard valid addr-specs.
		{"user@example.com", true},
		{"user+tag@example.co.uk", true},
		{"user.name@sub.domain.com", true},
		// Quoted local part — valid per RFC 5322 §3.4.1.
		{`"quoted user"@example.com`, true},
		{`"user@domain"@example.com`, true},
		// Domain literal — valid per RFC 5322 §3.4.1.
		{"user@[127.0.0.1]", true},
		// Empty / no address.
		{"", false},
		{"notanemail", false},
		// Missing local part or domain.
		{"@nodomain.com", false},
		{"noatsign.com", false},
		// Multiple @ signs.
		{"double@@domain.com", false},
		// Dot violations in local part.
		{".user@example.com", false},
		{"user.@example.com", false},
		{"user..name@example.com", false},
		// Display-name format — rejected; input must be plain addr-spec.
		{"Display Name <user@example.com>", false},
		// Angle-addr without display name — also rejected.
		{"<user@example.com>", false},
		// Whitespace in local part — not an atext character.
		{"user name@example.com", false},
	}

	for _, tc := range cases {
		t.Run(tc.email, func(t *testing.T) {
			if got := isValidSyntax(tc.email); got != tc.want {
				t.Errorf("isValidSyntax(%q) = %v, want %v", tc.email, got, tc.want)
			}
		})
	}
}

// ---- suggestDomain ----------------------------------------------------------

func TestSuggestDomain(t *testing.T) {
	cases := []struct {
		domain     string
		wantNil    bool
		wantResult string
	}{
		// Exact matches → no suggestion.
		{"gmail.com", true, ""},
		{"outlook.com", true, ""},
		{"yahoo.com", true, ""},
		// Common typos → should suggest.
		{"gmial.com", false, "gmail.com"},
		{"gamil.com", false, "gmail.com"},
		{"outllook.com", false, "outlook.com"},
		{"yaho.com", false, "yahoo.com"},
		{"gmaill.com", false, "gmail.com"},
		// Too far away → no suggestion.
		{"completely-unknown-domain.io", true, ""},
		{"xn--nxasmq6b.com", true, ""},
	}

	for _, tc := range cases {
		t.Run(tc.domain, func(t *testing.T) {
			got := suggestDomain(tc.domain)
			if tc.wantNil {
				if got != nil {
					t.Errorf("suggestDomain(%q) = %q, want nil", tc.domain, *got)
				}
				return
			}
			if got == nil {
				t.Fatalf("suggestDomain(%q) = nil, want %q", tc.domain, tc.wantResult)
			}
			if *got != tc.wantResult {
				t.Errorf("suggestDomain(%q) = %q, want %q", tc.domain, *got, tc.wantResult)
			}
		})
	}
}

// ---- ValidateEmail (syntax path, no network) --------------------------------

func TestValidateEmail_InvalidSyntax(t *testing.T) {
	svc := newTestService()

	cases := []string{"notanemail", "", "@nodomain.com", "double@@test.com"}
	for _, email := range cases {
		t.Run(email, func(t *testing.T) {
			result := svc.ValidateEmail(context.Background(), email)
			if result.Valid {
				t.Errorf("expected Valid=false for %q", email)
			}
			if result.SyntaxValid {
				t.Errorf("expected SyntaxValid=false for %q", email)
			}
		})
	}
}

// ---- ValidateEmail (integration: real DNS) ----------------------------------

func TestValidateEmail_ValidGmail(t *testing.T) {
	svc := newTestService()

	result := svc.ValidateEmail(context.Background(), "user@gmail.com")

	if !result.SyntaxValid {
		t.Error("expected SyntaxValid=true")
	}
	if !result.MxValid {
		t.Error("expected MxValid=true for gmail.com (has MX records)")
	}
	if !result.Valid {
		t.Error("expected Valid=true")
	}
	if *result.Domain != "gmail.com" {
		t.Errorf("expected Domain=gmail.com, got %q", *result.Domain)
	}
	if result.Suggestion != nil {
		t.Errorf("expected Suggestion=nil for known-good domain, got %q", *result.Suggestion)
	}
}

func TestValidateEmail_TypoDomain(t *testing.T) {
	svc := newTestService()

	// gmial.com is a real domain that likely has no MX records; we care about
	// the suggestion field, not the MX result.
	result := svc.ValidateEmail(context.Background(), "user@gmial.com")

	if !result.SyntaxValid {
		t.Error("expected SyntaxValid=true")
	}
	if result.Suggestion == nil {
		t.Fatal("expected a non-nil suggestion for gmial.com")
	}
	if *result.Suggestion != "gmail.com" {
		t.Errorf("expected suggestion=gmail.com, got %q", *result.Suggestion)
	}
}

func TestValidateEmail_GmailPlusNormalized(t *testing.T) {
	svc := newTestService()

	result := svc.ValidateEmail(context.Background(), "User.Name+tag@gmail.com")

	if !result.SyntaxValid {
		t.Error("expected SyntaxValid=true")
	}
	if *result.Normalized == "User.Name+tag@gmail.com" {
		t.Error("expected email to be normalized (Gmail strips dots and plus-tags)")
	}
}

// ---Batch



// ---- ValidateEmailBatch -----------------------------------------------------

func TestBatchRequest_Valid_OneEmail(t *testing.T) {
	req := BatchRequest{Emails: []string{"user@gmail.com"}}

	if err := httpx.Validate.Struct(&req); err != nil {
		t.Errorf("expected no error for single valid email, got %v", err)
	}
}

func TestBatchRequest_Valid_MaxEmails(t *testing.T) {
	emails := make([]string, 50)
	for i := range emails {
		emails[i] = "user@gmail.com"
	}

	req := BatchRequest{Emails: emails}

	if err := httpx.Validate.Struct(&req); err != nil {
		t.Errorf("expected no error for 50 emails, got %v", err)
	}
}

func TestBatchRequest_Empty_Fails(t *testing.T) {
	req := BatchRequest{Emails: []string{}}

	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for empty emails array")
	}
}

func TestBatchRequest_OverLimit_Fails(t *testing.T) {
	emails := make([]string, 51)
	for i := range emails {
		emails[i] = "user@gmail.com"
	}

	req := BatchRequest{Emails: emails}

	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for >50 emails")
	}
}

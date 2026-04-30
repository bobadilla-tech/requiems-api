package normalize

import (
	"testing"

	normalizer "github.com/bobadilla-tech/go-email-normalizer"
)

// containsChange reports whether target appears in the changes slice.
func containsChange(changes []normalizer.Change, target normalizer.Change) bool {
	for _, c := range changes {
		if c == target {
			return true
		}
	}
	return false
}

// --- Valid email tests ---

func TestService_Normalize_OriginalPreserved(t *testing.T) {
	svc := NewService()

	input := "User@Example.com"
	result, err := svc.Normalize(input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Original != input {
		t.Errorf("expected Original %q, got %q", input, result.Original)
	}
}

func TestService_Normalize_SplitsLocalAndDomain(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("user@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Local != "user" {
		t.Errorf("expected Local %q, got %q", "user", result.Local)
	}
	if result.Domain != "example.com" {
		t.Errorf("expected Domain %q, got %q", "example.com", result.Domain)
	}
}

func TestService_Normalize_LowercasesDomain(t *testing.T) {
	svc := NewService()

	// For unknown providers the local part is preserved (case-sensitive per
	// RFC 5321); only the domain is lowercased.
	result, err := svc.Normalize("User@Example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Normalized != "User@example.com" {
		t.Errorf("expected normalized %q, got %q", "User@example.com", result.Normalized)
	}
	if !containsChange(result.Changes, normalizer.ChangeLowercase) {
		t.Errorf("expected ChangeLowercase in changes, got %v", result.Changes)
	}
}

func TestService_Normalize_NoChangesForAlreadyNormalisedEmail(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("user@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Normalized != "user@example.com" {
		t.Errorf("expected normalized to equal input, got %q", result.Normalized)
	}
	if len(result.Changes) != 0 {
		t.Errorf("expected no changes for already-normalised email, got %v", result.Changes)
	}
	if result.Changes == nil {
		t.Error("expected Changes to be an empty slice, not nil")
	}
}

func TestService_Normalize_GmailRemovesDots(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("te.st.user@gmail.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Normalized != "testuser@gmail.com" {
		t.Errorf("expected normalized %q, got %q", "testuser@gmail.com", result.Normalized)
	}
	if !containsChange(result.Changes, normalizer.ChangeRemovedDots) {
		t.Errorf("expected ChangeRemovedDots in changes, got %v", result.Changes)
	}
}

func TestService_Normalize_GmailRemovesPlusTag(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("testuser+spam@gmail.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Normalized != "testuser@gmail.com" {
		t.Errorf("expected normalized %q, got %q", "testuser@gmail.com", result.Normalized)
	}
	if !containsChange(result.Changes, normalizer.ChangeRemovedPlusTag) {
		t.Errorf("expected ChangeRemovedPlusTag in changes, got %v", result.Changes)
	}
}

func TestService_Normalize_GooglemailCanonicalisedToGmail(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("user@googlemail.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Normalized != "user@gmail.com" {
		t.Errorf("expected normalized %q, got %q", "user@gmail.com", result.Normalized)
	}
	if !containsChange(result.Changes, normalizer.ChangeCanonicalisedDomain) {
		t.Errorf("expected ChangeCanonicalisedDomain in changes, got %v", result.Changes)
	}
}

func TestService_Normalize_WhitespaceTrimmed(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("  user@example.com  ")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Normalized != "user@example.com" {
		t.Errorf("expected whitespace to be trimmed, got %q", result.Normalized)
	}
	if !containsChange(result.Changes, normalizer.ChangeTrimmedWhitespace) {
		t.Errorf("expected ChangeTrimmedWhitespace in changes, got %v", result.Changes)
	}
}

func TestService_Normalize_NormalizedFieldPopulated(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("user@example.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Normalized == "" {
		t.Error("expected non-empty Normalized field")
	}
}

func TestService_Normalize_LocalAndDomainMatchNormalized(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("Test.User+tag@Gmail.com")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result.Local+"@"+result.Domain != result.Normalized {
		t.Errorf("expected Local@Domain to equal Normalized %q, got %q@%q",
			result.Normalized, result.Local, result.Domain)
	}
}

// --- Invalid email tests ---

func TestService_Normalize_InvalidEmail_NoAtSign(t *testing.T) {
	svc := NewService()

	_, err := svc.Normalize("not-an-email")
	if err == nil {
		t.Error("expected an error for input without '@', got nil")
	}
}

func TestService_Normalize_InvalidEmail_EmptyString(t *testing.T) {
	svc := NewService()

	_, err := svc.Normalize("")
	if err == nil {
		t.Error("expected an error for empty input, got nil")
	}
}

func TestService_Normalize_InvalidEmail_MissingLocal(t *testing.T) {
	svc := NewService()

	_, err := svc.Normalize("@example.com")
	if err == nil {
		t.Error("expected an error for input missing local part, got nil")
	}
}

func TestService_Normalize_InvalidEmail_MissingDomain(t *testing.T) {
	svc := NewService()

	_, err := svc.Normalize("user@")
	if err == nil {
		t.Error("expected an error for input missing domain, got nil")
	}
}

func TestService_Normalize_InvalidEmail_DotlessDomain(t *testing.T) {
	svc := NewService()

	_, err := svc.Normalize("user@gmailcom")
	if err == nil {
		t.Error("expected an error for dotless domain 'user@gmailcom', got nil")
	}
}

func TestService_Normalize_InvalidEmail_OnlyAtSign(t *testing.T) {
	svc := NewService()

	_, err := svc.Normalize("@")
	if err == nil {
		t.Error("expected an error for bare '@', got nil")
	}
}

func TestService_Normalize_InvalidEmail_ReturnsZeroValue(t *testing.T) {
	svc := NewService()

	result, err := svc.Normalize("not-an-email")
	if err == nil {
		t.Fatal("expected an error, got nil")
	}
	if result.Original != "" || result.Normalized != "" || result.Local != "" || result.Domain != "" {
		t.Errorf("expected zero-value EmailNormalization on error, got %+v", result)
	}
}

func TestService_NormalizeBatch_OrderAndValidity(t *testing.T) {
	svc := NewService()

	got := svc.NormalizeBatch([]string{"user@example.com", "not-an-email", "te.st@gmail.com"})
	if got.Total != 3 {
		t.Fatalf("total: want 3, got %d", got.Total)
	}
	if len(got.Results) != 3 {
		t.Fatalf("results len: want 3, got %d", len(got.Results))
	}
	if !got.Results[0].Valid || got.Results[0].Normalized != "user@example.com" {
		t.Errorf("result[0]: want valid normalized user@example.com, got %+v", got.Results[0])
	}
	if got.Results[1].Valid || got.Results[1].Message == "" {
		t.Errorf("result[1]: want invalid with message, got %+v", got.Results[1])
	}
	if !got.Results[2].Valid || got.Results[2].Normalized != "test@gmail.com" {
		t.Errorf("result[2]: want gmail normalized, got %+v", got.Results[2])
	}
}

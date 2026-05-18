package normalize

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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
	t.Parallel()
	svc := NewService()

	input := "User@Example.com"
	result, err := svc.Normalize(input)
	require.NoError(t, err)

	assert.Equal(t, input, result.Original)
}

func TestService_Normalize_SplitsLocalAndDomain(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("user@example.com")
	require.NoError(t, err)

	assert.Equal(t, "user", result.Local)
	assert.Equal(t, "example.com", result.Domain)
}

func TestService_Normalize_LowercasesDomain(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// For unknown providers the local part is preserved (case-sensitive per
	// RFC 5321); only the domain is lowercased.
	result, err := svc.Normalize("User@Example.com")
	require.NoError(t, err)

	assert.Equal(t, "User@example.com", result.Normalized)
	assert.True(t, containsChange(result.Changes, normalizer.ChangeLowercase), "expected ChangeLowercase in changes, got %v", result.Changes)
}

func TestService_Normalize_NoChangesForAlreadyNormalisedEmail(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("user@example.com")
	require.NoError(t, err)

	assert.Equal(t, "user@example.com", result.Normalized)
	assert.Len(t, result.Changes, 0)
	assert.NotNil(t, result.Changes)
}

func TestService_Normalize_GmailRemovesDots(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("te.st.user@gmail.com")
	require.NoError(t, err)

	assert.Equal(t, "testuser@gmail.com", result.Normalized)
	assert.True(t, containsChange(result.Changes, normalizer.ChangeRemovedDots), "expected ChangeRemovedDots in changes, got %v", result.Changes)
}

func TestService_Normalize_GmailRemovesPlusTag(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("testuser+spam@gmail.com")
	require.NoError(t, err)

	assert.Equal(t, "testuser@gmail.com", result.Normalized)
	assert.True(t, containsChange(result.Changes, normalizer.ChangeRemovedPlusTag), "expected ChangeRemovedPlusTag in changes, got %v", result.Changes)
}

func TestService_Normalize_GooglemailCanonicalisedToGmail(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("user@googlemail.com")
	require.NoError(t, err)

	assert.Equal(t, "user@gmail.com", result.Normalized)
	assert.True(t, containsChange(result.Changes, normalizer.ChangeCanonicalisedDomain), "expected ChangeCanonicalisedDomain in changes, got %v", result.Changes)
}

func TestService_Normalize_WhitespaceTrimmed(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("  user@example.com  ")
	require.NoError(t, err)

	assert.Equal(t, "user@example.com", result.Normalized)
	assert.True(t, containsChange(result.Changes, normalizer.ChangeTrimmedWhitespace), "expected ChangeTrimmedWhitespace in changes, got %v", result.Changes)
}

func TestService_Normalize_NormalizedFieldPopulated(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("user@example.com")
	require.NoError(t, err)

	assert.NotEmpty(t, result.Normalized)
}

func TestService_Normalize_LocalAndDomainMatchNormalized(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("Test.User+tag@Gmail.com")
	require.NoError(t, err)

	assert.Equal(t, result.Normalized, result.Local+"@"+result.Domain)
}

// --- Invalid email tests ---

func TestService_Normalize_InvalidEmail_NoAtSign(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Normalize("not-an-email")
	assert.Error(t, err)
}

func TestService_Normalize_InvalidEmail_EmptyString(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Normalize("")
	assert.Error(t, err)
}

func TestService_Normalize_InvalidEmail_MissingLocal(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Normalize("@example.com")
	assert.Error(t, err)
}

func TestService_Normalize_InvalidEmail_MissingDomain(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Normalize("user@")
	assert.Error(t, err)
}

func TestService_Normalize_InvalidEmail_DotlessDomain(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Normalize("user@gmailcom")
	assert.Error(t, err)
}

func TestService_Normalize_InvalidEmail_OnlyAtSign(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Normalize("@")
	assert.Error(t, err)
}

func TestService_Normalize_InvalidEmail_ReturnsZeroValue(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Normalize("not-an-email")
	require.Error(t, err)
	assert.True(t, result.Original == "" && result.Normalized == "" && result.Local == "" && result.Domain == "", "expected zero-value EmailNormalization on error, got %+v", result)
}

func TestService_NormalizeBatch_OrderAndValidity(t *testing.T) {
	t.Parallel()
	svc := NewService()

	got := svc.NormalizeBatch([]string{"user@example.com", "not-an-email", "te.st@gmail.com"})
	require.Equal(t, 3, len(got))
	require.Len(t, got, 3)
	assert.True(t, got[0].Valid && got[0].Normalized == "user@example.com", "result[0]: want valid normalized user@example.com, got %+v", got[0])
	assert.True(t, !got[1].Valid && got[1].Message != "", "result[1]: want invalid with message, got %+v", got[1])
	assert.True(t, got[2].Valid && got[2].Normalized == "test@gmail.com", "result[2]: want gmail normalized, got %+v", got[2])
}

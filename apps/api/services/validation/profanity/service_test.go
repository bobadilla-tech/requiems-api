package profanity

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Check_NoProfanity(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Check(context.Background(), "Hello, world!")

	assert.False(t, result.HasProfanity)
	assert.Equal(t, "Hello, world!", result.Censored)
	assert.Empty(t, result.FlaggedWords)
}

func TestService_Check_WithProfanity(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Check(context.Background(), "What the fuck is this shit")

	assert.True(t, result.HasProfanity)
	assert.Equal(t, "What the **** is this ****", result.Censored)
	assert.Len(t, result.FlaggedWords, 2)
}

func TestService_Check_CaseInsensitive(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// go-away detects "shit" as a substring of "BULLSHIT" and censors only
	// the matched portion; the flagged canonical word is "shit".
	result := svc.Check(context.Background(), "This is BULLSHIT")

	assert.True(t, result.HasProfanity)
	if len(result.FlaggedWords) != 1 || result.FlaggedWords[0] != "shit" {
		t.Errorf("expected flagged word [\"shit\"], got %v", result.FlaggedWords)
	}
}

func TestService_Check_DeduplicatesFlaggedWords(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Check(context.Background(), "shit shit shit")

	assert.Len(t, result.FlaggedWords, 1)
}

func TestService_Check_EmptyFlaggedWordsSlice(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Check(context.Background(), "clean text here")

	// FlaggedWords must be an empty slice, not nil (for consistent JSON serialisation).
	if result.FlaggedWords == nil {
		t.Error("expected FlaggedWords to be an empty slice, not nil")
	}
}

func TestService_Check_EmptyText(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Check(context.Background(), "")

	assert.False(t, result.HasProfanity)
	assert.Equal(t, "", result.Censored)
}

func TestService_Check_LeetSpeak(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// go-away handles leet-speak obfuscation out of the box.
	result := svc.Check(context.Background(), "F   u   C  k th1$ $h!t")

	assert.True(t, result.HasProfanity)
	assert.NotEmpty(t, result.FlaggedWords)
}

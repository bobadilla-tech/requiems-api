package spellcheck

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Check_NoMistakes(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Check("This is a test")
	require.NoError(t, err)

	assert.Equal(t, "This is a test", result.Corrected)
	assert.Len(t, result.Corrections, 0)
}

func TestService_Check_MisspelledWords(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Check("Ths is a tset")
	require.NoError(t, err)

	assert.NotEmpty(t, result.Corrections)

	foundThs := false
	foundTset := false
	for _, c := range result.Corrections {
		if c.Original == "Ths" && c.Position == 0 && c.Suggested != "" {
			foundThs = true
		}
		if c.Original == "tset" && c.Position == 9 && c.Suggested != "" {
			foundTset = true
		}
	}

	assert.True(t, foundThs, "expected correction for Ths at position 0; got %+v", result.Corrections)
	assert.True(t, foundTset, "expected correction for tset at position 9; got %+v", result.Corrections)
}

func TestService_Check_EmptyText(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// Empty text is not a valid request (validate:"required" enforces that at
	// the HTTP layer), but the service itself should return a safe empty result.
	result, err := svc.Check("")
	require.NoError(t, err)

	assert.Equal(t, "", result.Corrected)
	assert.Len(t, result.Corrections, 0)
}

func TestService_Check_CorrectedTextReflectsFixes(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Check("Ths is a tset")
	require.NoError(t, err)

	assert.False(t, result.Corrected == "Ths is a tset", "expected corrected text to differ from misspelled input")
}

func TestService_Check_CorrectionsSliceNotNil(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Check("Hello world")
	require.NoError(t, err)

	assert.NotNil(t, result.Corrections)
}

func TestService_Check_PositionIsRuneOffset(t *testing.T) {
	t.Parallel()
	svc := NewService()
	// "é" is a single rune but 2 UTF-8 bytes.
	// "tset" starts at rune index 2 (é=1, space=1) but byte index 3.
	result, err := svc.Check("é tset")
	require.NoError(t, err)
	require.NotEmpty(t, result.Corrections)
	assert.Equal(t, 2, result.Corrections[0].Position)
}

func TestMatchCase_LowerInput(t *testing.T) {
	t.Parallel()
	got := matchCase("abc", "suggested")
	assert.Equal(t, "suggested", got)
}

func TestMatchCase_CapitalisedInput(t *testing.T) {
	t.Parallel()
	got := matchCase("Abc", "suggested")
	assert.Equal(t, "Suggested", got)
}

func TestMatchCase_AllUpperInput(t *testing.T) {
	t.Parallel()
	got := matchCase("ABC", "suggested")
	assert.Equal(t, "SUGGESTED", got)
}

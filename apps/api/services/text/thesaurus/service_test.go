package thesaurus

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Lookup_KnownWord(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Lookup("happy")
	require.NoError(t, err)

	assert.Equal(t, "happy", result.Word)
	assert.NotEmpty(t, result.Synonyms)
	assert.NotEmpty(t, result.Antonyms)
}

func TestService_Lookup_CaseInsensitive(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []string{"HAPPY", "Happy", "hApPy", "happy"}
	for _, input := range tests {
		t.Run(input, func(t *testing.T) {
			t.Parallel()
			result, err := svc.Lookup(input)
			require.NoError(t, err)
			assert.Equal(t, "happy", result.Word)
		})
	}
}

func TestService_Lookup_UnknownWord(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Lookup("zzyzx")
	require.Error(t, err)
}

func TestService_Lookup_ReturnsNonNilSlices(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result, err := svc.Lookup("happy")
	require.NoError(t, err)

	assert.NotNil(t, result.Synonyms)
	assert.NotNil(t, result.Antonyms)
}

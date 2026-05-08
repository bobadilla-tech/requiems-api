package detectlanguage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService_Detect_French(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Detect("Bonjour, comment ça va?")

	assert.Equal(t, "French", result.Language)
	assert.Equal(t, "fr", result.Code)
	assert.True(t, result.Confidence > 0)
}

func TestService_Detect_English(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Detect("The quick brown fox jumps over the lazy dog")

	assert.Equal(t, "English", result.Language)
	assert.Equal(t, "en", result.Code)
	assert.True(t, result.Confidence > 0)
}

func TestService_Detect_Spanish(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Detect("El rápido zorro marrón salta sobre el perro perezoso")

	assert.Equal(t, "Spanish", result.Language)
	assert.Equal(t, "es", result.Code)
	assert.True(t, result.Confidence > 0)
}

func TestService_Detect_ConfidenceRange(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Detect("This is a longer English sentence to ensure reliable detection.")

	assert.True(t, result.Confidence >= 0 && result.Confidence <= 1, "confidence %f is outside expected range [0, 1]", result.Confidence)
}

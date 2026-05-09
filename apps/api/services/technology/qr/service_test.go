package qr

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Generate_ValidData(t *testing.T) {
	t.Parallel()
	svc := NewService()

	png, err := svc.Generate("https://example.com", 256, "medium")
	require.NoError(t, err)

	assert.NotEmpty(t, png)

	// Verify PNG signature (\x89PNG\r\n\x1a\n)
	if len(png) < 8 || string(png[:4]) != "\x89PNG" {
		t.Error("expected valid PNG signature")
	}
}

func TestService_Generate_SmallSize(t *testing.T) {
	t.Parallel()
	svc := NewService()

	png, err := svc.Generate("hello", 50, "")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
}

func TestService_Generate_LargeSize(t *testing.T) {
	t.Parallel()
	svc := NewService()

	png, err := svc.Generate("https://example.com/very/long/path?foo=bar&baz=qux", 1000, "low")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
}

func TestService_Generate_HighestRecovery(t *testing.T) {
	t.Parallel()
	svc := NewService()

	png, err := svc.Generate("https://example.com", 256, "highest")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
}

func TestService_Generate_AllRecoveryLevels(t *testing.T) {
	t.Parallel()
	svc := NewService()

	levels := []string{"low", "medium", "high", "highest"}
	for _, level := range levels {
		png, err := svc.Generate("test", 256, level)
		if err != nil {
			t.Errorf("unexpected error for recovery=%q: %v", level, err)
			continue
		}
		assert.NotEmpty(t, png)
	}
}

func TestService_Generate_EmptyData(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.Generate("", 256, "medium")
	assert.Error(t, err)
}

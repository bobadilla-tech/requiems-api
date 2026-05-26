package qr

import (
	"encoding/base64"
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

func TestService_GenerateBatch(t *testing.T) {
	t.Parallel()
	svc := NewService()

	items := []Query{
		{Data: "https://example.com", Size: 256},
		{Data: "hello world", Size: 128},
	}
	results := svc.GenerateBatch(items)

	require.Len(t, results, 2)
	for i, item := range results {
		assert.NotEmpty(t, item.Image, "expected image at index %d", i)
		assert.Empty(t, item.Error, "expected no error at index %d", i)

		decoded, err := base64.StdEncoding.DecodeString(item.Image)
		require.NoError(t, err, "expected valid base64 at index %d", i)
		if len(decoded) < 4 || string(decoded[:4]) != "\x89PNG" {
			t.Errorf("expected PNG signature at index %d", i)
		}
	}
	assert.Equal(t, "https://example.com", results[0].Data)
	assert.Equal(t, "hello world", results[1].Data)
}

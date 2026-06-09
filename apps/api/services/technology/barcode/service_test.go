package barcode

import (
	"context"
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Generate_Code128(t *testing.T) {
	t.Parallel()
	svc := NewService()

	png, width, height, err := svc.Generate("123456789", "code128")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
	assert.Equal(t, string(png[:4]), "\x89PNG")
	assert.Equal(t, defaultWidth, width)
	assert.Equal(t, defaultHeight, height)
}

func TestService_Generate_Code93(t *testing.T) {
	t.Parallel()
	svc := NewService()

	png, _, _, err := svc.Generate("HELLO", "code93")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
}

func TestService_Generate_Code39(t *testing.T) {
	t.Parallel()
	svc := NewService()

	png, _, _, err := svc.Generate("HELLO123", "code39")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
}

func TestService_Generate_EAN8(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// 7 digits (checksum auto-calculated)
	png, _, _, err := svc.Generate("1234567", "ean8")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
}

func TestService_Generate_EAN13(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// 12 digits (checksum auto-calculated)
	png, _, _, err := svc.Generate("123456789012", "ean13")
	require.NoError(t, err)

	assert.NotEmpty(t, png)
}

func TestService_Generate_EAN8_InvalidLength(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, _, _, err := svc.Generate("123", "ean8")
	require.Error(t, err)
}

func TestService_Generate_EAN13_InvalidLength(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, _, _, err := svc.Generate("123456789", "ean13")
	require.Error(t, err)
}

func TestService_Generate_UnsupportedType(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, _, _, err := svc.Generate("hello", "qrcode")
	require.Error(t, err)
}

func TestService_Generate_EmptyData_Code128(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, _, _, err := svc.Generate("", "code128")
	require.Error(t, err)
}

// ── GenerateBatch ──────────────────────────────────────────────────────────

func TestService_GenerateBatch_AllValid(t *testing.T) {
	t.Parallel()
	svc := NewService()

	items := []BatchItem{
		{Data: "123456789", Type: "code128"},
		{Data: "HELLO", Type: "code93"},
		{Data: "HELLO123", Type: "code39"},
		{Data: "1234567", Type: "ean8"},
		{Data: "123456789012", Type: "ean13"},
	}

	results := svc.GenerateBatch(context.Background(), items)

	require.Len(t, results, 5)
	for i, r := range results {
		assert.True(t, r.Success, "item %d (%s) should succeed", i, items[i].Type)
		assert.NotEmpty(t, r.Image, "item %d should have a base64 image", i)
		assert.Empty(t, r.Error, "item %d should have no error", i)
		assert.Equal(t, defaultWidth, r.Width, "item %d width", i)
		assert.Equal(t, defaultHeight, r.Height, "item %d height", i)
		assert.Equal(t, items[i].Type, r.Type, "item %d type should match input", i)
	}
}

func TestService_GenerateBatch_PartialFailure(t *testing.T) {
	t.Parallel()
	svc := NewService()

	items := []BatchItem{
		{Data: "123456789", Type: "code128"},  // valid
		{Data: "123", Type: "ean8"},           // invalid: wrong digit count
		{Data: "123456789012", Type: "ean13"}, // valid
		{Data: "HELLO", Type: "qrcode"},       // invalid: unsupported type
	}

	results := svc.GenerateBatch(context.Background(), items)

	require.Len(t, results, 4)

	assert.True(t, results[0].Success)
	assert.NotEmpty(t, results[0].Image)
	assert.Empty(t, results[0].Error)

	assert.False(t, results[1].Success)
	assert.Empty(t, results[1].Image)
	assert.NotEmpty(t, results[1].Error)
	assert.Equal(t, 0, results[1].Width)
	assert.Equal(t, 0, results[1].Height)

	assert.True(t, results[2].Success)
	assert.NotEmpty(t, results[2].Image)

	assert.False(t, results[3].Success)
	assert.NotEmpty(t, results[3].Error)
}

func TestService_GenerateBatch_AllFail(t *testing.T) {
	t.Parallel()
	svc := NewService()

	items := []BatchItem{
		{Data: "123", Type: "ean8"},     // wrong digit count
		{Data: "abc", Type: "ean13"},    // non-numeric
		{Data: "HELLO", Type: "qrcode"}, // unsupported type
	}

	results := svc.GenerateBatch(context.Background(), items)

	require.Len(t, results, 3)
	for i, r := range results {
		assert.False(t, r.Success, "item %d should fail", i)
		assert.Empty(t, r.Image, "item %d image should be empty", i)
		assert.NotEmpty(t, r.Error, "item %d should have an error message", i)
		assert.Equal(t, 0, r.Width)
		assert.Equal(t, 0, r.Height)
	}
}

func TestService_GenerateBatch_SingleItem(t *testing.T) {
	t.Parallel()
	svc := NewService()

	results := svc.GenerateBatch(context.Background(), []BatchItem{
		{Data: "123456789", Type: "code128"},
	})

	require.Len(t, results, 1)
	assert.True(t, results[0].Success)
	assert.NotEmpty(t, results[0].Image)
}

func TestService_GenerateBatch_MaxItems(t *testing.T) {
	t.Parallel()
	svc := NewService()

	items := make([]BatchItem, maxBatchSize)
	for i := range items {
		items[i] = BatchItem{Data: "123456789", Type: "code128"}
	}

	results := svc.GenerateBatch(context.Background(), items)

	require.Len(t, results, maxBatchSize)
	for i, r := range results {
		assert.True(t, r.Success, "item %d should succeed", i)
	}
}

func TestService_GenerateBatch_EAN13_ValidLengths(t *testing.T) {
	t.Parallel()
	svc := NewService()

	results := svc.GenerateBatch(context.Background(), []BatchItem{
		{Data: "123456789012", Type: "ean13"},  // 12 digits
		{Data: "1234567890128", Type: "ean13"}, // 13 digits
	})

	require.Len(t, results, 2)
	assert.True(t, results[0].Success)
	assert.True(t, results[1].Success)
}

func TestService_GenerateBatch_ImageDecodesAsValidPNG(t *testing.T) {
	t.Parallel()
	svc := NewService()

	results := svc.GenerateBatch(context.Background(), []BatchItem{
		{Data: "123456789", Type: "code128"},
	})

	require.Len(t, results, 1)
	require.True(t, results[0].Success)

	decoded, err := base64.StdEncoding.DecodeString(results[0].Image)
	require.NoError(t, err)
	assert.Equal(t, "\x89PNG", string(decoded[:4]))
}

func TestService_GenerateBatch_EmptyDataInBatch(t *testing.T) {
	t.Parallel()
	svc := NewService()

	results := svc.GenerateBatch(context.Background(), []BatchItem{
		{Data: "", Type: "code128"},
	})

	require.Len(t, results, 1)
	assert.False(t, results[0].Success)
	assert.NotEmpty(t, results[0].Error)
}

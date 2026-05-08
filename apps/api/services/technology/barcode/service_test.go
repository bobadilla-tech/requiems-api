package barcode

import (
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

package swift

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
)

func TestSanitizeSWIFT_Valid8Char(t *testing.T) {
	t.Parallel()
	got, err := sanitizeSWIFT("DEUTDEDB")
	require.NoError(t, err)
	assert.Equal(t, "DEUTDEDBXXX", got)
}

func TestSanitizeSWIFT_Valid11Char(t *testing.T) {
	t.Parallel()
	got, err := sanitizeSWIFT("DEUTDEDB001")
	require.NoError(t, err)
	assert.Equal(t, "DEUTDEDB001", got)
}

func TestSanitizeSWIFT_Lowercase(t *testing.T) {
	t.Parallel()
	got, err := sanitizeSWIFT("deutdedb")
	require.NoError(t, err)
	assert.Equal(t, "DEUTDEDBXXX", got)
}

func TestSanitizeSWIFT_WithSpaces(t *testing.T) {
	t.Parallel()
	got, err := sanitizeSWIFT("  DEUTDEDB  ")
	require.NoError(t, err)
	assert.Equal(t, "DEUTDEDBXXX", got)
}

func TestSanitizeSWIFT_PrimaryOfficeXXX(t *testing.T) {
	t.Parallel()
	got, err := sanitizeSWIFT("DEUTDEDBXXX")
	require.NoError(t, err)
	assert.Equal(t, "DEUTDEDBXXX", got)
}

func TestSanitizeSWIFT_AlphanumericLocation(t *testing.T) {
	t.Parallel()
	// Location code with digit is valid (chars 7-8 are alphanumeric).
	got, err := sanitizeSWIFT("DEUTDE2B")
	require.NoError(t, err)
	assert.Equal(t, "DEUTDE2BXXX", got)
}

func TestSanitizeSWIFT_TooShort(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("DEUTDE")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_TooLong(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("DEUTDEDB001X")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_9Chars(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("DEUTDEDB0")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_10Chars(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("DEUTDEDB01")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_Empty(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_DigitInBankCode(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("1EUTDEDB")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_DigitInCountryCode(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("DEUT1EDB")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_InvalidBranchCode(t *testing.T) {
	t.Parallel()
	_, err := sanitizeSWIFT("DEUTDEDB0-1")
	assertAppError(t, err)
}

func TestSanitizeSWIFT_8Char_AppendXXX(t *testing.T) {
	t.Parallel()
	got, err := sanitizeSWIFT("CHASUS33")
	require.NoError(t, err)
	assert.Len(t, got, 11)
	assert.Equal(t, "XXX", got[8:])
}

func TestSanitizeAlphaCode_Valid(t *testing.T) {
	t.Parallel()
	got, err := sanitizeAlphaCode("de", 2, "country code")
	require.NoError(t, err)
	require.Equal(t, "DE", got)
}

func TestSanitizeAlphaCode_InvalidLength(t *testing.T) {
	t.Parallel()
	_, err := sanitizeAlphaCode("D", 2, "country code")
	assertAppError(t, err)
}

func TestSanitizeAlphaCode_InvalidChars(t *testing.T) {
	t.Parallel()
	_, err := sanitizeAlphaCode("D1", 2, "country code")
	assertAppError(t, err)
}

// assertAppError verifies that err is a bad_request *svcerr.Error (KindInvalid).
// All sanitizeSWIFT validation errors return this kind and code.
func assertAppError(t *testing.T, err error) {
	t.Helper()
	require.Error(t, err)
	var se *svcerr.Error
	require.ErrorAs(t, err, &se)
	assert.Equal(t, svcerr.KindInvalid, se.Kind)
	assert.Equal(t, "bad_request", se.Code)
}

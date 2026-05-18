package bin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// ---- sanitizeBIN ----

func TestSanitizeBIN_Valid6Digit(t *testing.T) {
	t.Parallel()
	got, err := sanitizeBIN("424242")
	require.NoError(t, err)
	assert.Equal(t, "424242", got)
}

func TestSanitizeBIN_Valid8Digit(t *testing.T) {
	t.Parallel()
	got, err := sanitizeBIN("42424242")
	require.NoError(t, err)
	assert.Equal(t, "42424242", got)
}

func TestSanitizeBIN_StripsDashes(t *testing.T) {
	t.Parallel()
	got, err := sanitizeBIN("4242-42")
	require.NoError(t, err)
	assert.Equal(t, "424242", got)
}

func TestSanitizeBIN_StripsSpaces(t *testing.T) {
	t.Parallel()
	got, err := sanitizeBIN("  4242 42  ")
	require.NoError(t, err)
	assert.Equal(t, "424242", got)
}

func TestSanitizeBIN_TooShort(t *testing.T) {
	t.Parallel()
	_, err := sanitizeBIN("42424")
	require.Error(t, err)
}

func TestSanitizeBIN_TooLong(t *testing.T) {
	t.Parallel()
	_, err := sanitizeBIN("424242424")
	require.Error(t, err)
}

func TestSanitizeBIN_Empty(t *testing.T) {
	t.Parallel()
	_, err := sanitizeBIN("")
	require.Error(t, err)
}

func TestSanitizeBIN_NonDigits(t *testing.T) {
	t.Parallel()
	_, err := sanitizeBIN("abcdef")
	require.Error(t, err)
}

func TestSanitizeBIN_MixedAlphaDigits(t *testing.T) {
	t.Parallel()
	_, err := sanitizeBIN("4242ab")
	require.Error(t, err)
}

// ---- luhnValid ----

func TestLuhnValid_KnownValid(t *testing.T) {
	t.Parallel()
	// 424242 — well-known Visa test BIN with valid Luhn
	assert.True(t, luhnValid("424242"), "expected luhnValid(424242) = true")
}

func TestLuhnValid_KnownInvalid(t *testing.T) {
	t.Parallel()
	assert.False(t, luhnValid("123456"), "expected luhnValid(123456) = false")
}

func TestLuhnValid_AllZeros(t *testing.T) {
	t.Parallel()
	// 000000 → sum = 0, 0 % 10 = 0 → valid
	assert.True(t, luhnValid("000000"), "expected luhnValid(000000) = true")
}

func TestLuhnValid_8DigitValid(t *testing.T) {
	t.Parallel()
	assert.True(t, luhnValid("42424242"), "expected luhnValid(42424242) = true")
}

// ---- detectScheme ----

func TestDetectScheme_Visa(t *testing.T) {
	t.Parallel()
	cases := []string{"424242", "400000", "499999"}
	for _, bin := range cases {
		assert.Equal(t, "visa", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Mastercard5Series(t *testing.T) {
	t.Parallel()
	cases := map[string]string{
		"510000": "mastercard",
		"520000": "mastercard",
		"530000": "mastercard",
		"540000": "mastercard",
		"550000": "mastercard",
	}
	for bin, want := range cases {
		assert.Equal(t, want, detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Mastercard2Series(t *testing.T) {
	t.Parallel()
	cases := []string{"222100", "272000", "250000"}
	for _, bin := range cases {
		assert.Equal(t, "mastercard", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Mastercard2SeriesBoundaryLow(t *testing.T) {
	t.Parallel()
	// 2220xx is NOT Mastercard (range starts at 2221)
	assert.NotEqual(t, "mastercard", detectScheme("222099"), "detectScheme(222099) should not be mastercard")
}

func TestDetectScheme_Mastercard2SeriesBoundaryHigh(t *testing.T) {
	t.Parallel()
	// 2721xx is NOT Mastercard (range ends at 2720)
	assert.NotEqual(t, "mastercard", detectScheme("272100"), "detectScheme(272100) should not be mastercard")
}

func TestDetectScheme_Amex(t *testing.T) {
	t.Parallel()
	cases := []string{"340000", "370000", "378282"}
	for _, bin := range cases {
		assert.Equal(t, "amex", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Discover6011(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "discover", detectScheme("601100"), "detectScheme(601100)")
}

func TestDetectScheme_Discover622Inside(t *testing.T) {
	t.Parallel()
	cases := []string{"622126", "622500", "622925"}
	for _, bin := range cases {
		assert.Equal(t, "discover", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Discover622BoundaryLow(t *testing.T) {
	t.Parallel()
	// 622125 is below the Discover range → UnionPay
	assert.NotEqual(t, "discover", detectScheme("622125"), "detectScheme(622125) should not be discover")
}

func TestDetectScheme_Discover622BoundaryHigh(t *testing.T) {
	t.Parallel()
	// 622926 is above the Discover range → UnionPay
	assert.NotEqual(t, "discover", detectScheme("622926"), "detectScheme(622926) should not be discover")
}

func TestDetectScheme_Discover65(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "discover", detectScheme("650000"), "detectScheme(650000)")
}

func TestDetectScheme_JCB(t *testing.T) {
	t.Parallel()
	cases := []string{"352800", "358900", "356000"}
	for _, bin := range cases {
		assert.Equal(t, "jcb", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Diners(t *testing.T) {
	t.Parallel()
	cases := []string{"300000", "305999", "360000", "380000"}
	for _, bin := range cases {
		assert.Equal(t, "diners", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_UnionPay(t *testing.T) {
	t.Parallel()
	cases := []string{"620000", "810000"}
	for _, bin := range cases {
		assert.Equal(t, "unionpay", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Maestro(t *testing.T) {
	t.Parallel()
	cases := []string{"630400", "675900", "676100", "676200", "676300"}
	for _, bin := range cases {
		assert.Equal(t, "maestro", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Mir(t *testing.T) {
	t.Parallel()
	cases := []string{"220000", "220100", "220200", "220300", "220400"}
	for _, bin := range cases {
		assert.Equal(t, "mir", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_MirVsMastercardBoundary(t *testing.T) {
	t.Parallel()
	// 2205xx is between Mir (≤2204) and Mastercard 2-series (≥2221): neither
	got := detectScheme("220500")
	assert.NotEqual(t, "mir", got, "detectScheme(220500) = %q, expected neither mir nor mastercard", got)
	assert.NotEqual(t, "mastercard", got, "detectScheme(220500) = %q, expected neither mir nor mastercard", got)
}

func TestDetectScheme_RuPay(t *testing.T) {
	t.Parallel()
	cases := []string{"600000", "652100", "652200"}
	for _, bin := range cases {
		assert.Equal(t, "rupay", detectScheme(bin), "detectScheme(%s)", bin)
	}
}

func TestDetectScheme_Unknown(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", detectScheme("999999"), "detectScheme(999999)")
}

func TestDetectScheme_TooShort(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", detectScheme("42"), "detectScheme(42)")
}

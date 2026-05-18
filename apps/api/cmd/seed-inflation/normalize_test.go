package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalise_ValidCountryCode(t *testing.T) {
	t.Parallel()

	r := RawInflationRecord{
		CountryCode: " us ",
		CountryName: " United States ",
		Year:        2020,
		Rate:        2.123456789,
	}

	got := normalise(r)

	assert.Equal(t, "US", got.CountryCode)
	assert.Equal(t, "United States", got.CountryName)
	// Rate is rounded to 4 decimal places.
	assert.Equal(t, 2.1235, got.Rate)
}

func TestNormalise_RegionalAggregateCleared(t *testing.T) {
	t.Parallel()

	// World Bank uses codes like "EAP" (East Asia Pacific), "EMU", "1A" for
	// regional aggregates. The normalise function clears codes whose length is
	// not exactly 2 characters. Note: 2-character codes like "1W" are NOT
	// cleared (only codes with length ≠ 2 are discarded).
	tests := []struct {
		code    string
		cleared bool
	}{
		{"EAP", true}, // 3 chars
		{"EMU", true}, // 3 chars
		{"", true},    // empty
		{"1W", false}, // 2 chars — retained (numeric chars allowed)
	}

	for _, tt := range tests {
		t.Run(tt.code, func(t *testing.T) {
			t.Parallel()
			r := RawInflationRecord{CountryCode: tt.code, CountryName: "Region"}
			got := normalise(r)
			if tt.cleared {
				assert.Equal(t, "", got.CountryCode)
			} else {
				assert.NotEmpty(t, got.CountryCode)
			}
		})
	}
}

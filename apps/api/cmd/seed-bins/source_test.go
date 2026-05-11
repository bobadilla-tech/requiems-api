package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIsValidBINPrefix(t *testing.T) {
	t.Parallel()

	tests := []struct {
		bin  string
		want bool
	}{
		{"411111", true},      // 6-digit valid
		{"41111111", true},    // 8-digit valid
		{"4111", false},       // too short
		{"4111111111", false}, // 10 digits — neither 6 nor 8
		{"41111A", false},     // non-numeric
		{"", false},           // empty
		{"ABCDEF", false},     // all letters
	}

	for _, tt := range tests {
		t.Run(tt.bin, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, isValidBINPrefix(tt.bin))
		})
	}
}

func TestParseIannuttall(t *testing.T) {
	t.Parallel()

	// Columns: bin,brand,type,category,issuer,alpha_2,alpha_3,country,latitude,longitude,bank_phone,bank_url
	csv := "bin,brand,type,category,issuer,alpha_2,alpha_3,country,latitude,longitude,bank_phone,bank_url\n" +
		"411111,VISA,Credit,Classic,Some Bank,US,USA,United States,37.09,-95.71,+1-800-123-4567,https://example.com\n" +
		"INVALID,VISA,Credit,Classic,Skip Me,US,USA,United States,0,0,,\n"

	records, err := parseIannuttall(strings.NewReader(csv), "test", 0.75)
	require.NoError(t, err)
	require.Len(t, records, 1)

	r := records[0]
	assert.Equal(t, "411111", r.BINPrefix)
	assert.Equal(t, "VISA", r.Scheme)
	assert.Equal(t, "US", r.CountryCode)
	assert.Equal(t, "test", r.Source)
	assert.Equal(t, 0.75, r.Confidence)
}

func TestParseVenelinkochev(t *testing.T) {
	t.Parallel()

	// Columns: BIN,Brand,Type,Category,Issuer,IssuerPhone,IssuerUrl,isoCode2,isoCode3,CountryName
	csv := "BIN,Brand,Type,Category,Issuer,IssuerPhone,IssuerUrl,isoCode2,isoCode3,CountryName\n" +
		"411111,Visa,Credit,Classic,My Bank,+1555000,https://mybank.com,US,USA,United States\n" +
		"BADINP,Visa,Credit,Classic,Skip,,,US,USA,United States\n"

	records, err := parseVenelinkochev(strings.NewReader(csv), "venelinkochev", 0.80)
	require.NoError(t, err)
	require.Len(t, records, 1)

	r := records[0]
	assert.Equal(t, "411111", r.BINPrefix)
	assert.Equal(t, "My Bank", r.IssuerName)
	assert.Equal(t, "United States", r.CountryName)
}

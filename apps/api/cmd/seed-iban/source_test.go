package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseOptInt(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  int
	}{
		{"", 0},
		{"N/A", 0},
		{"-", 0},
		{"0", 0},
		{"4", 4},
		{"32767", 32767},   // max int16
		{"-32768", -32768}, // min int16
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tt.want, parseOptInt(tt.input))
		})
	}
}

func TestRawIBANCountry_BankAndAccountOffsets(t *testing.T) {
	t.Parallel()

	// Germany: BBAN is 18 chars (bank=0..7, branch=8..12, account=13..17)
	de := RawIBANCountry{
		CountryCode:   "DE",
		IBANLength:    22,
		BBANLength:    18,
		BankIDStart:   0,
		BankIDEnd:     7,
		BranchIDStart: 8,
		BranchIDEnd:   12,
	}

	assert.Equal(t, 0, de.BankOffset())
	assert.Equal(t, 8, de.BankLength()) // 7-0+1
	assert.Equal(t, 13, de.AccountOffset()) // BranchIDEnd(12)+1
	assert.Equal(t, 5, de.AccountLength())  // BBANLength(18) - 1 - AccountOffset(13) + 1
}

func TestRawIBANCountry_NoBranchCode(t *testing.T) {
	t.Parallel()

	// A country with no branch code: account starts right after the bank code.
	r := RawIBANCountry{
		BBANLength:    14,
		BankIDStart:   0,
		BankIDEnd:     3,
		BranchIDStart: 0,
		BranchIDEnd:   0, // no distinct branch (start == end, not > BankIDEnd)
	}

	assert.Equal(t, 4, r.AccountOffset())  // BankIDEnd(3)+1
	assert.Equal(t, 10, r.AccountLength()) // BBANLength(14) - 1 - AccountOffset(4) + 1
}

func TestFetchAndParse_LocalMockServer(t *testing.T) {
	t.Parallel()

	// registry.txt format: pipe-separated, first line is header.
	// Columns (0-indexed): 0=code, 1=name, 4=bban_format, 6=bban_len, 10=iban_len,
	//   11=bank_start, 12=bank_end, 13=branch_start, 14=branch_end, 16=sepa
	registry := "code|name|x|x|bban_format|x|bban_len|x|x|x|iban_len|bank_start|bank_end|branch_start|branch_end|x|sepa\n" +
		"DE|Germany|x|x|8!n10!n|x|18|x|x|x|22|0|7|8|12|x|1\n" +
		"US|United States|x|x|9!n|x|9|x|x|x|13|0|3|N/A|N/A|x|0\n"

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(registry))
	}))
	defer srv.Close()

	countries, err := fetchAndParse(srv.URL)
	require.NoError(t, err)
	require.Len(t, countries, 2)

	de := countries[0]
	assert.Equal(t, "DE", de.CountryCode)
	assert.Equal(t, 22, de.IBANLength)
	assert.Equal(t, 18, de.BBANLength)
	assert.True(t, de.SEPAMember, "expected DE to be a SEPA member")

	us := countries[1]
	assert.False(t, us.SEPAMember, "expected US not to be a SEPA member")
}

func TestFetchAndParse_NoValidRows(t *testing.T) {
	t.Parallel()

	// Valid header but all data rows are too short — expect an error.
	registry := "code|name|x|x|bban_format|x|bban_len|x|x|x|iban_len|bank_start|bank_end|branch_start|branch_end|x|sepa\n" +
		"DE|Germany\n" // too few fields

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(registry))
	}))
	defer srv.Close()

	_, err := fetchAndParse(srv.URL)
	require.Error(t, err)
}

func TestFetchAndParse_HTTPError(t *testing.T) {
	t.Parallel()

	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	_, err := fetchAndParse(srv.URL)
	require.Error(t, err)
}

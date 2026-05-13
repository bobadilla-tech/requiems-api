package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNormaliseScheme(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"VISA", "visa"},
		{"visa", "visa"},
		{"MASTERCARD", "mastercard"},
		{"Master Card", "mastercard"},
		{"AMEX", "amex"},
		{"American Express", "amex"},
		{"DISCOVER", "discover"},
		{"JCB", "jcb"},
		{"DINERS CLUB", "diners"},
		{"UNIONPAY", "unionpay"},
		{"Union Pay", "unionpay"},
		{"MAESTRO", "maestro"},
		{"MIR", "mir"},
		{"RUPAY", "rupay"},
		{"PRIVATE LABEL", "private_label"},
		{"unknown", "unknown"},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, normaliseScheme(tt.input))
		})
	}
}

func TestNormaliseCardType(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"CREDIT", "credit"},
		{"Credit", "credit"},
		{"DEBIT", "debit"},
		{"PREPAID", "prepaid"},
		{"Pre-Paid", "prepaid"},
		{"CHARGE", "charge"},
		{"Charge Card", "charge"},
		{"other", "other"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, normaliseCardType(tt.input))
		})
	}
}

func TestNormaliseCardLevel(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input string
		want  string
	}{
		{"CLASSIC", "classic"},
		{"Normal", "classic"},
		{"GOLD", "gold"},
		{"Premier", "gold"},
		{"PLATINUM", "platinum"},
		{"World", "platinum"},
		{"INFINITE", "infinite"},
		{"BLACK", "infinite"},
		{"BUSINESS", "business"},
		{"CORPORATE", "corporate"},
		{"SIGNATURE", "signature"},
		{"STANDARD", "standard"},
		{"unknown", "unknown"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, normaliseCardLevel(tt.input))
		})
	}
}

func TestDetectScheme(t *testing.T) {
	t.Parallel()

	tests := []struct {
		bin  string
		want string
	}{
		{"411111", "visa"}, // Visa starts with 4
		{"412345", "visa"},
		{"510000", "mastercard"}, // Mastercard 51-55
		{"555555", "mastercard"},
		{"222100", "mastercard"}, // Mastercard 2-series
		{"340000", "amex"},       // Amex 34
		{"370000", "amex"},       // Amex 37
		{"352800", "jcb"},        // JCB 3528-3589
		{"300000", "diners"},     // Diners 3000-3059
		{"360000", "diners"},     // Diners 36
		{"601100", "discover"},   // 6011 prefix → Discover (takes precedence over RuPay 60)
		{"600000", "rupay"},      // RuPay 60 (not 6011)
		{"621000", "unionpay"},   // UnionPay 62
		{"220000", "mir"},        // Mir 2200-2204
		{"63", ""},               // too short
	}

	for _, tt := range tests {
		t.Run(tt.bin, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, detectScheme(tt.bin))
		})
	}
}

func TestNormalise_PrepaidSetsFlag(t *testing.T) {
	t.Parallel()

	r := RawBINRecord{
		BINPrefix:   "411111",
		Scheme:      "VISA",
		CardType:    "PREPAID",
		CountryCode: "US",
		Confidence:  0.75,
	}

	got := normalise(r)
	require.True(t, got.Prepaid, "expected Prepaid = true for card type PREPAID")
	require.Equal(t, "visa", got.Scheme)
}

func TestNormalise_InvalidCountryCodeCleared(t *testing.T) {
	t.Parallel()

	r := RawBINRecord{
		BINPrefix:   "411111",
		CountryCode: "USA", // 3 letters — not valid 2-letter ISO
		CountryName: "United States",
		Confidence:  0.5,
	}

	got := normalise(r)
	require.Equal(t, "", got.CountryCode)
	require.Equal(t, "", got.CountryName)
}

func TestMergeRecords_HigherConfidenceWins(t *testing.T) {
	t.Parallel()

	records := []RawBINRecord{
		{BINPrefix: "411111", IssuerName: "Low Conf Bank", Confidence: 0.5, Source: "src-a"},
		{BINPrefix: "411111", IssuerName: "High Conf Bank", Confidence: 0.9, Source: "src-b"},
	}

	merged := mergeRecords(records)
	r, ok := merged["411111"]
	require.True(t, ok, "expected merged record for BIN 411111")
	require.Equal(t, "High Conf Bank", r.IssuerName)
}

func TestMergeRecords_FillsEmptyFieldsFromLoser(t *testing.T) {
	t.Parallel()

	records := []RawBINRecord{
		// winner by confidence but missing IssuerURL
		{BINPrefix: "411111", IssuerName: "Winner Bank", IssuerURL: "", Confidence: 0.9, Source: "src-a"},
		// loser has IssuerURL
		{BINPrefix: "411111", IssuerName: "", IssuerURL: "https://loser.com", Confidence: 0.5, Source: "src-b"},
	}

	merged := mergeRecords(records)
	r := merged["411111"]
	require.Equal(t, "https://loser.com", r.IssuerURL)
}

func TestCombineSources(t *testing.T) {
	t.Parallel()

	tests := []struct {
		a, b string
		want string
	}{
		{"src-a", "src-a", "src-a"},             // same → no duplicate
		{"src-a", "src-b", "src-a,src-b"},       // two distinct sources
		{"src-a,src-b", "src-b", "src-a,src-b"}, // already contains src-b
	}

	for _, tt := range tests {
		t.Run(tt.a+"/"+tt.b, func(t *testing.T) {
			t.Parallel()
			require.Equal(t, tt.want, combineSources(tt.a, tt.b))
		})
	}
}

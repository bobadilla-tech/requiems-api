package iban

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// mockRow implementa pgx.Row
type mockRow struct {
	row countryRow
	err error
}

func (r *mockRow) Scan(dest ...any) error {
	// If an error occurs, return the error
	if r.err != nil {
		return r.err
	}

	// Fill the destinations in the same order as the SELECT
	*dest[0].(*string) = r.row.name
	*dest[1].(*int16) = r.row.ibanLength
	*dest[2].(*int16) = r.row.bankOffset
	*dest[3].(*int16) = r.row.bankLength
	*dest[4].(*int16) = r.row.accountOffset
	*dest[5].(*int16) = r.row.accountLength

	return nil
}

// mockQuerier implements querier
type mockQuerier struct {
	rows map[string]countryRow // country code -> data
	err  error                 // error to return
}

func (m *mockQuerier) QueryRow(ctx context.Context, sql string, args ...any) pgx.Row {
	if m.err != nil {
		return &mockRow{err: m.err}
	}

	code := args[0].(string)
	row, ok := m.rows[code]

	if !ok {
		return &mockRow{err: pgx.ErrNoRows}
	}

	return &mockRow{row: row}
}

// ---- normalizeIBAN ----

func TestNormalizeIBAN_StripsSpaces(t *testing.T) {
	t.Parallel()
	got := normalizeIBAN("DE89 3704 0044 0532 0130 00")
	assert.Equal(t, "DE89370400440532013000", got)
}

func TestNormalizeIBAN_Uppercases(t *testing.T) {
	t.Parallel()
	got := normalizeIBAN("de89370400440532013000")
	assert.Equal(t, "DE89370400440532013000", got)
}

func TestNormalizeIBAN_TrimsTrimSpace(t *testing.T) {
	t.Parallel()
	got := normalizeIBAN("  DE89370400440532013000  ")
	assert.Equal(t, "DE89370400440532013000", got)
}

// ---- basicFormatOK ----

func TestBasicFormatOK_ValidDE(t *testing.T) {
	t.Parallel()
	assert.True(t, basicFormatOK("DE89370400440532013000"), "expected true for valid German IBAN")
}

func TestBasicFormatOK_ValidGB(t *testing.T) {
	t.Parallel()
	assert.True(t, basicFormatOK("GB82WEST12345698765432"), "expected true for valid UK IBAN")
}

func TestBasicFormatOK_TooShort(t *testing.T) {
	t.Parallel()
	assert.False(t, basicFormatOK("DE89"), "expected false for 4-char input")
}

func TestBasicFormatOK_Empty(t *testing.T) {
	t.Parallel()
	assert.False(t, basicFormatOK(""), "expected false for empty string")
}

func TestBasicFormatOK_DigitInCountryCode(t *testing.T) {
	t.Parallel()
	assert.False(t, basicFormatOK("1E89370400440532013000"), "expected false when first char is a digit")
}

func TestBasicFormatOK_LetterInCheckDigits(t *testing.T) {
	t.Parallel()
	assert.False(t, basicFormatOK("DEAB370400440532013000"), "expected false when check digits contain letters")
}

func TestBasicFormatOK_SpecialCharacter(t *testing.T) {
	t.Parallel()
	assert.False(t, basicFormatOK("DE89!70400440532013000"), "expected false for input containing '!'")
}

// ---- validateChecksum ----

var validIBANs = []struct {
	iban    string
	country string
}{
	{"DE89370400440532013000", "Germany"},
	{"GB82WEST12345698765432", "United Kingdom"},
	{"FR7630006000011234567890189", "France"},
	{"NL91ABNA0417164300", "Netherlands"},
	{"CH9300762011623852957", "Switzerland"},
	{"AT611904300234573201", "Austria"},
	{"BE68539007547034", "Belgium"},
	{"PL61109010140000071219812874", "Poland"},
}

func TestValidateChecksum_KnownValidIBANs(t *testing.T) {
	t.Parallel()
	for _, tc := range validIBANs {
		assert.True(t, validateChecksum(tc.iban), "validateChecksum(%s) = false, expected true (%s)", tc.iban, tc.country)
	}
}

func TestValidateChecksum_WrongCheckDigits(t *testing.T) {
	t.Parallel()
	// DE89... with check digits changed to 00 — should fail.
	assert.False(t, validateChecksum("DE00370400440532013000"), "expected false for IBAN with wrong check digits (DE00...)")
}

func TestValidateChecksum_TransposedDigits(t *testing.T) {
	t.Parallel()
	// Swapping two adjacent digits in the BBAN breaks the checksum.
	assert.False(t, validateChecksum("DE89370400440352013000"), "expected false for IBAN with transposed digits in BBAN")
}

func TestValidateChecksum_AllZeroCheckDigits(t *testing.T) {
	t.Parallel()
	assert.False(t, validateChecksum("GB00WEST12345698765432"), "expected false for 00 check digits")
}

// ---- mod97 ----

func TestMod97_DEExample(t *testing.T) {
	t.Parallel()
	// DE89370400440532013000 rearranged + letters replaced:
	// BBAN+CC+CD → 370400440532013000DE89 → 370400440532013000131489
	assert.Equal(t, 1, mod97("370400440532013000131489"))
}

func TestMod97_NLExample(t *testing.T) {
	t.Parallel()
	// NL91ABNA0417164300 rearranged → ABNA0417164300NL91
	// A=10 B=11 N=23 A=10 | 0417164300 | N=23 L=21 | 91
	// → "101123100417164300232191"
	assert.Equal(t, 1, mod97("101123100417164300232191"))
}

// ---- extract ----

func TestExtract_HappyPath(t *testing.T) {
	t.Parallel()
	got := extract("ABCDEFGH", 2, 3)
	assert.Equal(t, "CDE", got)
}

func TestExtract_FromStart(t *testing.T) {
	t.Parallel()
	got := extract("37040044ABCDE", 0, 8)
	assert.Equal(t, "37040044", got)
}

func TestExtract_ZeroLength(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", extract("ABCDEF", 0, 0), "extract with 0 length should be empty")
}

func TestExtract_OutOfBounds(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", extract("ABC", 1, 10), "out-of-bounds extract should be empty")
}

func TestExtract_NegativeOffset(t *testing.T) {
	t.Parallel()
	assert.Equal(t, "", extract("ABC", -1, 2), "negative offset extract should be empty")
}

// --- Parse ---

func TestService_Parse(t *testing.T) {
	t.Parallel()
	svc := &Service{db: &mockQuerier{rows: map[string]countryRow{
		"GB": {name: "United Kingdom", ibanLength: 22, bankOffset: 0, bankLength: 4, accountOffset: 4, accountLength: 14},
		"DE": {name: "Germany", ibanLength: 22, bankOffset: 0, bankLength: 8, accountOffset: 8, accountLength: 10},
		"ES": {name: "Spain", ibanLength: 24, bankOffset: 0, bankLength: 4, accountOffset: 4, accountLength: 10},
	}}}

	resp, err := svc.Parse(context.Background(), "GB29NWBK60161331926819")

	require.NoError(t, err)
	assert.True(t, resp.Valid)
	assert.Equal(t, "United Kingdom", resp.Country)
	assert.Equal(t, "NWBK", resp.BankCode)
}

// ---- ParseBatch ----

func TestService_ParseBatch_MixedResults(t *testing.T) {
	t.Parallel()
	svc := &Service{db: &mockQuerier{
		rows: map[string]countryRow{
			"GB": {name: "United Kingdom", ibanLength: 22, bankOffset: 0, bankLength: 4, accountOffset: 4, accountLength: 14},
			"DE": {name: "Germany", ibanLength: 22, bankOffset: 0, bankLength: 8, accountOffset: 8, accountLength: 10},
		},
	}}

	numbers := []string{
		"GB29NWBK60161331926819", // valid
		"DE89370400440532013000", // valid
		"XX89370400440532013000", // invalid — unknown country
	}

	resp, err := svc.ParseBatch(context.Background(), numbers)

	require.NoError(t, err)
	assert.Equal(t, 3, len(resp))
	assert.True(t, resp[0].Valid, "expected result[0] valid=true")
	assert.True(t, resp[1].Valid, "expected result[1] valid=true")
	assert.False(t, resp[2].Valid, "expected result[2] valid=false")
}

func TestService_ParseBatch_DBError(t *testing.T) {
	t.Parallel()
	svc := &Service{db: &mockQuerier{err: errors.New("database unreachable")}}

	numbers := []string{"GB29NWBK60161331926819"}

	resp, err := svc.ParseBatch(context.Background(), numbers)

	require.Error(t, err)
	assert.Equal(t, 0, len(resp))
}

func TestService_ParseBatch_OrderPreserved(t *testing.T) {
	t.Parallel()
	svc := &Service{db: &mockQuerier{rows: map[string]countryRow{
		"GB": {name: "United Kingdom", ibanLength: 22},
		"DE": {name: "Germany", ibanLength: 22},
		"ES": {name: "Spain", ibanLength: 24},
	}}}

	numbers := []string{
		"GB29NWBK60161331926819",
		"DE89370400440532013000",
		"ES9121000418450200051332",
	}

	resp, err := svc.ParseBatch(context.Background(), numbers)

	require.NoError(t, err)
	assert.Equal(t, "GB29NWBK60161331926819", resp[0].IBAN)
	assert.Equal(t, "DE89370400440532013000", resp[1].IBAN)
	assert.Equal(t, "ES9121000418450200051332", resp[2].IBAN)
}

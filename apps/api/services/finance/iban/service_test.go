package iban

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
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
	got := normalizeIBAN("DE89 3704 0044 0532 0130 00")
	want := "DE89370400440532013000"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNormalizeIBAN_Uppercases(t *testing.T) {
	got := normalizeIBAN("de89370400440532013000")
	want := "DE89370400440532013000"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestNormalizeIBAN_TrimsTrimSpace(t *testing.T) {
	got := normalizeIBAN("  DE89370400440532013000  ")
	want := "DE89370400440532013000"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

// ---- basicFormatOK ----

func TestBasicFormatOK_ValidDE(t *testing.T) {
	if !basicFormatOK("DE89370400440532013000") {
		t.Error("expected true for valid German IBAN")
	}
}

func TestBasicFormatOK_ValidGB(t *testing.T) {
	if !basicFormatOK("GB82WEST12345698765432") {
		t.Error("expected true for valid UK IBAN")
	}
}

func TestBasicFormatOK_TooShort(t *testing.T) {
	if basicFormatOK("DE89") {
		t.Error("expected false for 4-char input")
	}
}

func TestBasicFormatOK_Empty(t *testing.T) {
	if basicFormatOK("") {
		t.Error("expected false for empty string")
	}
}

func TestBasicFormatOK_DigitInCountryCode(t *testing.T) {
	if basicFormatOK("1E89370400440532013000") {
		t.Error("expected false when first char is a digit")
	}
}

func TestBasicFormatOK_LetterInCheckDigits(t *testing.T) {
	if basicFormatOK("DEAB370400440532013000") {
		t.Error("expected false when check digits contain letters")
	}
}

func TestBasicFormatOK_SpecialCharacter(t *testing.T) {
	if basicFormatOK("DE89!70400440532013000") {
		t.Error("expected false for input containing '!'")
	}
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
	for _, tc := range validIBANs {
		if !validateChecksum(tc.iban) {
			t.Errorf("validateChecksum(%s) = false, expected true (%s)", tc.iban, tc.country)
		}
	}
}

func TestValidateChecksum_WrongCheckDigits(t *testing.T) {
	// DE89... with check digits changed to 00 — should fail.
	if validateChecksum("DE00370400440532013000") {
		t.Error("expected false for IBAN with wrong check digits (DE00...)")
	}
}

func TestValidateChecksum_TransposedDigits(t *testing.T) {
	// Swapping two adjacent digits in the BBAN breaks the checksum.
	if validateChecksum("DE89370400440352013000") {
		t.Error("expected false for IBAN with transposed digits in BBAN")
	}
}

func TestValidateChecksum_AllZeroCheckDigits(t *testing.T) {
	if validateChecksum("GB00WEST12345698765432") {
		t.Error("expected false for 00 check digits")
	}
}

// ---- mod97 ----

func TestMod97_DEExample(t *testing.T) {
	// DE89370400440532013000 rearranged + letters replaced:
	// BBAN+CC+CD → 370400440532013000DE89 → 370400440532013000131489
	if got := mod97("370400440532013000131489"); got != 1 {
		t.Errorf("mod97 = %d, want 1", got)
	}
}

func TestMod97_NLExample(t *testing.T) {
	// NL91ABNA0417164300 rearranged → ABNA0417164300NL91
	// A=10 B=11 N=23 A=10 | 0417164300 | N=23 L=21 | 91
	// → "101123100417164300232191"
	if got := mod97("101123100417164300232191"); got != 1 {
		t.Errorf("mod97 = %d, want 1", got)
	}
}

// ---- extract ----

func TestExtract_HappyPath(t *testing.T) {
	got := extract("ABCDEFGH", 2, 3)
	if got != "CDE" {
		t.Errorf("extract = %q, want %q", got, "CDE")
	}
}

func TestExtract_FromStart(t *testing.T) {
	got := extract("37040044ABCDE", 0, 8)
	if got != "37040044" {
		t.Errorf("extract = %q, want %q", got, "37040044")
	}
}

func TestExtract_ZeroLength(t *testing.T) {
	if got := extract("ABCDEF", 0, 0); got != "" {
		t.Errorf("extract with 0 length should be empty, got %q", got)
	}
}

func TestExtract_OutOfBounds(t *testing.T) {
	if got := extract("ABC", 1, 10); got != "" {
		t.Errorf("out-of-bounds extract should be empty, got %q", got)
	}
}

func TestExtract_NegativeOffset(t *testing.T) {
	if got := extract("ABC", -1, 2); got != "" {
		t.Errorf("negative offset extract should be empty, got %q", got)
	}
}

// --- Parse ---

func TestService_Parse(t *testing.T) {
	svc := &Service{db: &mockQuerier{rows: map[string]countryRow{
		"GB": {name: "United Kingdom", ibanLength: 22, bankOffset: 0, bankLength: 4, accountOffset: 4, accountLength: 14},
		"DE": {name: "Germany", ibanLength: 22, bankOffset: 0, bankLength: 8, accountOffset: 8, accountLength: 10},
		"ES": {name: "Spain", ibanLength: 24, bankOffset: 0, bankLength: 4, accountOffset: 4, accountLength: 10},
	}}}

	resp, err := svc.Parse(context.Background(), "GB29NWBK60161331926819")

	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if !resp.Valid {
		t.Errorf("Expected valid=true")
	}

	if resp.Country != "United Kingdom" {
		t.Errorf("expected country United Kingdom, got %q", resp.Country)
	}

	if resp.BankCode != "NWBK" {
		t.Errorf("expected bank_code NWBK, got %q", resp.BankCode)
	}
}

// ---- ParseBatch ----

func TestService_ParseBatch_MixedResults(t *testing.T) {
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

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if resp.Total != 3 {
		t.Errorf("expected total=3, got %d", resp.Total)
	}
	if !resp.Results[0].Valid {
		t.Error("expected result[0] valid=true")
	}
	if !resp.Results[1].Valid {
		t.Error("expected result[1] valid=true")
	}
	if resp.Results[2].Valid {
		t.Error("expected result[2] valid=false")
	}

}

func TestService_ParseBatch_DBError(t *testing.T) {
	svc := &Service{db: &mockQuerier{err: errors.New("database unreachable")}}

	numbers := []string{"GB29NWBK60161331926819"}

	resp, err := svc.ParseBatch(context.Background(), numbers)

	if err == nil {
		t.Fatalf("Expected error, got nil")
	}

	if resp.Total != 0 {
		t.Errorf("Expected empty response, got %+v", resp)
	}
}

func TestService_ParseBatch_OrderPreserved(t *testing.T) {
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

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.Results[0].IBAN != "GB29NWBK60161331926819" {
		t.Errorf("Expected GB29N... at position 0, got %s", resp.Results[0].IBAN)
	}

	if resp.Results[1].IBAN != "DE89370400440532013000" {
		t.Errorf("Expected DE893... at position 1, got %s", resp.Results[1].IBAN)
	}

	if resp.Results[2].IBAN != "ES9121000418450200051332" {
		t.Errorf("Expected ES912... at position 2, got %s", resp.Results[2].IBAN)
	}
}

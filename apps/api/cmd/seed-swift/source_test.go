package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFetchAndParse_LocalFile(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	filePath := filepath.Join(dir, "swift.csv")

	csv := "swift_code,bank_name,city,country_name\nDEUTDEDB,Deutsche Bank,Frankfurt,Germany\nCHASUS33XXX,JPMorgan Chase,New York,United States\n"
	err := os.WriteFile(filePath, []byte(csv), 0o600)
	require.NoError(t, err)

	records, err := fetchAndParse(filePath)
	require.NoError(t, err)
	require.Len(t, records, 2)

	assert.Equal(t, "DEUTDEDBXXX", records[0].SwiftCode)
	assert.Equal(t, "XXX", records[1].BranchCode)
}

func TestFetchAndParse_NoValidRecords(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	filePath := filepath.Join(dir, "swift-empty.csv")

	csv := "swift_code,bank_name,city,country_name\nINVALID,No Bank,Nowhere,No Country\n"
	err := os.WriteFile(filePath, []byte(csv), 0o600)
	require.NoError(t, err)

	_, err = fetchAndParse(filePath)
	require.Error(t, err)
}

func TestParseRow_Accepts8And11CharCodes(t *testing.T) {
	t.Parallel()

	cols := colIndices{swift: 0, bankName: 1, city: 2, countryName: 3}

	r8, ok := parseRow([]string{"DEUTDEDB", "Deutsche Bank", "Frankfurt", "Germany"}, cols)
	require.True(t, ok, "expected 8-char SWIFT row to parse")
	assert.Equal(t, "DEUTDEDBXXX", r8.SwiftCode)

	r11, ok := parseRow([]string{"CHASUS33XXX", "JPMorgan Chase", "New York", "United States"}, cols)
	require.True(t, ok, "expected 11-char SWIFT row to parse")
	assert.Equal(t, "CHASUS33XXX", r11.SwiftCode)
}

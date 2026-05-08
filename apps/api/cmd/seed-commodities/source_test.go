package main

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseFREDRow(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		record  []string
		wantVal float64
		wantOK  bool
	}{
		{name: "valid daily row", record: []string{"2020-01-15", "55.42"}, wantVal: 55.42, wantOK: true},
		{name: "valid annual row", record: []string{"2010-01-01", "80.00"}, wantVal: 80.00, wantOK: true},
		{name: "dot placeholder", record: []string{"2020-06-01", "."}, wantOK: false},
		{name: "empty value", record: []string{"2020-06-01", ""}, wantOK: false},
		{name: "negative value", record: []string{"2020-06-01", "-5.0"}, wantOK: false},
		{name: "non-numeric value", record: []string{"2020-06-01", "N/A"}, wantOK: false},
		{name: "too few fields", record: []string{"2020-06-01"}, wantOK: false},
		// Years outside [1960, now] are rejected.
		{name: "year too early", record: []string{"1959-01-01", "10.0"}, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			val, _, ok := parseFREDRow(tt.record)
			require.Equal(t, tt.wantOK, ok)
			if ok {
				assert.Equal(t, tt.wantVal, val)
			}
		})
	}
}

func TestParseYahooClose(t *testing.T) {
	t.Parallel()

	closes := []interface{}{
		float64(1234.56),
		nil,
		json.Number("99.99"),
		float64(0),    // zero — rejected
		float64(-1.0), // negative — rejected
	}

	tests := []struct {
		name    string
		idx     int
		wantVal float64
		wantOK  bool
	}{
		{name: "float64 value", idx: 0, wantVal: 1234.56, wantOK: true},
		{name: "nil entry", idx: 1, wantOK: false},
		{name: "json.Number", idx: 2, wantVal: 99.99, wantOK: true},
		{name: "zero", idx: 3, wantOK: false},
		{name: "negative", idx: 4, wantOK: false},
		{name: "out of bounds", idx: 99, wantOK: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			val, ok := parseYahooClose(closes, tt.idx)
			require.Equal(t, tt.wantOK, ok)
			if ok {
				assert.Equal(t, tt.wantVal, val)
			}
		})
	}
}

func TestParseYear(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input   string
		want    int
		wantErr bool
	}{
		{input: "2020-01-15", want: 2020},
		{input: "1985", want: 1985},
		{input: "198", wantErr: true}, // too short
		{input: "abcd", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			t.Parallel()
			got, err := parseYear(tt.input)
			if tt.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestBuildRecords(t *testing.T) {
	t.Parallel()

	cfg := CommodityConfig{
		Slug:       "oil",
		Name:       "Crude Oil",
		Unit:       "barrel",
		Currency:   "USD",
		ConvFactor: 1.0,
	}

	byYear := map[int]*yearAcc{
		2020: {sum: 40.0, count: 2}, // avg = 20.0
	}

	records := buildRecords(cfg, byYear)
	require.Len(t, records, 1)

	r := records[0]
	assert.Equal(t, "oil", r.Slug)
	assert.Equal(t, 2020, r.Year)
	assert.Equal(t, 20.0, r.Price)
}

func TestAccumulate(t *testing.T) {
	t.Parallel()

	m := make(map[int]*yearAcc)
	accumulate(m, 2021, 10.0)
	accumulate(m, 2021, 20.0)
	accumulate(m, 2022, 5.0)

	assert.Equal(t, 2, m[2021].count)
	assert.Equal(t, 30.0, m[2021].sum)
	assert.Equal(t, 1, m[2022].count)
	assert.Equal(t, 5.0, m[2022].sum)
}

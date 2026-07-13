package units

import (
	"context"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_Convert(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name        string
		from        string
		to          string
		value       float64
		wantResult  float64
		wantFormula string
		wantErr     error
	}{
		// Length
		{
			name: "miles to km",
			from: "miles", to: "km", value: 10,
			wantResult:  16.09344,
			wantFormula: "miles × 1.609344",
		},
		{
			name: "km to miles",
			from: "km", to: "miles", value: 1,
			wantResult:  0.621371,
			wantFormula: "km × 0.621371",
		},
		{
			name: "m to cm",
			from: "m", to: "cm", value: 1,
			wantResult:  100,
			wantFormula: "m × 100",
		},
		{
			name: "ft to in",
			from: "ft", to: "in", value: 1,
			wantResult:  12,
			wantFormula: "ft × 12",
		},
		{
			name: "same unit (m to m)",
			from: "m", to: "m", value: 5,
			wantResult:  5,
			wantFormula: "m × 1",
		},
		// Weight
		{
			name: "kg to lb",
			from: "kg", to: "lb", value: 1,
			wantResult:  2.204624,
			wantFormula: "kg × 2.204624",
		},
		{
			name: "oz to g",
			from: "oz", to: "g", value: 1,
			wantResult:  28.3495,
			wantFormula: "oz × 28.3495",
		},
		// Volume
		{
			name: "l to ml",
			from: "l", to: "ml", value: 1,
			wantResult:  1000,
			wantFormula: "l × 1000",
		},
		{
			name: "gal to l",
			from: "gal", to: "l", value: 1,
			wantResult:  3.78541,
			wantFormula: "gal × 3.78541",
		},
		// Temperature
		{
			name: "celsius to fahrenheit",
			from: "c", to: "f", value: 100,
			wantResult:  212,
			wantFormula: "°C × 9/5 + 32",
		},
		{
			name: "fahrenheit to celsius",
			from: "f", to: "c", value: 32,
			wantResult:  0,
			wantFormula: "(°F − 32) × 5/9",
		},
		{
			name: "celsius to kelvin",
			from: "c", to: "k", value: 0,
			wantResult:  273.15,
			wantFormula: "°C + 273.15",
		},
		{
			name: "kelvin to celsius",
			from: "k", to: "c", value: 273.15,
			wantResult:  0,
			wantFormula: "K − 273.15",
		},
		// Area
		{
			name: "m2 to ft2",
			from: "m2", to: "ft2", value: 1,
			wantResult:  10.763915,
			wantFormula: "m2 × 10.763915",
		},
		// Speed
		{
			name: "mph to km_h",
			from: "mph", to: "km_h", value: 60,
			wantResult:  96.5604,
			wantFormula: "mph × 1.60934",
		},
		// Temperature — remaining conversions
		{
			name: "fahrenheit to kelvin",
			from: "f", to: "k", value: 32,
			wantResult:  273.15,
			wantFormula: "(°F − 32) × 5/9 + 273.15",
		},
		{
			name: "kelvin to fahrenheit",
			from: "k", to: "f", value: 373.15,
			wantResult:  212,
			wantFormula: "(K − 273.15) × 9/5 + 32",
		},
		{
			name: "same temperature unit (c to c)",
			from: "c", to: "c", value: 25,
			wantResult:  25,
			wantFormula: "c (no conversion needed)",
		},
		// Errors
		{
			name: "unknown from unit",
			from: "furlong", to: "km", value: 1,
			wantErr: ErrUnknownUnit,
		},
		{
			name: "unknown to unit",
			from: "km", to: "parsec", value: 1,
			wantErr: ErrUnknownUnit,
		},
		{
			name: "incompatible units",
			from: "miles", to: "kg", value: 1,
			wantErr: ErrIncompatibleUnits,
		},
		// Long-form aliases (demo + common API DX)
		{
			name: "meters alias to kilometers alias",
			from: "meters", to: "kilometers", value: 1000,
			wantResult:  1,
			wantFormula: "m × 0.001",
		},
		{
			name: "celsius alias to fahrenheit alias",
			from: "celsius", to: "fahrenheit", value: 0,
			wantResult:  32,
			wantFormula: "°C × 9/5 + 32",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := svc.Convert(tt.from, tt.to, tt.value)

			if tt.wantErr != nil {
				assert.ErrorIs(t, err, tt.wantErr)
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tt.wantResult, got.Result)

			assert.Equal(t, tt.wantFormula, got.Formula)
		})
	}
}

func TestAliases_ResolveToKnownUnits(t *testing.T) {
	t.Parallel()

	for alias, canonical := range aliases {
		if _, ok := units[canonical]; !ok {
			t.Errorf("alias %q maps to unknown unit %q", alias, canonical)
		}
	}
}

func TestService_Units(t *testing.T) {
	t.Parallel()
	svc := NewService()
	got := svc.Units()

	categories := map[string][]string{
		"length":      got.Length,
		"weight":      got.Weight,
		"volume":      got.Volume,
		"temperature": got.Temperature,
		"area":        got.Area,
		"speed":       got.Speed,
	}

	expectedMembers := map[string][]string{
		"length":      {"m", "km", "miles", "ft", "in", "cm", "mm", "yd", "nmi"},
		"weight":      {"g", "kg", "lb", "oz", "mg", "t", "stone"},
		"volume":      {"ml", "l", "gal", "cup", "pt", "qt", "tsp", "tbsp", "fl_oz"},
		"temperature": {"c", "f", "k"},
		"area":        {"m2", "km2", "ft2", "in2", "cm2", "mm2", "yd2", "acre", "ha"},
		"speed":       {"km_h", "mph", "knots", "m_s"},
	}

	for cat, members := range expectedMembers {
		got := categories[cat]
		assert.Len(t, got, len(members))
		for _, key := range members {
			if !slices.Contains(got, key) {
				t.Errorf("%s: missing unit %q", cat, key)
			}
		}
		// Verify sorted order.
		for i := 1; i < len(got); i++ {
			if got[i] < got[i-1] {
				t.Errorf("%s: not sorted at index %d: %q before %q", cat, i, got[i-1], got[i])
			}
		}
	}
}

func TestService_ConvertBatch(t *testing.T) {
	t.Parallel()
	svc := NewService()

	operations := []BatchItem{
		{
			From:  "miles",
			To:    "km",
			Value: new(float64(10)),
		},
		{
			From:  "kg",
			To:    "lb",
			Value: new(float64(5)),
		},
		{
			From:  "c",
			To:    "g",
			Value: new(float64(25)),
		},
	}

	results := svc.ConvertBatch(context.Background(), operations)

	require.Equal(t, 3, len(results))
	require.True(t, results[0].Success)
	require.True(t, results[1].Success)
	require.False(t, results[2].Success)
}

func TestService_ConvertBatch_PreservedOrder(t *testing.T) {
	t.Parallel()
	svc := NewService()

	operations := []BatchItem{
		{
			From:  "miles",
			To:    "km",
			Value: new(float64(10)),
		},
		{
			From:  "kg",
			To:    "lb",
			Value: new(float64(5)),
		},
		{
			From:  "c",
			To:    "g",
			Value: new(float64(25)),
		},
	}

	results := svc.ConvertBatch(context.Background(), operations)

	require.Equal(t, "miles", results[0].From)
	require.Equal(t, "km", results[0].To)

	require.Equal(t, "kg", results[1].From)
	require.Equal(t, "lb", results[1].To)

	require.Equal(t, "c", results[2].From)
	require.Equal(t, "g", results[2].To)
}

func TestService_ConvertBatch_NilValue(t *testing.T) {
	t.Parallel()
	svc := NewService()

	operations := []BatchItem{
		{
			From:  "miles",
			To:    "km",
			Value: nil, // malformed — missing value
		},
	}

	results := svc.ConvertBatch(context.Background(), operations)

	require.Len(t, results, 1)
	assert.False(t, results[0].Success)
	assert.Equal(t, "value is required", results[0].Error)
}

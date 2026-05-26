package numbase

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_ConvertBatch(t *testing.T) {
	t.Parallel()
	svc := NewService()

	items := []ConvertQuery{
		{From: 10, To: 16, Value: "255"},
		{From: 2, To: 10, Value: "11111111"},
		{From: 10, To: 16, Value: "xyz"},
	}
	results := svc.ConvertBatch(items)

	require.Len(t, results, 3)

	assert.Equal(t, "ff", results[0].Result)
	assert.Empty(t, results[0].Error)

	assert.Equal(t, "255", results[1].Result)
	assert.Empty(t, results[1].Error)

	assert.Empty(t, results[2].Result)
	assert.NotEmpty(t, results[2].Error)
}

func TestService_Convert(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name     string
		value    string
		fromBase int
		toBase   int
		want     string
		wantErr  error
	}{
		{
			name:     "decimal to hex",
			value:    "255",
			fromBase: 10,
			toBase:   16,
			want:     "ff",
		},
		{
			name:     "hex to decimal",
			value:    "ff",
			fromBase: 16,
			toBase:   10,
			want:     "255",
		},
		{
			name:     "hex with 0x prefix to decimal",
			value:    "0xff",
			fromBase: 16,
			toBase:   10,
			want:     "255",
		},
		{
			name:     "decimal to binary",
			value:    "255",
			fromBase: 10,
			toBase:   2,
			want:     "11111111",
		},
		{
			name:     "binary to decimal",
			value:    "11111111",
			fromBase: 2,
			toBase:   10,
			want:     "255",
		},
		{
			name:     "binary with 0b prefix to decimal",
			value:    "0b11111111",
			fromBase: 2,
			toBase:   10,
			want:     "255",
		},
		{
			name:     "decimal to octal",
			value:    "255",
			fromBase: 10,
			toBase:   8,
			want:     "377",
		},
		{
			name:     "octal to decimal",
			value:    "377",
			fromBase: 8,
			toBase:   10,
			want:     "255",
		},
		{
			name:     "octal with 0o prefix to decimal",
			value:    "0o377",
			fromBase: 8,
			toBase:   10,
			want:     "255",
		},
		{
			name:     "hex to binary",
			value:    "ff",
			fromBase: 16,
			toBase:   2,
			want:     "11111111",
		},
		{
			name:     "zero",
			value:    "0",
			fromBase: 10,
			toBase:   16,
			want:     "0",
		},
		{
			name:     "negative decimal to hex",
			value:    "-255",
			fromBase: 10,
			toBase:   16,
			want:     "-ff",
		},
		{
			name:     "same base",
			value:    "42",
			fromBase: 10,
			toBase:   10,
			want:     "42",
		},
		{
			name:     "invalid from base",
			value:    "255",
			fromBase: 3,
			toBase:   10,
			wantErr:  ErrInvalidBase,
		},
		{
			name:     "invalid to base",
			value:    "255",
			fromBase: 10,
			toBase:   5,
			wantErr:  ErrInvalidBase,
		},
		{
			name:     "invalid value for base",
			value:    "xyz",
			fromBase: 10,
			toBase:   16,
			wantErr:  ErrInvalidValue,
		},
		{
			name:     "binary value with decimal digits",
			value:    "29",
			fromBase: 2,
			toBase:   10,
			wantErr:  ErrInvalidValue,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := svc.Convert(tt.value, tt.fromBase, tt.toBase)

			if tt.wantErr != nil {
				require.Error(t, err)
				assert.True(t, errors.Is(err, tt.wantErr), "expected error %v, got %v", tt.wantErr, err)
				return
			}

			require.NoError(t, err)

			assert.Equal(t, tt.want, got.Result)
			assert.Equal(t, tt.value, got.Input)
			assert.Equal(t, tt.fromBase, got.From)
			assert.Equal(t, tt.toBase, got.To)
		})
	}
}

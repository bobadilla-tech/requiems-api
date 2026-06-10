package base64 //nolint:revive // package name matches the service it tests; renaming would obscure intent

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
)

func TestService_Encode(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name    string
		value   string
		variant string
		want    string
	}{
		{
			name:  "standard encoding",
			value: "Hello, world!",
			want:  "SGVsbG8sIHdvcmxkIQ==",
		},
		{
			name:    "standard encoding explicit",
			value:   "Hello, world!",
			variant: "standard",
			want:    "SGVsbG8sIHdvcmxkIQ==",
		},
		{
			name:    "url-safe encoding",
			value:   "Hello, world!",
			variant: "url",
			want:    "SGVsbG8sIHdvcmxkIQ==",
		},
		{
			name:    "url-safe encoding with url-unsafe characters",
			value:   "\xfb\xff\xfe",
			variant: "url",
			want:    "-__-",
		},
		{
			name:    "standard encoding with url-unsafe characters",
			value:   "\xfb\xff\xfe",
			variant: "standard",
			want:    "+//+",
		},
		{
			name:  "empty string",
			value: "",
			want:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := svc.Encode(tt.value, tt.variant)
			assert.Equal(t, tt.want, got.Result)
		})
	}
}

func TestService_Decode(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name    string
		value   string
		variant string
		want    string
		wantErr bool
	}{
		{
			name:  "standard decoding",
			value: "SGVsbG8sIHdvcmxkIQ==",
			want:  "Hello, world!",
		},
		{
			name:    "standard decoding explicit",
			value:   "SGVsbG8sIHdvcmxkIQ==",
			variant: "standard",
			want:    "Hello, world!",
		},
		{
			name:    "url-safe decoding",
			value:   "SGVsbG8sIHdvcmxkIQ==",
			variant: "url",
			want:    "Hello, world!",
		},
		{
			name:    "url-safe encoding with url-unsafe characters",
			value:   "-__-",
			variant: "url",
			want:    "\xfb\xff\xfe",
		},
		{
			name:    "invalid standard base64",
			value:   "not-valid-base64!!!",
			wantErr: true,
		},
		{
			name:    "invalid url base64",
			value:   "not+valid+base64!!!",
			variant: "url",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got, err := svc.Decode(tt.value, tt.variant)

			if tt.wantErr {
				require.Error(t, err)
				var se *svcerr.Error
				require.ErrorAs(t, err, &se, "expected *svcerr.Error, got %T", err)
				assert.Equal(t, svcerr.KindUnknown, se.Kind)
				return
			}

			require.NoError(t, err)
			assert.Equal(t, tt.want, got.Result)
		})
	}
}

func TestService_EncodeBatch(t *testing.T) {
	t.Parallel()

	svc := NewService()

	tests := []struct {
		name    string
		values  []string
		variant string
		want    []Result
	}{
		{
			name:   "standard encoding",
			values: []string{"Hello", "World"},
			want:   []Result{{Result: "SGVsbG8="}, {Result: "V29ybGQ="}},
		},
		{
			name:    "url encoding",
			values:  []string{"\xfb\xff\xfe"},
			variant: "url",
			want:    []Result{{Result: "-__-"}},
		},
		{
			name:   "empty batch",
			values: []string{},
			want:   []Result{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := svc.EncodeBatch(tt.values, tt.variant)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestService_DecodeBatch(t *testing.T) {
	t.Parallel()

	svc := NewService()

	tests := []struct {
		name    string
		values  []string
		variant string
		want    []Result
	}{
		{
			name:   "standard decoding",
			values: []string{"SGVsbG8=", "V29ybGQ="},
			want:   []Result{{Result: "Hello"}, {Result: "World"}},
		},
		{
			name:    "url decoding",
			values:  []string{"-__-"},
			variant: "url",
			want:    []Result{{Result: "\xfb\xff\xfe"}},
		},
		{
			name: "partial failure",
			values: []string{
				"SGVsbG8=",
				"not-valid-base64!!!",
				"V29ybGQ=",
			},
			want: []Result{{Result: "Hello"}, {Result: ""}, {Result: "World"}},
		},
		{
			name:   "empty batch",
			values: []string{},
			want:   []Result{},
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := svc.DecodeBatch(tt.values, tt.variant)

			assert.Equal(t, tt.want, got)
		})
	}
}

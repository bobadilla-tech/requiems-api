package inflation

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

// ---- Request validation ----

func TestRequest_Valid_US(t *testing.T) {
	t.Parallel()
	req := Request{Country: "US"}
	err := httpx.Validate.Struct(&req)
	assert.NoError(t, err)
}

func TestRequest_Valid_GB(t *testing.T) {
	t.Parallel()
	req := Request{Country: "GB"}
	err := httpx.Validate.Struct(&req)
	assert.NoError(t, err)
}

func TestRequest_Empty_Country_Fails(t *testing.T) {
	t.Parallel()
	req := Request{}
	err := httpx.Validate.Struct(&req)
	require.Error(t, err)
}

func TestRequest_Invalid_Country_ZZZ_Fails(t *testing.T) {
	t.Parallel()
	req := Request{Country: "ZZZ"}
	err := httpx.Validate.Struct(&req)
	require.Error(t, err)
}

func TestRequest_Lowercase_Fails(t *testing.T) {
	t.Parallel()
	// iso3166_1_alpha2 requires uppercase; the transport layer uppercases before binding.
	req := Request{Country: "us"}
	err := httpx.Validate.Struct(&req)
	require.Error(t, err)
}

// ---- BatchRequest validation ----

func TestBatchRequest_Valid_OneCountry(t *testing.T) {
	t.Parallel()
	req := BatchRequest{Countries: []string{"US"}}
	err := httpx.Validate.Struct(&req)
	assert.NoError(t, err)
}

func TestBatchRequest_Valid_MaxCountries(t *testing.T) {
	t.Parallel()
	// 50 countries is the allowed maximum.
	countries := make([]string, 50)
	for i := range countries {
		countries[i] = "US"
	}
	req := BatchRequest{Countries: countries}
	err := httpx.Validate.Struct(&req)
	assert.NoError(t, err)
}

func TestBatchRequest_Empty_Fails(t *testing.T) {
	t.Parallel()
	// An empty countries array must be rejected.
	req := BatchRequest{Countries: []string{}}
	err := httpx.Validate.Struct(&req)
	require.Error(t, err)
}

func TestBatchRequest_OverLimit_Fails(t *testing.T) {
	t.Parallel()
	// 51 countries exceeds the max of 50.
	countries := make([]string, 51)
	for i := range countries {
		countries[i] = "US"
	}
	req := BatchRequest{Countries: countries}
	err := httpx.Validate.Struct(&req)
	require.Error(t, err)
}

func TestBatchRequest_InvalidCode_Fails(t *testing.T) {
	t.Parallel()
	// Each item in the array is also validated as iso3166_1_alpha2.
	req := BatchRequest{Countries: []string{"US", "ZZZ"}}
	err := httpx.Validate.Struct(&req)
	require.Error(t, err)
}

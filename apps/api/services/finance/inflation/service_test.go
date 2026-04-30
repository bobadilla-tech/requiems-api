package inflation

import (
	"testing"

	"requiems-api/platform/httpx"
)

// ---- Request validation ----

func TestRequest_Valid_US(t *testing.T) {
	req := Request{Country: "US"}
	if err := httpx.Validate.Struct(&req); err != nil {
		t.Errorf("expected no error for valid country US, got %v", err)
	}
}

func TestRequest_Valid_GB(t *testing.T) {
	req := Request{Country: "GB"}
	if err := httpx.Validate.Struct(&req); err != nil {
		t.Errorf("expected no error for valid country GB, got %v", err)
	}
}

func TestRequest_Empty_Country_Fails(t *testing.T) {
	req := Request{}
	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for empty country, got nil")
	}
}

func TestRequest_Invalid_Country_ZZZ_Fails(t *testing.T) {
	req := Request{Country: "ZZZ"}
	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for invalid country code ZZZ, got nil")
	}
}

func TestRequest_Lowercase_Fails(t *testing.T) {
	// iso3166_1_alpha2 requires uppercase; the transport layer uppercases before binding.
	req := Request{Country: "us"}
	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for lowercase country code, got nil")
	}
}

// ---- BatchRequest validation ----

func TestBatchRequest_Valid_OneCountry(t *testing.T) {
	req := BatchRequest{Countries: []string{"US"}}
	if err := httpx.Validate.Struct(&req); err != nil {
		t.Errorf("expected no error for single valid country, got %v", err)
	}
}

func TestBatchRequest_Valid_MaxCountries(t *testing.T) {
	// 50 countries is the allowed maximum.
	countries := make([]string, 50)
	for i := range countries {
		countries[i] = "US"
	}
	req := BatchRequest{Countries: countries}
	if err := httpx.Validate.Struct(&req); err != nil {
		t.Errorf("expected no error for 50 countries, got %v", err)
	}
}

func TestBatchRequest_Empty_Fails(t *testing.T) {
	// An empty countries array must be rejected.
	req := BatchRequest{Countries: []string{}}
	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for empty countries array, got nil")
	}
}

func TestBatchRequest_OverLimit_Fails(t *testing.T) {
	// 51 countries exceeds the max of 50.
	countries := make([]string, 51)
	for i := range countries {
		countries[i] = "US"
	}
	req := BatchRequest{Countries: countries}
	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for 51 countries, got nil")
	}
}

func TestBatchRequest_InvalidCode_Fails(t *testing.T) {
	// Each item in the array is also validated as iso3166_1_alpha2.
	req := BatchRequest{Countries: []string{"US", "ZZZ"}}
	if err := httpx.Validate.Struct(&req); err == nil {
		t.Error("expected error for invalid country code ZZZ in batch, got nil")
	}
}

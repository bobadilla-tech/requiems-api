package base64 //nolint:revive // package name matches the service domain it implements

import (
	"encoding/base64"

	"requiems-api/platform/svcerr"
)

// Service provides Base64 encode and decode operations.
type Service struct{}

// NewService creates a new Base64 Service.
func NewService() *Service { return &Service{} }

// Encode encodes value using the specified variant ("standard" or "url").
// When variant is empty it defaults to standard Base64 encoding.
func (s *Service) Encode(value, variant string) Result {
	if variant == "url" {
		return Result{Result: base64.URLEncoding.EncodeToString([]byte(value))}
	}
	return Result{Result: base64.StdEncoding.EncodeToString([]byte(value))}
}

// Decode decodes a Base64-encoded string using the specified variant ("standard"
// or "url"). When variant is empty it defaults to standard Base64 encoding.
// Returns an error for invalid Base64 input.
func (s *Service) Decode(value, variant string) (Result, error) {
	var (
		decoded []byte
		err     error
	)

	if variant == "url" {
		decoded, err = base64.URLEncoding.DecodeString(value)
	} else {
		decoded, err = base64.StdEncoding.DecodeString(value)
	}

	if err != nil {
		return Result{}, svcerr.Unknown("invalid_base64", "value is not valid base64")
	}

	return Result{Result: string(decoded)}, nil
}

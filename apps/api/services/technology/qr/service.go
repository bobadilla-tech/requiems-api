package qr

import (
	qrcode "github.com/skip2/go-qrcode"
)

// Base64Response is the JSON response payload returned by GET /qr/base64.
type Base64Response struct {
	Image  string `json:"image"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}


// Service generates QR codes.
type Service struct{}

// NewService returns a new Service instance.
func NewService() *Service {
	return &Service{}
}

// recoveryLevel maps the user-supplied string to a qrcode.RecoveryLevel.
// An empty string (or any unrecognised value) falls back to Medium.
func recoveryLevel(s string) qrcode.RecoveryLevel {
	switch s {
	case "low":
		return qrcode.Low
	case "high":
		return qrcode.High
	case "highest":
		return qrcode.Highest
	default:
		return qrcode.Medium
	}
}

// Generate returns the raw PNG bytes for a QR code encoding data at the
// given pixel size and error-correction level.
// Accepted recovery values: "low", "medium", "high", "highest" (default: "medium").
func (s *Service) Generate(data string, size int, recovery string) ([]byte, error) {
	return qrcode.Encode(data, recoveryLevel(recovery), size)
}

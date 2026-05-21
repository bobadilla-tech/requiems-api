package qr

import (
	"encoding/base64"

	qrcode "github.com/skip2/go-qrcode"
)

// Base64Response is the JSON response payload returned by GET /qr/base64.
type Base64Response struct {
	Image  string `json:"image"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Query is the per-item input for the batch QR base64 endpoint.
type Query struct {
	Data     string `json:"data"     validate:"required"`
	Size     int    `json:"size"     validate:"omitempty,min=50,max=1000"`
	Recovery string `json:"recovery" validate:"omitempty,oneof=low medium high highest"`
}

// BatchBase64Item is the per-item result returned by GenerateBatch.
type BatchBase64Item struct {
	Data   string `json:"data"`
	Image  string `json:"image,omitempty"`
	Width  int    `json:"width,omitempty"`
	Height int    `json:"height,omitempty"`
	Error  string `json:"error,omitempty"`
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

// GenerateBatch generates a base64-encoded QR code for each item in input order.
// A Size of 0 defaults to 256. Per-item errors are absorbed in-band.
func (s *Service) GenerateBatch(items []Query) []BatchBase64Item {
	results := make([]BatchBase64Item, len(items))
	for i, item := range items {
		size := item.Size
		if size == 0 {
			size = 256
		}
		png, err := s.Generate(item.Data, size, item.Recovery)
		if err != nil {
			results[i] = BatchBase64Item{Data: item.Data, Error: err.Error()}
		} else {
			results[i] = BatchBase64Item{
				Data:   item.Data,
				Image:  base64.StdEncoding.EncodeToString(png),
				Width:  size,
				Height: size,
			}
		}
	}
	return results
}

// Generate returns the raw PNG bytes for a QR code encoding data at the
// given pixel size and error-correction level.
// Accepted recovery values: "low", "medium", "high", "highest" (default: "medium").
func (s *Service) Generate(data string, size int, recovery string) ([]byte, error) {
	return qrcode.Encode(data, recoveryLevel(recovery), size)
}

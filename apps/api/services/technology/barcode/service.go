package barcode

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"image/png"
	"sync"
	"unicode"

	bcode "github.com/boombuler/barcode"
	"github.com/boombuler/barcode/code128"
	"github.com/boombuler/barcode/code39"
	"github.com/boombuler/barcode/code93"
	"github.com/boombuler/barcode/ean"
)

const (
	defaultWidth  = 300
	defaultHeight = 100

	maxBatchSize    = 20
	maxBatchWorkers = 10
)

// Base64Response is the JSON response payload returned by GET /barcode/base64.
type Base64Response struct {
	Image  string `json:"image"`
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Service generates barcodes.
type Service struct{}

// NewService returns a new Service instance.
func NewService() *Service {
	return &Service{}
}

// Generate returns raw PNG bytes for a barcode encoding data of the given type.
// Supported types: code128, code93, code39, ean8, ean13.
func (s *Service) Generate(data, barcodeType string) (imgBytes []byte, width, height int, err error) {
	bc, err := encode(data, barcodeType)
	if err != nil {
		return nil, 0, 0, err
	}

	// Scale to a readable size.
	scaled, err := bcode.Scale(bc, defaultWidth, defaultHeight)
	if err != nil {
		return nil, 0, 0, err
	}

	var buf bytes.Buffer
	if err := png.Encode(&buf, scaled); err != nil {
		return nil, 0, 0, err
	}

	return buf.Bytes(), defaultWidth, defaultHeight, nil
}

// encode dispatches to the appropriate barcode encoder.
func encode(data, barcodeType string) (bcode.Barcode, error) {
	switch barcodeType {
	case "code128":
		return code128.Encode(data)
	case "code93":
		return code93.Encode(data, true, false)
	case "code39":
		return code39.Encode(data, true, false)
	case "ean8":
		if len(data) != 7 && len(data) != 8 {
			return nil, errors.New("ean8 requires 7 or 8 digits")
		}
		if !isNumeric(data) {
			return nil, errors.New("ean8 requires numeric digits only")
		}
		return ean.Encode(data)
	case "ean13":
		if len(data) != 12 && len(data) != 13 {
			return nil, errors.New("ean13 requires 12 or 13 digits")
		}
		if !isNumeric(data) {
			return nil, errors.New("ean13 requires numeric digits only")
		}
		return ean.Encode(data)
	default:
		return nil, errors.New("unsupported barcode type")
	}
}

// isNumeric reports whether s consists entirely of decimal digit characters.
func isNumeric(s string) bool {
	for _, r := range s {
		if !unicode.IsDigit(r) {
			return false
		}
	}
	return true
}

// BatchItem is a single barcode request within a batch.
type BatchItem struct {
	Data string `json:"data" validate:"required"`
	Type string `json:"type" validate:"required,oneof=code128 code93 code39 ean8 ean13"`
}

// BatchRequest is the JSON body for POST /barcode/batch.
type BatchRequest struct {
	Items []BatchItem `json:"items" validate:"required,min=1,max=20,dive"`
}

// BatchResultItem is the result for a single item in a batch.
type BatchResultItem struct {
	Image   string `json:"image"`
	Type    string `json:"type"`
	Width   int    `json:"width"`
	Height  int    `json:"height"`
	Success bool   `json:"success"`
	Error   string `json:"error"`
}

// GenerateBatch encodes multiple barcodes concurrently. Results are returned
// in the same order as the input slice. Per-item errors are absorbed in-band
// (Success: false, Error: <reason>) so a single bad item never blocks the rest.
func (s *Service) GenerateBatch(ctx context.Context, items []BatchItem) []BatchResultItem {
	results := make([]BatchResultItem, len(items))

	sem := make(chan struct{}, maxBatchWorkers)
	var wg sync.WaitGroup

	for i, item := range items {
		wg.Add(1)
		sem <- struct{}{}

		go func(i int, item BatchItem) {
			defer wg.Done()
			defer func() { <-sem }()

			pngData, width, height, err := s.Generate(item.Data, item.Type)
			if err != nil {
				results[i] = BatchResultItem{
					Type:    item.Type,
					Success: false,
					Error:   err.Error(),
				}
				return
			}

			results[i] = BatchResultItem{
				Image:   base64.StdEncoding.EncodeToString(pngData),
				Type:    item.Type,
				Width:   width,
				Height:  height,
				Success: true,
			}
		}(i, item)
	}

	wg.Wait()

	return results
}

package barcode

import (
	"bytes"
	"errors"
	"image/png"
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
)

// Base64Response is the JSON response payload returned by GET /barcode/base64.
type Base64Response struct {
	Image  string `json:"image"`
	Type   string `json:"type"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

func (Base64Response) IsData() {}

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

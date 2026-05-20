package numbase

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ErrInvalidBase is returned when from or to is not a supported base (2, 8, 10, 16).
var ErrInvalidBase = errors.New("unsupported base: must be 2, 8, 10, or 16")

// ErrInvalidValue is returned when value cannot be parsed in the given base.
var ErrInvalidValue = errors.New("value is not valid for the given base")

// validBase reports whether b is a supported number base.
func validBase(b int) bool {
	return b == 2 || b == 8 || b == 10 || b == 16
}

// stripPrefix removes the optional base prefix (0x, 0b, 0o) from value when it
// matches fromBase. A leading minus sign is preserved across the stripping.
func stripPrefix(value string, fromBase int) string {
	s := value
	neg := strings.HasPrefix(s, "-")
	if neg {
		s = s[1:]
	}

	lower := strings.ToLower(s)
	switch fromBase {
	case 16:
		if strings.HasPrefix(lower, "0x") {
			s = s[2:]
		}
	case 2:
		if strings.HasPrefix(lower, "0b") {
			s = s[2:]
		}
	case 8:
		if strings.HasPrefix(lower, "0o") {
			s = s[2:]
		}
	}

	if neg {
		return "-" + s
	}
	return s
}

// Result is the response returned by the base conversion endpoint.
type Result struct {
	Input  string `json:"input"`
	From   int    `json:"from"`
	To     int    `json:"to"`
	Result string `json:"result"`
}

// Service provides number base conversion operations.
type Service struct{}

// NewService creates a new base conversion Service.
func NewService() *Service { return &Service{} }

// Convert parses value as a signed integer in fromBase and formats it in toBase.
// Supported bases are 2 (binary), 8 (octal), 10 (decimal), and 16 (hexadecimal).
// Common prefixes such as 0x, 0b, and 0o are accepted for the respective bases.
func (s *Service) Convert(value string, fromBase, toBase int) (Result, error) {
	if !validBase(fromBase) {
		return Result{}, fmt.Errorf("%w: %d", ErrInvalidBase, fromBase)
	}

	if !validBase(toBase) {
		return Result{}, fmt.Errorf("%w: %d", ErrInvalidBase, toBase)
	}

	n, err := strconv.ParseInt(stripPrefix(value, fromBase), fromBase, 64)
	if err != nil {
		return Result{}, ErrInvalidValue
	}

	return Result{
		Input:  value,
		From:   fromBase,
		To:     toBase,
		Result: strconv.FormatInt(n, toBase),
	}, nil
}

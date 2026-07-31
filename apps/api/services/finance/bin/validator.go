package bin

import (
	"strings"

	"requiems-api/platform/svcerr"
)

// sanitizeBIN strips common separators and validates that the result is 6–8
// decimal digits.
func sanitizeBIN(raw string) (string, error) {
	cleaned := strings.Map(func(r rune) rune {
		if r == '-' || r == ' ' {
			return -1
		}
		return r
	}, strings.TrimSpace(raw))

	if len(cleaned) < 6 || len(cleaned) > 8 {
		return "", svcerr.Invalid("bad_request", "BIN must be between 6 and 8 digits")
	}

	for _, ch := range cleaned {
		if ch < '0' || ch > '9' {
			return "", svcerr.Invalid("bad_request", "BIN must contain digits only")
		}
	}

	return cleaned, nil
}

// luhnValid runs the Luhn algorithm on the digit string.
// For a BIN (6–8 digits) this is a partial check on the prefix only.
func luhnValid(s string) bool {
	sum := 0
	nDigits := len(s)
	parity := nDigits % 2

	for i, ch := range s {
		if ch < '0' || ch > '9' {
			return false
		}
		digit := int(ch - '0')
		if i%2 == parity {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}
		sum += digit
	}
	return sum%10 == 0
}

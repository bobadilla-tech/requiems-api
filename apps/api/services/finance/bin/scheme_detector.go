package bin

// detectScheme derives the card scheme from a BIN prefix using the canonical
// ISO/IEC 7812 prefix ranges. Ranges are checked from most specific to least
// specific to avoid false matches on overlapping prefixes.
func detectScheme(bin string) string {
	if len(bin) < 4 {
		return ""
	}

	n2 := atoiN(bin, 2)
	n3 := atoiN(bin, 3)
	n4 := atoiN(bin, 4)
	n6 := atoiN(bin, 6)

	switch {
	// Mir: 2200–2204 — must come before Mastercard 2-series
	case n4 >= 2200 && n4 <= 2204:
		return "mir"

	// Mastercard 2-series: 2221–2720
	case n4 >= 2221 && n4 <= 2720:
		return "mastercard"

	// Amex: 34, 37 — must come before Visa (both start with 3x)
	case n2 == 34 || n2 == 37:
		return "amex"

	// JCB: 3528–3589
	case n4 >= 3528 && n4 <= 3589:
		return "jcb"

	// Diners Club: 300–305, 36, 38
	case (n4 >= 3000 && n4 <= 3059) || n2 == 36 || n2 == 38:
		return "diners"

	// Visa: starts with 4
	case bin[0] == '4':
		return "visa"

	// Mastercard 5-series: 51–55
	case n2 >= 51 && n2 <= 55:
		return "mastercard"

	// Maestro specific prefixes — check before UnionPay (overlapping 6x space)
	case n4 == 6304 || n4 == 6759 || n4 == 6761 || n4 == 6762 || n4 == 6763:
		return "maestro"

	// Discover: 6011
	case n4 == 6011:
		return "discover"

	// Discover: 622126–622925 — must come before UnionPay 62xx
	case n6 >= 622126 && n6 <= 622925:
		return "discover"

	// RuPay: 6521, 6522 — must come before Discover 65xx range
	case n4 == 6521 || n4 == 6522:
		return "rupay"

	// Discover: 644–649 and 65xx
	case (n3 >= 644 && n3 <= 649) || n2 == 65:
		return "discover"

	// RuPay: 60 — check before UnionPay 62
	case n2 == 60:
		return "rupay"

	// UnionPay: 62, 81
	case n2 == 62 || n2 == 81:
		return "unionpay"
	}

	return ""
}

// atoiN converts the first n digits of s to an integer. Returns 0 if s has
// fewer than n characters or contains non-digit bytes.
func atoiN(s string, n int) int {
	if len(s) < n {
		return 0
	}
	v := 0
	for i := range n {
		b := s[i]
		if b < '0' || b > '9' {
			return 0
		}
		v = v*10 + int(b-'0')
	}
	return v
}

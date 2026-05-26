package main

import (
	"math"
	"strings"
)

// normalise applies canonical mapping to a raw record's fields and adjusts the
// confidence score based on data completeness.
func normalise(r RawBINRecord) RawBINRecord {
	r.BINPrefix = strings.TrimSpace(r.BINPrefix)
	r.Scheme = normaliseScheme(r.Scheme)
	normType := normaliseCardType(r.CardType)
	r.CardType = normType
	if normType == "prepaid" {
		r.Prepaid = true
	}
	r.CardLevel = normaliseCardLevel(r.CardLevel)
	r.CountryCode = strings.ToUpper(strings.TrimSpace(r.CountryCode))
	r.CountryName = strings.TrimSpace(r.CountryName)
	r.IssuerName = strings.TrimSpace(r.IssuerName)
	r.IssuerURL = strings.TrimSpace(r.IssuerURL)
	r.IssuerPhone = strings.TrimSpace(r.IssuerPhone)

	// Discard invalid country codes.
	if len(r.CountryCode) != 2 {
		r.CountryCode = ""
		r.CountryName = ""
	}

	// If source didn't include a scheme, derive it from the prefix.
	if r.Scheme == "" {
		r.Scheme = detectScheme(r.BINPrefix)
	}

	// Adjust confidence based on data completeness.
	r.Confidence = computeConfidence(r)

	return r
}

// normaliseScheme maps source scheme strings to lowercase canonical names.
func normaliseScheme(s string) string {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "VISA":
		return "visa"
	case "MASTERCARD", "MASTER CARD", "MASTER":
		return "mastercard"
	case "AMERICAN EXPRESS", "AMEX", "AMERICANEXPRESS":
		return "amex"
	case "DISCOVER", "DISCOVER CARD":
		return "discover"
	case "JCB":
		return "jcb"
	case "DINERS CLUB", "DINERS", "DINERS CLUB INTERNATIONAL":
		return "diners"
	case "UNIONPAY", "UNION PAY", "CHINA UNIONPAY", "CUP":
		return "unionpay"
	case "MAESTRO":
		return "maestro"
	case "MIR":
		return "mir"
	case "RUPAY", "RU PAY":
		return "rupay"
	case "PRIVATE LABEL", "PRIVATELABEL":
		return "private_label"
	default:
		return strings.ToLower(strings.TrimSpace(s))
	}
}

// normaliseCardType maps source type strings to canonical values.
func normaliseCardType(s string) string {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "CREDIT":
		return "credit"
	case "DEBIT":
		return "debit"
	case "PREPAID", "PRE-PAID", "PREPAID DEBIT", "DEBIT (PREPAID)":
		return "prepaid"
	case "CHARGE", "CHARGE CARD":
		return "charge"
	default:
		return strings.ToLower(strings.TrimSpace(s))
	}
}

// normaliseCardLevel maps source category strings to canonical levels.
func normaliseCardLevel(s string) string {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "CLASSIC", "NORMAL", "REGULAR", "BASIC", "TRADITIONAL":
		return "classic"
	case "STANDARD":
		return "standard"
	case "GOLD", "PREMIER", "PREFERRED":
		return "gold"
	case "PLATINUM", "WORLD", "PREMIUM", "TITANIUM", "TITANIO":
		return "platinum"
	case "INFINITE", "WORLD ELITE", "BLACK", "ELITE", "ULTRA":
		return "infinite"
	case "BUSINESS", "PURCHASING":
		return "business"
	case "CORPORATE":
		return "corporate"
	case "SIGNATURE":
		return "signature"
	default:
		return strings.ToLower(strings.TrimSpace(s))
	}
}

// computeConfidence calculates a confidence score (0.00–1.00) for a record
// based on its data completeness and scheme validity.
func computeConfidence(r RawBINRecord) float64 {
	score := r.Confidence

	// Bonus for having a scheme that matches prefix detection rules.
	if r.Scheme != "" && r.Scheme == detectScheme(r.BINPrefix) {
		score += 0.05
	}

	// Bonus for having a bank/issuer name.
	if r.IssuerName != "" {
		score += 0.03
	}

	// Bonus for having country info.
	if r.CountryCode != "" {
		score += 0.02
	}

	return math.Min(1.00, math.Round(score*100)/100)
}

// mergeRecords deduplicates records by BIN prefix. When two sources provide the
// same prefix, the higher-confidence record wins field-by-field. Source names
// are combined and an agreement bonus is applied.
func mergeRecords(records []RawBINRecord) map[string]RawBINRecord {
	merged := make(map[string]RawBINRecord, len(records))

	for _, incoming := range records {
		existing, exists := merged[incoming.BINPrefix]
		if !exists {
			merged[incoming.BINPrefix] = incoming
			continue
		}

		// Combine source names if different.
		sources := combineSources(existing.Source, incoming.Source)

		// Pick the winner by confidence; use the higher-confidence record as base,
		// then fill in any empty fields from the lower-confidence one.
		winner, loser := existing, incoming
		if incoming.Confidence > existing.Confidence {
			winner, loser = incoming, existing
		}

		// Fill empty fields from the loser (secondary source).
		if winner.IssuerName == "" {
			winner.IssuerName = loser.IssuerName
		}
		if winner.IssuerURL == "" {
			winner.IssuerURL = loser.IssuerURL
		}
		if winner.IssuerPhone == "" {
			winner.IssuerPhone = loser.IssuerPhone
		}
		if winner.CountryCode == "" {
			winner.CountryCode = loser.CountryCode
		}
		if winner.CountryName == "" {
			winner.CountryName = loser.CountryName
		}
		if winner.CardType == "" {
			winner.CardType = loser.CardType
		}
		if winner.CardLevel == "" {
			winner.CardLevel = loser.CardLevel
		}

		// Multi-source agreement bonus: +0.10, capped at 1.00.
		winner.Confidence = math.Min(1.00, winner.Confidence+0.10)
		winner.Source = sources

		merged[incoming.BINPrefix] = winner
	}

	return merged
}

// combineSources returns a deduplicated comma-separated list of source names.
func combineSources(a, b string) string {
	if a == b {
		return a
	}
	existing := make(map[string]bool)
	for s := range strings.SplitSeq(a, ",") {
		existing[strings.TrimSpace(s)] = true
	}
	for s := range strings.SplitSeq(b, ",") {
		s = strings.TrimSpace(s)
		if s != "" && !existing[s] {
			existing[s] = true
			a += "," + s
		}
	}
	return a
}

// detectScheme derives the card scheme from a BIN prefix using ISO/IEC 7812
// prefix ranges. Ranges are checked from most specific to least specific to
// avoid false matches on overlapping prefixes.
func detectScheme(bin string) string {
	if len(bin) < 4 {
		return ""
	}

	n2 := atoi2(bin)
	n4 := atoi4(bin)
	n6 := atoi6(bin)

	switch {
	// Mir: 2200–2204 — must come before Mastercard 2-series (2221–2720)
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
	case (n4/10 >= 644 && n4/10 <= 649) || n2 == 65:
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

// atoi2 returns the first 2 digits of s as an int. Returns 0 on error.
func atoi2(s string) int {
	if len(s) < 2 {
		return 0
	}
	return digitVal(s[0])*10 + digitVal(s[1])
}

// atoi4 returns the first 4 digits of s as an int. Returns 0 on error.
func atoi4(s string) int {
	if len(s) < 4 {
		return 0
	}
	return digitVal(s[0])*1000 + digitVal(s[1])*100 + digitVal(s[2])*10 + digitVal(s[3])
}

// atoi6 returns the first 6 digits of s as an int. Returns 0 on error.
func atoi6(s string) int {
	if len(s) < 6 {
		return 0
	}
	v := 0
	for i := range 6 {
		v = v*10 + digitVal(s[i])
	}
	return v
}

func digitVal(b byte) int {
	if b < '0' || b > '9' {
		return 0
	}
	return int(b - '0')
}

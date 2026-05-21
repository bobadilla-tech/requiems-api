package password

import (
	"crypto/rand"
	"math/big"
)

const (
	charsetLower   = "abcdefghijklmnopqrstuvwxyz"
	charsetUpper   = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	charsetNumbers = "0123456789"
	charsetSymbols = "!@#$%^&*()-_=+[]{}|;:,.<>?"
)

// Password is the response payload for the password generator.
type Password struct {
	Password string `json:"password"` //nolint:gosec // intentional: this is the generated password value, not a secret
	Length   int    `json:"length"`
	Strength string `json:"strength"`
}

// PasswordQuery is the per-item input for the batch password endpoint.
type PasswordQuery struct {
	Length    int  `json:"length"    validate:"omitempty,min=8,max=128"`
	Uppercase bool `json:"uppercase"`
	Numbers   bool `json:"numbers"`
	Symbols   bool `json:"symbols"`
}

// BatchResult is the per-item result returned by GenerateBatch.
type BatchResult struct {
	Result *Password `json:"result,omitempty"`
	Error  string    `json:"error,omitempty"`
}

// Service generates cryptographically secure passwords.
type Service struct{}

// NewService returns a new Service instance.
func NewService() *Service {
	return &Service{}
}

// GenerateBatch generates a password for each item and returns results in input order.
// A Length of 0 defaults to 16. Per-item errors are absorbed in-band.
func (s *Service) GenerateBatch(items []PasswordQuery) []BatchResult {
	results := make([]BatchResult, len(items))
	for i, item := range items {
		length := item.Length
		if length == 0 {
			length = 16
		}
		p, err := s.Generate(length, item.Uppercase, item.Numbers, item.Symbols)
		if err != nil {
			results[i] = BatchResult{Error: err.Error()}
		} else {
			results[i] = BatchResult{Result: &p}
		}
	}
	return results
}

// Generate builds a random password from the requested character sets and
// calculates a strength label based on length and character variety.
func (s *Service) Generate(length int, useUppercase, useNumbers, useSymbols bool) (Password, error) {
	alphabet := charsetLower

	// Collect one guaranteed character per enabled charset to ensure variety.
	// Capacity covers all four possible charsets: lower, upper, numbers, symbols.
	const maxCharsets = 4
	required := make([]byte, 0, maxCharsets)

	rc, err := randomByte(charsetLower)
	if err != nil {
		return Password{}, err
	}

	required = append(required, rc)

	if useUppercase {
		alphabet += charsetUpper

		rc, err = randomByte(charsetUpper)
		if err != nil {
			return Password{}, err
		}

		required = append(required, rc)
	}

	if useNumbers {
		alphabet += charsetNumbers

		rc, err = randomByte(charsetNumbers)
		if err != nil {
			return Password{}, err
		}

		required = append(required, rc)
	}

	if useSymbols {
		alphabet += charsetSymbols

		rc, err = randomByte(charsetSymbols)
		if err != nil {
			return Password{}, err
		}

		required = append(required, rc)
	}

	// Fill remaining positions from the full alphabet.
	buf := make([]byte, length)
	copy(buf, required)

	for i := len(required); i < length; i++ {
		buf[i], err = randomByte(alphabet)
		if err != nil {
			return Password{}, err
		}
	}

	// Shuffle using Fisher-Yates with crypto/rand so guaranteed chars are not
	// always at predictable positions.
	if err := shuffle(buf); err != nil {
		return Password{}, err
	}

	return Password{
		Password: string(buf),
		Length:   length,
		Strength: calculateStrength(length, useUppercase, useNumbers, useSymbols),
	}, nil
}

// randomByte returns a cryptographically random byte from alphabet.
func randomByte(alphabet string) (byte, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(alphabet))))
	if err != nil {
		return 0, err
	}

	return alphabet[n.Int64()], nil
}

// shuffle performs an in-place Fisher-Yates shuffle using crypto/rand.
func shuffle(buf []byte) error {
	for i := len(buf) - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return err
		}

		buf[i], buf[j.Int64()] = buf[j.Int64()], buf[i]
	}

	return nil
}

// calculateStrength returns a human-readable strength label based on length
// and the number of enabled character sets.
//
// Scoring (note: the handler enforces length ≥ 8 via validation):
//   - +1 if length ≥ 8  (always true when called through the HTTP handler)
//   - +1 if length ≥ 16
//   - +1 per enabled optional charset (uppercase, numbers, symbols)
//
// Labels:
//   - weak:   score ≤ 1
//   - medium: score 2–3
//   - strong: score ≥ 4
func calculateStrength(length int, uppercase, numbers, symbols bool) string {
	score := 0

	if length >= 8 {
		score++
	}

	if length >= 16 {
		score++
	}

	if uppercase {
		score++
	}

	if numbers {
		score++
	}

	if symbols {
		score++
	}

	switch {
	case score >= 4:
		return "strong"
	case score >= 2:
		return "medium"
	default:
		return "weak"
	}
}

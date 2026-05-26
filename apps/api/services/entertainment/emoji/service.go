package emoji

import (
	"crypto/rand"
	"math/big"
	"strings"
)

// Emoji represents a single emoji with its metadata.
type Emoji struct {
	Emoji    string `json:"emoji"`
	Name     string `json:"name"`
	Category string `json:"category"`
	Unicode  string `json:"unicode"`
}

// List represents a collection of emoji search results.
type List struct {
	Items []Emoji `json:"items"`
	Total int     `json:"total"`
}

// Service provides emoji lookup and search operations.
type Service struct{}

// NewService returns a new Service.
func NewService() *Service { return &Service{} }

// Random returns a randomly selected emoji using a cryptographically secure
// random number generator.
func (s *Service) Random() Emoji {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(emojis))))
	if err != nil {
		// Fallback to first emoji on the extremely unlikely crypto/rand failure.
		return emojis[0]
	}
	return emojis[n.Int64()]
}

// GetByName returns the emoji matching the given name (snake_case).
// Returns the emoji and true if found, or a zero value and false if not.
func (s *Service) GetByName(name string) (Emoji, bool) {
	name = strings.ToLower(name)
	for _, e := range emojis {
		if e.Name == name {
			return e, true
		}
	}
	return Emoji{}, false
}

// Search returns all emojis whose name contains the query string
// (case-insensitive). Returns a List with matching results.
func (s *Service) Search(query string) List {
	query = strings.ToLower(query)
	matches := make([]Emoji, 0)
	for _, e := range emojis {
		if strings.Contains(e.Name, query) || strings.Contains(strings.ToLower(e.Category), query) {
			matches = append(matches, e)
		}
	}
	return List{Items: matches, Total: len(matches)}
}

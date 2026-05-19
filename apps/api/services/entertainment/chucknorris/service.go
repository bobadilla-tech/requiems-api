package chucknorris

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// Fact represents a single Chuck Norris fact/joke.
type Fact struct {
	ID   string `json:"id"`
	Fact string `json:"fact"`
}


// Service provides Chuck Norris fact operations.
type Service struct{}

// NewService returns a new Service.
func NewService() *Service { return &Service{} }

// Random returns a randomly selected Chuck Norris fact using a
// cryptographically secure random number generator.
func (s *Service) Random() Fact {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(len(facts))))
	if err != nil {
		// Fallback to first fact on the extremely unlikely crypto/rand failure.
		return Fact{ID: "cn_0", Fact: facts[0]}
	}
	idx := n.Int64()
	return Fact{
		ID:   fmt.Sprintf("cn_%d", idx),
		Fact: facts[idx],
	}
}

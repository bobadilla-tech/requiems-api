package normalize

import (
	"strings"

	normalizer "github.com/bobadilla-tech/go-email-normalizer"
)

type EmailNormalization struct {
	Original   string              `json:"original"`
	Normalized string              `json:"normalized"`
	Local      string              `json:"local"`
	Domain     string              `json:"domain"`
	Changes    []normalizer.Change `json:"changes"`
}

type EmailNormalizationBatchItem struct {
	Original   string              `json:"original"`
	Normalized string              `json:"normalized,omitempty"`
	Local      string              `json:"local,omitempty"`
	Domain     string              `json:"domain,omitempty"`
	Changes    []normalizer.Change `json:"changes,omitempty"`
	Valid      bool                `json:"valid"`
	Message    string              `json:"message,omitempty"`
}

type Service struct {
	n normalizer.Normalizer
}

func NewService() *Service {
	return &Service{
		n: *normalizer.NewNormalizer(),
	}
}

func (s *Service) Normalize(email string) (EmailNormalization, error) {
	result, err := s.n.Normalize2(email)
	if err != nil {
		return EmailNormalization{}, err
	}

	changes := result.Changes
	if changes == nil {
		changes = []normalizer.Change{}
	}

	local, domain, ok := strings.Cut(result.Normalized, "@")
	if !ok {
		return EmailNormalization{
			Original:   email,
			Normalized: result.Normalized,
			Changes:    changes,
		}, nil
	}

	return EmailNormalization{
		Original:   email,
		Normalized: result.Normalized,
		Local:      local,
		Domain:     domain,
		Changes:    changes,
	}, nil
}

// NormalizeBatch normalizes each address in order. Invalid addresses yield
// valid=false and a message; the overall response is always a full result slice.
func (s *Service) NormalizeBatch(emails []string) []EmailNormalizationBatchItem {
	results := make([]EmailNormalizationBatchItem, len(emails))

	for i, email := range emails {
		res, err := s.Normalize(email)
		if err != nil {
			results[i] = EmailNormalizationBatchItem{
				Original: email,
				Valid:    false,
				Message:  err.Error(),
			}
			continue
		}

		results[i] = EmailNormalizationBatchItem{
			Original:   res.Original,
			Normalized: res.Normalized,
			Local:      res.Local,
			Domain:     res.Domain,
			Changes:    res.Changes,
			Valid:      true,
		}
	}

	return results
}

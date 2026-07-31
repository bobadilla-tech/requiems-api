package convformat

import (
	"fmt"

	"requiems-api/platform/svcerr"
)

const maxContentSize = 512 * 1024 // 512 KB

// Response holds the converted output.
type Response struct {
	Result string `json:"result"`
}

// Service converts content between supported data formats.
type Service struct{}

// NewService creates a new format conversion Service.
func NewService() *Service { return &Service{} }

// Convert converts content from one format to another.
func (s *Service) Convert(req Request) (Response, error) {
	if len(req.Content) > maxContentSize {
		return Response{}, svcerr.Invalid("content_too_large", fmt.Sprintf("content exceeds maximum allowed size of %d bytes", maxContentSize))
	}

	if req.From == req.To {
		return Response{Result: req.Content}, nil
	}

	// Parse input to intermediate representation.
	intermediate, err := parseInput(req.From, req.Content)
	if err != nil {
		return Response{}, err
	}

	// Serialize intermediate to output format.
	result, err := serializeOutput(req.To, intermediate)
	if err != nil {
		return Response{}, err
	}

	return Response{Result: result}, nil
}

// BatchFormatItem is the result for a single item in a batch format conversion request.
type BatchFormatItem struct {
	From   string `json:"from"`
	To     string `json:"to"`
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

// ConvertBatch converts each item and returns results in input order.
// Items that fail conversion return an in-band error.
func (s *Service) ConvertBatch(items []Request) []BatchFormatItem {
	results := make([]BatchFormatItem, len(items))
	for i, item := range items {
		r, err := s.Convert(item)
		if err != nil {
			results[i] = BatchFormatItem{From: item.From, To: item.To, Error: err.Error()}
		} else {
			results[i] = BatchFormatItem{From: item.From, To: item.To, Result: r.Result}
		}
	}
	return results
}

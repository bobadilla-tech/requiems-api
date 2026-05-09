package markdown

import (
	"bytes"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

// Convert renders markdown to HTML following the CommonMark spec.
// When sanitize is true, raw HTML blocks and inline HTML in the markdown input
// are stripped from the output instead of being passed through.
func (s *Service) Convert(markdown string, sanitize bool) (Response, error) {
	opts := []goldmark.Option{
		goldmark.WithExtensions(extension.GFM),
	}

	if !sanitize {
		opts = append(opts, goldmark.WithRendererOptions(html.WithUnsafe()))
	}

	md := goldmark.New(opts...)

	var buf bytes.Buffer
	if err := md.Convert([]byte(markdown), &buf); err != nil {
		return Response{}, err
	}

	return Response{HTML: strings.TrimRight(buf.String(), "\n")}, nil
}

// ConvertBatch renders multiple markdown strings to HTML following the CommonMark spec.
// Each markdown input is processed independently using Convert.
// When sanitize is true, raw HTML blocks and inline HTML in each markdown input
// are stripped from the output instead of being passed through.
func (s *Service) ConvertBatch(markdowns []string, sanitize bool) (BatchResponse, error) {
	results := make([]Response, 0, len(markdowns))

	for _, md := range markdowns {
		resp, err := s.Convert(md, sanitize)
		if err != nil {
			return BatchResponse{}, err
		}

		results = append(results, resp)
	}

	return BatchResponse{
		Results: results,
	}, nil
}

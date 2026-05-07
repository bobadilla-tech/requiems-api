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

// no paralelo pq no hace I/O
// no hace HTTP
// no consulta DB
// todo ocurre en memoria

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


//version concurrente
// func (s *Service) ConvertBatch(req BatchRequest) (BatchResponse, error) {
// 	results := make([]Response, len(req.Markdowns))

// 	errCh := make(chan error, len(req.Markdowns))

// 	var wg sync.WaitGroup

// 	for i, md := range req.Markdowns {
// 		wg.Add(1)

// 		go func(i int, md string) {
// 			defer wg.Done()

// 			resp, err := s.Convert(md, req.Sanitize)
// 			if err != nil {
// 				errCh <- err
// 				return
// 			}

// 			results[i] = resp
// 		}(i, md)
// 	}

// 	wg.Wait()
// 	close(errCh)

// 	if err := <-errCh; err != nil {
// 		return BatchResponse{}, err
// 	}

// 	return BatchResponse{
// 		Results: results,
// 	}, nil
// }
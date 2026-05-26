package lorem

import (
	lorelai "github.com/bobadilla-tech/lorelai/pkg"
)

// Lorem is the response payload for the lorem generator.
type Lorem struct {
	Text       string `json:"text"`
	Paragraphs int    `json:"paragraphs"`
	WordCount  int    `json:"word_count"`
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Generate(paragraphs, sentences int) Lorem {
	lorem := lorelai.ClassicGenerate(paragraphs, sentences)

	return Lorem{
		Text:       lorem.Text,
		Paragraphs: lorem.Paragraphs,
		WordCount:  lorem.WordCount,
	}
}

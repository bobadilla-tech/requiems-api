package lorem

import (
	lorelai "github.com/bobadilla-tech/lorelai/pkg"
)

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

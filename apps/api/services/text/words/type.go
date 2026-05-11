package words

type Word struct {
	ID           int    `json:"id"`
	Word         string `json:"word"`
	Definition   string `json:"definition"`
	PartOfSpeech string `json:"part_of_speech,omitempty"`
}

func (Word) IsData() {}

// Definition represents a single definition entry for a word.
type Definition struct {
	PartOfSpeech string `json:"partOfSpeech"`
	Definition   string `json:"definition"`
	Example      string `json:"example,omitempty"`
}

// DictionaryEntry is the response payload for the dictionary endpoint.
type DictionaryEntry struct {
	Word        string       `json:"word"`
	Phonetic    string       `json:"phonetic,omitempty"`
	Definitions []Definition `json:"definitions"`
	Synonyms    []string     `json:"synonyms"`
}

func (DictionaryEntry) IsData() {}

type BatchRequest struct {
	Items []string `json:"items" validate:"required,min=1,max=50,dive,required"`
}

type BatchItem struct {
	Word  string           `json:"word"`
	Found bool             `json:"found"`
	Entry *DictionaryEntry `json:"entry,omitempty"`
	Error string           `json:"error,omitempty"`
}

type BatchResponse struct {
	Results []BatchItem `json:"results"`
	Total   int         `json:"total"`
}

func (BatchResponse) IsData() {}
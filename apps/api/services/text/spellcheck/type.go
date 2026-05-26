package spellcheck

import "requiems-api/platform/httpx"

// Correction describes a single spelling mistake and its suggested fix.
// Suggested holds the best replacement (used in the auto-corrected text).
// Suggestions holds the top-3 alternatives returned by LanguageTool so
// callers can surface a picker instead of applying the first option blindly.
type Correction struct {
	Original    string   `json:"original"`
	Suggested   string   `json:"suggested"`
	Suggestions []string `json:"suggestions"`
	Position    int      `json:"position"`
}

// Result is the response payload for the spell check endpoint.
type Result struct {
	Corrected   string       `json:"corrected"`
	Corrections []Correction `json:"corrections"`
}

func (Result) IsData() {}

// BatchCheckRequest is the body for checking multiple texts at once.
type BatchCheckRequest struct {
	Texts []string `json:"texts" validate:"required,min=1,max=50,dive,required"`
}

// BatchCheckResponse is the standard batch envelope for spell-check results.
type BatchCheckResponse = httpx.BatchResponse[Result]

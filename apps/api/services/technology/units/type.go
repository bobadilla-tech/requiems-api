package units

type Result struct {
	From    string  `json:"from"`
	To      string  `json:"to"`
	Input   float64 `json:"input"`
	Result  float64 `json:"result"`
	Formula string  `json:"formula"`
}

func (Result) IsData() {}

// Results maps each measurement category to its supported unit keys.
type Results struct {
	Length      []string `json:"length"`
	Weight      []string `json:"weight"`
	Volume      []string `json:"volume"`
	Temperature []string `json:"temperature"`
	Area        []string `json:"area"`
	Speed       []string `json:"speed"`
}

func (Results) IsData() {}

// BatchItem represents a single unit conversion operation.
type BatchItem struct {
	From  string   `json:"from" validate:"required"`
	To    string   `json:"to"  validate:"required"`
	Value *float64 `json:"value" validate:"required"`
}

// BatchRequest represents a request containing multiple conversion operations.
type BatchRequest struct {
	Operations []BatchItem `json:"operations" validate:"required,min=1,max=50,dive"`
}

// BatchResponse represents the result of a single batch conversion operation.
type BatchResponse struct {
	From    string `json:"from"`
	To      string `json:"to"`
	Success bool   `json:"success"`
	Error   string `json:"error,omitempty"`
	Data    Result `json:"data"`
}

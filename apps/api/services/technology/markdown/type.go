package markdown

// Request is the input for the markdown-to-HTML conversion endpoint.
type Request struct {
	Markdown string `json:"markdown" validate:"required"`
	Sanitize bool   `json:"sanitize"`
}

// Response holds the converted HTML output.
type Response struct {
	HTML string `json:"html"`
}

func (Response) IsData() {}

// BatchRequest is the input for batch markdown-to-HTML conversion.
type BatchRequest struct {
	Markdowns []string `json:"markdowns" validate:"required,min=1,max=50,dive,required"`
	Sanitize  bool     `json:"sanitize"`
}

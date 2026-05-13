package randomuser

// Holds the user's postal address fields.
type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	State   string `json:"state"`
	Zip     string `json:"zip"`
	Country string `json:"country"`
}

// Response model for the random-user endpoint.
type User struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Phone   string  `json:"phone"`
	Address Address `json:"address"`
	Avatar  string  `json:"avatar"`
}

// IsData marks User as a valid httpx response data type.
func (User) IsData() {}

// BatchGenerateRequest is the body for generating multiple random users at once.
type BatchGenerateRequest struct {
	Count int `json:"count" validate:"required,min=1,max=50"`
}

// BatchGenerateResponse is the response for a batch random user generation request.
type BatchGenerateResponse struct {
	Results []User `json:"results"`
	Total   int    `json:"total"`
}

// IsData marks BatchGenerateResponse as a valid httpx response data type.
func (BatchGenerateResponse) IsData() {}

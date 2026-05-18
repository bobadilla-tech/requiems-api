package timezone

// Request holds the query parameters for the timezone endpoint.
// Either lat+lon together, or city must be provided.
type Request struct {
	Lat  float64 `query:"lat" validate:"min=-90,max=90"`
	Lon  float64 `query:"lon" validate:"min=-180,max=180"`
	City string  `query:"city"`
}

// Info represents the timezone response.
type Info struct {
	Timezone    string `json:"timezone"`
	Offset      string `json:"offset"`
	CurrentTime string `json:"current_time"`
	IsDST       bool   `json:"is_dst"`
}

func (Info) IsData() {}

type BatchRequest struct {
	Cities []string `json:"cities" validate:"required,min=1,max=50"`
}

type BatchResult struct {
	City string `json:"city"`
	Info *Info  `json:"info,omitempty"`
}

type BatchResponse struct {
	Results []BatchResult `json:"results"`
}

func (BatchResponse) IsData() {}

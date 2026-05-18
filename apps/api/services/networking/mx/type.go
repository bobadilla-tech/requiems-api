package mx

// Record represents a single MX record entry.
type Record struct {
	Host     string `json:"host"`
	Priority uint16 `json:"priority"`
}

// LookupResponse is the JSON payload returned by the MX lookup endpoint.
type LookupResponse struct {
	Domain  string   `json:"domain"`
	Records []Record `json:"records"`
}

// BatchRequest is the body for validating multiple domains at once.
type BatchRequest struct {
	Domains []string `json:"domains" validate:"required,min=1,max=50,dive,required,fqdn"`
}

// BatchLookupItem represents the result of a single domain lookup
// within a batch request. It contains the domain name, whether it
// was found, any error encountered, and the associated lookup data.
type BatchLookupItem struct {
	Domain string         `json:"domain"`
	Found  bool           `json:"found"`
	Error  string         `json:"error,omitempty"`
	Data   LookupResponse `json:"data,omitempty"`
}

func (LookupResponse) IsData() {}

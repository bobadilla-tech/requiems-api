package whois

import "requiems-api/platform/httpx"

// LookupResponse is the JSON payload returned by the WHOIS endpoint.
type LookupResponse struct {
	Domain      string   `json:"domain"`
	Registrar   string   `json:"registrar,omitempty"`
	NameServers []string `json:"name_servers,omitempty"`
	Status      []string `json:"status,omitempty"`
	CreatedDate string   `json:"created_date,omitempty"`
	UpdatedDate string   `json:"updated_date,omitempty"`
	ExpiryDate  string   `json:"expiry_date,omitempty"`
	DNSSec      bool     `json:"dnssec"`
}

func (LookupResponse) IsData() {}

type BatchLookupRequest struct {
	Domains []string `json:"domains" validate:"required,min=1,max=50,dive,required,hostname_rfc1123"`
}

type BatchLookupItem struct {
	Domain string         `json:"domain"`
	Found  bool           `json:"found"`
	Error  string         `json:"error,omitempty"`
	Data   LookupResponse `json:"data,omitempty"`
}

type BatchLookupResponse = httpx.BatchResponse[BatchLookupItem]

package whois

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
	Domains []string `json:"domains" validate:"required,min=1,dive,required"`
}

type BatchLookupResponse struct {
	Results []LookupResponse `json:"results"`
}

func (BatchLookupResponse) IsData() {}
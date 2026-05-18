package vpn

import "github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"

// IPCheckResponse is the JSON payload returned by the VPN/proxy detection endpoint.
type IPCheckResponse struct {
	// IP is the analysed address.
	IP string `json:"ip"`
	// IsVPN is true when the address belongs to a known VPN provider.
	IsVPN bool `json:"is_vpn"`
	// IsProxy is true when the address is a known public or web proxy.
	IsProxy bool `json:"is_proxy"`
	// IsTor is true when the address is a known Tor exit node, detected via the
	// IP2Proxy database or the Tor Project's DNSBL.
	IsTor bool `json:"is_tor"`
	// IsHosting is true when the address belongs to a data-centre or hosting
	// provider range (DCH in IP2Proxy terminology).
	IsHosting bool `json:"is_hosting"`
	// Score is the raw threat score used to derive Threat.
	// Tor contributes 3, VPN or Proxy each contribute 2, Hosting contributes 1.
	Score int `json:"score"`
	// Threat is the threat level derived from Score:
	// 0 → None, 1 → Low, 2–3 → Medium, 4–5 → High, ≥6 → Critical.
	Threat ipi.ThreatLevel `json:"threat"`
	// FraudScore is populated when using an IP2Proxy database of tier PX5 or
	// higher. It ranges from 0 (no fraud risk) to 100 (high fraud risk). Zero
	// means the value is unavailable for the current database tier.
	FraudScore int `json:"fraud_score"`
	// AsnOrg is the name of the organisation that owns the Autonomous System
	// containing the address (e.g. "DIGITALOCEAN-ASN"). Empty when the ASN
	// lookup returns no record.
	AsnOrg string `json:"asn_org"`
}

func (IPCheckResponse) IsData() {}


type BatchRequest struct {
	IPs []string `json:"ips" validate:"required,min=1,max=50,dive,required,ip"`
}


type BatchResponse struct {
	Results []IPCheckResponse `json:"results"`
}

func (BatchResponse) IsData() {}
package domain

import (
	"context"
	"net"
	"strings"
	"sync"
	"time"
)

// MXRecord holds an MX record entry.
type MXRecord struct {
	Host     string `json:"host"`
	Priority uint16 `json:"priority"`
}

// DNSRecords holds the DNS records for a domain.
type DNSRecords struct {
	A     []string   `json:"a"`
	AAAA  []string   `json:"aaaa"`
	MX    []MXRecord `json:"mx"`
	NS    []string   `json:"ns"`
	TXT   []string   `json:"txt"`
	CNAME string     `json:"cname,omitempty"`
}

// InfoResponse is the response for a domain info request.
type InfoResponse struct {
	Domain    string     `json:"domain"`
	Available bool       `json:"available"`
	DNS       DNSRecords `json:"dns"`
}

// Service handles domain information and availability checks.
type Service struct{}

// NewService creates a new domain Service.
func NewService() *Service {
	return &Service{}
}

// GetInfo returns DNS records and availability status for a domain.
func (s *Service) GetInfo(ctx context.Context, domainName string) InfoResponse {
	resp := InfoResponse{
		Domain: domainName,
		DNS: DNSRecords{
			A:    []string{},
			AAAA: []string{},
			MX:   []MXRecord{},
			NS:   []string{},
			TXT:  []string{},
		},
	}

	type ipResult struct {
		a    []string
		aaaa []string
	}
	type mxResult struct{ records []MXRecord }
	type nsResult struct {
		records []string
		err     error
	}
	type txtResult struct{ records []string }
	type cnameResult struct{ cname string }

	var (
		ipCh    = make(chan ipResult, 1)
		mxCh    = make(chan mxResult, 1)
		nsCh    = make(chan nsResult, 1)
		txtCh   = make(chan txtResult, 1)
		cnameCh = make(chan cnameResult, 1)
	)

	var wg sync.WaitGroup
	wg.Add(5)

	// Lookup IP addresses (A and AAAA records).
	go func() {
		defer wg.Done()
		var r ipResult
		addrs, _ := net.DefaultResolver.LookupIPAddr(ctx, domainName)
		for _, addr := range addrs {
			if addr.IP.To4() != nil {
				r.a = append(r.a, addr.IP.String())
			} else {
				r.aaaa = append(r.aaaa, addr.IP.String())
			}
		}
		ipCh <- r
	}()

	// Lookup MX records.
	go func() {
		defer wg.Done()
		var r mxResult
		mxs, _ := net.DefaultResolver.LookupMX(ctx, domainName)
		for _, mx := range mxs {
			r.records = append(r.records, MXRecord{Host: mx.Host, Priority: mx.Pref})
		}
		mxCh <- r
	}()

	// Lookup NS records (also used for availability check).
	go func() {
		defer wg.Done()
		var r nsResult
		nss, err := net.DefaultResolver.LookupNS(ctx, domainName)
		r.err = err
		for _, ns := range nss {
			r.records = append(r.records, ns.Host)
		}
		nsCh <- r
	}()

	// Lookup TXT records.
	go func() {
		defer wg.Done()
		var r txtResult
		r.records, _ = net.DefaultResolver.LookupTXT(ctx, domainName)
		txtCh <- r
	}()

	// Lookup CNAME. LookupCNAME always returns the canonical name (the
	// input domain with a trailing dot when no alias exists), so only
	// populate the field when a real alias is present.
	go func() {
		defer wg.Done()
		var r cnameResult
		cname, _ := net.DefaultResolver.LookupCNAME(ctx, domainName)
		canonical := strings.TrimSuffix(cname, ".")
		if canonical != "" && canonical != domainName {
			r.cname = cname
		}
		cnameCh <- r
	}()

	wg.Wait()

	ip := <-ipCh
	if len(ip.a) > 0 {
		resp.DNS.A = ip.a
	}
	if len(ip.aaaa) > 0 {
		resp.DNS.AAAA = ip.aaaa
	}

	mx := <-mxCh
	if len(mx.records) > 0 {
		resp.DNS.MX = mx.records
	}

	ns := <-nsCh
	if len(ns.records) > 0 {
		resp.DNS.NS = ns.records
	}
	// A domain is considered available (unregistered) when the NS lookup
	// returns NXDOMAIN, meaning no authoritative name servers are delegated.
	resp.Available = isNXDomain(ns.err)

	txt := <-txtCh
	if len(txt.records) > 0 {
		resp.DNS.TXT = txt.records
	}

	resp.DNS.CNAME = (<-cnameCh).cname

	return resp
}

// BatchDomainItem is the result for a single domain in a batch info request.
type BatchDomainItem struct {
	Domain string        `json:"domain"`
	Result *InfoResponse `json:"result,omitempty"`
	Error  string        `json:"error,omitempty"`
}

// GetInfoBatch performs DNS lookups for multiple domains concurrently.
// Domains that fail format validation return an in-band error.
// Max concurrency is 5 — each domain already fans out 5 DNS queries internally.
func (s *Service) GetInfoBatch(ctx context.Context, domains []string) []BatchDomainItem {
	const maxWorkers = 5
	results := make([]BatchDomainItem, len(domains))

	sem := make(chan struct{}, maxWorkers)
	var wg sync.WaitGroup

	for i, d := range domains {
		if !domainRe.MatchString(d) {
			results[i] = BatchDomainItem{Domain: d, Error: "invalid domain format"}
			continue
		}

		wg.Add(1)
		sem <- struct{}{}
		go func(i int, domain string) {
			defer wg.Done()
			defer func() { <-sem }()

			itemCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()

			info := s.GetInfo(itemCtx, domain)
			results[i] = BatchDomainItem{Domain: domain, Result: &info}
		}(i, d)
	}

	wg.Wait()
	return results
}

// isNXDomain reports whether err is a DNS "domain not found" error.
func isNXDomain(err error) bool {
	if err == nil {
		return false
	}
	dnsErr, ok := err.(*net.DNSError)
	return ok && dnsErr.IsNotFound
}

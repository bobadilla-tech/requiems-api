package httpx

import (
	"net"
	"net/http"
	"strings"
)

// trustedProxyCIDRs are the hops allowed to set X-Forwarded-For/X-Real-IP:
// RFC1918 + loopback (the local Caddy, reached over the docker-compose
// network) and Cloudflare's published edge ranges (ahead of Phase 3 item 6's
// origin lockdown landing — defense in depth once Caddy proxies directly).
// Static snapshot, not fetched dynamically; same "placeholder, revisit"
// caveat as this codebase's other hardcoded infra constants — Cloudflare's
// ranges change rarely but do change.
var trustedProxyCIDRs = mustParseCIDRs(
	// RFC1918 + loopback
	"10.0.0.0/8", "172.16.0.0/12", "192.168.0.0/16", "127.0.0.0/8", "::1/128",
	// Cloudflare IPv4 (source: https://www.cloudflare.com/ips-v4)
	"173.245.48.0/20", "103.21.244.0/22", "103.22.200.0/22", "103.31.4.0/22",
	"141.101.64.0/18", "108.162.192.0/18", "190.93.240.0/20", "188.114.96.0/20",
	"197.234.240.0/22", "198.41.128.0/17", "162.158.0.0/15", "104.16.0.0/13",
	"104.24.0.0/14", "172.64.0.0/13", "131.0.72.0/22",
	// Cloudflare IPv6 (source: https://www.cloudflare.com/ips-v6)
	"2400:cb00::/32", "2606:4700::/32", "2803:f800::/32", "2405:b500::/32",
	"2405:8100::/32", "2a06:98c0::/29", "2c0f:f248::/32",
)

func mustParseCIDRs(cidrs ...string) []*net.IPNet {
	nets := make([]*net.IPNet, 0, len(cidrs))
	for _, cidr := range cidrs {
		_, ipNet, err := net.ParseCIDR(cidr)
		if err != nil {
			panic("httpx: invalid trusted proxy CIDR " + cidr + ": " + err.Error())
		}
		nets = append(nets, ipNet)
	}
	return nets
}

func isTrustedProxy(ip net.IP) bool {
	if ip == nil {
		return false
	}
	for _, ipNet := range trustedProxyCIDRs {
		if ipNet.Contains(ip) {
			return true
		}
	}
	return false
}

// ClientIP returns the caller's real IP, honoring X-Forwarded-For/X-Real-IP
// only when the immediate connection (RemoteAddr) comes from a trusted proxy
// hop (Cloudflare's edge or the local Caddy). Otherwise a spoofed forwarded
// header from an untrusted, direct connection is ignored and RemoteAddr — the
// actual TCP peer — is returned instead.
func ClientIP(r *http.Request) string {
	remoteHost := r.RemoteAddr
	if host, _, err := net.SplitHostPort(remoteHost); err == nil {
		remoteHost = host
	}
	remoteIP := net.ParseIP(remoteHost)

	if !isTrustedProxy(remoteIP) {
		return remoteHost
	}

	if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
		if before, _, ok := strings.Cut(xff, ","); ok {
			return strings.TrimSpace(before)
		}
		return strings.TrimSpace(xff)
	}
	if xri := r.Header.Get("X-Real-IP"); xri != "" {
		return strings.TrimSpace(xri)
	}
	return remoteHost
}

package httpx

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestClientIP_TrustedProxyHonorsForwardedHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	req.RemoteAddr = "172.20.0.5:54321" // docker-compose bridge range, trusted
	req.Header.Set("X-Forwarded-For", "203.0.113.7, 172.20.0.5")

	got := ClientIP(req)
	if got != "203.0.113.7" {
		t.Fatalf("ClientIP() = %q, want %q", got, "203.0.113.7")
	}
}

func TestClientIP_TrustedProxyHonorsCloudflareRange(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	req.RemoteAddr = "104.16.1.1:443" // within Cloudflare's published range
	req.Header.Set("X-Forwarded-For", "198.51.100.9")

	got := ClientIP(req)
	if got != "198.51.100.9" {
		t.Fatalf("ClientIP() = %q, want %q", got, "198.51.100.9")
	}
}

func TestClientIP_UntrustedRemoteIgnoresSpoofedHeader(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	req.RemoteAddr = "203.0.113.55:12345" // public, not a trusted proxy hop
	req.Header.Set("X-Forwarded-For", "6.6.6.6")

	got := ClientIP(req)
	if got != "203.0.113.55" {
		t.Fatalf("ClientIP() = %q, want the actual RemoteAddr %q (spoofed header must be ignored)", got, "203.0.113.55")
	}
}

func TestClientIP_UntrustedRemoteIgnoresSpoofedXRealIP(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	req.RemoteAddr = "203.0.113.55:12345"
	req.Header.Set("X-Real-IP", "6.6.6.6")

	got := ClientIP(req)
	if got != "203.0.113.55" {
		t.Fatalf("ClientIP() = %q, want %q", got, "203.0.113.55")
	}
}

func TestClientIP_TrustedProxyWithNoForwardedHeaderFallsBackToRemoteAddr(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	req.RemoteAddr = "172.20.0.5:54321"

	got := ClientIP(req)
	if got != "172.20.0.5" {
		t.Fatalf("ClientIP() = %q, want %q", got, "172.20.0.5")
	}
}

func TestClientIP_LoopbackIsTrusted(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", http.NoBody)
	req.RemoteAddr = "127.0.0.1:8080"
	req.Header.Set("X-Forwarded-For", "203.0.113.7")

	got := ClientIP(req)
	if got != "203.0.113.7" {
		t.Fatalf("ClientIP() = %q, want %q", got, "203.0.113.7")
	}
}

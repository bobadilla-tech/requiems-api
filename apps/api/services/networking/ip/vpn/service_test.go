package vpn

import (
	"context"
	"net"
	"testing"

	"github.com/bobadilla-tech/go-ip-intelligence/v2/ipi"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func newTestClient() (*ipi.Client, error) {
	return ipi.New(
		ipi.WithDatabasePath(""),
		ipi.WithASNDatabasePath(""),
		ipi.WithCityDatabasePath(""),
	)
}

func TestService_CheckIP(t *testing.T) {
	t.Parallel()
	client, err := newTestClient()
	if err != nil {
		t.Skipf("VPN service not available: %v", err)
	}
	svc := NewService(client)

	tests := []struct {
		name    string
		ip      string
		wantIP  string
		wantErr bool
	}{
		{
			name:    "valid IPv4",
			ip:      "8.8.8.8",
			wantIP:  "8.8.8.8",
			wantErr: false,
		},
		{
			name:    "valid IPv6",
			ip:      "2001:4860:4860::8888",
			wantIP:  "2001:4860:4860::8888",
			wantErr: false,
		},
		{
			name:    "another IPv4",
			ip:      "1.1.1.1",
			wantIP:  "1.1.1.1",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ip := net.ParseIP(tt.ip)
			require.NotNil(t, ip, "failed to parse IP: %s", tt.ip)

			result, err := svc.CheckIP(context.Background(), ip)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			require.NoError(t, err)

			assert.Equal(t, tt.wantIP, result.IP)
		})
	}
}

func TestService_CheckIP_ResponseFields(t *testing.T) {
	t.Parallel()
	client, err := newTestClient()
	if err != nil {
		t.Skipf("VPN service not available: %v", err)
	}
	svc := NewService(client)

	ip := net.ParseIP("8.8.8.8")
	result, err := svc.CheckIP(context.Background(), ip)
	require.NoError(t, err)

	assert.NotEmpty(t, result.IP)

	assert.True(t, result.Score >= 0, "expected non-negative score, got %d", result.Score)

	validThreats := map[string]bool{
		"none":     true,
		"low":      true,
		"medium":   true,
		"high":     true,
		"critical": true,
	}
	assert.True(t, validThreats[result.Threat.String()], "invalid threat level: %s", result.Threat)

	assert.True(t, result.FraudScore >= 0 && result.FraudScore <= 100, "fraud_score out of range: %d", result.FraudScore)
}

func TestService_CheckBatch(t *testing.T) {
	t.Parallel()

	client, err := newTestClient()
	if err != nil {
		t.Skipf("VPN service not available: %v", err)
	}

	svc := NewService(client)

	tests := []struct {
		name        string
		ips         []string
		expectedLen int
	}{
		{
			name: "multiple valid IPs",
			ips: []string{
				"8.8.8.8",
				"1.1.1.1",
				"2001:4860:4860::8888",
			},
			expectedLen: 3,
		},
		{
			name: "includes invalid IP",
			ips: []string{
				"8.8.8.8",
				"invalid-ip",
				"1.1.1.1",
			},
			expectedLen: 3,
		},
		{
			name:        "empty batch",
			ips:         []string{},
			expectedLen: 0,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			results, err := svc.CheckBatch(context.Background(), tt.ips)

			require.NoError(t, err)
			require.Len(t, results, tt.expectedLen)

			for i, res := range results {
				assert.Equal(t, tt.ips[i], res.IP)
			}
		})
	}
}

func TestService_CheckBatch_PreservesOrder(t *testing.T) {
	t.Parallel()

	client, err := newTestClient()
	if err != nil {
		t.Skipf("VPN service not available: %v", err)
	}

	svc := NewService(client)

	ips := []string{
		"8.8.8.8",
		"1.1.1.1",
		"2001:4860:4860::8888",
	}

	results, err := svc.CheckBatch(context.Background(), ips)

	require.NoError(t, err)
	require.Len(t, results, len(ips))

	for i := range ips {
		assert.Equal(t, ips[i], results[i].IP)
	}
}

func TestService_CheckBatch_InvalidIPs(t *testing.T) {
	t.Parallel()

	client, err := newTestClient()
	if err != nil {
		t.Skipf("VPN service not available: %v", err)
	}

	svc := NewService(client)

	ips := []string{
		"invalid-ip",
		"",
		"999.999.999.999",
	}

	results, err := svc.CheckBatch(context.Background(), ips)

	require.NoError(t, err)
	require.Len(t, results, len(ips))

	for i, res := range results {
		assert.Equal(t, ips[i], res.IP)

		assert.False(t, res.IsVPN)
		assert.False(t, res.IsProxy)
		assert.False(t, res.IsTor)
		assert.False(t, res.IsHosting)
		assert.Equal(t, 0, res.Score)
	}
}

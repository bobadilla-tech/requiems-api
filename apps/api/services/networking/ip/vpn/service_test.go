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

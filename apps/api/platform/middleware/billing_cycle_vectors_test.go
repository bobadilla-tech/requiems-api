package middleware

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// billingCycleVector is shared with the Rails BillingCycle test
// (apps/dashboard/test/lib/billing_cycle_test.rb). next_start is only asserted
// by Rails: Go has no next-boundary helper (cycle.AddDate(0, 1, 0) is the
// known 29-31 overflow tracked as D1 in the usage plan).
type billingCycleVector struct {
	Name      string    `json:"name"`
	Anchor    time.Time `json:"anchor"`
	Now       time.Time `json:"now"`
	Start     time.Time `json:"start"`
	NextStart time.Time `json:"next_start"`
}

func TestCycleStartMatchesSharedVectors(t *testing.T) {
	raw, err := os.ReadFile("testdata/billing_cycle_vectors.json")
	require.NoError(t, err)

	var vectors []billingCycleVector
	require.NoError(t, json.Unmarshal(raw, &vectors))
	require.NotEmpty(t, vectors)

	for _, v := range vectors {
		t.Run(v.Name, func(t *testing.T) {
			require.True(t, v.Start.Equal(cycleStart(v.Anchor, v.Now)),
				"cycleStart(%s, %s) = %s, want %s", v.Anchor, v.Now, cycleStart(v.Anchor, v.Now), v.Start)
		})
	}
}

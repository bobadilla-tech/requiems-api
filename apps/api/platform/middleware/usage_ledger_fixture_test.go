package middleware

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// TestSumUsageMatchesSharedLedgerFixture runs sumUsage (the query that
// bootstraps the Redis quota counter) over the same rows the Rails
// UsageSummary test seeds (apps/dashboard/test/services/usage_summary_test.rb),
// so the dashboard's "used" and the enforced baseline cannot drift apart.
func TestSumUsageMatchesSharedLedgerFixture(t *testing.T) {
	raw, err := os.ReadFile("testdata/usage_ledger_fixture.json")
	require.NoError(t, err)

	var fixture struct {
		CycleStart   time.Time `json:"cycle_start"`
		ExpectedUsed int64     `json:"expected_used"`
		Rows         []struct {
			Owner       string    `json:"owner"`
			UsedAt      time.Time `json:"used_at"`
			CreditsUsed *int      `json:"credits_used"`
		} `json:"rows"`
	}
	require.NoError(t, json.Unmarshal(raw, &fixture))

	pool := setupUsageTestDB(t)
	ctx := context.Background()

	const subjectID, otherID, apiKeyID = int64(910001), int64(910002), int64(910001)
	owners := map[string]int64{"subject": subjectID, "other": otherID}

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id IN ($1, $2)`, subjectID, otherID)
	})

	for _, row := range fixture.Rows {
		_, err := pool.Exec(ctx, `
			INSERT INTO usage_logs (user_id, api_key_id, endpoint, credits_used, request_method, status_code, used_at, usage_date)
			VALUES ($1, $2, 'fixture', $3, 'GET', 200, $4, $5)
		`, owners[row.Owner], apiKeyID+owners[row.Owner]-subjectID, row.CreditsUsed, row.UsedAt, row.UsedAt.Format("2006-01-02"))
		require.NoError(t, err)
	}

	u := NewUsageQuota(pool, nil, nil, nil)
	sum, err := u.sumUsage(ctx, subjectID, fixture.CycleStart)
	require.NoError(t, err)
	require.Equal(t, fixture.ExpectedUsed, sum)
}

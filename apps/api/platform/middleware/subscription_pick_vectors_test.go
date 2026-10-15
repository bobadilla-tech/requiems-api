package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// subscriptionPickVector is shared with the Rails BillingCycle test
// (apps/dashboard/test/lib/billing_cycle_test.rb): both sides must pick the
// same subscription, and so the same plan and cycle anchor, from the same rows.
type subscriptionPickVector struct {
	Name         string    `json:"name"`
	KeyCreatedAt time.Time `json:"key_created_at"`
	Subscription []struct {
		PlanName           string     `json:"plan_name"`
		Status             string     `json:"status"`
		CurrentPeriodStart *time.Time `json:"current_period_start"`
		UpdatedAt          time.Time  `json:"updated_at"`
	} `json:"subscriptions"`
	ExpectPlan   string    `json:"expect_plan"`
	ExpectAnchor time.Time `json:"expect_anchor"`
}

func TestAPIKeyAuth_SubscriptionPickMatchesSharedVectors(t *testing.T) {
	raw, err := os.ReadFile("testdata/subscription_pick_vectors.json")
	require.NoError(t, err)

	var vectors []subscriptionPickVector
	require.NoError(t, json.Unmarshal(raw, &vectors))
	require.NotEmpty(t, vectors)

	pool := setupAPIKeyAuthTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)
	ctx := context.Background()

	// The Rails schema allows a NULL period start (webhooks can leave it
	// unset); the shared test table defaults to NOW() NOT NULL.
	_, err = pool.Exec(ctx, `ALTER TABLE subscriptions ALTER COLUMN current_period_start DROP NOT NULL`)
	require.NoError(t, err)

	for _, v := range vectors {
		t.Run(v.Name, func(t *testing.T) {
			prefix, fullKey := testKey(t)
			userID := time.Now().UnixNano()
			insertAPIKey(t, pool, prefix, fullKey, userID, true, false)

			_, err := pool.Exec(ctx, `UPDATE api_keys SET created_at = $1 WHERE key_prefix = $2`, v.KeyCreatedAt, prefix)
			require.NoError(t, err)

			for _, s := range v.Subscription {
				_, err := pool.Exec(ctx, `
					INSERT INTO subscriptions (user_id, plan_name, status, current_period_start, updated_at)
					VALUES ($1, $2, $3, $4, $5)
				`, userID, s.PlanName, s.Status, s.CurrentPeriodStart, s.UpdatedAt)
				require.NoError(t, err)
			}
			t.Cleanup(func() { _, _ = pool.Exec(ctx, `DELETE FROM subscriptions WHERE user_id = $1`, userID) })

			var got APIKeyPrincipal
			next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				got, _ = APIKeyPrincipalFromContext(r.Context())
				w.WriteHeader(http.StatusOK)
			})

			w := doRequest(t, NewAPIKeyAuth(pool, rdb, time.Minute).Middleware()(next), fullKey)
			require.Equal(t, http.StatusOK, w.Code)
			require.Equal(t, v.ExpectPlan, got.Plan)
			require.True(t, v.ExpectAnchor.Equal(got.CurrentPeriodStart),
				"anchor = %s, want %s", got.CurrentPeriodStart, v.ExpectAnchor)
		})
	}
}

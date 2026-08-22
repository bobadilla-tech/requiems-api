package middleware

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"

	"requiems-api/platform/httpx"
)

// apiKeyHeader is the header the (eventually retired) Cloudflare Worker
// gateway validates against KV today; this middleware validates the same
// header directly against Postgres for requests that reach Go without going
// through the Worker.
const apiKeyHeader = "requiems-api-key"

// keyPrefixLength mirrors the Worker's extractKeyPrefix
// (apps/workers/shared/src/api-key-generator.ts) and Rails'
// ApiKeyGenerator.extract_prefix (apps/dashboard/app/services/api_key_generator.rb):
// the first 12 characters of the full key.
const keyPrefixLength = 12

// apiKeyCacheKeyPrefix is part of the Redis key contract shared with Rails'
// revocation invalidation path (ApiKey#revoke!). Rails must issue
// `DEL apikey:{key_prefix}` against a raw, unnamespaced Redis connection —
// NOT Rails.cache, which prefixes keys with "rails_cache:" and would
// silently no-op. Changing this prefix requires updating the Rails side too.
const apiKeyCacheKeyPrefix = "apikey:"

type apiKeyContextKey struct{}

// APIKeyPrincipal is the authenticated identity attached to the request
// context once APIKeyAuth has validated a key.
type APIKeyPrincipal struct {
	UserID int64
	Plan   string

	// APIKeyID identifies the specific key used for this request (api_keys.id).
	// Rate limiting (ratelimit.go) is keyed per-key, not per-user.
	APIKeyID int64

	// CurrentPeriodStart anchors the user's monthly billing cycle for quota
	// tracking (usage.go): subscriptions.current_period_start, falling back
	// to the key's own created_at when the user has no subscription row
	// (free-tier users, who may never have one). The quota middleware derives
	// the *current* cycle boundary from this anchor's day-of-month, mirroring
	// the legacy Worker's getResetTime (apps/workers/auth-gateway/src/requests.ts)
	// — it is not itself the current cycle's start once more than one cycle
	// has elapsed since it was read.
	CurrentPeriodStart time.Time
}

// APIKeyPrincipalFromContext returns the principal APIKeyAuth stored on the
// request context, if the middleware ran and the key was valid.
func APIKeyPrincipalFromContext(ctx context.Context) (APIKeyPrincipal, bool) {
	p, ok := ctx.Value(apiKeyContextKey{}).(APIKeyPrincipal)
	return p, ok
}

// cachedAPIKey is the Redis-cached verification result, keyed by key_prefix
// (not by a hash of the raw key: Rails discards the raw key after creation,
// so a raw-key-derived cache key could never be invalidated on revoke).
//
// The cache entry is a *candidate*, not an authorization: it still carries
// the matched row's key_hash, and every request — cache hit or miss — must
// bcrypt-compare the presented key against it before the principal it names
// is used. Skipping that check on a hit would mean anyone who learns a
// 12-character key_prefix (not meant to be secret — it's what masked_key
// displays) could authenticate as that key's owner for the rest of the TTL
// without ever presenting the real secret.
type cachedAPIKey struct {
	UserID             int64     `json:"user_id"`
	Plan               string    `json:"plan"`
	Revoked            bool      `json:"revoked"`
	APIKeyID           int64     `json:"api_key_id"`
	CurrentPeriodStart time.Time `json:"current_period_start"`
	KeyHash            string    `json:"key_hash"`
}

// APIKeyAuth validates the requiems-api-key header via candidate-then-verify
// against Postgres api_keys.key_hash (bcrypt), caching the verified result in
// Redis by key_prefix. It fails closed: any lookup path that cannot
// affirmatively resolve the key (Redis error AND Postgres error, or no
// matching/non-revoked candidate) rejects the request.
type APIKeyAuth struct {
	pool *pgxpool.Pool
	rdb  *redis.Client
	ttl  time.Duration
}

// NewAPIKeyAuth builds an APIKeyAuth. ttl controls how long a verified
// key_prefix result is cached before a fresh Postgres lookup is required; it
// also bounds how long a revoked key can keep authenticating if Rails'
// revocation DEL (see apiKeyCacheKeyPrefix) fails, so keep it short.
func NewAPIKeyAuth(pool *pgxpool.Pool, rdb *redis.Client, ttl time.Duration) *APIKeyAuth {
	return &APIKeyAuth{pool: pool, rdb: rdb, ttl: ttl}
}

// Middleware returns the enforcing http middleware. See the package-level
// note on this file's mounting in app.go for what traffic it actually gates.
func (a *APIKeyAuth) Middleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			presented := r.Header.Get(apiKeyHeader)

			if len(presented) < keyPrefixLength {
				unauthorized(w)
				return
			}

			prefix := presented[:keyPrefixLength]

			result, found := a.lookupCache(r.Context(), prefix)

			// A cache hit only names a candidate key_hash; it is not itself
			// proof the presented key is correct. Verify it here, and fall
			// through to Postgres on a mismatch — this is what lets a
			// second key sharing the same prefix (a real, if rare,
			// collision) still resolve correctly instead of being rejected
			// or, worse, silently authenticated as the wrong key.
			if found && bcrypt.CompareHashAndPassword([]byte(result.KeyHash), []byte(presented)) != nil {
				found = false
			}

			if !found {
				var err error

				result, found, err = a.lookupDB(r.Context(), prefix, presented)

				if err != nil {
					unauthorized(w)
					return
				}

				if found {
					a.storeCache(r.Context(), prefix, result)
				}
			}

			if !found || result.Revoked {
				unauthorized(w)
				return
			}

			ctx := context.WithValue(r.Context(), apiKeyContextKey{}, APIKeyPrincipal{
				UserID:             result.UserID,
				Plan:               result.Plan,
				APIKeyID:           result.APIKeyID,
				CurrentPeriodStart: result.CurrentPeriodStart,
			})

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func unauthorized(w http.ResponseWriter) {
	httpx.Error(w, http.StatusUnauthorized, "unauthorized", "Invalid API key")
}

// lookupCache returns the cached result for prefix. A Redis error (including
// "unreachable") or cache miss both just report found=false, falling through
// to lookupDB — a Redis outage degrades to a direct Postgres lookup rather
// than failing the request outright.
func (a *APIKeyAuth) lookupCache(ctx context.Context, prefix string) (cachedAPIKey, bool) {
	if a.rdb == nil {
		return cachedAPIKey{}, false
	}

	val, err := a.rdb.Get(ctx, apiKeyCacheKeyPrefix+prefix).Result()
	if err != nil {
		return cachedAPIKey{}, false
	}

	var cached cachedAPIKey
	if err := json.Unmarshal([]byte(val), &cached); err != nil {
		return cachedAPIKey{}, false
	}

	return cached, true
}

func (a *APIKeyAuth) storeCache(ctx context.Context, prefix string, result cachedAPIKey) {
	if a.rdb == nil {
		return
	}

	b, err := json.Marshal(result)
	if err != nil {
		return
	}

	// Best-effort: a failed cache write just means the next request pays the
	// Postgres+bcrypt cost again, not a correctness issue.
	a.rdb.Set(ctx, apiKeyCacheKeyPrefix+prefix, string(b), a.ttl)
}

// lookupDB runs the candidate-then-verify query: select every api_keys row
// sharing prefix (in practice almost always exactly one, but never assumed),
// then bcrypt.compare the presented key against each candidate's key_hash
// until one matches. Plan is resolved from the user's active subscription,
// defaulting to "free" to match User#current_plan in Rails.
//
// Only a genuine query error (e.g. Postgres unreachable) is returned as an
// error; "no candidate matched" is a normal not-found result, not an error.
func (a *APIKeyAuth) lookupDB(ctx context.Context, prefix, presented string) (cachedAPIKey, bool, error) {
	rows, err := a.pool.Query(ctx, `
		SELECT api_keys.id, api_keys.key_hash, api_keys.user_id, api_keys.active, api_keys.revoked_at,
		       COALESCE(subscriptions.plan_name, 'free') AS plan,
		       COALESCE(subscriptions.current_period_start, api_keys.created_at) AS current_period_start
		FROM api_keys
		LEFT JOIN subscriptions ON subscriptions.user_id = api_keys.user_id
		WHERE api_keys.key_prefix = $1
	`, prefix)
	if err != nil {
		return cachedAPIKey{}, false, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			apiKeyID           int64
			keyHash            string
			userID             int64
			active             bool
			revokedAt          *time.Time
			plan               string
			currentPeriodStart time.Time
		)

		if err := rows.Scan(&apiKeyID, &keyHash, &userID, &active, &revokedAt, &plan, &currentPeriodStart); err != nil {
			return cachedAPIKey{}, false, err
		}

		if bcrypt.CompareHashAndPassword([]byte(keyHash), []byte(presented)) != nil {
			continue
		}

		return cachedAPIKey{
			UserID:             userID,
			Plan:               plan,
			Revoked:            !active || revokedAt != nil,
			APIKeyID:           apiKeyID,
			CurrentPeriodStart: currentPeriodStart,
			KeyHash:            keyHash,
		}, true, rows.Err()
	}

	if err := rows.Err(); err != nil {
		return cachedAPIKey{}, false, err
	}

	return cachedAPIKey{}, false, nil
}

package middleware

import (
	"bytes"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
)

// usageHarness wires a UsageQuota to the test Postgres/Redis with a user on a
// roomy plan, and exposes the two things these tests compare: the Redis
// counter that enforces the quota and the usage_logs ledger.
type usageHarness struct {
	t         *testing.T
	pool      *pgxpool.Pool
	rdb       *redis.Client
	uq        *UsageQuota
	logs      *bytes.Buffer
	principal APIKeyPrincipal
	key       string
}

type ledgerRow struct {
	Endpoint string
	Status   int
	Credits  *int
	UsedAt   time.Time
	Latency  int
}

func newUsageHarness(t *testing.T) *usageHarness {
	t.Helper()

	return newUsageHarnessWithLimit(t, 100000)
}

func newUsageHarnessWithLimit(t *testing.T, limit int64) *usageHarness {
	t.Helper()

	pool := setupUsageTestDB(t)
	rdb := setupAPIKeyAuthTestRedis(t)

	planID := "usage-err-plan-" + randomHex(t, 4)
	insertPlan(t, pool, planID, ptr(limit), nil)

	logs := &bytes.Buffer{}
	uq := NewUsageQuota(pool, rdb, NewPlanCache(pool), slog.New(slog.NewTextHandler(logs, nil)))

	userID := time.Now().UnixNano()
	principal := APIKeyPrincipal{UserID: userID, Plan: planID, APIKeyID: userID, CurrentPeriodStart: time.Now()}
	key := fmt.Sprintf("usage:%d:%d", userID, cycleStart(principal.CurrentPeriodStart, time.Now()).Unix())

	t.Cleanup(func() {
		_, _ = pool.Exec(context.Background(), `DELETE FROM usage_logs WHERE user_id = $1`, userID)
		rdb.Del(context.Background(), key)

		for _, day := range []time.Time{time.Now().UTC(), time.Now().UTC().AddDate(0, 0, -1)} {
			rdb.HDel(context.Background(), "route_miss:"+day.Format("2006-01-02"), strconv.FormatInt(userID, 10))
		}
	})

	return &usageHarness{t: t, pool: pool, rdb: rdb, uq: uq, logs: logs, principal: principal, key: key}
}

func (h *usageHarness) request(method, path string) *http.Request {
	req := httptest.NewRequest(method, path, http.NoBody)
	return req.WithContext(context.WithValue(req.Context(), apiKeyContextKey{}, h.principal))
}

func (h *usageHarness) serve(handler http.Handler, req *http.Request) *httptest.ResponseRecorder {
	h.t.Helper()

	w := httptest.NewRecorder()
	h.uq.Middleware()(handler).ServeHTTP(w, req)

	return w
}

// warmUp makes the quota counter exist so later assertions compare against a
// real pre-request value (a cold key would be bootstrapped from the ledger).
func (h *usageHarness) warmUp() int64 {
	h.t.Helper()

	w := h.serve(okHandler(), h.request(http.MethodGet, "/v1/text/warmup"))
	require.Equal(h.t, http.StatusOK, w.Code)

	return h.counter()
}

func (h *usageHarness) counter() int64 {
	h.t.Helper()

	raw, err := h.rdb.Get(context.Background(), h.key).Result()
	if err == redis.Nil {
		return -1
	}
	require.NoError(h.t, err)

	n, err := strconv.ParseInt(raw, 10, 64)
	require.NoError(h.t, err)

	return n
}

func (h *usageHarness) ledger() []ledgerRow {
	h.t.Helper()

	rows, err := h.pool.Query(context.Background(),
		`SELECT endpoint, status_code, credits_used, used_at, COALESCE(response_time_ms, 0) FROM usage_logs WHERE user_id = $1 ORDER BY id`, h.principal.UserID)
	require.NoError(h.t, err)
	defer rows.Close()

	var out []ledgerRow
	for rows.Next() {
		var r ledgerRow
		require.NoError(h.t, rows.Scan(&r.Endpoint, &r.Status, &r.Credits, &r.UsedAt, &r.Latency))
		out = append(out, r)
	}
	require.NoError(h.t, rows.Err())

	return out
}

func (h *usageHarness) last() ledgerRow {
	h.t.Helper()

	rows := h.ledger()
	require.NotEmpty(h.t, rows)

	return rows[len(rows)-1]
}

// routeMissCount sums today's and yesterday's buckets so a run that crosses
// midnight UTC still sees what it wrote.
func (h *usageHarness) routeMissCount() int64 {
	h.t.Helper()

	var total int64

	for _, day := range []time.Time{time.Now().UTC(), time.Now().UTC().AddDate(0, 0, -1)} {
		raw, err := h.rdb.HGet(context.Background(), "route_miss:"+day.Format("2006-01-02"), strconv.FormatInt(h.principal.UserID, 10)).Result()
		if err == redis.Nil {
			continue
		}
		require.NoError(h.t, err)

		n, err := strconv.ParseInt(raw, 10, 64)
		require.NoError(h.t, err)

		total += n
	}

	return total
}

func respondWith(status int, headers map[string]string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(status)
	})
}

func creditsOf(t *testing.T, row ledgerRow) int {
	t.Helper()

	require.NotNil(t, row.Credits)

	return *row.Credits
}

func TestUsageQuota_ErrorsAreLoggedButNeverCharged(t *testing.T) {
	for _, status := range []int{400, 401, 404, 409, 422, 429, 500, 502, 503} {
		t.Run(strconv.Itoa(status), func(t *testing.T) {
			h := newUsageHarness(t)
			before := h.warmUp()

			w := h.serve(respondWith(status, nil), h.request(http.MethodGet, "/v1/places/XX"))

			require.Equal(t, status, w.Code)
			require.Equal(t, before, h.counter(), "the reserved credit must be refunded")

			row := h.last()
			require.Equal(t, status, row.Status, "the error is still logged")
			require.Equal(t, 0, creditsOf(t, row))
			require.Len(t, h.ledger(), 2)
		})
	}
}

func TestUsageQuota_ErrorOverridesUsageCountHeader(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	h.serve(respondWith(http.StatusUnprocessableEntity, map[string]string{"X-Usage-Count": "7"}), h.request(http.MethodPost, "/v1/text/words/batch"))

	require.Equal(t, before, h.counter())
	require.Equal(t, 0, creditsOf(t, h.last()))
}

func TestUsageQuota_DictionaryErrorRefundsBothReservedCredits(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	h.serve(respondWith(http.StatusUnprocessableEntity, nil), h.request(http.MethodGet, "/v1/text/dictionary/word"))

	require.Equal(t, before, h.counter(), "requested 2, response 422: the counter drops by 2")
	require.Equal(t, 0, creditsOf(t, h.last()))
}

func TestUsageQuota_SuccessStillChargesAndRedirectsStillCharge(t *testing.T) {
	t.Run("dictionary success bills 2", func(t *testing.T) {
		h := newUsageHarness(t)
		before := h.warmUp()

		h.serve(respondWith(http.StatusOK, map[string]string{"X-Usage-Count": "2"}), h.request(http.MethodGet, "/v1/text/dictionary/word"))

		require.Equal(t, before+2, h.counter())
		require.Equal(t, 2, creditsOf(t, h.last()))
	})

	t.Run("batch success bills N", func(t *testing.T) {
		h := newUsageHarness(t)
		before := h.warmUp()

		h.serve(respondWith(http.StatusOK, map[string]string{"X-Usage-Count": "5"}), h.request(http.MethodPost, "/v1/text/words/batch"))

		require.Equal(t, before+5, h.counter())
		require.Equal(t, 5, creditsOf(t, h.last()))
	})

	t.Run("3xx is charged as before", func(t *testing.T) {
		h := newUsageHarness(t)
		before := h.warmUp()

		h.serve(respondWith(http.StatusFound, nil), h.request(http.MethodGet, "/v1/text/advice"))

		require.Equal(t, before+1, h.counter())
		row := h.last()
		require.Equal(t, http.StatusFound, row.Status)
		require.Equal(t, 1, creditsOf(t, row))
	})
}

func TestUsageQuota_EmptyBatchIsFree(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	h.serve(respondWith(http.StatusOK, map[string]string{"X-Usage-Count": "0"}), h.request(http.MethodPost, "/v1/text/words/batch"))

	require.Equal(t, before, h.counter(), "0 items: the reserved credit is returned")
	row := h.last()
	require.Equal(t, http.StatusOK, row.Status)
	require.Equal(t, 0, creditsOf(t, row))
}

func TestUsageQuota_HandlerThatNeverWritesIsLoggedAs200(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	w := h.serve(http.HandlerFunc(func(http.ResponseWriter, *http.Request) {}), h.request(http.MethodGet, "/v1/text/advice"))

	require.Equal(t, http.StatusOK, w.Code)
	require.Equal(t, before+1, h.counter())
	row := h.last()
	require.Equal(t, http.StatusOK, row.Status, "raw status 0 must be normalised")
	require.Equal(t, 1, creditsOf(t, row))
}

func TestUsageQuota_PanicBeforeAnyWriteReturns500AndRefunds(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	w := h.serve(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { panic("boom") }), h.request(http.MethodGet, "/v1/text/advice"))

	require.Equal(t, http.StatusInternalServerError, w.Code)
	require.Contains(t, w.Body.String(), "internal_error")
	require.Contains(t, w.Header().Get("Content-Type"), "application/json")
	require.Equal(t, before, h.counter(), "a recovered panic leaves the counter at its pre-request value (single refund)")

	row := h.last()
	require.Equal(t, http.StatusInternalServerError, row.Status)
	require.Equal(t, 0, creditsOf(t, row))
	require.Contains(t, h.logs.String(), "handler panicked")
}

func TestUsageQuota_PanicAfterWriteHeaderKeepsClientStatusButLogs500(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	w := h.serve(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("X-Usage-Count", "4")
		w.WriteHeader(http.StatusOK)
		panic("late boom")
	}), h.request(http.MethodGet, "/v1/text/advice"))

	require.Equal(t, http.StatusOK, w.Code, "the committed response cannot be changed")
	require.Equal(t, before, h.counter())

	row := h.last()
	require.Equal(t, http.StatusInternalServerError, row.Status)
	require.Equal(t, 0, creditsOf(t, row))
}

func TestUsageQuota_AbortHandlerPanicIsAccountedThenRepanicked(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	require.PanicsWithValue(t, http.ErrAbortHandler, func() {
		h.serve(http.HandlerFunc(func(http.ResponseWriter, *http.Request) { panic(http.ErrAbortHandler) }), h.request(http.MethodGet, "/v1/text/advice"))
	})

	require.Equal(t, before, h.counter())
	row := h.last()
	require.Equal(t, http.StatusInternalServerError, row.Status)
	require.Equal(t, 0, creditsOf(t, row))
}

func TestUsageQuota_RefundSurvivesClientDisconnect(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	ctx, cancel := context.WithCancel(h.request(http.MethodGet, "/v1/text/advice").Context())
	req := h.request(http.MethodGet, "/v1/text/advice").WithContext(ctx)

	h.serve(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		cancel() // the client went away
		w.WriteHeader(http.StatusInternalServerError)
	}), req)

	require.Equal(t, before, h.counter(), "refund must not depend on the request context")
	require.Equal(t, 0, creditsOf(t, h.last()))
}

func TestUsageQuota_FailedRefundIsLoggedAndColdBootstrapSelfHeals(t *testing.T) {
	h := newUsageHarness(t)
	h.warmUp()

	// The counter expires mid-request, so the refund finds no key.
	h.serve(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		h.rdb.Del(context.Background(), h.key)
		w.WriteHeader(http.StatusUnprocessableEntity)
	}), h.request(http.MethodGet, "/v1/text/advice"))

	require.Contains(t, h.logs.String(), "failed to reconcile Redis credits", "a failed refund is logged, never silent")
	require.Equal(t, int64(-1), h.counter())

	// The next request bootstraps from the ledger (1 credit warm-up + 0 for the
	// error) and adds its own credit.
	h.serve(okHandler(), h.request(http.MethodGet, "/v1/text/advice"))
	require.Equal(t, int64(2), h.counter())
}

func TestUsageQuota_UsedAtIsTheRequestStart(t *testing.T) {
	h := newUsageHarness(t)
	h.warmUp()

	start := time.Now()
	h.serve(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		time.Sleep(400 * time.Millisecond)
		w.WriteHeader(http.StatusOK)
	}), h.request(http.MethodGet, "/v1/text/slow"))

	row := h.last()
	require.GreaterOrEqual(t, row.Latency, 350, "latency is the handler's own time")
	require.Less(t, row.UsedAt.Sub(start), 200*time.Millisecond, "used_at is the request start, not when the handler finished")
}

func TestUsageQuota_RouteMissSkipsQuotaLedgerAndHandler(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	var matched []string
	h.uq.SetRouteMatcher(func(method, path string) bool {
		matched = append(matched, method+" "+path)
		return false
	})

	reached := false
	w := h.serve(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		reached = true
		w.WriteHeader(http.StatusNotFound)
	}), h.request(http.MethodGet, "/v1/nonexistent"))

	require.True(t, reached, "the router still answers the request")
	require.Equal(t, http.StatusNotFound, w.Code)
	require.Equal(t, before, h.counter(), "no reservation for a route that does not exist")
	require.Len(t, h.ledger(), 1, "no ledger row for a route miss")
	require.Equal(t, []string{"GET /v1/nonexistent"}, matched)
	require.EqualValues(t, 1, h.routeMissCount(), "scanner traffic stays visible in a per-user bucket")
}

func TestUsageQuota_RouteMissSurvivesARedisOutage(t *testing.T) {
	h := newUsageHarness(t)
	h.uq.rdb = unreachableRedis(t)
	h.uq.SetRouteMatcher(func(string, string) bool { return false })

	reached := false
	started := time.Now()
	w := h.serve(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		reached = true
		w.WriteHeader(http.StatusNotFound)
	}), h.request(http.MethodGet, "/v1/nonexistent"))

	require.True(t, reached)
	require.Equal(t, http.StatusNotFound, w.Code, "reserving would have returned 503; a miss reserves nothing, and a failing bucket never fails the request")
	require.Less(t, time.Since(started), 2*time.Second)
	require.Empty(t, h.ledger())
}

func TestUsageQuota_RouteMissIsNotA429ForAnOverQuotaUser(t *testing.T) {
	h := newUsageHarnessWithLimit(t, 2)
	h.serve(okHandler(), h.request(http.MethodGet, "/v1/text/a"))
	h.serve(okHandler(), h.request(http.MethodGet, "/v1/text/b"))
	require.Equal(t, int64(2), h.counter())

	over := h.serve(okHandler(), h.request(http.MethodGet, "/v1/text/c"))
	require.Equal(t, http.StatusTooManyRequests, over.Code, "a real route past the limit is rejected")

	h.uq.SetRouteMatcher(func(string, string) bool { return false })
	w := h.serve(respondWith(http.StatusNotFound, nil), h.request(http.MethodGet, "/v1/typo"))

	require.Equal(t, http.StatusNotFound, w.Code, "a typo must not turn into a 429")
	require.Equal(t, int64(2), h.counter())
	require.Len(t, h.ledger(), 2)
}

// routedHarness serves a real chi router (middleware at the root, one mounted
// service) so the pre-match is exercised against actual chi dispatch.
func (h *usageHarness) routedHarness() *chi.Mux {
	root := chi.NewRouter()
	root.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), apiKeyContextKey{}, h.principal)))
		})
	})
	root.Use(h.uq.Middleware())

	svc := chi.NewRouter()
	svc.NotFound(NotFoundHandler())
	svc.MethodNotAllowed(MethodNotAllowedHandler())
	svc.Get("/echo/{word}", func(w http.ResponseWriter, _ *http.Request) { w.WriteHeader(http.StatusOK) })
	root.Mount("/v1/svc", svc)

	root.NotFound(NotFoundHandler())
	root.MethodNotAllowed(MethodNotAllowedHandler())

	h.uq.SetRouteMatcher(func(method, path string) bool {
		return root.Match(chi.NewRouteContext(), method, path)
	})

	return root
}

func TestUsageQuota_PreMatchAgreesWithChiDispatch(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()
	root := h.routedHarness()

	do := func(method, target string) *httptest.ResponseRecorder {
		w := httptest.NewRecorder()
		root.ServeHTTP(w, httptest.NewRequest(method, target, http.NoBody))

		return w
	}

	t.Run("a real route with an encoded slash is recorded and charged", func(t *testing.T) {
		w := do(http.MethodGet, "/v1/svc/echo/a%2Fb")

		require.Equal(t, http.StatusOK, w.Code)
		require.Equal(t, before+1, h.counter())
		require.Equal(t, "/v1/svc/echo/a/b", h.last().Endpoint)
	})

	rowsBefore := len(h.ledger())
	counterBefore := h.counter()

	misses := []struct {
		name, method, target string
		status               int
		allow                string
	}{
		{"unknown route", http.MethodGet, "/v1/svc/zzz", http.StatusNotFound, ""},
		{"trailing slash is a miss like chi", http.MethodGet, "/v1/svc/echo/x/", http.StatusNotFound, ""},
		{"missing parameter", http.MethodGet, "/v1/svc/echo", http.StatusNotFound, ""},
		{"wrong method", http.MethodPost, "/v1/svc/echo/x", http.StatusMethodNotAllowed, "GET"},
		{"HEAD on a GET route", http.MethodHead, "/v1/svc/echo/x", http.StatusMethodNotAllowed, "GET"},
		{"unknown top-level", http.MethodGet, "/v1/nothing", http.StatusNotFound, ""},
	}

	for _, tt := range misses {
		t.Run(tt.name, func(t *testing.T) {
			w := do(tt.method, tt.target)

			require.Equal(t, tt.status, w.Code)
			require.Equal(t, tt.allow, w.Header().Get("Allow"), "405 keeps its Allow header")
			require.Equal(t, counterBefore, h.counter(), "no quota movement")
			require.Len(t, h.ledger(), rowsBefore, "no ledger row")
		})
	}

	require.EqualValues(t, len(misses), h.routeMissCount())
}

func TestUsageQuota_ConcurrentRequestsAndMatcherSwaps(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()
	root := h.routedHarness()
	matcher := *h.uq.matcher.Load()

	const workers, perWorker = 8, 5

	var wg sync.WaitGroup

	stop := make(chan struct{})
	swapDone := make(chan struct{})

	go func() {
		defer close(swapDone)

		for {
			select {
			case <-stop:
				return
			default:
				h.uq.SetRouteMatcher(nil) // fail open for a moment...
				h.uq.SetRouteMatcher(matcher)
			}
		}
	}()

	for w := range workers {
		wg.Add(1)

		go func() {
			defer wg.Done()

			for i := range perWorker {
				rec := httptest.NewRecorder()
				root.ServeHTTP(rec, httptest.NewRequest(http.MethodGet, fmt.Sprintf("/v1/svc/echo/w%d-%d", w, i), http.NoBody))
				if rec.Code != http.StatusOK {
					t.Errorf("worker %d request %d: status %d", w, i, rec.Code)
				}
			}
		}()
	}

	wg.Wait()
	close(stop)
	<-swapDone

	require.Equal(t, before+workers*perWorker, h.counter())
	require.Len(t, h.ledger(), 1+workers*perWorker)
}

func TestUsageQuota_MethodNotAllowedHandlerWithoutAMatcherStillAnswers405(t *testing.T) {
	w := httptest.NewRecorder()
	MethodNotAllowedHandler()(w, httptest.NewRequest(http.MethodPost, "/v1/x", http.NoBody))

	require.Equal(t, http.StatusMethodNotAllowed, w.Code)
	require.Empty(t, w.Header().Get("Allow"))
}

func TestUsageQuota_MatcherUsesRawPathWhenSet(t *testing.T) {
	h := newUsageHarness(t)

	var got string
	h.uq.SetRouteMatcher(func(_, path string) bool {
		got = path
		return true
	})

	req := h.request(http.MethodGet, "/v1/text/a%2Fb")
	require.Equal(t, "/v1/text/a/b", req.URL.Path)
	require.Equal(t, "/v1/text/a%2Fb", req.URL.RawPath)

	h.serve(okHandler(), req)

	require.Equal(t, "/v1/text/a%2Fb", got, "chi routes on RawPath, so the matcher must too")
}

func TestUsageQuota_MatcherFailsOpen(t *testing.T) {
	t.Run("no matcher set", func(t *testing.T) {
		h := newUsageHarness(t)
		before := h.warmUp()

		h.serve(okHandler(), h.request(http.MethodGet, "/v1/text/advice"))

		require.Equal(t, before+1, h.counter())
		require.Len(t, h.ledger(), 2)
	})

	t.Run("matcher panics", func(t *testing.T) {
		h := newUsageHarness(t)
		before := h.warmUp()
		h.uq.SetRouteMatcher(func(string, string) bool { panic("matcher bug") })

		h.serve(okHandler(), h.request(http.MethodGet, "/v1/text/advice"))

		require.Equal(t, before+1, h.counter(), "a broken matcher must not drop real usage")
		require.Len(t, h.ledger(), 2)
	})

	t.Run("matcher cleared again", func(t *testing.T) {
		h := newUsageHarness(t)
		h.uq.SetRouteMatcher(func(string, string) bool { return false })
		h.uq.SetRouteMatcher(nil)
		before := h.warmUp()

		require.Equal(t, int64(1), before)
		require.Len(t, h.ledger(), 1)
	})
}

func TestUsageQuota_FlaggedRouteMissIsRefundedWithoutALedgerRow(t *testing.T) {
	for name, handler := range map[string]http.Handler{
		"404": NotFoundHandler(),
		"405": MethodNotAllowedHandler(),
	} {
		t.Run(name, func(t *testing.T) {
			h := newUsageHarness(t)
			before := h.warmUp()

			w := h.serve(handler, h.request(http.MethodGet, "/v1/text/nonexistent"))

			expected := http.StatusNotFound
			if name == "405" {
				expected = http.StatusMethodNotAllowed
			}

			require.Equal(t, expected, w.Code)
			require.Equal(t, before, h.counter(), "the fallback refunds the reservation")
			require.Len(t, h.ledger(), 1, "and writes no ledger row")
			require.EqualValues(t, 1, h.routeMissCount())
		})
	}
}

func TestUsageQuota_HandlerLevel404IsLoggedNotFlagged(t *testing.T) {
	h := newUsageHarness(t)
	before := h.warmUp()

	h.serve(respondWith(http.StatusNotFound, nil), h.request(http.MethodGet, "/v1/places/ZZ"))

	require.Equal(t, before, h.counter())
	row := h.last()
	require.Equal(t, http.StatusNotFound, row.Status, "a real route that answered 404 is logged (free)")
	require.Equal(t, 0, creditsOf(t, row))
	require.EqualValues(t, 0, h.routeMissCount())
}

func TestHeaderCredits_ZeroIsDistinctFromMissing(t *testing.T) {
	zero := http.Header{}
	zero.Set("X-Usage-Count", "0")
	require.EqualValues(t, 0, headerCredits(zero))

	require.EqualValues(t, 1, headerCredits(http.Header{}))

	negative := http.Header{}
	negative.Set("X-Usage-Count", "-1")
	require.EqualValues(t, 1, headerCredits(negative))
}

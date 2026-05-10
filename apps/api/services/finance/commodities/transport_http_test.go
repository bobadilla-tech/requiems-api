package commodities

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// stubGetterHTTP implements Getter for transport tests.
type stubGetterHTTP struct {
	result CommodityPrice
	err    error
}

func (s *stubGetterHTTP) Get(_ context.Context, slug string) (CommodityPrice, error) {
	if s.err != nil {
		return CommodityPrice{}, s.err
	}
	r := s.result
	r.Commodity = slug
	return r, nil
}

func setupRouter(g Getter) chi.Router {
	r := chi.NewRouter()
	registerCommodityRoutes(r, g)
	return r
}

func decodeResponse(t *testing.T, w *httptest.ResponseRecorder) httpx.Response[CommodityPrice] {
	t.Helper()
	var resp httpx.Response[CommodityPrice]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

// ---- tests ----

func TestCommodity_KnownSlug_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubGetterHTTP{result: CommodityPrice{
		Name:      "Gold",
		Price:     2386.33,
		Unit:      "oz",
		Currency:  "USD",
		Change24h: 23.01,
		Historical: []HistoricalPrice{
			{Period: "2023", Price: 1940.54},
		},
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/commodities/gold", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d: %s", w.Code, w.Body.String())

	resp := decodeResponse(t, w)
	assert.Equal(t, "gold", resp.Data.Commodity)
	assert.Equal(t, 2386.33, resp.Data.Price)
	assert.NotEmpty(t, resp.Metadata.Timestamp, "expected metadata.timestamp to be set")
}

func TestCommodity_ResponseEnvelope(t *testing.T) {
	t.Parallel()
	svc := &stubGetterHTTP{result: CommodityPrice{Price: 2386.33}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/commodities/gold", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code, "expected 200, got %d", w.Code)

	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)
	assert.Contains(t, raw, "data", "response must have a 'data' key")
	assert.Contains(t, raw, "metadata", "response must have a 'metadata' key")
}

func TestCommodity_UnknownSlug_Returns404(t *testing.T) {
	t.Parallel()
	svc := &stubGetterHTTP{err: svcerr.NotFound("not_found", "commodity not found")}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/commodities/unobtainium", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code, "expected 404, got %d: %s", w.Code, w.Body.String())
}

func TestCommodity_InternalError_Returns500(t *testing.T) {
	t.Parallel()
	svc := &stubGetterHTTP{err: errors.New("internal server error")}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/commodities/gold", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code, "expected 500, got %d: %s", w.Code, w.Body.String())
}

func TestCommodity_HistoricalFieldPresent(t *testing.T) {
	t.Parallel()
	svc := &stubGetterHTTP{result: CommodityPrice{
		Price: 2386.33,
		Historical: []HistoricalPrice{
			{Period: "2023", Price: 1940.54},
			{Period: "2022", Price: 1800.12},
		},
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/commodities/gold", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.Len(t, resp.Data.Historical, 2)
}

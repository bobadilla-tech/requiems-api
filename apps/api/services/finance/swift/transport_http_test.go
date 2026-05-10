package swift

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

// stubService implements Looker for transport tests. It returns a fixed result
// or a fixed error on every call, keeping tests DB-free and fast.
type stubService struct {
	result LookupResponse
	list   ListResponse
	err    error
}

func (s *stubService) Lookup(_ context.Context, code string) (LookupResponse, error) {
	if s.err != nil {
		return LookupResponse{}, s.err
	}
	r := s.result
	r.SwiftCode = code
	return r, nil
}

func (s *stubService) List(_ context.Context, _ ListFilter) (ListResponse, error) {
	if s.err != nil {
		return ListResponse{}, s.err
	}
	return s.list, nil
}

// setupRouter wires up a stub service into a chi router for handler testing.
func setupRouter(svc Looker) chi.Router {
	r := chi.NewRouter()
	registerSWIFTRoutes(r, svc)
	return r
}

// ---- helper ----

func decodeResponse(t *testing.T, body *httptest.ResponseRecorder) httpx.Response[LookupResponse] {
	t.Helper()
	var resp httpx.Response[LookupResponse]
	err := json.NewDecoder(body.Body).Decode(&resp)
	require.NoError(t, err)
	return resp
}

// ---- tests ----

func TestSWIFT_KnownCode_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{
		BankCode:     "DEUT",
		CountryCode:  "DE",
		LocationCode: "DB",
		BranchCode:   "XXX",
		BankName:     "Deutsche Bank",
		City:         "Frankfurt",
		CountryName:  "Germany",
		IsPrimary:    true,
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/swift/DEUTDEDBXXX", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	assert.Equal(t, "DEUTDEDBXXX", resp.Data.SwiftCode)
	assert.Equal(t, "Deutsche Bank", resp.Data.BankName)
}

func TestSWIFT_List_Returns200(t *testing.T) {
	t.Parallel()
	svc := &stubService{list: ListResponse{
		Items: []LookupResponse{
			{SwiftCode: "DEUTDEDBXXX", BankCode: "DEUT", CountryCode: "DE", BankName: "Deutsche Bank"},
			{SwiftCode: "CHASUS33XXX", BankCode: "CHAS", CountryCode: "US", BankName: "JPMorgan Chase"},
		},
		Limit:    50,
		Offset:   0,
		Returned: 2,
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/swift?country_code=DE&limit=50&offset=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ListResponse]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	require.Equal(t, 2, resp.Data.Returned)
}

func TestSWIFT_List_InvalidLimit_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubService{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift?limit=abc", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSWIFT_List_InvalidOffset_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubService{}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift?offset=abc", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSWIFT_List_ServiceAppError_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: svcerr.Invalid("bad_request", "invalid filter")}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift?country_code=D1", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSWIFT_List_ServiceGenericError_Returns500(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: errors.New("db down")}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSWIFT_ResponseEnvelope(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{BankCode: "DEUT"}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift/DEUTDEDBXXX", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)
	_, ok := raw["data"]
	assert.True(t, ok, "response must have a 'data' key")
	_, ok = raw["metadata"]
	assert.True(t, ok, "response must have a 'metadata' key")
}

func TestSWIFT_MetadataTimestamp(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift/CHASUS33XXX", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.NotEmpty(t, resp.Metadata.Timestamp)
}

func TestSWIFT_UnknownCode_Returns404(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: svcerr.NotFound("not_found", "SWIFT code not found")}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/swift/ZZZZZZZZXXX", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

func TestSWIFT_BadFormat_Returns400(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: svcerr.Invalid("bad_request", "SWIFT code must be 8 or 11 characters")}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/swift/INVALID", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSWIFT_DBError_Returns500(t *testing.T) {
	t.Parallel()
	svc := &stubService{err: errors.New("db unavailable")}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/swift/DEUTDEDBXXX", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestSWIFT_PrimaryOffice_IsPrimaryTrue(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{BranchCode: "XXX", IsPrimary: true}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift/DEUTDEDBXXX", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.True(t, resp.Data.IsPrimary)
}

func TestSWIFT_BranchOffice_IsPrimaryFalse(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{BranchCode: "001", IsPrimary: false}}
	r := setupRouter(svc)

	req := httptest.NewRequest(http.MethodGet, "/swift/DEUTDEDB001", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	resp := decodeResponse(t, w)
	assert.False(t, resp.Data.IsPrimary)
}

func TestSWIFT_AllResponseFieldsPresent(t *testing.T) {
	t.Parallel()
	svc := &stubService{result: LookupResponse{
		BankCode:     "CHAS",
		CountryCode:  "US",
		LocationCode: "33",
		BranchCode:   "XXX",
		BankName:     "JPMorgan Chase Bank",
		City:         "New York",
		CountryName:  "United States",
		IsPrimary:    true,
	}}

	r := setupRouter(svc)
	req := httptest.NewRequest(http.MethodGet, "/swift/CHASUS33XXX", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	resp := decodeResponse(t, w)
	d := resp.Data

	checks := map[string]bool{
		"swift_code non-empty":    d.SwiftCode != "",
		"bank_code non-empty":     d.BankCode != "",
		"country_code non-empty":  d.CountryCode != "",
		"location_code non-empty": d.LocationCode != "",
		"branch_code non-empty":   d.BranchCode != "",
		"bank_name non-empty":     d.BankName != "",
		"city non-empty":          d.City != "",
		"country_name non-empty":  d.CountryName != "",
		"is_primary is true":      d.IsPrimary,
	}

	for name, ok := range checks {
		assert.True(t, ok, "field check failed: %s", name)
	}
}

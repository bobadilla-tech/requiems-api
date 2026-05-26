package timezoneip

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	ipinfo "requiems-api/services/networking/ip/info"
	"requiems-api/services/places/timezone"
)

type stubIPInfo struct {
	r   ipinfo.LookupResponse
	err error
}

func (s *stubIPInfo) CheckInfo(_ context.Context, _ string) (ipinfo.LookupResponse, error) {
	return s.r, s.err
}

type stubTimezone struct {
	r   *timezone.Info
	err error
}

func (s *stubTimezone) GetTimezoneByCoords(_, _ float64) (*timezone.Info, error) { return s.r, s.err }
func (s *stubTimezone) GetTimezoneByCity(_ string) (*timezone.Info, error)       { return s.r, s.err }

func setupRouter(i *stubIPInfo, t *stubTimezone) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService(i, t))
	return r
}

func get(t *testing.T, r chi.Router, path string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestTimezoneIP_ValidPublicIP(t *testing.T) {
	t.Parallel()
	tzInfo := &timezone.Info{
		Timezone:    "Europe/Berlin",
		Offset:      "+02:00",
		CurrentTime: "2026-05-25T19:00:00Z",
		IsDST:       true,
	}
	r := setupRouter(
		&stubIPInfo{r: ipinfo.LookupResponse{IP: "203.0.113.42", City: "Berlin", CountryCode: "DE"}},
		&stubTimezone{r: tzInfo},
	)
	w := get(t, r, "/timezone/from-ip/203.0.113.42")
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	require.NotNil(t, resp.Data.Timezone)
	assert.Equal(t, "Europe/Berlin", *resp.Data.Timezone)
	assert.Equal(t, "DE", resp.Data.CountryCode)
	require.NotNil(t, resp.Data.CurrentTime)
}

func TestTimezoneIP_UnknownIP(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubIPInfo{err: fmt.Errorf("not found")},
		&stubTimezone{},
	)
	w := get(t, r, "/timezone/from-ip/192.0.2.1")
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestTimezoneIP_IPFoundTimezoneFailure(t *testing.T) {
	t.Parallel()
	r := setupRouter(
		&stubIPInfo{r: ipinfo.LookupResponse{City: "Nowhere", CountryCode: "ZZ"}},
		&stubTimezone{err: fmt.Errorf("unknown city")},
	)
	w := get(t, r, "/timezone/from-ip/1.2.3.4")
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Nil(t, resp.Data.Timezone)
}

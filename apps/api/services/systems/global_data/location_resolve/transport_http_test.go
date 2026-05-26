package locationresolve

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/services/places/geocode"
	"requiems-api/services/places/holidays"
	"requiems-api/services/places/timezone"
)

type stubGeocoder struct {
	geocodeR   geocode.GeocodeResponse
	geocodeErr error
	revR       geocode.ReverseGeocodeResponse
	revErr     error
}

func (s *stubGeocoder) Geocode(_ context.Context, _ string) (geocode.GeocodeResponse, error) {
	return s.geocodeR, s.geocodeErr
}

func (s *stubGeocoder) ReverseGeocode(_ context.Context, _, _ float64) (geocode.ReverseGeocodeResponse, error) {
	return s.revR, s.revErr
}

type stubTimezone struct {
	r   *timezone.Info
	err error
}

func (s *stubTimezone) GetTimezoneByCoords(_, _ float64) (*timezone.Info, error) { return s.r, s.err }

type stubHolidays struct {
	r   holidays.HolidayList
	err error
}

func (s *stubHolidays) GetHolidays(_ string, _ int) (holidays.HolidayList, error) {
	return s.r, s.err
}

type stubWorkingDays struct{ n int }

func (s *stubWorkingDays) GetWorkingDays(_, _ time.Time, _, _ string) int { return s.n }

func buildSvc(g *stubGeocoder, tz *stubTimezone, h *stubHolidays, w *stubWorkingDays) *Service {
	return NewService(g, tz, h, w)
}

func setupLocationRouter(g *stubGeocoder, tz *stubTimezone, h *stubHolidays, w *stubWorkingDays) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, buildSvc(g, tz, h, w))
	return r
}

func post(t *testing.T, r chi.Router, body any) *httptest.ResponseRecorder {
	t.Helper()
	b, err := json.Marshal(body)
	require.NoError(t, err)
	req := httptest.NewRequest(http.MethodPost, "/location/resolve", bytes.NewReader(b))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestLocationResolve_ByAddress(t *testing.T) {
	t.Parallel()
	tzInfo := &timezone.Info{Timezone: "Europe/Paris", Offset: "+02:00", CurrentTime: "2026-05-26T12:00:00Z"}
	r := setupLocationRouter(
		&stubGeocoder{geocodeR: geocode.GeocodeResponse{
			Address: "Paris, France", City: "Paris", Country: "FR", Lat: 48.85, Lon: 2.35,
		}},
		&stubTimezone{r: tzInfo},
		&stubHolidays{r: holidays.HolidayList{Holidays: []holidays.Holiday{{Date: "2026-07-14", Name: "Bastille Day"}}}},
		&stubWorkingDays{n: 21},
	)
	w := post(t, r, map[string]any{"address": "Paris, France"})
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "Paris, France", resp.Data.Address)
	assert.Equal(t, "Paris", resp.Data.City)
	require.NotNil(t, resp.Data.Timezone)
	assert.Equal(t, "Europe/Paris", *resp.Data.Timezone)
	assert.Equal(t, 21, resp.Data.WorkingDaysThisMonth)
	assert.Empty(t, resp.Data.Flags)
}

func TestLocationResolve_ByCoordinates(t *testing.T) {
	t.Parallel()
	tzInfo := &timezone.Info{Timezone: "America/New_York", Offset: "-04:00", CurrentTime: "2026-05-26T08:00:00Z"}
	r := setupLocationRouter(
		&stubGeocoder{revR: geocode.ReverseGeocodeResponse{
			Address: "New York, US", City: "New York", Country: "US", Lat: 40.71, Lon: -74.00,
		}},
		&stubTimezone{r: tzInfo},
		&stubHolidays{r: holidays.HolidayList{Holidays: []holidays.Holiday{}}},
		&stubWorkingDays{n: 20},
	)
	w := post(t, r, map[string]any{
		"coordinates": map[string]any{"lat": 40.71, "lng": -74.00},
	})
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "New York, US", resp.Data.Address)
	require.NotNil(t, resp.Data.Timezone)
	assert.Equal(t, "America/New_York", *resp.Data.Timezone)
	assert.Empty(t, resp.Data.Flags)
}

func TestLocationResolve_MissingInput(t *testing.T) {
	t.Parallel()
	r := setupLocationRouter(
		&stubGeocoder{},
		&stubTimezone{},
		&stubHolidays{},
		&stubWorkingDays{},
	)
	w := post(t, r, map[string]any{})
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestLocationResolve_TimezoneUnavailable(t *testing.T) {
	t.Parallel()
	r := setupLocationRouter(
		&stubGeocoder{geocodeR: geocode.GeocodeResponse{
			Address: "Berlin, DE", City: "Berlin", Country: "DE", Lat: 52.52, Lon: 13.40,
		}},
		&stubTimezone{err: fmt.Errorf("tz db unavailable")},
		&stubHolidays{r: holidays.HolidayList{Holidays: []holidays.Holiday{}}},
		&stubWorkingDays{n: 22},
	)
	w := post(t, r, map[string]any{"address": "Berlin, DE"})
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Nil(t, resp.Data.Timezone)
	assert.Contains(t, resp.Data.Flags, "timezone_unavailable")
}

func TestLocationResolve_CalendarUnavailable(t *testing.T) {
	t.Parallel()
	tzInfo := &timezone.Info{Timezone: "Europe/Berlin", Offset: "+02:00", CurrentTime: "2026-05-26T12:00:00Z"}
	r := setupLocationRouter(
		&stubGeocoder{geocodeR: geocode.GeocodeResponse{
			Address: "Berlin, DE", City: "Berlin", Country: "DE", Lat: 52.52, Lon: 13.40,
		}},
		&stubTimezone{r: tzInfo},
		&stubHolidays{err: fmt.Errorf("holiday api down")},
		&stubWorkingDays{n: 22},
	)
	w := post(t, r, map[string]any{"address": "Berlin, DE"})
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Nil(t, resp.Data.NextHoliday)
	assert.Contains(t, resp.Data.Flags, "calendar_unavailable")
}

func TestLocationResolve_GeocodeFail(t *testing.T) {
	t.Parallel()
	r := setupLocationRouter(
		&stubGeocoder{geocodeErr: fmt.Errorf("not found")},
		&stubTimezone{},
		&stubHolidays{},
		&stubWorkingDays{},
	)
	w := post(t, r, map[string]any{"address": "nowhere xyz"})
	assert.NotEqual(t, http.StatusOK, w.Code)
}

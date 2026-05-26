package businesscalendar

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/services/places/holidays"
)

type stubHolidays struct {
	r   holidays.HolidayList
	err error
}

func (s *stubHolidays) GetHolidays(_ string, _ int) (holidays.HolidayList, error) {
	return s.r, s.err
}

type stubWorkingDays struct{ n int }

func (s *stubWorkingDays) GetWorkingDays(_, _ time.Time, _, _ string) int { return s.n }

func setupRouter(h *stubHolidays, w *stubWorkingDays) chi.Router {
	r := chi.NewRouter()
	RegisterRoutes(r, NewService(h, w))
	return r
}

func get(t *testing.T, r chi.Router, path string) *httptest.ResponseRecorder {
	t.Helper()
	req := httptest.NewRequest(http.MethodGet, path, http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func deHolidays() holidays.HolidayList {
	return holidays.HolidayList{
		Country:  "DE",
		Year:     2026,
		Holidays: []holidays.Holiday{{Name: "New Year", Date: "2026-01-01"}},
		Total:    1,
	}
}

func TestBusinessCalendar_ValidCountryAndMonth(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubHolidays{r: deHolidays()}, &stubWorkingDays{n: 21})
	w := get(t, r, "/business-calendar/DE?year=2026&month=1")
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "DE", resp.Data.CountryCode)
	assert.Equal(t, 2026, resp.Data.Year)
	assert.NotNil(t, resp.Data.Month)
	assert.Greater(t, resp.Data.WorkingDays, 0)
	assert.NotNil(t, resp.Data.TotalDays)
}

func TestBusinessCalendar_YearScopeNoMonth(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubHolidays{r: deHolidays()}, &stubWorkingDays{n: 252})
	w := get(t, r, "/business-calendar/DE?year=2026")
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Nil(t, resp.Data.Month)
	assert.Nil(t, resp.Data.TotalDays)
	assert.Equal(t, 252, resp.Data.WorkingDays)
}

func TestBusinessCalendar_NoHolidaysThisMonth(t *testing.T) {
	t.Parallel()
	// holidays only in January, request for February
	list := holidays.HolidayList{Country: "DE", Year: 2026,
		Holidays: []holidays.Holiday{{Name: "New Year", Date: "2026-01-01"}}}
	r := setupRouter(&stubHolidays{r: list}, &stubWorkingDays{n: 20})
	w := get(t, r, "/business-calendar/DE?year=2026&month=2")
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 0, resp.Data.HolidayCount)
	assert.Empty(t, resp.Data.Holidays)
}

func TestBusinessCalendar_UnknownCountry(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubHolidays{}, &stubWorkingDays{})
	// Too long — not a valid alpha-2
	w := get(t, r, "/business-calendar/XYZZY?year=2026")
	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBusinessCalendar_MissingCountry(t *testing.T) {
	t.Parallel()
	r := setupRouter(&stubHolidays{}, &stubWorkingDays{})
	w := get(t, r, "/business-calendar/?year=2026")
	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestBusinessCalendar_NextHolidayNull(t *testing.T) {
	t.Parallel()
	// past holiday only — next_holiday should be null
	list := holidays.HolidayList{Country: "DE", Year: 2000,
		Holidays: []holidays.Holiday{{Name: "Old Year", Date: "2000-01-01"}}}
	r := setupRouter(&stubHolidays{r: list}, &stubWorkingDays{n: 20})
	w := get(t, r, "/business-calendar/DE?year=2000")
	require.Equal(t, http.StatusOK, w.Code)
	var resp httpx.Response[Result]
	require.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Nil(t, resp.Data.NextHoliday)
}

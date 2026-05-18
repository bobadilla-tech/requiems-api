package holidays

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetHolidays_Countries(t *testing.T) {
	t.Parallel()
	cases := []struct {
		name    string
		country string
		year    int
	}{
		{name: "US_2025", country: "US", year: 2025},
		{name: "GB_2025", country: "GB", year: 2025},
		{name: "Japan", country: "JP", year: 2025},
		{name: "Germany", country: "DE", year: 2025},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			svc := NewService()

			resp, err := svc.GetHolidays(tc.country, tc.year)
			require.NoError(t, err)

			assert.Equal(t, tc.country, resp.Country)
			assert.Equal(t, tc.year, resp.Year)
			assert.NotEmpty(t, resp.Holidays)
		})
	}
}

func TestService_GetHolidays_NewYear(t *testing.T) {
	t.Parallel()
	svc := NewService()

	resp, err := svc.GetHolidays("US", 2025)
	require.NoError(t, err)

	found := false
	for _, h := range resp.Holidays {
		if h.Name == "New Year's Day" && h.Date == "2025-01-01" {
			found = true
			break
		}
	}
	if !found {
		t.Error("expected New Year's Day 2025-01-01 in US holidays")
	}
}

func TestService_GetHolidays_DateFormat(t *testing.T) {
	t.Parallel()
	svc := NewService()

	resp, err := svc.GetHolidays("US", 2025)
	require.NoError(t, err)

	for _, h := range resp.Holidays {
		if len(h.Date) != 10 {
			t.Errorf("expected date in YYYY-MM-DD format, got %q", h.Date)
			continue
		}
		if h.Date[4] != '-' || h.Date[7] != '-' {
			t.Errorf("expected date in YYYY-MM-DD format, got %q", h.Date)
		}
	}
}

func TestService_GetHolidays_InvalidCountry(t *testing.T) {
	t.Parallel()
	svc := NewService()

	_, err := svc.GetHolidays("XX", 2025)
	require.Error(t, err)
	if !strings.Contains(err.Error(), "no holidays found") {
		t.Errorf("expected error message to contain 'no holidays found', got %q", err.Error())
	}
}

// ---- batch service tests ----

func TestService_GetHolidaysBatch_AllFound(t *testing.T) {
	t.Parallel()
	svc := NewService()

	queries := []BatchQuery{
		{Country: "US", Year: 2025},
		{Country: "GB", Year: 2025},
	}

	resp := svc.GetHolidaysBatch(queries)

	assert.Equal(t, 2, len(resp))
	for _, item := range resp {
		if !item.Found {
			t.Errorf("expected found=true for %s %d", item.Country, item.Year)
		}
		if len(item.Holidays) == 0 {
			t.Errorf("expected non-empty holidays for %s %d", item.Country, item.Year)
		}
	}
}

func TestService_GetHolidaysBatch_PartialFailure(t *testing.T) {
	t.Parallel()
	svc := NewService()

	queries := []BatchQuery{
		{Country: "US", Year: 2025},
		{Country: "AQ", Year: 2025}, // valid ISO code (Antarctica) but no holiday data — triggers found:false
	}

	resp := svc.GetHolidaysBatch(queries)

	assert.Equal(t, 2, len(resp))
	if !resp[0].Found {
		t.Error("expected Results[0] (US) to be found")
	}
	if resp[1].Found {
		t.Error("expected Results[1] (ZZ) to be not found")
	}
	if len(resp[1].Holidays) != 0 {
		t.Error("expected empty holidays slice for not-found item")
	}
}

func TestService_GetHolidaysBatch_PreservesOrder(t *testing.T) {
	t.Parallel()
	svc := NewService()

	queries := []BatchQuery{
		{Country: "DE", Year: 2025},
		{Country: "AR", Year: 2024},
		{Country: "JP", Year: 2025},
	}

	resp := svc.GetHolidaysBatch(queries)

	assert.Equal(t, "DE", resp[0].Country)
	assert.Equal(t, "AR", resp[1].Country)
	assert.Equal(t, "JP", resp[2].Country)
}

package holidays

import (
	"strings"
	"testing"
)

func TestService_GetHolidays_Countries(t *testing.T) {
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
			svc := NewService()

			resp, err := svc.GetHolidays(tc.country, tc.year)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if resp.Country != tc.country {
				t.Errorf("expected country %q, got %q", tc.country, resp.Country)
			}
			if resp.Year != tc.year {
				t.Errorf("expected year %d, got %d", tc.year, resp.Year)
			}
			if len(resp.Holidays) == 0 {
				t.Error("expected non-empty holidays list")
			}
		})
	}
}

func TestService_GetHolidays_NewYear(t *testing.T) {
	svc := NewService()

	resp, err := svc.GetHolidays("US", 2025)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
	svc := NewService()

	resp, err := svc.GetHolidays("US", 2025)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

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
	svc := NewService()

	_, err := svc.GetHolidays("XX", 2025)
	if err == nil {
		t.Fatal("expected error for invalid country code, got nil")
	}
	if !strings.Contains(err.Error(), "no holidays found") {
		t.Errorf("expected error message to contain 'no holidays found', got %q", err.Error())
	}
}

// ---- batch service tests ----

func TestService_GetHolidaysBatch_AllFound(t *testing.T) {
	svc := NewService()

	queries := []BatchQuery{
		{Country: "US", Year: 2025},
		{Country: "GB", Year: 2025},
	}

	resp := svc.GetHolidaysBatch(queries)

	if resp.Total != 2 {
		t.Errorf("expected total 2, got %d", resp.Total)
	}
	for _, item := range resp.Results {
		if !item.Found {
			t.Errorf("expected found=true for %s %d", item.Country, item.Year)
		}
		if len(item.Holidays) == 0 {
			t.Errorf("expected non-empty holidays for %s %d", item.Country, item.Year)
		}
	}
}

func TestService_GetHolidaysBatch_PartialFailure(t *testing.T) {
	svc := NewService()

	queries := []BatchQuery{
		{Country: "US", Year: 2025},
		{Country: "AQ", Year: 2025}, // valid ISO code (Antarctica) but no holiday data — triggers found:false
	}

	resp := svc.GetHolidaysBatch(queries)

	if resp.Total != 2 {
		t.Errorf("expected total 2, got %d", resp.Total)
	}
	if !resp.Results[0].Found {
		t.Error("expected Results[0] (US) to be found")
	}
	if resp.Results[1].Found {
		t.Error("expected Results[1] (ZZ) to be not found")
	}
	if len(resp.Results[1].Holidays) != 0 {
		t.Error("expected empty holidays slice for not-found item")
	}
}

func TestService_GetHolidaysBatch_PreservesOrder(t *testing.T) {
	svc := NewService()

	queries := []BatchQuery{
		{Country: "DE", Year: 2025},
		{Country: "AR", Year: 2024},
		{Country: "JP", Year: 2025},
	}

	resp := svc.GetHolidaysBatch(queries)

	if resp.Results[0].Country != "DE" {
		t.Errorf("expected Results[0] country DE, got %s", resp.Results[0].Country)
	}
	if resp.Results[1].Country != "AR" {
		t.Errorf("expected Results[1] country AR, got %s", resp.Results[1].Country)
	}
	if resp.Results[2].Country != "JP" {
		t.Errorf("expected Results[2] country JP, got %s", resp.Results[2].Country)
	}
}

package workingdays

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mustDate(t *testing.T, s string) time.Time {
	t.Helper()
	d, err := time.Parse("2006-01-02", s)
	require.NoErrorf(t, err, "invalid fixture date %q", s)
	return d
}

type fixture struct {
	name     string
	country  string
	region   string
	start    string
	end      string
	expected int
	note     string
}

var fixtures = []fixture{
	// ---------------------------------------------------------------
	// BRAZIL
	// ---------------------------------------------------------------
	{
		name: "brazil/good_friday_plus_tiradentes", country: "BR", region: "",
		start: "2025-04-14", end: "2025-04-21", expected: 4,
		note: "Good Friday (18-Apr, national) + Tiradentes (21-Apr, fixed) in the same week",
	},
	{
		name: "brazil/independence_day_on_sunday", country: "BR", region: "",
		start: "2025-09-01", end: "2025-09-07", expected: 5,
		note: "7-Sep falls on Sunday; Brazil has NO substitute-day rule, so no reduction",
	},
	{
		name: "brazil/carnaval_observing_state", country: "BR", region: "BR-RJ",
		start: "2025-03-03", end: "2025-03-04", expected: 0,
		note: "Carnaval Mon/Tue — treated as holiday in most states (e.g. Rio). Region code fixed to full ISO 3166-2 (BR-RJ).",
	},
	// REMOVED: brazil/carnaval_non_observing_city
	// The "XX-NONOBS" region code was invented and doesn't exist in the
	// library's data. Need a real BR-<state> code for a non-observing
	// state/city before this vector means anything. Re-add once confirmed.
	{
		name: "brazil/christmas_fixed_weekday", country: "BR", region: "",
		start: "2025-12-22", end: "2025-12-25", expected: 3,
		note: "Christmas fixed, falls on a Thursday in 2025",
	},
	{
		name: "brazil/black_awareness_day", country: "BR", region: "",
		start: "2025-11-17", end: "2025-11-21", expected: 4,
		note: "20-Nov became a NATIONAL holiday only in 2024 — likely gap in stale libraries",
	},

	// ---------------------------------------------------------------
	// JAPAN — no subdivisions in this country, unaffected by the region-code bug.
	// These remain genuine library gaps (substitution rule not implemented).
	// ---------------------------------------------------------------
	{
		name: "japan/golden_week_substitute", country: "JP", region: "",
		start: "2025-05-01", end: "2025-05-06", expected: 2,
		note: "KNOWN GAP: Greenery Day (4-May, Sunday) has no substitute holiday in the library. Expected to still FAIL until fixed upstream or overridden in service.go.",
	},
	{
		name: "japan/showa_day_fixed", country: "JP", region: "",
		start: "2025-04-28", end: "2025-05-02", expected: 4,
		note: "Showa Day fixed date, falls on a Tuesday",
	},
	{
		name: "japan/happy_monday_sports_day", country: "JP", region: "",
		start: "2025-10-10", end: "2025-10-14", expected: 2,
		note: "Sports Day = 2nd Monday of October ('Happy Monday' floating rule)",
	},
	{
		name: "japan/foundation_day_fixed", country: "JP", region: "",
		start: "2025-02-10", end: "2025-02-14", expected: 4,
		note: "Simple fixed national holiday, weekday baseline",
	},
	{
		name: "japan/happy_monday_marine_day", country: "JP", region: "",
		start: "2025-07-18", end: "2025-07-22", expected: 2,
		note: "Marine Day = 3rd Monday of July, confirms floating-rule consistency",
	},

	// ---------------------------------------------------------------
	// UK
	// ---------------------------------------------------------------
	{
		name: "uk/good_friday_easter_monday", country: "GB", region: "",
		start: "2025-04-14", end: "2025-04-21", expected: 4,
		note: "TODO-VERIFY: empty region = union of national + ALL UK subdivisions. " +
			"Scotland does NOT observe Easter Monday as a bank holiday — confirm this doesn't " +
			"change the union count before trusting this number.",
	},
	{
		name: "uk/st_andrews_substitute_scotland", country: "GB", region: "GB-SCT",
		start: "2025-11-28", end: "2025-12-02", expected: 2,
		note: "KNOWN GAP (separate from region-code bug): St Andrew's Day substitute (30-Nov Sun -> 1-Dec Mon) not implemented for Scotland. Expected to still FAIL.",
	},
	{
		name: "uk/st_andrews_not_observed_england", country: "GB", region: "GB-ENG",
		start: "2025-11-28", end: "2025-12-02", expected: 3,
		note: "England/Wales never observe St Andrew's Day, no substitution involved — should pass regardless of the Scotland gap",
	},
	{
		name: "uk/christmas_boxing_day", country: "GB", region: "",
		start: "2025-12-22", end: "2025-12-26", expected: 3,
		note: "TODO-VERIFY: empty region = union across all UK subdivisions; confirm no nation-specific extra holiday falls in this week before trusting the count",
	},
	{
		name: "uk/early_may_bank_holiday", country: "GB", region: "",
		start: "2025-05-05", end: "2025-05-09", expected: 4,
		note: "Early May bank holiday is UK-wide, no subdivision variance expected this week",
	},

	// ---------------------------------------------------------------
	// SPAIN
	// ---------------------------------------------------------------
	{
		name: "spain/no_subdivision_is_union_not_national_only", country: "ES", region: "",
		start: "2025-04-28", end: "2025-05-02", expected: 3,
		note: "IMPORTANT SEMANTIC NOTE: empty Subdivision returns the UNION of national + " +
			"every region's holidays, not a national-only calendar. Madrid Day (2-May) is the " +
			"only regional holiday that week, so this equals the ES-MD result below (3), NOT 4. " +
			"There is currently no way to request 'national holidays only' from this library.",
	},
	{
		name: "spain/labour_day_plus_madrid_day", country: "ES", region: "ES-MD",
		start: "2025-04-28", end: "2025-05-02", expected: 3,
		note: "Madrid adds 2-May (Comunidad de Madrid day) on top of 1-May national — region code fixed to full ISO 3166-2",
	},
	{
		name: "spain/easter_monday_catalonia", country: "ES", region: "ES-CT",
		start: "2025-04-17", end: "2025-04-21", expected: 1,
		note: "Good Friday (national) + Easter Monday (regional, Catalonia only) — region code fixed",
	},
	{
		name: "spain/holy_week_madrid", country: "ES", region: "ES-MD",
		start: "2025-04-17", end: "2025-04-21", expected: 1,
		note: "Madrid observes Jueves Santo (17-Apr, regional) + Good Friday (18-Apr, national). " +
			"Does NOT observe Easter Monday (21-Apr) — that one's Catalonia/Balearics/etc. only. " +
			"Corrected from an earlier fixture that missed Jueves Santo entirely.",
	},
	{
		name: "spain/constitution_day_on_saturday", country: "ES", region: "",
		start: "2025-12-01", end: "2025-12-08", expected: 5,
		note: "6-Dec falls Saturday, no shift to a weekday. No known regional holiday overlaps this week, so union == national-only here.",
	},
	{
		name: "spain/christmas_union_equals_catalonia", country: "ES", region: "",
		start: "2025-12-22", end: "2025-12-26", expected: 3,
		note: "SEMANTIC NOTE (same as above): empty Subdivision = union of all regions. " +
			"Catalonia's Sant Esteve (26-Dec) is the only regional holiday that week, so this " +
			"equals the ES-CT result below (3), NOT a plain national-only 4.",
	},
	{
		name: "spain/christmas_plus_sant_esteve_catalonia", country: "ES", region: "ES-CT",
		start: "2025-12-22", end: "2025-12-26", expected: 3,
		note: "Catalonia adds Sant Esteve (26-Dec) — region code fixed",
	},

	// ---------------------------------------------------------------
	// USA — no subdivisions used here, unaffected by the region-code bug.
	// ---------------------------------------------------------------
	{
		name: "usa/july4_on_weekday", country: "US", region: "",
		start: "2025-06-30", end: "2025-07-04", expected: 4,
		note: "Independence Day falls on a Friday, no substitution needed",
	},
	{
		name: "usa/july4_observed_friday", country: "US", region: "",
		start: "2026-06-29", end: "2026-07-05", expected: 4,
		note: "KNOWN GAP: 4-Jul-2026 falls Saturday, library returned 0 holidays in range — 'observed' rule not implemented for 2026. Expected to still FAIL.",
	},
	{
		name: "usa/labor_day_nth_monday", country: "US", region: "",
		start: "2025-09-01", end: "2025-09-05", expected: 4,
		note: "Labor Day = 1st Monday of September (floating rule)",
	},
	{
		name: "usa/veterans_day_fixed", country: "US", region: "",
		start: "2025-11-10", end: "2025-11-14", expected: 4,
		note: "Fixed federal holiday, weekday baseline",
	},
	{
		name: "usa/thanksgiving_friday_not_federal", country: "US", region: "",
		start: "2025-11-24", end: "2025-11-28", expected: 4,
		note: "Thanksgiving is Thursday only; the following Friday is NOT a federal holiday",
	},
}

func TestWorkingDaysFixtures(t *testing.T) {
	service := NewService()
	for _, f := range fixtures {
		f := f
		t.Run(f.name, func(t *testing.T) {
			from := mustDate(t, f.start)
			to := mustDate(t, f.end)

			got := service.GetWorkingDays(from, to, f.country, f.region)

			assert.Equalf(t, f.expected, got,
				"GetWorkingDays(%s..%s, country=%s, region=%s)\nnote: %s",
				f.start, f.end, f.country, f.region, f.note)
		})
	}
}

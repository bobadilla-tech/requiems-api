package workingdays

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetWorkingDaysBatch(t *testing.T) {
	t.Parallel()

	svc := NewService()

	from := time.Date(2026, 1, 1, 0, 0, 0, 0, time.UTC)
	to := time.Date(2026, 1, 10, 0, 0, 0, 0, time.UTC)

	items := []Request{
		{
			From:        from,
			To:          to,
			Country:     "",
			Subdivision: "",
		},
		{
			From:        from,
			To:          to,
			Country:     "PE",
			Subdivision: "",
		},
		{
			From:        from,
			To:          to,
			Country:     "US",
			Subdivision: "CA",
		},
	}

	results := svc.GetWorkingDaysBatch(items)

	require.Len(t, results, 3)

	// Validate structure consistency
	for i, r := range results {
		assert.NotEmpty(t, r.From)
		assert.NotEmpty(t, r.To)

		assert.Equal(t, items[i].Country, r.Country)
		assert.Equal(t, items[i].Subdivision, r.Subdivision)
		assert.GreaterOrEqual(t, r.WorkingDays, 0)
	}
}

func TestService_GetWorkingDaysBatch_Empty(t *testing.T) {
	t.Parallel()

	svc := NewService()

	results := svc.GetWorkingDaysBatch([]Request{})

	require.Empty(t, results)
}

func TestService_GetWorkingDaysBatch_SingleItem(t *testing.T) {
	t.Parallel()

	svc := NewService()

	from := time.Now().AddDate(0, 0, -10)
	to := time.Now()

	items := []Request{
		{
			From: from,
			To:   to,
		},
	}

	results := svc.GetWorkingDaysBatch(items)

	require.Len(t, results, 1)
	assert.GreaterOrEqual(t, results[0].WorkingDays, 0)
}

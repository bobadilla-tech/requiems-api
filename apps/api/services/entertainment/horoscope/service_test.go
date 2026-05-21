package horoscope

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_DailyBatch(t *testing.T) {
	t.Parallel()
	svc := NewService()

	signs := []string{
		"taurus", "aries", "gemini", "cancer",
	}

	results, err := svc.DailyBatch(signs)

	require.NoError(t, err)
	require.Equal(t, 4, len(results))

	for i, sign := range signs {
		assert.Equal(t, sign, results[i].Sign)
		assert.NotEmpty(t, results[i].Horoscope)
		assert.NotEmpty(t, results[i].Mood)
		assert.NotZero(t, results[i].LuckyNumber)
		assert.Equal(t, time.Now().UTC().Format("2006-01-02"), results[i].Date)
	}
}

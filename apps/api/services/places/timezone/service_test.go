package timezone

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_GetTimezoneBatch(t *testing.T) {
	t.Parallel()

	svc, err := NewService()
	require.NoError(t, err)

	tests := []struct {
		name     string
		cities   []string
		wantSize int
	}{
		{
			name: "multiple valid cities",
			cities: []string{
				"lima",
				"tokyo",
				"new york",
			},
			wantSize: 3,
		},
		{
			name: "single city",
			cities: []string{
				"paris",
			},
			wantSize: 1,
		},
		{
			name: "mixed valid and invalid cities",
			cities: []string{
				"lima",
				"invalid-city",
				"tokyo",
			},
			wantSize: 3,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			result, err := svc.GetTimezoneBatch(
				context.Background(),
				tt.cities,
			)

			require.NoError(t, err)

			require.Len(t, result.Results, tt.wantSize)

			for i, city := range tt.cities {
				assert.Equal(t, city, result.Results[i].City)
			}
		})
	}
}

func TestService_GetTimezoneBatch_PreservesOrder(t *testing.T) {
	t.Parallel()

	svc, err := NewService()
	require.NoError(t, err)

	cities := []string{
		"tokyo",
		"lima",
		"new york",
		"paris",
	}

	result, err := svc.GetTimezoneBatch(
		context.Background(),
		cities,
	)

	require.NoError(t, err)

	require.Len(t, result.Results, len(cities))

	for i := range cities {
		assert.Equal(t, cities[i], result.Results[i].City)
	}
}

func TestService_GetTimezoneBatch_InvalidCity(t *testing.T) {
	t.Parallel()

	svc, err := NewService()
	require.NoError(t, err)

	result, err := svc.GetTimezoneBatch(
		context.Background(),
		[]string{
			"invalid-city",
		},
	)

	require.NoError(t, err)

	require.Len(t, result.Results, 1)

	assert.Equal(t, "invalid-city", result.Results[0].City)
	assert.Nil(t, result.Results[0].Info)
}

func TestService_GetTimezoneBatch_ValidCityInfo(t *testing.T) {
	t.Parallel()

	svc, err := NewService()
	require.NoError(t, err)

	result, err := svc.GetTimezoneBatch(
		context.Background(),
		[]string{
			"lima",
		},
	)

	require.NoError(t, err)

	require.Len(t, result.Results, 1)

	item := result.Results[0]

	assert.Equal(t, "lima", item.City)

	require.NotNil(t, item.Info)

	assert.NotEmpty(t, item.Info.Timezone)
	assert.NotEmpty(t, item.Info.Offset)
	assert.NotEmpty(t, item.Info.CurrentTime)
}

func TestService_GetTimezoneBatch_EmptyInput(t *testing.T) {
	t.Parallel()

	svc, err := NewService()
	require.NoError(t, err)

	result, err := svc.GetTimezoneBatch(
		context.Background(),
		[]string{},
	)

	require.NoError(t, err)

	assert.Empty(t, result.Results)
}

package commodities

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

// stubGetter implements Getter for service-layer unit tests.
type stubGetter struct {
	result CommodityPrice
	err    error
}

func (s *stubGetter) Get(_ context.Context, slug string) (CommodityPrice, error) {
	if s.err != nil {
		return CommodityPrice{}, s.err
	}
	r := s.result
	r.Commodity = slug
	return r, nil
}

// ---- HistoricalPrice ----

func TestHistoricalPrice_FieldsPresent(t *testing.T) {
	t.Parallel()
	h := HistoricalPrice{Period: "2023", Price: 1940.54}
	assert.Equal(t, "2023", h.Period)
	assert.Equal(t, 1940.54, h.Price)
}

// ---- CommodityPrice ----

func TestCommodityPrice_IsData(t *testing.T) {
	t.Parallel()
	// IsData() must be callable — verifies the interface is satisfied.
	var c CommodityPrice
	c.IsData()
}

func TestCommodityPrice_FullResponse(t *testing.T) {
	t.Parallel()
	cp := CommodityPrice{
		Commodity: "gold",
		Name:      "Gold",
		Price:     2386.33,
		Unit:      "oz",
		Currency:  "USD",
		Change24h: 23.01,
		Historical: []HistoricalPrice{
			{Period: "2023", Price: 1940.54},
			{Period: "2022", Price: 1800.12},
		},
	}

	assert.Equal(t, "gold", cp.Commodity)
	assert.Equal(t, 2386.33, cp.Price)
	assert.Len(t, cp.Historical, 2)
}

// ---- Getter stub ----

func TestStubGetter_ReturnsCommodity(t *testing.T) {
	t.Parallel()
	stub := &stubGetter{result: CommodityPrice{
		Name:  "Gold",
		Price: 2386.33,
		Unit:  "oz",
	}}

	result, err := stub.Get(context.Background(), "gold")
	require.NoError(t, err)
	assert.Equal(t, "gold", result.Commodity)
}

func TestStubGetter_PropagatesError(t *testing.T) {
	t.Parallel()
	stub := &stubGetter{err: &httpx.AppError{
		Status:  404,
		Code:    "not_found",
		Message: "commodity not found",
	}}

	_, err := stub.Get(context.Background(), "unknown")
	require.Error(t, err)
}

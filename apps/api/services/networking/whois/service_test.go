package whois

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestService_LookupBatch verifies that LookupBatch successfully
// processes multiple domains and returns valid WHOIS data.
func TestService_LookupBatch(t *testing.T) {
	svc := &Service{
		q: &fakeQuerier{
			result: sampleWHOIS,
		},
	}

	domains := []string{
		"example.com",
		"google.com",
	}

	resp, err := svc.LookupBatch(context.Background(), domains)

	assert.NoError(t, err)
	assert.Equal(t, 2, resp.Total)
	assert.Len(t, resp.Results, 2)

	for i, item := range resp.Results {
		assert.Equal(t, domains[i], item.Domain)
		assert.True(t, item.Found)
		assert.Empty(t, item.Error)

		assert.Equal(t, domains[i], item.Data.Domain)
		assert.NotEmpty(t, item.Data.Registrar)
		assert.NotEmpty(t, item.Data.CreatedDate)
	}
}

// TestService_LookupBatch_NotFound verifies that LookupBatch
// correctly marks domains as not found when WHOIS data does not exist.
func TestService_LookupBatch_NotFound(t *testing.T) {
	svc := &Service{
		q: &fakeQuerier{
			result: notFoundWHOIS,
		},
	}

	domains := []string{
		"doesnotexist.com",
	}

	resp, err := svc.LookupBatch(context.Background(), domains)

	assert.NoError(t, err)
	assert.Equal(t, 1, resp.Total)
	assert.Len(t, resp.Results, 1)

	item := resp.Results[0]

	assert.Equal(t, "doesnotexist.com", item.Domain)
	assert.False(t, item.Found)
	assert.NotEmpty(t, item.Error)
}

// TestService_LookupBatch_Empty verifies that LookupBatch
// handles an empty domain list without errors.
func TestService_LookupBatch_Empty(t *testing.T) {
	svc := &Service{
		q: &fakeQuerier{
			result: sampleWHOIS,
		},
	}

	resp, err := svc.LookupBatch(context.Background(), []string{})

	assert.NoError(t, err)
	assert.Equal(t, 0, resp.Total)
	assert.Empty(t, resp.Results)
}

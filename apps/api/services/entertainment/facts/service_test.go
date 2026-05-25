package facts
	import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
func TestService_RandomBatch_AllValid(t *testing.T) {
	t.Parallel()
	svc := NewService()

	categories := []string{"science", "history", "technology", "nature", "space", "food"}
	resp := svc.RandomBatch(context.Background(), categories)

	require.Len(t, resp.Results, len(categories))
	for i, item := range resp.Results {
		assert.Equal(t, categories[i], item.Category)
		assert.NotEmpty(t, item.Fact)
		assert.NotEmpty(t, item.Source)
		assert.Empty(t, item.Error)
	}
}

func TestService_RandomBatch_InvalidCategory_InBand(t *testing.T) {
	t.Parallel()
	svc := NewService()

	resp := svc.RandomBatch(context.Background(), []string{"science", "dragons", "space"})

	require.Len(t, resp.Results, 3)

	assert.NotEmpty(t, resp.Results[0].Fact)
	assert.Empty(t, resp.Results[0].Error)

	assert.Empty(t, resp.Results[1].Fact)
	assert.NotEmpty(t, resp.Results[1].Error) // absorbed in-band

	assert.NotEmpty(t, resp.Results[2].Fact)
	assert.Empty(t, resp.Results[2].Error)
}

func TestService_RandomBatch_PreservesOrder(t *testing.T) {
	t.Parallel()
	svc := NewService()

	categories := []string{"food", "science", "history", "nature"}
	for range 20 {
		resp := svc.RandomBatch(context.Background(), categories)
		for i, item := range resp.Results {
			assert.Equal(t, categories[i], item.Category)
		}
	}
}

func TestService_RandomBatch_AllInvalid(t *testing.T) {
	t.Parallel()
	svc := NewService()

	resp := svc.RandomBatch(context.Background(), []string{"foo", "bar"})

	for _, item := range resp.Results {
		assert.Empty(t, item.Fact)
		assert.NotEmpty(t, item.Error)
	}
}

func TestService_RandomBatch_DuplicateCategories(t *testing.T) {
	t.Parallel()
	svc := NewService()

	resp := svc.RandomBatch(context.Background(), []string{"science", "science", "science"})

	require.Len(t, resp.Results, 3)
	for _, item := range resp.Results {
		assert.Equal(t, "science", item.Category)
		assert.NotEmpty(t, item.Fact)
	}
}
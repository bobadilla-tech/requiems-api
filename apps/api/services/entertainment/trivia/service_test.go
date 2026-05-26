package trivia

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestService_RandomBatch_AllSucceed(t *testing.T) {
	t.Parallel()
	svc := NewService()

	filters := []Request{
		{Category: "science", Difficulty: "hard"},
		{Category: "history", Difficulty: "medium"},
	}

	results := svc.RandomBatch(filters)

	require.Equal(t, len(filters), len(results))
	for i, filter := range filters {
		assert.Equal(t, filter.Category, results[i].Category)
		assert.Equal(t, filter.Difficulty, results[i].Difficulty)
		assert.Empty(t, results[i].Error)
		assert.NotEmpty(t, results[i].Data.Question)
		assert.NotEmpty(t, results[i].Data.Answer)
		assert.NotEmpty(t, results[i].Data.Options)
	}
}

func TestService_RandomBatch_PreservedOrder(t *testing.T) {
	t.Parallel()
	svc := NewService()

	filters := []Request{
		{Category: "__invalid__", Difficulty: "__invalid__"}, // fails
		{Category: "science", Difficulty: "hard"},            // succeeds
		{Category: "history", Difficulty: "medium"},          // succeeds
	}

	results := svc.RandomBatch(filters)

	require.Equal(t, len(filters), len(results))
	for i, filter := range filters {
		assert.Equal(t, filter.Category, results[i].Category)
		assert.Equal(t, filter.Difficulty, results[i].Difficulty)
	}

	assert.NotEmpty(t, results[0].Error)
	assert.Empty(t, results[0].Data.Question)

	assert.Empty(t, results[1].Error)
	assert.NotEmpty(t, results[1].Data.Question)

	assert.Empty(t, results[2].Error)
	assert.NotEmpty(t, results[2].Data.Question)
}

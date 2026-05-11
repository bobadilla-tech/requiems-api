package sudoku

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGenerate_InvalidDifficulty(t *testing.T) {
	t.Parallel()
	svc := NewService()
	p, err := svc.Generate("impossible")
	require.Error(t, err)
	assert.Equal(t, Puzzle{}, p)
}

func TestGenerateBatch_InvalidDifficultyPropagates(t *testing.T) {
	t.Parallel()
	svc := NewService()
	_, err := svc.GenerateBatch([]string{"easy", "impossible"})
	require.Error(t, err)
}

func TestGenerateBatch_AllValid(t *testing.T) {
	t.Parallel()
	svc := NewService()
	results, err := svc.GenerateBatch([]string{"easy", "medium", "hard"})
	require.NoError(t, err)
	require.Len(t, results, 3)
	assert.Equal(t, "easy", results[0].Difficulty)
	assert.Equal(t, "medium", results[1].Difficulty)
	assert.Equal(t, "hard", results[2].Difficulty)
}

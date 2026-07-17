package sentiment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyze_Positive(t *testing.T) {
	t.Parallel()

	result := testSvc.Analyze("I love this product! It's amazing.")

	assert.Equal(t, "positive", result.Sentiment)
	assert.True(t, result.Score > 0.5, "expected score > 0.5, got %.2f", result.Score)
	assert.True(t, result.Breakdown.Positive > result.Breakdown.Negative, "expected positive > negative in breakdown, got pos=%.2f neg=%.2f", result.Breakdown.Positive, result.Breakdown.Negative)
}

func TestAnalyze_Negative(t *testing.T) {
	t.Parallel()
	result := testSvc.Analyze("This is terrible and awful. I hate it.")

	assert.Equal(t, "negative", result.Sentiment)
	assert.True(t, result.Score > 0.5, "expected score > 0.5, got %.2f", result.Score)
	assert.True(t, result.Breakdown.Negative > result.Breakdown.Positive, "expected negative > positive in breakdown, got pos=%.2f neg=%.2f", result.Breakdown.Positive, result.Breakdown.Negative)
}

func TestAnalyze_Neutral(t *testing.T) {
	t.Parallel()
	result := testSvc.Analyze("The document is on the table.")

	assert.Equal(t, "neutral", result.Sentiment)
	assert.Equal(t, 1.0, result.Score)
	assert.Equal(t, 1.0, result.Breakdown.Neutral)
}

func TestAnalyze_BreakdownSumsToOne(t *testing.T) {
	t.Parallel()

	texts := []string{
		"I love this product! It's amazing.",
		"This is terrible and awful.",
		"The document is on the table.",
		"It's okay but not great. Some issues exist.",
	}

	for _, text := range texts {
		result := testSvc.Analyze(text)
		sum := result.Breakdown.Positive + result.Breakdown.Negative + result.Breakdown.Neutral
		// Allow a tolerance of 0.02 for floating-point rounding.
		assert.True(t, sum >= 0.98 && sum <= 1.02, "breakdown values for %q sum to %.4f, want ~1.0", text, sum)
	}
}

func TestAnalyze_Negation(t *testing.T) {
	t.Parallel()

	// "not good" should score less positively than "good" alone.
	withNeg := testSvc.Analyze("This is not good.")
	withoutNeg := testSvc.Analyze("This is good.")

	assert.True(t, withNeg.Breakdown.Positive < withoutNeg.Breakdown.Positive, "negation should reduce positive score: negated=%.2f plain=%.2f", withNeg.Breakdown.Positive, withoutNeg.Breakdown.Positive)
}

func TestAnalyze_ContractionNegation(t *testing.T) {
	t.Parallel()

	result := testSvc.Analyze("I don't like this product.")
	assert.Equal(t, "negative", result.Sentiment)
	assert.True(t, result.Breakdown.Negative > result.Breakdown.Positive,
		"contraction negation should yield negative: neg=%.2f pos=%.2f",
		result.Breakdown.Negative, result.Breakdown.Positive)
}

func TestAnalyze_Intensifier(t *testing.T) {
	t.Parallel()

	// "very good" should score more positively than "good" alone.
	withIntensifier := testSvc.Analyze("This is very good.")
	without := testSvc.Analyze("This is good.")

	assert.True(t, withIntensifier.Breakdown.Positive > without.Breakdown.Positive, "intensifier should increase positive score: intensified=%.2f plain=%.2f", withIntensifier.Breakdown.Positive, without.Breakdown.Positive)
}

func TestAnalyze_ScoreMatchesDominantClass(t *testing.T) {
	t.Parallel()

	result := testSvc.Analyze("I love this product! It's amazing.")

	var dominant float64
	switch result.Sentiment {
	case "positive":
		dominant = result.Breakdown.Positive
	case "negative":
		dominant = result.Breakdown.Negative
	case "neutral":
		dominant = result.Breakdown.Neutral
	}

	assert.Equal(t, dominant, result.Score)
}

func TestAnalyzeBatch_ReturnsResultsInOrder(t *testing.T) {
	t.Parallel()

	texts := []string{
		"I love this product! It's amazing.",
		"This is terrible and I hate it.",
		"The document is on the table.",
	}

	results := testSvc.AnalyzeBatch(texts)

	assert.Len(t, results, 3)
	assert.Equal(t, "positive", results[0].Sentiment)
	assert.Equal(t, "negative", results[1].Sentiment)
	assert.Equal(t, "neutral", results[2].Sentiment)
}

func TestAnalyzeBatch_SingleItem(t *testing.T) {
	t.Parallel()

	results := testSvc.AnalyzeBatch([]string{"I love this!"})

	assert.Len(t, results, 1)
	assert.Equal(t, "positive", results[0].Sentiment)
}

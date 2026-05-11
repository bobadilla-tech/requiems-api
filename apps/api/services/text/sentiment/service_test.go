package sentiment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnalyze_Positive(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Analyze("I love this product! It's amazing.")

	assert.Equal(t, "positive", result.Sentiment)
	assert.True(t, result.Score > 0.5, "expected score > 0.5, got %.2f", result.Score)
	assert.True(t, result.Breakdown.Positive > result.Breakdown.Negative, "expected positive > negative in breakdown, got pos=%.2f neg=%.2f", result.Breakdown.Positive, result.Breakdown.Negative)
}

func TestAnalyze_Negative(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Analyze("This is terrible and awful. I hate it.")

	assert.Equal(t, "negative", result.Sentiment)
	assert.True(t, result.Score > 0.5, "expected score > 0.5, got %.2f", result.Score)
	assert.True(t, result.Breakdown.Negative > result.Breakdown.Positive, "expected negative > positive in breakdown, got pos=%.2f neg=%.2f", result.Breakdown.Positive, result.Breakdown.Negative)
}

func TestAnalyze_Neutral(t *testing.T) {
	t.Parallel()
	svc := NewService()
	result := svc.Analyze("The document is on the table.")

	assert.Equal(t, "neutral", result.Sentiment)
	assert.Equal(t, 1.0, result.Score)
	assert.Equal(t, 1.0, result.Breakdown.Neutral)
}

func TestAnalyze_BreakdownSumsToOne(t *testing.T) {
	t.Parallel()
	svc := NewService()

	texts := []string{
		"I love this product! It's amazing.",
		"This is terrible and awful.",
		"The document is on the table.",
		"It's okay but not great. Some issues exist.",
	}

	for _, text := range texts {
		result := svc.Analyze(text)
		sum := result.Breakdown.Positive + result.Breakdown.Negative + result.Breakdown.Neutral
		// Allow a tolerance of 0.02 for floating-point rounding.
		assert.True(t, sum >= 0.98 && sum <= 1.02, "breakdown values for %q sum to %.4f, want ~1.0", text, sum)
	}
}

func TestAnalyze_Negation(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// "not good" should score less positively than "good" alone.
	withNeg := svc.Analyze("This is not good.")
	withoutNeg := svc.Analyze("This is good.")

	assert.True(t, withNeg.Breakdown.Positive < withoutNeg.Breakdown.Positive, "negation should reduce positive score: negated=%.2f plain=%.2f", withNeg.Breakdown.Positive, withoutNeg.Breakdown.Positive)
}

func TestAnalyze_Intensifier(t *testing.T) {
	t.Parallel()
	svc := NewService()

	// "very good" should score more positively than "good" alone.
	withIntensifier := svc.Analyze("This is very good.")
	without := svc.Analyze("This is good.")

	assert.True(t, withIntensifier.Breakdown.Positive > without.Breakdown.Positive, "intensifier should increase positive score: intensified=%.2f plain=%.2f", withIntensifier.Breakdown.Positive, without.Breakdown.Positive)
}

func TestAnalyze_ScoreMatchesDominantClass(t *testing.T) {
	t.Parallel()
	svc := NewService()

	result := svc.Analyze("I love this product! It's amazing.")

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

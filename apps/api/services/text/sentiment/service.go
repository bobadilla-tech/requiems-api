package sentiment

import (
	"math"
	"strings"
	"unicode"
)

// neutralityWeight is added to the denominator when computing the breakdown
// to ensure no class ever reaches a probability of exactly 1 and to model
// the inherent ambiguity of natural language.
const neutralityWeight = 0.1

// lexicon maps lower-case words to their sentiment valence.
// Positive values indicate positive sentiment; negative values indicate
// negative sentiment. Magnitudes range from 0.0 (neutral) to 1.0 (extreme).
var lexicon = map[string]float64{
	// Strongly positive
	"love": 0.9, "adore": 0.9, "amazing": 0.9, "excellent": 0.85,
	"wonderful": 0.85, "fantastic": 0.85, "awesome": 0.85, "perfect": 0.9,
	"brilliant": 0.85, "outstanding": 0.85, "superb": 0.85, "exceptional": 0.85,
	"extraordinary": 0.85, "magnificent": 0.85, "fabulous": 0.85, "terrific": 0.8,
	"delightful": 0.8, "thrilled": 0.8, "joyful": 0.8, "ecstatic": 0.9,
	"elated": 0.8, "overjoyed": 0.9, "exhilarating": 0.85,

	// Moderately positive
	"great": 0.8, "good": 0.7, "nice": 0.65, "happy": 0.75, "joy": 0.8,
	"beautiful": 0.75, "impressive": 0.75, "remarkable": 0.75, "captivating": 0.75,
	"charming": 0.7, "pleasing": 0.7, "pleased": 0.7, "glad": 0.65,
	"satisfied": 0.65, "enjoy": 0.7, "enjoyed": 0.7, "enjoying": 0.7,
	"exciting": 0.75, "excited": 0.75, "grateful": 0.75, "thankful": 0.75,
	"appreciate": 0.7, "appreciated": 0.7, "inspiring": 0.75, "inspired": 0.7,
	"innovative": 0.7, "recommend": 0.7, "recommended": 0.7, "reliable": 0.65,
	"trustworthy": 0.7, "efficient": 0.65, "effective": 0.65, "professional": 0.65,
	"helpful": 0.65, "useful": 0.6, "positive": 0.6, "best": 0.85,
	"better": 0.65, "prefer": 0.55, "quality": 0.7, "powerful": 0.65,
	"safe": 0.65, "secure": 0.65, "fun": 0.7, "friendly": 0.65,
	"comfortable": 0.65, "easy": 0.55, "smooth": 0.6, "fast": 0.6,
	"quick": 0.55, "clean": 0.6, "sturdy": 0.6, "robust": 0.65,
	"generous": 0.7, "warm": 0.6, "kind": 0.65, "caring": 0.7,
	"honest": 0.65, "authentic": 0.65, "genuine": 0.65, "smart": 0.65,
	"clever": 0.65, "like": 0.5, "liked": 0.5, "charmed": 0.7,
	"refreshing": 0.65, "flawless": 0.85, "polished": 0.65,

	// Strongly negative
	"hate": -0.9, "terrible": -0.9, "awful": -0.85, "horrible": -0.9,
	"disgusting": -0.9, "dreadful": -0.85, "hideous": -0.85, "atrocious": -0.9,
	"appalling": -0.9, "abysmal": -0.9, "worthless": -0.85, "scam": -0.85,
	"fraud": -0.85, "evil": -0.9, "cruel": -0.85, "violent": -0.85,
	"toxic": -0.85, "dangerous": -0.8, "harmful": -0.8, "offensive": -0.8,
	"despise": -0.9, "loathe": -0.9, "detest": -0.9, "abhor": -0.9,

	// Moderately negative
	"bad": -0.7, "poor": -0.65, "disappointing": -0.75, "disappointed": -0.7,
	"ugly": -0.7, "useless": -0.75, "broken": -0.7, "defective": -0.75,
	"damaged": -0.7, "annoying": -0.7, "annoyed": -0.65, "frustrating": -0.75,
	"frustrated": -0.7, "angry": -0.75, "upset": -0.65, "sad": -0.7,
	"depressed": -0.8, "unhappy": -0.7, "miserable": -0.85, "painful": -0.75,
	"difficult": -0.5, "problem": -0.6, "problems": -0.6, "issue": -0.5,
	"issues": -0.5, "failure": -0.8, "fail": -0.7, "failed": -0.7,
	"wrong": -0.65, "incorrect": -0.6, "inaccurate": -0.6, "false": -0.6,
	"fake": -0.7, "waste": -0.7, "slow": -0.55, "dirty": -0.65,
	"messy": -0.55, "complicated": -0.55, "confusing": -0.6, "unreliable": -0.7,
	"unsafe": -0.75, "rude": -0.75, "mean": -0.7, "lazy": -0.6,
	"corrupt": -0.85, "lie": -0.8, "lied": -0.8, "lying": -0.8,
	"mistake": -0.65, "error": -0.6, "mediocre": -0.5, "overpriced": -0.6,
	"regret": -0.7, "regretted": -0.7, "shame": -0.7, "shameful": -0.8,
	"hostile": -0.75, "aggressive": -0.7, "worst": -0.9, "dislike": -0.65,
	"disliked": -0.65,
}

// negationWords, when encountered, invert the valence of the next three tokens.
var negationWords = map[string]bool{
	"not": true, "never": true, "no": true, "nobody": true,
	"nothing": true, "neither": true, "nor": true, "nowhere": true,
	"hardly": true, "barely": true, "scarcely": true,
	"without": true, "cannot": true,
	// Contractions — tokenize handles stripping punctuation, so these
	// also need to be listed in their "clean" forms where applicable.
	"dont": true, "didnt": true, "wont": true, "cant": true,
	"couldnt": true, "wouldnt": true, "shouldnt": true,
	"isnt": true, "arent": true, "wasnt": true, "werent": true,
	"doesnt": true, "havent": true, "hasnt": true, "hadnt": true,
}

// intensifiers multiply the valence of the following word.
var intensifiers = map[string]float64{
	"very": 1.3, "extremely": 1.5, "really": 1.2, "absolutely": 1.5,
	"totally": 1.4, "completely": 1.4, "utterly": 1.5, "incredibly": 1.5,
	"especially": 1.3, "particularly": 1.3, "quite": 1.1, "highly": 1.3,
	"deeply": 1.3, "truly": 1.3, "genuinely": 1.2, "super": 1.4,
	"most": 1.3,
}

// Breakdown contains the proportional score for each sentiment class.
type Breakdown struct {
	Positive float64 `json:"positive"`
	Negative float64 `json:"negative"`
	Neutral  float64 `json:"neutral"`
}

// Result is the response payload for the sentiment endpoint.
type Result struct {
	Sentiment string    `json:"sentiment"`
	Score     float64   `json:"score"`
	Breakdown Breakdown `json:"breakdown"`
}

// Service performs sentiment analysis.
type Service struct{}

// NewService returns a new sentiment Service.
func NewService() *Service { return &Service{} }

// Analyze computes the sentiment of the given text, returning a label
// ("positive", "negative", or "neutral"), a confidence score, and a
// probability breakdown across the three classes.
func (s *Service) Analyze(text string) Result {
	tokens := tokenize(text)

	posScore, negScore := scoreTokens(tokens)
	total := posScore + negScore

	if total == 0 {
		return Result{
			Sentiment: "neutral",
			Score:     1.0,
			Breakdown: Breakdown{Neutral: 1.0},
		}
	}

	// Divide each class by (total + neutralityWeight) so that even a single
	// strongly positive word does not push the positive score all the way to 1.
	denom := total + neutralityWeight
	pos := round2(posScore / denom)
	neg := round2(negScore / denom)
	// Derive neutral as the remainder to guarantee the three values sum to 1.
	neu := round2(1 - pos - neg)

	sentiment := "neutral"
	score := neu
	if pos >= neg && pos > neu {
		sentiment = "positive"
		score = pos
	} else if neg > pos && neg > neu {
		sentiment = "negative"
		score = neg
	}

	return Result{
		Sentiment: sentiment,
		Score:     score,
		Breakdown: Breakdown{
			Positive: pos,
			Negative: neg,
			Neutral:  neu,
		},
	}
}

// scoreTokens iterates over tokens, applying negation and intensifier
// modifiers, and returns separate positive and negative valence sums.
func scoreTokens(tokens []string) (posSum, negSum float64) {
	// negationLeft tracks how many subsequent tokens are still under a negation.
	negationLeft := 0
	// intensifierMult carries an intensifier multiplier into the next token.
	intensifierMult := 1.0

	for _, token := range tokens {
		if negationLeft > 0 {
			negationLeft--
		}

		// Update state for the current token before applying it as a valence.
		if negationWords[token] {
			negationLeft = 3
			intensifierMult = 1.0
			continue
		}

		if mult, ok := intensifiers[token]; ok {
			intensifierMult = mult
			continue
		}

		valence, ok := lexicon[token]
		if !ok {
			intensifierMult = 1.0
			continue
		}

		valence *= intensifierMult
		intensifierMult = 1.0

		if negationLeft > 0 {
			valence = -valence
		}

		if valence > 0 {
			posSum += valence
		} else {
			negSum += math.Abs(valence)
		}
	}

	return posSum, negSum
}

// tokenize lower-cases text and splits it into alphanumeric tokens,
// stripping punctuation and whitespace.
func tokenize(text string) []string {
	var tokens []string
	var buf strings.Builder

	for _, r := range strings.ToLower(text) {
		switch {
		case unicode.IsLetter(r) || unicode.IsDigit(r):
			buf.WriteRune(r)
		case r == '\'' || r == '\u2019':
			// skip apostrophes so contractions stay intact: "don’t" → "dont"
			continue
		case buf.Len() > 0:
			tokens = append(tokens, buf.String())
			buf.Reset()
		}
	}
	if buf.Len() > 0 {
		tokens = append(tokens, buf.String())
	}
	return tokens
}

// round2 rounds f to two decimal places.
func round2(f float64) float64 {
	return math.Round(f*100) / 100
}

// AnalyzeBatch analyzes a slice of texts and returns results in the same order.
func (s *Service) AnalyzeBatch(texts []string) []Result {
	results := make([]Result, len(texts))
	for i, text := range texts {
		results[i] = s.Analyze(text)
	}
	return results
}

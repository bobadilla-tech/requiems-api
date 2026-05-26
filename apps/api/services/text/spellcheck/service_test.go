package spellcheck

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/svcerr"
)

// newMockLT starts a test HTTP server that returns the given matches,
// simulating the LanguageTool /v2/check endpoint.
// It asserts the HTTP contract: the caller must use POST /v2/check.
func newMockLT(t *testing.T, matches []ltMatch) *httptest.Server {
	t.Helper()
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost || r.URL.Path != "/v2/check" {
			http.Error(w, "unexpected request: want POST /v2/check", http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(ltResponse{Matches: matches})
	}))
}

func TestService_Check_NoMistakes(t *testing.T) {
	t.Parallel()
	srv := newMockLT(t, []ltMatch{})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("This is a test")

	require.NoError(t, err)
	assert.Equal(t, "This is a test", result.Corrected)
	assert.Empty(t, result.Corrections)
}

func TestService_Check_MisspelledWords(t *testing.T) {
	t.Parallel()
	srv := newMockLT(t, []ltMatch{
		{Offset: 0, Length: 3, Replacements: []ltReplacement{{Value: "This"}}},
		{Offset: 9, Length: 4, Replacements: []ltReplacement{{Value: "test"}}},
	})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("Ths is a tset")

	require.NoError(t, err)
	require.Len(t, result.Corrections, 2)
	assert.Equal(t, "Ths", result.Corrections[0].Original)
	assert.Equal(t, "This", result.Corrections[0].Suggested)
	assert.Equal(t, 0, result.Corrections[0].Position)
	assert.Equal(t, "tset", result.Corrections[1].Original)
	assert.Equal(t, "test", result.Corrections[1].Suggested)
	assert.Equal(t, 9, result.Corrections[1].Position)
}

func TestService_Check_CorrectedTextReflectsFixes(t *testing.T) {
	t.Parallel()
	srv := newMockLT(t, []ltMatch{
		{Offset: 0, Length: 3, Replacements: []ltReplacement{{Value: "This"}}},
		{Offset: 9, Length: 4, Replacements: []ltReplacement{{Value: "test"}}},
	})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("Ths is a tset")

	require.NoError(t, err)
	assert.Equal(t, "This is a test", result.Corrected)
}

func TestService_Check_CorrectionsSliceNotNil(t *testing.T) {
	t.Parallel()
	srv := newMockLT(t, []ltMatch{})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("Hello world")

	require.NoError(t, err)
	assert.NotNil(t, result.Corrections)
}

func TestService_Check_Unreachable(t *testing.T) {
	t.Parallel()
	// Start and immediately close a server to get a guaranteed-dead URL
	// without relying on a fixed port that might be in use.
	dead := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	dead.Close()

	_, err := NewService(dead.URL).Check("hello")
	var svcErr *svcerr.Error
	require.ErrorAs(t, err, &svcErr)
	assert.Equal(t, svcerr.KindUpstream, svcErr.Kind)
	assert.Equal(t, "upstream_error", svcErr.Code)
}

func TestService_Check_PositionIsRuneOffset(t *testing.T) {
	t.Parallel()
	// "é" is one rune but two UTF-8 bytes. "tset" starts at rune index 2.
	srv := newMockLT(t, []ltMatch{
		{Offset: 2, Length: 4, Replacements: []ltReplacement{{Value: "test"}}},
	})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("é tset")

	require.NoError(t, err)
	require.Len(t, result.Corrections, 1)
	assert.Equal(t, 2, result.Corrections[0].Position)
}

func TestService_Check_SuggestionsTopThree(t *testing.T) {
	t.Parallel()
	// LanguageTool returns 4 replacements; we expect only the first 3.
	srv := newMockLT(t, []ltMatch{
		{Offset: 0, Length: 3, Replacements: []ltReplacement{
			{Value: "This"}, {Value: "The"}, {Value: "Thus"}, {Value: "Those"},
		}},
	})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("Ths is fine")

	require.NoError(t, err)
	require.Len(t, result.Corrections, 1)
	assert.Equal(t, "This", result.Corrections[0].Suggested)
	assert.Equal(t, []string{"This", "The", "Thus"}, result.Corrections[0].Suggestions)
}

func TestService_Check_SuggestionsOnlyOne(t *testing.T) {
	t.Parallel()
	// When LanguageTool returns a single replacement, Suggestions has length 1.
	srv := newMockLT(t, []ltMatch{
		{Offset: 0, Length: 4, Replacements: []ltReplacement{{Value: "test"}}},
	})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("tset done")

	require.NoError(t, err)
	require.Len(t, result.Corrections, 1)
	assert.Equal(t, []string{"test"}, result.Corrections[0].Suggestions)
}

func TestService_Check_LanguageToolNonOKStatus(t *testing.T) {
	t.Parallel()
	// Simulate LanguageTool returning a 500 — Check must propagate this as an upstream error.
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}))
	defer srv.Close()

	_, err := NewService(srv.URL).Check("hello")
	var svcErr *svcerr.Error
	require.ErrorAs(t, err, &svcErr)
	assert.Equal(t, svcerr.KindUpstream, svcErr.Kind)
	assert.Equal(t, "upstream_error", svcErr.Code)
}

func TestService_Check_NegativeOffsetIgnored(t *testing.T) {
	t.Parallel()
	// A match with a negative offset is malformed — it must be silently skipped,
	// leaving the corrected text equal to the input.
	srv := newMockLT(t, []ltMatch{
		{Offset: -1, Length: 3, Replacements: []ltReplacement{{Value: "This"}}},
	})
	defer srv.Close()

	result, err := NewService(srv.URL).Check("Ths is fine")

	require.NoError(t, err)
	assert.Empty(t, result.Corrections)
	assert.Equal(t, "Ths is fine", result.Corrected)
}

func TestService_CheckBatch_OrderPreserved(t *testing.T) {
	t.Parallel()
	// No matches — Corrected equals the input, so we can verify order by content.
	srv := newMockLT(t, []ltMatch{})
	defer srv.Close()

	texts := []string{"first", "second", "third"}
	results := NewService(srv.URL).CheckBatch(texts)

	require.Len(t, results, 3)
	for i, want := range texts {
		assert.Equal(t, want, results[i].Corrected, "results[%d] out of order", i)
	}
}

func TestService_CheckBatch_LTUnreachable(t *testing.T) {
	t.Parallel()
	// Dead server — each item gets a zero-value Result; the batch never errors out.
	dead := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	dead.Close()

	results := NewService(dead.URL).CheckBatch([]string{"hello", "world"})

	require.Len(t, results, 2)
	for i, r := range results {
		assert.Equal(t, "", r.Corrected, "results[%d].Corrected should be zero-value", i)
	}
}

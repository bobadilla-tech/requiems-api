package emoji

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
)

func setupRouter() chi.Router {
	r := chi.NewRouter()
	svc := NewService()
	RegisterRoutes(r, svc)
	return r
}

func TestEmoji_Random(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/random", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Emoji]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	e := resp.Data
	assert.NotEmpty(t, e.Emoji)
	assert.NotEmpty(t, e.Name)
	assert.NotEmpty(t, e.Category)
	assert.NotEmpty(t, e.Unicode)
}

func TestEmoji_GetByName_Found(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/grinning_face", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Emoji]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	e := resp.Data
	assert.Equal(t, "grinning_face", e.Name)
	assert.Equal(t, "😀", e.Emoji)
	assert.Equal(t, "Smileys & Emotion", e.Category)
	assert.Equal(t, "U+1F600", e.Unicode)
}

func TestEmoji_GetByName_CaseInsensitive(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/GRINNING_FACE", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code, "expected status 200 for uppercase name, got %d", w.Code)

	var resp httpx.Response[Emoji]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, "grinning_face", resp.Data.Name)
}

func TestEmoji_GetByName_NotFound(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/does_not_exist", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

func TestEmoji_Search_WithResults(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/search?q=happy", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[List]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.GreaterOrEqual(t, resp.Data.Total, 0)
}

func TestEmoji_Search_NoQuery(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/search", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code, "expected status 422 for missing query, got %d", w.Code)
}

func TestEmoji_Search_ReturnsMatchingEmojis(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/search?q=smile", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[List]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.NotEqual(t, 0, resp.Data.Total, "expected at least one result for 'smile'")

	for _, e := range resp.Data.Items {
		assert.True(t, e.Emoji != "" && e.Name != "" && e.Category != "" && e.Unicode != "", "expected all fields to be non-empty for emoji: %+v", e)
	}
}

func TestEmoji_Search_EmptyResults(t *testing.T) {
	t.Parallel()
	r := setupRouter()

	req := httptest.NewRequest(http.MethodGet, "/emoji/search?q=zzzyyyxxx", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[List]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)

	assert.Equal(t, 0, resp.Data.Total, "expected 0 results for nonsense query, got %d", resp.Data.Total)
	assert.NotNil(t, resp.Data.Items, "expected non-nil items slice for empty results")
}

func TestService_Random_ReturnsValidEmoji(t *testing.T) {
	t.Parallel()
	svc := NewService()

	e := svc.Random()
	assert.NotEmpty(t, e.Emoji, "expected non-empty emoji from Random()")
	assert.NotEmpty(t, e.Name, "expected non-empty name from Random()")
}

func TestService_GetByName(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name      string
		input     string
		wantFound bool
	}{
		{name: "valid name", input: "grinning_face", wantFound: true},
		{name: "valid name uppercase", input: "GRINNING_FACE", wantFound: true},
		{name: "unknown name", input: "not_a_real_emoji", wantFound: false},
		{name: "empty string", input: "", wantFound: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			_, found := svc.GetByName(tt.input)
			assert.Equal(t, tt.wantFound, found, "GetByName(%q): got found=%v, want %v", tt.input, found, tt.wantFound)
		})
	}
}

func TestService_Search(t *testing.T) {
	t.Parallel()
	svc := NewService()

	tests := []struct {
		name      string
		query     string
		wantEmpty bool
	}{
		{name: "smile query returns results", query: "smile", wantEmpty: false},
		{name: "heart query returns results", query: "heart", wantEmpty: false},
		{name: "nonsense returns empty", query: "zzzyyyxxx", wantEmpty: true},
		{name: "category search", query: "food", wantEmpty: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			result := svc.Search(tt.query)
			isEmpty := result.Total == 0
			assert.Equal(t, tt.wantEmpty, isEmpty, "Search(%q): got empty=%v, want empty=%v (total=%d)", tt.query, isEmpty, tt.wantEmpty, result.Total)
			assert.Equal(t, result.Total, len(result.Items), "Search(%q): Total=%d does not match len(Items)=%d", tt.query, result.Total, len(result.Items))
		})
	}
}

package exercises

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"requiems-api/platform/httpx"
	"requiems-api/platform/svcerr"
)

// stubQuerier implements exerciseQuerier for HTTP handler tests.
type stubQuerier struct {
	listResult   ExerciseList
	getResult    Exercise
	randomResult Exercise
	stringResult StringList
	err          error
}

func (s *stubQuerier) List(_ context.Context, _ ListParams) (ExerciseList, error) {
	return s.listResult, s.err
}

func (s *stubQuerier) Get(_ context.Context, _ int) (Exercise, error) {
	if s.err != nil {
		return Exercise{}, s.err
	}
	return s.getResult, nil
}

func (s *stubQuerier) Random(_ context.Context, _ ListParams) (Exercise, error) {
	if s.err != nil {
		return Exercise{}, s.err
	}
	return s.randomResult, nil
}

func (s *stubQuerier) BodyParts(_ context.Context) (StringList, error) {
	return s.stringResult, s.err
}

func (s *stubQuerier) Equipment(_ context.Context) (StringList, error) {
	return s.stringResult, s.err
}

func (s *stubQuerier) Muscles(_ context.Context) (StringList, error) {
	return s.stringResult, s.err
}

func setupTestRouter(q exerciseQuerier) chi.Router {
	r := chi.NewRouter()
	registerExerciseRoutes(r, q)
	return r
}

// ---- GET /exercises ----

func TestListExercises_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{
		listResult: ExerciseList{
			Items:   []Exercise{{ID: 1, Name: "squat"}},
			Total:   1,
			Page:    1,
			PerPage: 20,
		},
	}

	r := setupTestRouter(stub)
	req := httptest.NewRequest(http.MethodGet, "/exercises", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ExerciseList]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, 1, resp.Data.Total)
	assert.Len(t, resp.Data.Items, 1)
}

func TestListExercises_ResponseEnvelope(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{listResult: ExerciseList{Items: []Exercise{}}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var raw map[string]json.RawMessage
	err := json.NewDecoder(w.Body).Decode(&raw)
	require.NoError(t, err)
	if _, ok := raw["data"]; !ok {
		t.Error("response must have a 'data' key")
	}
	if _, ok := raw["metadata"]; !ok {
		t.Error("response must have a 'metadata' key")
	}
}

func TestListExercises_ServiceError_Returns500(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{err: errors.New("db down")}

	r := setupTestRouter(stub)
	req := httptest.NewRequest(http.MethodGet, "/exercises", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListExercises_InvalidPerPage_Returns400(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises?per_page=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

// ---- GET /exercises/random ----

func TestRandomExercise_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{randomResult: Exercise{ID: 7, Name: "deadlift"}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/random", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Exercise]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "deadlift", resp.Data.Name)
}

func TestRandomExercise_NoMatch_Returns404(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{err: svcerr.NotFound("not_found", "no exercises found")}

	r := setupTestRouter(stub)
	req := httptest.NewRequest(http.MethodGet, "/exercises/random?body_part=unknown", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

// ---- GET /exercises/{id} ----

func TestGetExercise_ValidID_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{getResult: Exercise{ID: 3, Name: "bench press"}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/3", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Exercise]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, "bench press", resp.Data.Name)
}

func TestGetExercise_NonNumericID_Returns400(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/abc", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetExercise_ZeroID_Returns400(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetExercise_NotFound_Returns404(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{err: svcerr.NotFound("not_found", "exercise not found")}

	r := setupTestRouter(stub)
	req := httptest.NewRequest(http.MethodGet, "/exercises/9999", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusNotFound, w.Code)
}

// ---- GET /body-parts, /equipment, /muscles ----

func TestBodyParts_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{stringResult: StringList{Items: []string{"chest", "back"}, Total: 2}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/body-parts", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[StringList]
	err := json.NewDecoder(w.Body).Decode(&resp)
	require.NoError(t, err)
	assert.Equal(t, 2, resp.Data.Total)
}

func TestEquipment_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{stringResult: StringList{Items: []string{"barbell", "dumbbell"}, Total: 2}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/equipment", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestMuscles_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{stringResult: StringList{Items: []string{"biceps", "triceps"}, Total: 2}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/muscles", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestMetadata_ServiceError_Returns500(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name string
		path string
	}{
		{"body-parts", "/body-parts"},
		{"equipment", "/equipment"},
		{"muscles", "/muscles"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			stub := &stubQuerier{err: errors.New("db error")}

			r := setupTestRouter(stub)
			req := httptest.NewRequest(http.MethodGet, tt.path, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	}
}

package exercises

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"

	"requiems-api/platform/httpx"
)

// stubQuerier implements exerciseQuerier for HTTP handler tests.
type stubQuerier struct {
	listResult   ExerciseList
	getResult    Exercise
	randomResult Exercise
	batchResult  []Exercise
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

func (s *stubQuerier) GetBatch(_ context.Context, _ []int) ([]Exercise, error) {
	return s.batchResult, s.err
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

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[ExerciseList]
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
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
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&raw))
	assert.Contains(t, raw, "data", "response must have a 'data' key")
	assert.Contains(t, raw, "metadata", "response must have a 'metadata' key")
}

func TestListExercises_ServiceError_Returns500(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{err: &httpx.AppError{
		Status: http.StatusInternalServerError, Code: "internal_error", Message: "db down",
	}}

	r := setupTestRouter(stub)
	req := httptest.NewRequest(http.MethodGet, "/exercises", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
}

func TestListExercises_InvalidPerPage_Returns400(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises?per_page=0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

// ---- GET /exercises/random ----

func TestRandomExercise_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{randomResult: Exercise{ID: 7, Name: "deadlift"}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/random", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Exercise]
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "deadlift", resp.Data.Name)
}

func TestRandomExercise_NoMatch_Returns404(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{err: &httpx.AppError{
		Status: http.StatusNotFound, Code: "not_found", Message: "no exercises found",
	}}

	r := setupTestRouter(stub)
	req := httptest.NewRequest(http.MethodGet, "/exercises/random?body_part=unknown", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// ---- GET /exercises/{id} ----

func TestGetExercise_ValidID_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{getResult: Exercise{ID: 3, Name: "bench press"}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/3", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[Exercise]
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, "bench press", resp.Data.Name)
}

func TestGetExercise_NonNumericID_Returns400(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/abc", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetExercise_ZeroID_Returns400(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/exercises/0", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestGetExercise_NotFound_Returns404(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{err: &httpx.AppError{
		Status: http.StatusNotFound, Code: "not_found", Message: "exercise not found",
	}}

	r := setupTestRouter(stub)
	req := httptest.NewRequest(http.MethodGet, "/exercises/9999", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}

// ---- GET /body-parts, /equipment, /muscles ----

func TestBodyParts_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{stringResult: StringList{Items: []string{"chest", "back"}, Total: 2}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/body-parts", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[StringList]
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Equal(t, 2, resp.Data.Total)
}

func TestEquipment_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{stringResult: StringList{Items: []string{"barbell", "dumbbell"}, Total: 2}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/equipment", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestMuscles_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{stringResult: StringList{Items: []string{"biceps", "triceps"}, Total: 2}}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodGet, "/muscles", http.NoBody)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

// ---- POST /exercises/batch ----

func TestBatchExercises_Returns200(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{
		batchResult: []Exercise{
			{ID: 1, Name: "squat"},
			{ID: 7, Name: "deadlift"},
		},
	}
	r := setupTestRouter(stub)

	body := `{"ids": [1, 7]}`
	req := httptest.NewRequest(http.MethodPost, "/exercises/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var resp httpx.Response[BatchExerciseResponse]
	assert.NoError(t, json.NewDecoder(w.Body).Decode(&resp))
	assert.Len(t, resp.Data.Results, 2)
	assert.Equal(t, 2, resp.Data.Total)
	assert.Equal(t, "squat", resp.Data.Results[0].Name)
}

func TestBatchExercises_EmptyIDs_Returns422(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	body := `{"ids": []}`
	req := httptest.NewRequest(http.MethodPost, "/exercises/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatchExercises_TooManyIDs_Returns422(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	ids := make([]int, 51)
	for i := range ids {
		ids[i] = i + 1
	}
	bodyBytes, _ := json.Marshal(map[string][]int{"ids": ids})
	req := httptest.NewRequest(http.MethodPost, "/exercises/batch", strings.NewReader(string(bodyBytes)))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnprocessableEntity, w.Code)
}

func TestBatchExercises_InvalidBody_Returns400(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{}
	r := setupTestRouter(stub)

	req := httptest.NewRequest(http.MethodPost, "/exercises/batch", strings.NewReader(`not json`))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestBatchExercises_ServiceError_Returns500(t *testing.T) {
	t.Parallel()
	stub := &stubQuerier{err: &httpx.AppError{
		Status: http.StatusInternalServerError, Code: "internal_error", Message: "db down",
	}}
	r := setupTestRouter(stub)

	body := `{"ids": [1, 2, 3]}`
	req := httptest.NewRequest(http.MethodPost, "/exercises/batch", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
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
			stub := &stubQuerier{err: &httpx.AppError{
				Status: http.StatusInternalServerError, Code: "internal_error", Message: "db error",
			}}

			r := setupTestRouter(stub)
			req := httptest.NewRequest(http.MethodGet, tt.path, http.NoBody)
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			assert.Equal(t, http.StatusInternalServerError, w.Code)
		})
	}
}

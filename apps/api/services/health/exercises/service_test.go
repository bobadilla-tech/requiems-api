package exercises

import (
	"context"
	"errors"
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestExercise_FieldsPresent(t *testing.T) {
	t.Parallel()
	e := Exercise{
		ID:               1,
		Name:             "band shrug",
		BodyParts:        []string{"neck"},
		Equipment:        []string{"band"},
		TargetMuscles:    []string{"traps"},
		SecondaryMuscles: []string{"shoulders"},
		Instructions:     []string{"Stand with feet shoulder-width apart."},
	}

	assert.Equal(t, 1, e.ID)
	assert.Equal(t, "band shrug", e.Name)
	assert.Equal(t, []string{"neck"}, e.BodyParts)
	assert.Len(t, e.Instructions, 1)
}

func TestExerciseList_Pagination(t *testing.T) {
	t.Parallel()
	l := ExerciseList{
		Items:   []Exercise{{ID: 1, Name: "squat"}, {ID: 2, Name: "deadlift"}},
		Total:   50,
		Page:    2,
		PerPage: 20,
	}

	assert.Equal(t, 50, l.Total)
	assert.Equal(t, 2, l.Page)
	assert.Equal(t, 20, l.PerPage)
	assert.Len(t, l.Items, 2)
}

func TestStringList_ItemsAndTotal(t *testing.T) {
	t.Parallel()
	s := StringList{
		Items: []string{"chest", "back", "legs"},
		Total: 3,
	}

	assert.Equal(t, 3, s.Total)
	assert.Len(t, s.Items, 3)
}

// TestListParams_ZeroValue verifies the zero value of ListParams has empty
// filter fields. The page/per_page defaults (1/20) are applied by the HTTP
// handler before binding, not by the struct itself.
func TestListParams_ZeroValue(t *testing.T) {
	t.Parallel()
	var p ListParams

	assert.Empty(t, p.BodyPart)
	assert.Empty(t, p.Equipment)
	assert.Empty(t, p.Muscle)
	assert.Empty(t, p.Search)
	assert.Zero(t, p.Page)
	assert.Zero(t, p.PerPage)
}

// TestBatchGetRequest_IDsPreservedInOrder verifies that the IDs slice
// maintains input order — the contract that the SQL query (ORDER BY
// array_position) is expected to honour at the DB level.
func TestBatchGetRequest_IDsPreservedInOrder(t *testing.T) {
	t.Parallel()
	req := BatchGetRequest{IDs: []int{7, 1, 42}}

	assert.Equal(t, []int{7, 1, 42}, req.IDs)
	assert.Equal(t, 7, req.IDs[0])
	assert.Equal(t, 1, req.IDs[1])
	assert.Equal(t, 42, req.IDs[2])
}

// TestBatchGetRequest_SingleItem verifies that a one-element list is valid input.
func TestBatchGetRequest_SingleItem(t *testing.T) {
	t.Parallel()
	req := BatchGetRequest{IDs: []int{5}}

	assert.Len(t, req.IDs, 1)
	assert.Equal(t, 5, req.IDs[0])
}

// ---- mockDB / mockRows — test doubles for GetBatch unit tests ----

// mockDB implements dbPool without a live database connection.
type mockDB struct {
	rows     dbRows
	queryErr error
}

func (m *mockDB) Query(_ context.Context, _ string, _ ...any) (dbRows, error) {
	return m.rows, m.queryErr
}

func (m *mockDB) QueryRow(_ context.Context, _ string, _ ...any) pgx.Row {
	panic("QueryRow not expected in GetBatch tests")
}

// mockRows implements dbRows, feeding a fixed slice of exercises to the scanner.
type mockRows struct {
	data    []Exercise
	idx     int
	scanErr error
	rowsErr error
}

func (m *mockRows) Close()     {}
func (m *mockRows) Err() error { return m.rowsErr }

func (m *mockRows) Next() bool {
	return m.idx < len(m.data)
}

func (m *mockRows) Scan(dest ...any) error {
	if m.scanErr != nil {
		return m.scanErr
	}
	e := m.data[m.idx]
	m.idx++
	*dest[0].(*int) = e.ID
	*dest[1].(*string) = e.Name
	*dest[2].(*[]string) = e.BodyParts
	*dest[3].(*[]string) = e.Equipment
	*dest[4].(*[]string) = e.TargetMuscles
	*dest[5].(*[]string) = e.SecondaryMuscles
	*dest[6].(*[]string) = e.Instructions
	return nil
}

// ---- GetBatch unit tests ----

// TestGetBatch_HappyPath verifies that GetBatch returns exercises in input order
// when the database scan succeeds.
func TestGetBatch_HappyPath(t *testing.T) {
	t.Parallel()
	rows := &mockRows{data: []Exercise{
		{
			ID: 1, Name: "squat",
			BodyParts: []string{"upper legs"}, Equipment: []string{"barbell"},
			TargetMuscles: []string{"quads"}, SecondaryMuscles: []string{"glutes"},
			Instructions: []string{"Stand with feet shoulder-width apart."},
		},
		{
			ID: 7, Name: "deadlift",
			BodyParts: []string{"back"}, Equipment: []string{"barbell"},
			TargetMuscles: []string{"spine"}, SecondaryMuscles: []string{"hamstrings"},
			Instructions: []string{"Hinge at the hips."},
		},
	}}
	svc := &Service{db: &mockDB{rows: rows}}

	results, err := svc.GetBatch(context.Background(), []int{1, 7})
	require.NoError(t, err)
	require.Len(t, results, 2)
	assert.Equal(t, 1, results[0].ID)
	assert.Equal(t, "squat", results[0].Name)
	assert.Equal(t, 7, results[1].ID)
	assert.Equal(t, "deadlift", results[1].Name)
}

// TestGetBatch_QueryError verifies that a database query failure is propagated as an error.
func TestGetBatch_QueryError(t *testing.T) {
	t.Parallel()
	svc := &Service{db: &mockDB{queryErr: errors.New("connection refused")}}

	results, err := svc.GetBatch(context.Background(), []int{1})
	require.Error(t, err)
	assert.Nil(t, results)
}

// TestGetBatch_ScanError verifies that a row scan failure is propagated as an error.
func TestGetBatch_ScanError(t *testing.T) {
	t.Parallel()
	rows := &mockRows{
		data:    []Exercise{{ID: 1, Name: "squat"}},
		scanErr: errors.New("scan failed"),
	}
	svc := &Service{db: &mockDB{rows: rows}}

	results, err := svc.GetBatch(context.Background(), []int{1})
	require.Error(t, err)
	assert.Nil(t, results)
}

// TestGetBatch_RowsError verifies that an error surfaced by rows.Err() after
// the loop is propagated correctly.
func TestGetBatch_RowsError(t *testing.T) {
	t.Parallel()
	rows := &mockRows{rowsErr: errors.New("rows iteration error")}
	svc := &Service{db: &mockDB{rows: rows}}

	results, err := svc.GetBatch(context.Background(), []int{1})
	require.Error(t, err)
	assert.Nil(t, results)
}

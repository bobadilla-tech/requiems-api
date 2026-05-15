package exercises

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExercise_IsData(t *testing.T) {
	t.Parallel()
	var e Exercise
	e.IsData()
}

func TestExerciseList_IsData(t *testing.T) {
	t.Parallel()
	var l ExerciseList
	l.IsData()
}

func TestStringList_IsData(t *testing.T) {
	t.Parallel()
	var s StringList
	s.IsData()
}

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

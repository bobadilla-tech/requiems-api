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
	if len(e.BodyParts) != 1 || e.BodyParts[0] != "neck" {
		t.Errorf("unexpected body_parts: %v", e.BodyParts)
	}
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

	if p.BodyPart != "" || p.Equipment != "" || p.Muscle != "" || p.Search != "" {
		t.Error("expected all filter fields to be empty on zero value")
	}
	if p.Page != 0 || p.PerPage != 0 {
		t.Error("expected page and per_page to be zero on zero value")
	}
}

package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoadExercises_HappyPath(t *testing.T) {
	t.Parallel()

	exercises := []map[string]interface{}{
		{
			"exerciseId":       "ex001",
			"name":             "Push-up",
			"bodyParts":        []string{"chest"},
			"equipments":       []string{"body weight"},
			"targetMuscles":    []string{"pectorals"},
			"secondaryMuscles": []string{"triceps"},
			"instructions":     []string{"Step:1 Get into position.", "Step:2 Lower your body."},
			"gifUrl":           "https://example.com/push-up.gif", // must be silently discarded
		},
	}

	dir := writeExercisesJSON(t, exercises)
	records, err := loadExercises(dir)
	require.NoError(t, err)
	require.Len(t, records, 1)

	r := records[0]
	assert.Equal(t, "ex001", r.ExternalID)
	assert.Equal(t, "Push-up", r.Name)
	// Step prefix should be stripped from instructions.
	if assert.Len(t, r.Instructions, 2) {
		assert.Equal(t, "Get into position.", r.Instructions[0])
	}
}

func TestLoadExercises_SkipsMissingIDOrName(t *testing.T) {
	t.Parallel()

	exercises := []map[string]interface{}{
		{"exerciseId": "", "name": "No ID Exercise"},               // missing ID
		{"exerciseId": "ex002", "name": ""},                        // missing name
		{"exerciseId": "ex003", "name": "Valid", "bodyParts": nil}, // valid — nil fields normalised
	}

	dir := writeExercisesJSON(t, exercises)
	records, err := loadExercises(dir)
	require.NoError(t, err)
	require.Len(t, records, 1)
	assert.Equal(t, "ex003", records[0].ExternalID)
}

func TestLoadExercises_NilSlicesNormalised(t *testing.T) {
	t.Parallel()

	exercises := []map[string]any{
		{"exerciseId": "ex004", "name": "Squat"}, // omit all slice fields
	}

	dir := writeExercisesJSON(t, exercises)
	records, err := loadExercises(dir)
	require.NoError(t, err)
	r := records[0]

	// All slice fields should be non-nil empty slices, not nil.
	assert.NotNil(t, r.BodyParts)
	assert.NotNil(t, r.Equipment)
	assert.NotNil(t, r.TargetMuscles)
	assert.NotNil(t, r.SecondaryMuscles)
}

func TestLoadExercises_MissingFile(t *testing.T) {
	t.Parallel()

	_, err := loadExercises(t.TempDir())
	require.Error(t, err)
}

func TestLoadExercises_InvalidJSON(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()
	err := os.WriteFile(filepath.Join(dir, "exercises.json"), []byte("not-json"), 0o600)
	require.NoError(t, err)

	_, err = loadExercises(dir)
	require.Error(t, err)
}

// writeExercisesJSON marshals exercises to exercises.json in a temp dir and
// returns the dir path.
func writeExercisesJSON(t *testing.T, exercises interface{}) string {
	t.Helper()

	dir := t.TempDir()
	data, err := json.Marshal(exercises)
	require.NoError(t, err)
	err = os.WriteFile(filepath.Join(dir, "exercises.json"), data, 0o600)
	require.NoError(t, err)
	return dir
}

package exercises

import "requiems-api/platform/httpx"

// Exercise is the public representation of an exercise record.
// external_id is intentionally omitted — it is an internal upsert key only.
type Exercise struct {
	ID               int      `json:"id"`
	Name             string   `json:"name"`
	BodyParts        []string `json:"body_parts"`
	Equipment        []string `json:"equipment"`
	TargetMuscles    []string `json:"target_muscles"`
	SecondaryMuscles []string `json:"secondary_muscles"`
	Instructions     []string `json:"instructions"`
}

func (Exercise) IsData() {}

// ExerciseList wraps a paginated set of exercises.
type ExerciseList struct {
	Items   []Exercise `json:"items"`
	Total   int        `json:"total"`
	Page    int        `json:"page"`
	PerPage int        `json:"per_page"`
}

func (ExerciseList) IsData() {}

// StringList wraps a sorted list of unique string values (muscles, equipment, body parts).
type StringList struct {
	Items []string `json:"items"`
	Total int      `json:"total"`
}

func (StringList) IsData() {}

// ListParams holds the query parameters accepted by the list and random endpoints.
type ListParams struct {
	BodyPart  string `query:"body_part"`
	Equipment string `query:"equipment"`
	Muscle    string `query:"muscle"`
	Search    string `query:"search"`
	Page      int    `query:"page"     validate:"min=1"`
	PerPage   int    `query:"per_page" validate:"min=1,max=100"`
}

// BatchGetRequest is the body for fetching multiple exercises by ID.
type BatchGetRequest struct {
	IDs []int `json:"ids" validate:"required,min=1,max=50,dive,min=1"`
}

// BatchExerciseResponse is the response for a batch exercise lookup.
type BatchExerciseResponse = httpx.BatchResponse[Exercise]

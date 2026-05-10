package httpx

// FieldError describes a single field-level validation failure.
type FieldError struct {
	Field string `json:"field"`
	Rule  string `json:"rule"`
}

// ValidationFailure is returned by BindAndValidate when struct validation fails.
// Handlers can check for this type to distinguish validation errors from other
// decode failures.
type ValidationFailure struct {
	Fields []FieldError
}

func (e *ValidationFailure) Error() string { return "validation_failed" }


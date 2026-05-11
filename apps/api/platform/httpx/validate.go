package httpx

import (
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Validate is the package-level validator instance. Custom rules can be
// registered via Validate.RegisterValidation before first use.
var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())

	// Report errors using the JSON field name rather than the Go struct field name.
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})
}

// transformValidationErrors converts go-playground/validator errors into a
// structured []FieldError slice, using JSON field names set above.
func transformValidationErrors(errs validator.ValidationErrors) []FieldError {
	out := make([]FieldError, 0, len(errs))
	for _, e := range errs {
		out = append(out, FieldError{
			Field: e.Field(),
			Rule:  e.Tag(),
		})
	}
	return out
}

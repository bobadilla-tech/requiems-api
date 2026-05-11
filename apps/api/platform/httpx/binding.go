package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
)

// BindAndValidate decodes the JSON request body into dst and runs struct
// validation using the Validate instance. Unknown JSON fields cause a decode
// error. Returns *ValidationFailure for constraint violations, or a plain
// error for malformed JSON.
//
// Note: body size limiting is handled by the Handle wrapper via
// http.MaxBytesReader before this function is called.
func BindAndValidate(r *http.Request, dst any) error {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return fmt.Errorf("invalid JSON: %w", err)
	}

	return validateStruct(dst)
}

// BindQuery decodes URL query parameters into dst using `query:"name"` struct
// tags, then validates the struct. Only string, int*, float*, and bool field
// types are supported. Fields with no matching query param are left at their
// current value (defaults can be set before calling BindQuery).
func BindQuery(r *http.Request, dst any) error {
	v := reflect.ValueOf(dst)
	if v.Kind() != reflect.Pointer || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("BindQuery: dst must be a pointer to a struct")
	}

	v = v.Elem()
	t := v.Type()
	q := r.URL.Query()

	for i := range t.NumField() {
		field := t.Field(i)

		tag := field.Tag.Get("query")
		if tag == "" || tag == "-" {
			continue
		}

		raw := q.Get(tag)
		if raw == "" {
			continue
		}

		if err := setFieldValue(v.Field(i), field.Type.Kind(), raw, tag); err != nil {
			return err
		}
	}

	return validateStruct(dst)
}

// setFieldValue sets a struct field from a raw string value, converting to the
// appropriate type.
func setFieldValue(fv reflect.Value, kind reflect.Kind, raw, tag string) error {
	switch kind {
	case reflect.String:
		fv.SetString(raw)

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		n, err := strconv.ParseInt(raw, 10, 64)
		if err != nil {
			return fmt.Errorf("invalid value for %q: must be an integer", tag)
		}
		fv.SetInt(n)

	case reflect.Float32, reflect.Float64:
		f, err := strconv.ParseFloat(raw, 64)
		if err != nil {
			return fmt.Errorf("invalid value for %q: must be a number", tag)
		}
		fv.SetFloat(f)

	case reflect.Bool:
		b, err := strconv.ParseBool(raw)
		if err != nil {
			return fmt.Errorf("invalid value for %q: must be true or false", tag)
		}
		fv.SetBool(b)

	case reflect.Struct:
		if fv.Type() == reflect.TypeFor[time.Time]() {
			t, err := time.ParseInLocation("2006-01-02", raw, time.UTC)
			if err != nil {
				return fmt.Errorf("invalid value for %q: must be a date in YYYY-MM-DD format", tag)
			}
			fv.Set(reflect.ValueOf(t))
			return nil
		}
	}

	return nil
}

// validateStruct runs Validate.Struct on dst and converts ValidationErrors to
// *ValidationFailure.
func validateStruct(dst any) error {
	if err := Validate.Struct(dst); err != nil {
		if ve, ok := errors.AsType[validator.ValidationErrors](err); ok {
			return &ValidationFailure{Fields: transformValidationErrors(ve)}
		}

		return err
	}

	return nil
}

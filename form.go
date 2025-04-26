package bubbleform

import (
	"fmt"
	"reflect"

	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

type Form struct {
	data   interface{}
	fields map[string]Field
	errors map[string]error
}

func NewForm(data interface{}) *Form {
	return &Form{
		data:   data,
		fields: make(map[string]Field),
		errors: make(map[string]error),
	}
}

// Bind a field to a specific struct field
func (f *Form) Bind(key string, field Field) error {
	val := reflect.ValueOf(f.data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if val.Kind() != reflect.Struct {
		return fmt.Errorf("data must be a struct or a pointer to a struct")
	}

	fieldVal := val.FieldByName(key)
	if !fieldVal.IsValid() {
		return fmt.Errorf("no such field: %s", key)
	}

	f.fields[key] = field
	return nil
}

// Validate all form fields and populate the data struct
func (f *Form) Validate() bool {
	val := reflect.ValueOf(f.data)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	// Clear previous errors
	f.errors = make(map[string]error)
	allValid := true

	// First populate the data with current field values
	for key, field := range f.fields {
		fieldVal := val.FieldByName(key)
		if !fieldVal.IsValid() || !fieldVal.CanSet() {
			continue
		}

		// Field-level validation if available
		if validator, ok := field.(interface{ Validate() error }); ok {
			if err := validator.Validate(); err != nil {
				f.errors[key] = err
				field.SetError(err)
				allValid = false
				continue
			}
		}

		// Set value (currently handling strings only)
		if fieldVal.Kind() == reflect.String {
			fieldVal.SetString(field.Value())
		}
		// Add other type conversions as needed
	}

	// Then validate the entire struct
	if err := validate.Struct(f.data); err != nil {
		if validationErrors, ok := err.(validator.ValidationErrors); ok {
			for _, validationErr := range validationErrors {
				fieldName := validationErr.Field()
				if field, ok := f.fields[fieldName]; ok {
					errMsg := fmt.Errorf("%s: %s", validationErr.Tag(), validationErr.Error())
					f.errors[fieldName] = errMsg
					field.SetError(errMsg)
					allValid = false
				}
			}
		}
	}

	return allValid
}

// Submit validates and populates the data struct
func (f *Form) Submit() bool {
	return f.Validate()
}

// GetErrors returns all validation errors
func (f *Form) GetErrors() map[string]error {
	return f.errors
}

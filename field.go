package bubbleform

type Field interface {
	// Get the current value of the field
	Value() string

	// Set validation error on the field
	SetError(err error)

	// Optional method to reset field errors
	ClearError()

	// Optional method for field-level validation
	// Validate() error
}

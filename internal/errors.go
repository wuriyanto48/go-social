package internal

import "fmt"

// ErrorEmptyValue structure
type ErrorEmptyValue struct {
	field   string
	message string
}

// NewErrorEmptyValue function
func NewErrorEmptyValue(field string) *ErrorEmptyValue {
	return &ErrorEmptyValue{field: field, message: "%s cannot be emtpy"}
}

// Error function
func (e *ErrorEmptyValue) Error() string {
	return fmt.Sprintf(e.message, e.field)
}

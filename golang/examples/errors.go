package examples

import (
	"errors"
	"fmt"
)

type operationStringError struct {
	message string
	cause error
}
func (e operationStringError) Error() string {
	return fmt.Sprintf("%s: %s", e.message, e.cause)
}

// Examples of manipulating errors.
// It’s possible to use custom types as errors by
// implementing the Error() method on them.
func ErrorsExamples() {
	op := operation{
		code: 10,
		status: StatusError,
	}
	fmt.Printf("(errors) operation: %s\n", op.StringOrPanic())
}

func (op operation) StringOrPanic() string {
	status, err := op.String()

	if err != nil {
		if errors.Is(err, errors.ErrUnsupported) {
			panic(fmt.Errorf("error not supported: %w", err))
		}
		var customErr error = operationStringError{
			message: "operation cannot be converted to string",
			cause: err,
		}
		panic(customErr)			// panic: operation cannot be converted to string: status cannot be converted to string: invalid type
		//panic(err)				// panic: status cannot be converted to string: invalid type
		//panic(errors.Unwrap(err))	// panic: invalid type
	}
	
	return status
}

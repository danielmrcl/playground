package examples

import (
	"errors"
	"fmt"
)

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
		panic(err)					// panic: status cannot be converted to string: invalid type
		//panic(errors.Unwrap(err))	// panic: invalid type
	}
	
	return status
}

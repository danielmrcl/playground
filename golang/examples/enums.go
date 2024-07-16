package examples

import "fmt"

type Status uint8
const (
    StatusSuccess = iota // StatusSuccess = 0
    StatusError          // StatusError   = 1
)
func (s Status) String() (string, error) {
    switch s {
	case StatusSuccess:
		return "sucesso", nil
	case StatusError:
		return "error", nil
	default:
		return "", fmt.Errorf("invalid type")
	}
}
type operation struct {
	code uint16
	status Status
}
func (o operation) String() (string, error) {
	if status, err := o.status.String(); err == nil {
		return fmt.Sprintf("{%d %s}", o.code, status), nil
	} else {
		return "", fmt.Errorf("status cannot be converted to string: %w", err)
	}
}

// Enums: Go doesn’t have an enum type as a distinct
// language feature, but enums are simple to implement
// using existing language idioms.
// (type + const + map/switch)
func EnumsExamples() {
	fmt.Println("(enum) status:", operation{23, StatusSuccess})
	fmt.Println("(enum) status:", operation{24, StatusError})
}


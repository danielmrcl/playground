package examples

import "fmt"

type Status uint8
const (
    StatusSuccess = iota // StatusSuccess = 0
    StatusError          // StatusError   = 1
)
func (s Status) String() string {
    switch s {
	case StatusSuccess:
		return "sucesso"
	case StatusError:
		return "error"
	default:
		return "" // TODO: return error
	}
}
type operation struct {
	code uint16
	status Status
}
func (o operation) String() string {
	return fmt.Sprintf("{%d %s}", o.code, o.status.String())
}

// Enums: Go doesn’t have an enum type as a distinct
// language feature, but enums are simple to implement
// using existing language idioms.
// (type + const + map/switch)
func EnumsExamples() {
	fmt.Println("(enum) status:", operation{23, StatusSuccess})
	fmt.Println("(enum) status:", operation{24, StatusError})
}


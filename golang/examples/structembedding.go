package examples

import "fmt"

type update_action struct {
	operation
	new_value string
}

func (ua update_action) String() string {
	return fmt.Sprintf("{%d %s %s}", ua.code, ua.status, ua.new_value) // {1 sucesso Hello!}
}

// Struct Embedding:
// composition of types.
func StructEmbeddingExample() {
	ua := update_action{
		operation: operation{
			code: 1,
			status: StatusSuccess,
		},
		new_value: "Hello!",
	}

	fmt.Printf("(struct embedding) update_action: %s\n", ua)
}

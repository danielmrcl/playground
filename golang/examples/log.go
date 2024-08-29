package examples

import "fmt"

type logger struct {
	name string
}

func (log logger) log(message string) {
	fmt.Printf("(%s) %s\n", log.name, message)
}

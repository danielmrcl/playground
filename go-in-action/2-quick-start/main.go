package main

import (
	"log"
	"os"
	_ "danielmrcl.dev/go-in-action/matchers" // imported just to run its `init()` function.
	"danielmrcl.dev/go-in-action/search"
)

func init() {
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}

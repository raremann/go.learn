package main

import (
	"fmt"
	"os"
)

func main() {
	// variable without explicit initialization are set to '0' for the type, which is "" for string
	var s, sep string
	// No parenthesis allowed. for is the only loop construct in go.
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}
	fmt.Println(s)
}

// echo2 prints its command line args
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", " "
	// range produces the index and value tuple at each step
	// '_' is special and indicates that element of returned tuple will not be used in computation
	for i, arg := range os.Args[0:] {
		// += makes a new string on each execution. the old string will eventually be gc'ed
		s += sep + arg
		fmt.Println("index, value: ", i, arg)
	}
	fmt.Println(s)
}

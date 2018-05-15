// fib computes the Fibonacci number iteratively
package main

import (
	"flag"
	"fmt"
)

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		x, y = y, x+y
	}
	return x
}

var n = flag.Int("n", 1, "Sequence element to compute")

func main() {
	flag.Parse()
	answer := fib(*n)
	fmt.Printf("Fib sequnece element %d = %d\n", *n, answer)
}

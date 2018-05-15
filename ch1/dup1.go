// dup1 prints the text of each line that appears more than once on stdin
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map -- key must be type that supports == operations
	counts := make(map[string]int)
	// Scanner type reads and parses by newline
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		// First time counts[key] is accessed the value is '0' for it's type, int and 0.
		counts[input.Text()]++
	}
	// TODO: handle errors from input.Err()
	// map iteration order not defined; explicitly randomized to catch invalid assumptions
	for line, n := range counts {
		// No partheneses, ever. Braces required
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

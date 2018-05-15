// dup3 print the count and text of lines that appear
// more than once in the named input files

// Interestingly, this programs produces a difference result than dup1/dup2, seemingly do to
// the handling of a blank line at the end of the test files
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		// read all the bytes at once into data
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		// ReadFile returns void data; must convert it to string type
		for _, line := range strings.Split(string(data), "\n") {
			fmt.Println("line ", line)
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

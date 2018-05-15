// dup2 priints the count and test of lines that appear mor than once
// in the input. It reads from stdin or from a list of named files

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	dupedFiles := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, dupedFiles)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			// No exceptions in go. Must check error code
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			// NB: countLines call precedes declaration. Funcs and other package-level entities may be declared in any order
			countLines(f, counts, dupedFiles)
			// Do not leak
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s; from\n", n, line, dupedFiles[line])
		}
	}
}

// A map is a reference; when a map is passed to a function, the function recieves a copy
// of the ref, thus any change made in the func is visible to the caller.
func countLines(f *os.File, counts map[string]int, dupedFiles map[string]string) {
	input := bufio.NewScanner(f)
	// Streams input
	for input.Scan() {
		line := input.Text()
		counts[line]++
		dupedFiles[line] = dupedFiles[line] + "," + f.Name()
	}
	//TODO. Handle error from input.Err()
}

// basename removes directory components and a .suffix
// e.g, a => a, a.go => a, a/b/c.co => c, a/b.c.go => b.c
package main

import (
	"fmt"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		r := basename(arg)
		fmt.Println(r)
	}
}
func basename(s string) string {
	// Discard last '/' and everything before it.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	// Preserve everything before the last '.'.
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

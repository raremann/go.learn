// Nonempty is an example of an in-place slice algorithm.
package main

import "fmt"

// noempty returns a slice holding only the non-empty strings.
// The underlying arry is modified during the call.
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonemptyV2(strings []string) []string {
	out := strings[:0] // zero-length slice
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func main() {
	data := []string{"hi", "", "there", "friendo"}
	fmt.Printf("%q\n", nonempty(data)) // hi there friendo
	fmt.Printf("%q\n", data)           // hi there friendo friendo
	data = []string{"hi", "", "there", "friendo"}
	fmt.Printf("%q\n", nonemptyV2(data)) // hi there friendo
	fmt.Printf("%q\n", data)             // hi there friendo friendo

}

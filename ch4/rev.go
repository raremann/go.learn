// reverse reverses a slice of ints in place
package main

import "fmt"

// in place reversal
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// slices are not comparable, unlike arrays. Slices are indirect, values can change as underlying array changes.
func compSlice(x, y []string) bool {
	if len(x) != len(y) {
		return false
	}
	for i := range x {
		if x[i] != y[i] {
			return false
		}
	}
}

func main() {
	a := [...]int{1, 2, 3, 4}
	reverse(a[:])
	fmt.Println(a)
}

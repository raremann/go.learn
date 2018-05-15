// intsToString is similar to fmt.Sprint() but insert commas.
package main

import (
	"bytes"
	"fmt"
)
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteRune('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d",v)
	}
	buf.WriteRune(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int {1,2,3,4,5})) // "[1, 2, 3, 4, 5]"
}
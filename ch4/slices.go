// slices is a hodge-podge of code snippets about slices
package main

import (
	"fmt"
)

func main() {
	months()
}
func months() {
	//array of strings
	months := [...]string{1: "Jan", 2: "Feb", 3: "Mar", 4: "Apr", 5: "May",
		6: "June", 7: "Jul", 8: "Aug", 9: "Sep", 10: "Oct", 11: "Nov", 12: "Dec"}
	// Slices are a pointer to an array element together with a length and capacity
	q2 := months[4:7] // length 3, capacity 9
	// They can overlap
	summer := months[6:9] // length 3, capcity 7
	// Extending beyond length is legal
	endlessSummer := summer[:5]
	// But not beyond capacity
	// err := summer[:8]. Runtime Panic: slice out of range
	fmt.Println(q2)
	fmt.Println(summer)
	fmt.Println(endlessSummer)
}

// intro is a hodge-podge of code snippets
package main

import (
	"fmt"
)

func printArrayInfo() {
	// Array size is part of the type. Values initialized to 0 for the type.
	var a [3]int
	fmt.Printf("a len=%d, a[0]=%d a[len(a)-1]=%d\n", len(a), a[0], a[2])
	// elispes in the length cause the length to be determined by the number of initializers
	b := [...]int{1, 2, 3, 4}
	for i, v := range b {
		fmt.Printf("b[%d]=%d\n", i, v)
	}
	// Can specify a set of non-default values
	r := [...]int{98: -1, -2, -3}
	count := 0
	for i, v := range r {
		if v == 0 {
			count++
		} else {
			fmt.Println("Non zero element @", i, v)
		}
	}
	fmt.Println("r has this many zero elements:", count)
}

func printIndexValuePairs() {
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "∑", GBP: "£", RMB: "¶"}
	fmt.Println(RMB, symbol[RMB]) // "3, ¶"
}

func main() {
	printArrayInfo()
	printIndexValuePairs()
}

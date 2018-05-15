// appendInt uses built-in append functionality of slice
package main

import "fmt"

// ... are the signature of a variadic function
func appendInt(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		// There is room to grow, extend the slice.
		z = x[:zlen]
	} else {
		// There is not enough space. Allocate a new array.
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	//z[len(x)] = y
	copy(z[len(x):], y)
	return z
}

func main() {
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d capacity=%d\t%v\n", i, cap(y), y)
		x = y
	}

	var k []int
	k = appendInt(k, 1)
	k = appendInt(k, 2, 3)
	k = appendInt(k, 4, 5, 6)
	// ... supply a list of arguments from a slice
	k = appendInt(k, k...)
	fmt.Println(k)
}

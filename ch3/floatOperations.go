// floatOperations follows book and adds comments
package main

import (
	"fmt"
	"math"
)

func main() {
	// Dangers of float32. Use float64 always
	var f float32 = 16777216
	if f == f+1 {
		fmt.Println("Oh no float32 rounding error already @ ", f)
	}

	fmt.Println("printing floats")
	for x := 0; x < 8; x++ {
		fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	}

	fmt.Println("Overflows and NaNs")
	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) //	"0 -0 +Inf -Inf NaN"
}

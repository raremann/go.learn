package main

import (
	"fmt"
)

func main() {
	var x = 1<<1 | 1<<5 // 00100010
	var y = 1<<1 | 1<<2 // 00000110
	fmt.Printf("%08b\n", y)
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", x&y)  // bitwise AND 		00000010
	fmt.Printf("%08b\n", x|y)  // bitwise OR  		00100110
	fmt.Printf("%08b\n", x^y)  // bitwise XOR 		00100100
	fmt.Printf("%08b\n", x&^y) // bitwise AND NOT 	00100000

	fmt.Println("Non-zero bits of x")
	for i := uint(0); i < 8; i++ {
		if x&(1<<i) != 0 {
			fmt.Println(i)
		}
	}
}

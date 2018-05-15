// sha256 compares the SHA256 digest of two inputs
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
	r, c := bitDiff(c1, c2)
	fmt.Printf("%x,%d", r, c)
}

func bitDiff(d1, d2 [32]uint8) ([32]uint8, int8) {
	var ret [32]uint8
	var count int8
	for i, v := range d1 {
		ret[i] = v
		count++
		if v != d2[i] {
			ret[i] = 0
			count--
		}
	}
	return ret, count
}

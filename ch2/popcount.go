package popcount

// package level variables typically start life with an initialization expression but some structure, like tables do not fit well into an expression
// pc[i] is the populatin count of i
var pc [256]byte

// special init function automatically executed when program starts. cannot be called nor referenced. any number of them is allowed
// populate a table of 
func init() {
	for i:= range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// popcount returns the number of set bits of x
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
	pc[byte(x>>(1*8))] +
	pc[byte(x>>(2*8))] +
	pc[byte(x>>(3*8))] +
	pc[byte(x>>(4*8))] +
	pc[byte(x>>(5*8))] +
	pc[byte(x>>(6*8))] +
	pc[byte(x>>(7*8))]
	)
}

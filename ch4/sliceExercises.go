// Exercises 4.3-4.7

package main

import (
	"fmt"
)

func reverseV2(in *[]int) {
	for i, j := 0, len(*in)-1; i < j; i, j = i+1, j-1 {
		(*in)[i], (*in)[j] = (*in)[j], (*in)[i]
	}
}

func rotateV2(in []int, position int) {
	p1 := in[:position]
	p2 := in[position+1:]
	temp := make([]int, len(p1))
	copy(temp, p1)
	copy(in[:(len(p2))], p2)
	copy(in[len(p2)+1:], temp)
}

func dedup(in []string, dedupList ...string) {
	j := 0
	l := len(in)
	dupeMap := make(map[string]int8)
	for _, v := range dedupList {
		dupeMap[v] = 1
	}
	for i := 1; i < l; i++ {
		if in[i] != in[j] {
			// Test whether we ought to attempt deduping this string
			if len(dedupList) != 0 && dupeMap[in[i]] == 0 {
				continue
			}
			j++
			in[j] = in[i]
		}
	}
}

func squeeze(in []byte) {

}
func main() {
	v := []int{1, 2, 3, 7, 11}
	reverseV2(&v)
	fmt.Println(v)
	rotateV2(v, 2)
	fmt.Println(v)

	s := []string{"the", "the", "the", "cow", "j-", "j-", "j-", "jumped", "over", "the", "moooooooooooon"}
	dedup(s)
	fmt.Println(s)
}

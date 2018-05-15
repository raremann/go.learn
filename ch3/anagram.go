// isAnagram returns true if the arguments are anagrams
package main

import (
	"fmt"
	"os"
)

func main() {
	l := len(os.Args[1:])
	for i := 1; i < l; i += 2 {
		s1 := os.Args[i]
		s2 := os.Args[i+1]
		fmt.Printf("anagram(%s,%s)=%b\n", s1, s2, isAnagram(s1, s2))
	}
}
func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	runeMap := make(map[rune]int)
	for _, v := range s1 {
		runeMap[v]++
	}
	// for each rune in s2, check that there is a counterpart in the runeMap
	for _, w := range s2 {
		if runeMap[w] == 0 {
			return false
		}
		runeMap[w]--
	}
	return true
}

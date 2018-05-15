// Exercises 4.8 and 4.9

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

type category uint64

const (
	Control category = 1 << iota
	Digit
	Graphic
	Letter
	Lower
	Mark
	Number
	Print
	Punct
	Space
	Symbol
	Title
	Upper
)

func getUnicodeCategories(r rune) uint64 {
	var ret uint64
	if unicode.IsLetter(r) {
		ret = ret | uint64(Letter)
	}
	if unicode.IsDigit(r) {
		ret = ret | uint64(Digit)
	}
	if unicode.IsSymbol(r) {
		ret = ret | uint64(Symbol)
	}
	if unicode.IsControl(r) {
		ret = ret | uint64(Control)
	}
	if unicode.IsSpace(r) {
		ret = ret | uint64(Space)
	}
	return ret
}

func charcount() {
	counts := make(map[uint64]int)
	var utflen [utf8.UTFMax + 1]int
	invalid := 0

	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount:%v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		cat := getUnicodeCategories(r)
		counts[cat]++
		utflen[n]++
	}
	fmt.Printf("category\tcount\n")
	for c, n := range counts {
		fmt.Printf("%d\t%d\n", c, n)
	}

}

func main() {
	charcount()
}

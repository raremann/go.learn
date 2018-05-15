// comma inserts a common at each 10^3 place
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(comma(arg))
		fmt.Println(comma2(arg))
		fmt.Println(commaFloat(arg))
	}
}
func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	n := 0
	r := []rune(s)
	l := len(r)
	for i := 0; i < l; i++ {
		buf.WriteRune(r[i])
		n = (l - 1) - i
		if n%3 == 0 && n != 0 {
			buf.WriteByte(',')
		}

	}
	return buf.String()
}

func commaFloat(s string) string {
	d := strings.LastIndex(s, ".")
	var ret string
	if d != -1 {
		whole := comma2(s[:d-1])
		ret = whole + s[d:]
	} else {
		ret = comma2(s)
	}
	return ret
}

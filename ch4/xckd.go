package main

import (
	"fmt"

	"go.learn/ch4/xkcd"
)

func main() {
	//var c xkcd.Comic
	c, _ := xkcd.GetComic(1)
	fmt.Println(c)
}

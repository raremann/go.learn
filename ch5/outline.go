//Outline prints structure of HTML tree in outline using recursion
package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

type ElementType int

const (
	Head ElementType = iota
	Meta
	Link
	Script
	Body
	Div
	A
	Form
	Span
	Input
	Button
	Br
	Textarea
)

var elementNameMap map[string]ElementType

func initElementNameMap() {

	elementNameMap = map[string]ElementType{
		"head":     Head,
		"meta":     Meta,
		"link":     Link,
		"script":   Script,
		"body":     Body,
		"div":      Div,
		"a":        A,
		"form":     Form,
		"span":     Span,
		"input":    Input,
		"button":   Button,
		"br":       Br,
		"textarea": Textarea,
	}
}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("outline: %v\n", err)
	}
	initElementNameMap()
	elementCountMap := make(map[ElementType]int)
	outline(nil, doc, elementCountMap)
	fmt.Println(elementCountMap)
}

func outline(stack []string, n *html.Node, m map[ElementType]int) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag onto stack
		fmt.Printf("n.Data %s, nameMap[n.Data] %d\n", n.Data, elementNameMap[n.Data])
		m[elementNameMap[n.Data]]++
		fmt.Println(stack)

	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c, m)
	}
}

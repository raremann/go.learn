// Findlinks prtins the links in an HTML documetn read from standard input.

package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		log.Fatalf("findlinks: %v\n", err)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
	fmt.Println("Calling visit2")
	visit2(doc)
}

// visit appends to links each link found in n and returns the results.
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

// visit appends to links each link found in n and returns the results.
func visit2(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
			}
		}
	}
	printText(n)
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit2(c)
	}
}

//
func printText(n *html.Node) {
	if n.Type == html.TextNode {
		fmt.Println(n.Attr)
		fmt.Println(n.Data)
	}
}

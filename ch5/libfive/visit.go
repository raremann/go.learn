// Findlinks prtins the links in an HTML documetn read from standard input.

package libfive

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

/*
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

	fmt.Println("Printing text")
	printText(doc)
}
*/
// visit appends to links each link found in n and returns the results.
func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = Visit(links, c)
	}
	return links
}

// visit appends to links each link found in n and returns the results.
func Visit2(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				fmt.Println(a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Visit2(c)
	}
}

// printText print the body of TextNode, skipping all scripts
func PrintText(n *html.Node) {
	if n.Type == html.ElementNode && n.Data == "script" {
		return
	}
	if n.Type == html.TextNode {
		//	fmt.Println(n.Attr)
		fmt.Println("############################################")
		fmt.Println(strings.Trim(n.Data, "\n\n"))
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		PrintText(c)
	}
}

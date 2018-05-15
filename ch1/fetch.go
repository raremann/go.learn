// fetch prints the contents found at each specified url. curl-y
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func prefixUrl(url *string) {
	if !strings.HasPrefix(*url, "http") {
		*url = "http://" + *url
	}
}

func main() {
	for _, url := range os.Args[1:] {
		prefixUrl(&url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		bod, err := ioutil.ReadAll(resp.Body)
		// Do not leak
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v=n", url, err)
		}
		fmt.Printf("%s", bod)
	}
}

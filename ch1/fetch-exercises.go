// fetch prints the contents found at each specified url. curl-y
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func prefixUrl(url *string) {
	if !strings.HasPrefix(*url, "http") {
		fmt.Println("old url:", *url)
		*url = "http://" + *url
		fmt.Println("new url:", *url)
	}
}

func main() {
	for _, url := range os.Args[1:] {
		prefixUrl(&url)
		fmt.Println("url in main ", url)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Status: ", resp.Status)
		bytesWritten, err := io.Copy(os.Stdout, resp.Body)
		//bod, err := ioutil.ReadAll(resp.Body)
		// Do not leak
		fmt.Println("bytes ", bytesWritten)
		resp.Body.Close()
		if err != nil {

			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v=n", url, err)
		}
	}
}

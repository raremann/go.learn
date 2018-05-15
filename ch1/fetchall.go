// fetch prints the contents found at each specified url. curl-y
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func prefixUrl(url *string) {
	if !strings.HasPrefix(*url, "http") {
		fmt.Println("old url:", *url)
		*url = "http://" + *url
		fmt.Println("new url:", *url)
	}
}

func main() {
	start := time.Now()
	// Another instance of make; previous we did make(map...)
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // start a goroutine
	}
	for range os.Args[1:] {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch. What happens with error handling??
		return
	}
	// Send body to dev/null; just use copy to count response size
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // do not leak!
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}

// server1 is a minimal echo server
package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mutex sync.Mutex
var counter int

func main() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/stats", stats)
	log.Fatal(http.ListenAndServe("localhost:8001", nil))
}

//handler echoes the Path componet of the requested URL.
func stats(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	fmt.Fprintf(w, "Total number of requests = %d\n", counter++)
	mutex.Unlock()
}

func handler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	counter++
	mutex.Unlock()
	fmt.Fprintf(w, "url.path = %q\n", r.URL.Path)
}

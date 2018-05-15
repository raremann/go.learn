// surfaceServer is a HTTP server that returns SVG surface plots to the client
package main

import (
	"log"
	"net/http"

	"go.learn/ch3/surface"
)

func main() {
	http.HandleFunc("/", handleSurface)
	log.Fatal(http.ListenAndServe("localhost:8002", nil))
}

func handleSurface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")
	surface.Surface(w)
}

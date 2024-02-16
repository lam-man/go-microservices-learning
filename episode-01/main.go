package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	// Register a handler function for a specific path
	// Here, requests that not match any of the paths will be handled by the default handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Hello World")
		d, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}

		// Write back (send message) to clients
		fmt.Fprintf(w, "Hello %s\n", d)
	})

	http.HandleFunc("/goodbye", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Goodbey World")
	})

	// Register a handler for path /test
	// When a request is send to /test, you need to make sure that
	// the questcontains "contact"
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		d := r.Header["Contact"]
		if d == nil || len(d) == 0 {
			http.Error(w, "Ooops", http.StatusBadRequest)
			return
		}

		fmt.Fprintf(w, "Hello %s\n", d[0])
	})

	http.ListenAndServe(":9090", nil)
}

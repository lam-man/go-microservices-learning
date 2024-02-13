package main

import (
	"log"
	"net/http"
	"os"

	"episode-02/handlers"
)

func main() {
	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	hh := handlers.NewHello(l)

	// ServeMux is a HTTP request multiplexer, which matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.
	sm := http.NewServeMux()
	sm.Handle("/", hh)

	http.ListenAndServe(":9090", sm)
}

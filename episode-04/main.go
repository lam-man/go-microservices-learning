package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"episode-04/product-api/handlers"

	"github.com/nicholasjackson/env"
)

func main() {

	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)

	ph := handlers.NewProducts(l)

	// ServeMux is a HTTP request multiplexer, which matches the URL of each incoming request against a list of registered patterns and calls the handler for the pattern that most closely matches the URL.

	sm := http.NewServeMux()
	sm.Handle("/", ph)

	// Why we need a new server? Why cannot we use the default server?
	// We need to think about the timeout. If there are too many blocking requests (server takes more time to serve these requests), the default server will be blocked and the server will not be able to handle other requests. So, we need to set the timeout for the server.
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		WriteTimeout: 25 * time.Second,
		ReadTimeout:  25 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			l.Fatal(err)
		}
	}()

	// When the app received a signal, will it stop accepting new requests?
	// How does app stop accepting new requests?
	// What will happen to the client send requests to this server?

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Received terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	// why the shutodown will not be executed until we send a kill/interrupt signal?
	s.Shutdown(tc)

	// With basic settings in ep1, we didn't specify the servemux, so it uses the default servemux with http.ListenAndServe(":9090", nil)
	// http.ListenAndServe(":9090", sm)
}

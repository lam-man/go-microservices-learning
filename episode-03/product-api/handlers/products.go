package handlers

import (
	"episode-03/product-api/data"
	"log"
	"net/http"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// Need to use encoding/json package to convert the productList to JSON
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// Handle the GET request
	if r.Method == http.MethodGet {
		p.GetProducts(rw, r)
		return
	}

	// Handle the Update request

	// Catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	// Get the list of products
	lp := data.GetProducts()

	// How to return the above list to users?
	// Convert the list to JSON string
	// d, err := json.Marshal(lp)

	// Instead of using marshal from json package, we can use the encoder from json package
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}

	// rw.Write(d)
}

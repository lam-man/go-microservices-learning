package handlers

import (
	"log"
	"net/http"
	"time"
)

type Goodbye struct {
	l *log.Logger
}

func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	// sleep 30 seconds before execute the shutdown
	time.Sleep(20 * time.Second)
	rw.Write([]byte("Bye!"))
}

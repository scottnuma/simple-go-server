package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	s := &http.Server{
		Addr: os.Getenv("ADDRESS"),
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "would i")
		}),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

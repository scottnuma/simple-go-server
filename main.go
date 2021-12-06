package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	address := os.Getenv("ADDRESS")
	if address == "" {
		log.Println("ADDRESS not set")
		address = "0.0.0.0:8080"
	}
	log.Printf("serving on %s", address)

	s := &http.Server{
		Addr: address,
		Handler: http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			fmt.Fprintf(w, "would i")
		}),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Prints the current value of GOMAXPROCS (without modifying it)
		n := runtime.GOMAXPROCS(-1)
		log.Printf("GOMAXPROCS==%d\n", n)
		fmt.Fprintf(w, "GOMAXPROCS==%d\n", n)
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	err := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	log.Fatal(err)
}

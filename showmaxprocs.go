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
		// Prints the current value of NumCPU and GOMAXPROCS (without modifying them)
		cpus := runtime.NumCPU()
		maxprocs := runtime.GOMAXPROCS(-1)
		log.Printf("NumCPU==%d\nGOMAXPROCS==%d\n", cpus, maxprocs)
		fmt.Fprintf(w, "NumCPU==%d\nGOMAXPROCS==%d\n", cpus, maxprocs)
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

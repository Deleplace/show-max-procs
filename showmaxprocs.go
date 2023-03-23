package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"runtime"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Prints the current value of NumCPU and GOMAXPROCS (without modifying them)
		cpus := runtime.NumCPU()
		maxprocs := runtime.GOMAXPROCS(-1)
		log.Printf("NumCPU==%d\nGOMAXPROCS==%d\n", cpus, maxprocs)
		fmt.Fprintf(w, "NumCPU==%d\nGOMAXPROCS==%d\n", cpus, maxprocs)

		procCPUInfo, err := readFromProcCPUInfo()
		if err == nil {
			log.Printf("/proc/cpuinfo==%d\n", procCPUInfo)
			fmt.Fprintf(w, "/proc/cpuinfo==%d\n", procCPUInfo)
		} else {
			log.Printf("/proc/cpuinfo==unknown (%v)\n", err)
			fmt.Fprintf(w, "/proc/cpuinfo==unknown (%v)\n", err)
		}
	})

	http.HandleFunc("/proc/cpuinfo", func(w http.ResponseWriter, r *http.Request) {
		info, err := os.ReadFile("/proc/cpuinfo")
		if err != nil {
			fmt.Fprintf(w, "error reading /proc/cpuinfo: %v\n", err)
			return
		}
		w.Write(info)
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

var procPattern = regexp.MustCompile(`^processor\s*:`)

// readFromProcCPUInfo is intended to work on Linux
func readFromProcCPUInfo() (int, error) {
	info, err := os.ReadFile("/proc/cpuinfo")
	if err != nil {
		return 0, err
	}
	lines := bytes.Split(info, []byte("\n"))
	hits := 0
	for _, line := range lines {
		if procPattern.Match(line) {
			hits++
		}
	}
	return hits, nil
}

package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

var (
	originServer string
	port         string
)

func main() {
	// Parse command-line arguments.
	// This is done to dynamically set the port and origin server URL, instead of hardcoding them.
	flag.StringVar(&port, "port", "8080", "Port on which the proxy server will run")
	flag.StringVar(&originServer, "origin", "", "URL of the origin server")
	flag.Parse()

	// Check if the --origin flag is provided
	if originServer == "" {
		fmt.Println("Error: --origin flag is required")
		os.Exit(1)
	}

	// Start the proxy server
	fmt.Printf("Caching proxy server running on port %s, forwarding to %s\n", port, originServer)
	http.HandleFunc("/", handleRequest)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// handleRequest will forward the incoming req to the origin server and return the response
func handleRequest(w http.ResponseWriter, r *http.Request) {
	// Construct the target URL
	targetURL := originServer + r.URL.Path

	// Forward request to the origin server
	resp, err := http.Get(targetURL)
	if err != nil {
		http.Error(w, "Error contacting origin server", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	// Copy response headers
	for key, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(key, value)
		}
	}

	// Copy response status code
	w.WriteHeader(resp.StatusCode)

	// Copy response body
	io.Copy(w, resp.Body)
}

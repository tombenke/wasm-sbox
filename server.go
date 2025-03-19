package main

import (
	"flag"
	"log"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func runServer() {

	defer wg.Done()

	// Read config parameters
	port := flag.String("port", "8080", "The server port")
	dir := flag.String("dir", ".", "The base directory of content")
	flag.Parse()

	// Setup static file server
	fs := http.FileServer(http.Dir(*dir))

	// Start the HTTP server
	go func() {
		log.Printf("Server listens on %s port...", *port)
		if err := http.ListenAndServe(":"+*port, fs); err != nil {
			log.Fatalf("ListenAndServe: %v", err)
		}
	}()

	// Wait for a signal to stop the server (e.g., Ctrl+C)
	// You might want to add a channel and signal handling here
	// to gracefully shut down the server.
	select {} // This blocks indefinitely, allowing the server to run.
}

func main() {
	wg.Add(1)
	go runServer()
	wg.Wait()
}

package main

import (
	"flag"
	"fmt"
	"net/http"
)

func main() {
	// Define flags for the directory and port
	dir := flag.String("dir", ".", "the directory to serve files from")
	port := flag.String("port", "8080", "the port to serve on")
	ip := flag.String("ip", "192.168.0.148", "the IP address to serve on")
	flag.Parse()

	// Serve files from the specified directory
	fs := http.FileServer(http.Dir(*dir))
	http.Handle("/", fs)

	// Start the server on the specified IP and port
	address := *ip + ":" + *port
	fmt.Printf("Serving %s on http://%s\n", *dir, address)
	if err := http.ListenAndServe(address, nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
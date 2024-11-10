package main

import (
	"log"
	"net/http"
)

// buildServerMux returns the server mux that will be used to handle the file server and other endpoints
func buildServeMux() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.Handle("/app/", http.StripPrefix("/app", http.FileServer(http.Dir("."))))
	serveMux.HandleFunc("/healthz", handleHealthz)

	return serveMux
}

// Method to be called when /healthz endpoint is hit
func handleHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(200)
	w.Write([]byte("OK"))
}

// startServer kicks off the process and logs any errors
func startServer(serveMux *http.ServeMux, port string) {
	server := http.Server{Handler: serveMux, Addr: port}

	log.Printf("Serving files on port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func main() {
	serveMux := buildServeMux()
	startServer(serveMux, ":8080")
}

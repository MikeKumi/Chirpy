package main

import (
	"log"
	"net/http"
	"sync/atomic"
)

// buildServerMux returns the server mux that will be used to handle the file server and other endpoints
func buildServeMux() *http.ServeMux {
	const filepathRoot = "."
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
	}

	serveMux := http.NewServeMux()
	serveMux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/", http.FileServer(http.Dir(filepathRoot)))))
	serveMux.HandleFunc("GET /api/healthz", handlerHealthz)
	serveMux.HandleFunc("GET /admin/metrics", apiCfg.handlerMetrics)
	serveMux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	return serveMux
}

// Method to be called when /healthz endpoint is hit
func handlerHealthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}

// startServer kicks off the process and logs any errors
func startServer(serveMux *http.ServeMux, port string) {
	server := http.Server{Handler: serveMux, Addr: port}

	log.Printf("Serving files on port: %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func main() {
	const port = ":8080"
	serveMux := buildServeMux()
	startServer(serveMux, port)
}

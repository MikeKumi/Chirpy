package main

import (
	"fmt"
	"net/http"
)

func buildServeMux() *http.ServeMux {
	serveMux := http.NewServeMux()
	serveMux.Handle("/", http.FileServer(http.Dir("web")))
	return serveMux
}

func startServer(serveMux *http.ServeMux, port string) {
	server := http.Server{Handler: serveMux, Addr: port}

	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	serveMux := buildServeMux()
	startServer(serveMux, ":8080")
}

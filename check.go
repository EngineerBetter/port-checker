package main

import (
	"fmt"
	"net/http"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Up")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(ServeIndex))
	server := http.Server{Handler: mux}
	fmt.Println("Started")
	server.ListenAndServe()
}

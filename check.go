package main

import (
	"fmt"
	"net/http"
)

type IndexHandler struct {
}

func (h *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Up")
}

func main() {
	mux := http.NewServeMux()
	handler := new(IndexHandler)
	mux.Handle("/", handler)
	server := http.Server{Handler: mux}
	fmt.Println("Started")
	server.ListenAndServe()
}

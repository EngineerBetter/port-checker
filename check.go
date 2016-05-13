package main

import (
	"fmt"
	"net/http"
	"os"
)

func ServeIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Up")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(ServeIndex))
	server := http.Server{Handler: mux}
	fmt.Println("Starting")
	err := server.ListenAndServe()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

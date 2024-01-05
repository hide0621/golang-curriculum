package main

import (
	"fmt"
	"net/http"
)

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Serve HTTP")
}

func main() {
	fmt.Println("Hello World")

	handler := &myHandler{}
	http.Handle("/", handler)
	// http.HandleFunc("/", handler.ServeHTTP)
	http.ListenAndServe(":8080", nil)
}

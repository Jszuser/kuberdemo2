package main

import (
	"io"
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!")
}

func main() {
	http.HandleFunc("/", hello)
	port := ":8020"
	fmt.Printf("Starting a demo service on port %s\n", port);
	http.ListenAndServe(port, nil)
}

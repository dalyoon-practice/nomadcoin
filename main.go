package main

import (
	"fmt"
	"net/http"
)

const port string = ":4000"

func home(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World!")
}

func main() {
	http.HandleFunc("/", home)
	fmt.Printf("Server listening on port http://localhost%s\n", port)
	http.ListenAndServe(port, nil)
}

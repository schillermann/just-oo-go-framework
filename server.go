package main

import (
	"fmt"
	"io"
	"net/http"

	"schillermann/just-oo-go-framework/hello"
)

func handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Accept", "text/html")
	io.WriteString(w, hello.Hello())
}

func main() {
	portNumber := "9000"
	http.HandleFunc("/", handle)
	fmt.Println("Server listening on port ", portNumber)
	http.ListenAndServe(":"+portNumber, nil)
}

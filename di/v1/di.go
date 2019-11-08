package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	http.ListenAndServe(":10010", http.HandlerFunc(MyGreeterHandler))
}

func MyGreeterHandler(w http.ResponseWriter, t *http.Request) {
	Greet(w, "world")
}

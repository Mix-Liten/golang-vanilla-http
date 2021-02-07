package main

import (
	"fmt"
	"io"
	"net/http"
)

func indexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is index page")
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world")
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/hello", helloHandler)
	port := ":8080"
	fmt.Println("init in port", port)
	http.ListenAndServe(port, r)
}

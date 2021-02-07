package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

func notFoundHandler(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	io.WriteString(w, "Not found, ")
	io.WriteString(w, "找不到路徑, ")
	io.WriteString(w, req.URL.Path)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "this is index page")
}

func articleHandler(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	articleName := vars["name"]
	io.WriteString(w, "Articles: ")
	io.WriteString(w, articleName)
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/articles/{name}", articleHandler)
	port := ":8080"
	fmt.Println("init in port", port)
	http.ListenAndServe(port, r)
}

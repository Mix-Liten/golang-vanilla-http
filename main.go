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
	page := req.URL.Query().Get("p")
	ua := req.Header.Get("User-Agent")
	io.WriteString(w, "Articles: ")
	io.WriteString(w, articleName)
	io.WriteString(w, ", where page: ")
	io.WriteString(w, page)
	io.WriteString(w, "user-agent: ")
	io.WriteString(w, ua)
}

func main() {
	r := mux.NewRouter()
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/articles/{name}", articleHandler)
	port := ":8080"
	fmt.Println("init in port", port)
	http.ListenAndServe(port, r)
}

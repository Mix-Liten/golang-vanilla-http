package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
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

func errorPageHandler(w http.ResponseWriter, req *http.Request) {
	panic("test recovery middleware...")
	//log.Fatal("test recovery middleware...")
	//...do something
}

func articleHandler(w http.ResponseWriter, req *http.Request) {
	// /articles/test?p=3
	vars := mux.Vars(req)
	articleName := vars["name"]
	page := req.URL.Query().Get("p")
	ua := req.Header.Get("User-Agent")
	io.WriteString(w, "Articles: ")
	io.WriteString(w, articleName)
	io.WriteString(w, "\nwhere page: ")
	io.WriteString(w, page)
	io.WriteString(w, "\nuser-agent: ")
	io.WriteString(w, ua)
}

func serverMiddlewareHandler(w http.ResponseWriter, req *http.Request, next http.HandlerFunc) {
	w.Header().Set("SERVER", "VANILLA-SERVER")
	next(w, req)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.NotFoundHandler = http.HandlerFunc(notFoundHandler)
	r.HandleFunc("/articles/{name}", articleHandler)
	r.HandleFunc("/error", errorPageHandler)

	n := negroni.New(negroni.NewRecovery(), negroni.NewLogger())
	n.Use(negroni.HandlerFunc(serverMiddlewareHandler))
	n.UseHandler(r)

	port := ":8080"
	fmt.Println("init in port", port)
	http.ListenAndServe(port, n)
}

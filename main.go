package main

import (
    "log"
    "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/" {
        http.NotFound(w, r)
        return
    }
    w.Write([]byte("Hi you"))
}

func show(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("show"))
}

func create(w http.ResponseWriter, r *http.Request) {
    if r.Method != "POST" {
        w.Header().Set("Allow", "POST")
        http.Error(w, "Method Not Allowed", 405)
        return
    }
    w.Write([]byte("create"))
}

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/show", show)
    mux.HandleFunc("/snippet/create", create)

    log.Println("Starting server on :4000")
    err := http.ListenAndServe(":4000", mux)
    log.Fatal(err)
}

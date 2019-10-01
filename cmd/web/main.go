package main

import (
    "flag"
    "log"
    "net/http"
)

type Config struct {
    Addr      string
    StaticDir string
}

func main() {
    cfg := new(Config)
    flag.StringVar(&cfg.Addr, "addr", ":4000", "HHTP network address")
    flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
    flag.Parse()

    mux := http.NewServeMux()

    mux.HandleFunc("/", home)
    mux.HandleFunc("/snippet/show", show)
    mux.HandleFunc("/snippet/create", create)

    fileServer := http.FileServer(http.Dir(cfg.StaticDir))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    log.Printf("Starting server on %s", cfg.Addr)
    err := http.ListenAndServe(cfg.Addr, mux)
    log.Fatal(err)
}

package main

import (
    "flag"
    "log"
    "net/http"
    "os"
)

type Config struct {
    Addr      string
    StaticDir string
}

type Application struct {
    errorLog *log.Logger
    infoLog  *log.Logger
}

func main() {
    cfg := new(Config)
    flag.StringVar(&cfg.Addr, "addr", ":4000", "HHTP network address")
    flag.StringVar(&cfg.StaticDir, "static-dir", "./ui/static", "Path to static assets")
    flag.Parse()

    infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
    errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

    app := &Application{
        errorLog: errorLog,
        infoLog:  infoLog,
    }

    mux := http.NewServeMux()

    mux.HandleFunc("/", app.home)
    mux.HandleFunc("/snippet/show", app.show)
    mux.HandleFunc("/snippet/create", app.create)

    fileServer := http.FileServer(http.Dir(cfg.StaticDir))
    mux.Handle("/static/", http.StripPrefix("/static", fileServer))

    srv := &http.Server{
        Addr:     cfg.Addr,
        ErrorLog: errorLog,
        Handler:  mux,
    }

    infoLog.Printf("Starting server on %s", cfg.Addr)
    err := srv.ListenAndServe()
    errorLog.Fatal(err)
}

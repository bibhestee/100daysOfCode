package main

import (
  "flag"
  "log"
  "net/http"
  "os"
)

type application struct {
  errorLog *log.Logger
  infoLog *log.Logger
}


func main() {
  mux := http.NewServeMux()

  addr := flag.String("addr", ":4000", "HTTP network address")

  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
  }

  fileServer := http.FileServer(http.Dir("./ui/static"))

  mux.Handle("/static/", http.StripPrefix("/static", fileServer))
  mux.HandleFunc("/", app.home)
  mux.HandleFunc("/snippet/view", app.snippetView)
  mux.HandleFunc("/snippet/create", app.snippetCreate)

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: mux,
  }

  infoLog.Printf("Starting server on %s", *addr)
  err := srv.ListenAndServe()
  errorLog.Fatal(err)
}

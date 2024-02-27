package main

import (
  "database/sql"
  "flag"
  "log"
  "net/http"
  "os"
  _ "github.com/go-sql-driver/mysql"
)

type application struct {
  errorLog *log.Logger
  infoLog *log.Logger
}


func main() {

  addr := flag.String("addr", ":4000", "HTTP network address")
  dsn := flag.String("dsn", "web:web@/snippetbox?parseTime=true", "MySQL data source name")

  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

  app := &application{
    errorLog: errorLog,
    infoLog: infoLog,
  }

  // Connect db
  db, err := openDB(*dsn)
  if err != nil {
    errorLog.Fatal(err)
  }

  // Close db connection
  defer db.Close()

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
  }

  infoLog.Printf("Starting server on %s", *addr)
  err = srv.ListenAndServe()
  errorLog.Fatal(err)
}

// openDB() function wraps sql.Open() and returns a sql.DB connection pool for a given dns
func openDB(dsn string) (*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
     return nil, err
  }
  // Test the connection
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}

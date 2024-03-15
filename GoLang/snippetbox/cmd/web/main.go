package main

import (
  "crypto/tls"
  "database/sql"
  "flag"
  "html/template"
  "log"
  "net/http"
  "os"
  "github.com/go-playground/form/v4"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
  _ "github.com/go-sql-driver/mysql"
  "time"
  "github.com/alexedwards/scs/mysqlstore"
  "github.com/alexedwards/scs/v2"
)

type application struct {
  debug bool
  errorLog *log.Logger
  infoLog *log.Logger
  snippets models.SnippetModelInterface
  users models.UserModelInterface
  templateCache map[string]*template.Template
  formDecoder *form.Decoder
  sessionManager *scs.SessionManager
}


func main() {

  addr := flag.String("addr", ":4000", "HTTP network address")
  dsn := flag.String("dsn", "web:web@/snippetbox?parseTime=true", "MySQL data source name")
  debug := flag.Bool("debug", false, "Enable Debug mode")

  flag.Parse()

  infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
  errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)


  // Connect db
  db, err := openDB(*dsn)
  if err != nil {
    errorLog.Fatal(err)
  }

  // Close db connection
  defer db.Close()

  // Template cache
  templateCache, err := newTemplateCache()
  if err != nil {
    errorLog.Fatal(err)
  }

  formDecoder := form.NewDecoder()
  sessionManager := scs.New()

  sessionManager.Store = mysqlstore.New(db)
  sessionManager.Lifetime = 12 * time.Hour
  sessionManager.Cookie.Secure = true

  app := &application{
    debug: *debug,
    errorLog: errorLog,
    infoLog: infoLog,
    snippets: &models.SnippetModel{DB: db},
    users: &models.UserModel{DB: db},
    templateCache: templateCache,
    formDecoder: formDecoder,
    sessionManager: sessionManager,
  }


  tlsConfig := &tls.Config{
    MinVersion: tls.VersionTLS12,
    MaxVersion: tls.VersionTLS12,
    CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
  }

  srv := &http.Server{
    Addr: *addr,
    ErrorLog: errorLog,
    Handler: app.routes(),
    TLSConfig: tlsConfig,
    IdleTimeout: time.Minute,
    ReadTimeout: 5 * time.Second,
    WriteTimeout: 10 * time.Second,
  }

  infoLog.Printf("Starting server on %s", *addr)
  err = srv.ListenAndServeTLS("./tls/cert.pem", "./tls/key.pem")
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

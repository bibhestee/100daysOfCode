package main

import (
  "log"
  "net/http"
)

// Define a home handler function which writes a bite slice containing
// "Hello world from Snippetbox" as the response body
func home(res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Hello world from Snippetbox"))
}

func main() {
  // Use http.NewServeMux() function to initialize a new servemux, then
  // register the home function as the handler for the "/" URL pattern
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)

  // Use the http.ListenAndServe() function to start a new web server.
  // We pass in two parameters: TCP network address to listen on ( in this case ":4000")
  // and the servemux we just created.
  // If http.ListenAndServe() returns an error, we use the log.Fatal() function to log the error message and exit.
  // Note that any error returned by http.ListenAndServe() is always non-nil
  log.Print("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}

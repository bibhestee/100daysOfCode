package main

import (
  "log"
  "net/http"
)

// Define a home handler function which writes a byte slice containing
// Display the home page using "Hello world from Snippetbox" as the response body
func home(res http.ResponseWriter, req *http.Request) {
  // Check if the current URL path exactly matches "/"
  if req.URL.Path != "/" {
    http.NotFound(res, req)
    return
  }
  res.Write([]byte("Hello world from Snippetbox"))
}

// Display a specific snippet
func snippetView(res http.ResponseWriter, req *http.Request) {
  res.Write([]byte("Display a specific snippet..."))
}

// Create a new snippet
func snippetCreate(res http.ResponseWriter, req *http.Request) {
  // Restrict this route to act on POST requests only
  if req.Method != "POST" {
    // Use the res.Header().Set() method to add a new header
    // Add a new header that let the user know which request methods are supported for this route
    res.Header().Set("Allow", "POST")
    // Send a 405 Method not allowed response
    res.WriteHeader(405) // This method can only be called once
    res.Write([]byte("Method No Allowed"))
    return
  }
  res.Write([]byte("Create a new snippet..."))
}


func main() {
  // Use http.NewServeMux() function to initialize a new servemux, then
  // register the home function as the handler for the "/" URL pattern and the two new handlers
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  mux.HandleFunc("/snippet/view", snippetView)
  mux.HandleFunc("/snippet/create", snippetCreate)

  // Use the http.ListenAndServe() function to start a new web server.
  // We pass in two parameters: TCP network address to listen on ( in this case ":4000")
  // and the servemux we just created.
  // If http.ListenAndServe() returns an error, we use the log.Fatal() function to log the error message and exit.
  // Note that any error returned by http.ListenAndServe() is always non-nil
  log.Print("Starting server on :4000")
  err := http.ListenAndServe(":4000", mux)
  log.Fatal(err)
}

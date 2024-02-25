package main

import (
  "fmt"
  "net/http"
  "strconv"
)

func home(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path != "/" {
    http.NotFound(res, req)
    return
  }

  res.Write([]byte("Hello from Snippetbox"))
}


func snippetView(res http.ResponseWriter, req *http.Request) {
  id, err := strconv.Atoi(req.URL.Query().Get("id"))
  if err != nil || id < 1 {
    http.NotFound(res, req)
    return
  }

  fmt.Fprintf(res, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(res http.ResponseWriter, req *http.Request) {
  if req.Method != http.MethodPost {
    res.Header().Set("Allow", http.MethodPost)
    http.Error(res, "Method Not Allowed", http.StatusMethodNotAllowed)
    return
  }

  res.Write([]byte("Create a new snippet..."))
}

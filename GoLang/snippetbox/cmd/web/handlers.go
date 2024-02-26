package main

import (
  "fmt"
  "html/template"
  "net/http"
  "strconv"
)

func (app *application) home(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path != "/" {
    app.notFound(res)
    return
  }

  files := []string{
    "./ui/html/base.html",
    "./ui/html/partials/nav.html",
    "./ui/html/pages/home.html",
  }
  // Use the template.ParseFiles() function to read the template file into a template set.
  ts, err := template.ParseFiles(files...)
  if err != nil {
    app.errorLog.Print(err.Error())
    app.serverError(res, err)
    return
  }

  err = ts.ExecuteTemplate(res, "base", nil)
  if err != nil {
    app.errorLog.Print(err.Error())
    app.serverError(res, err)
  }
}


func (app *application) snippetView(res http.ResponseWriter, req *http.Request) {
  id, err := strconv.Atoi(req.URL.Query().Get("id"))
  if err != nil || id < 1 {
    app.notFound(res)
    return
  }

  fmt.Fprintf(res, "Display a specific snippet with ID %d...", id)
}

func (app *application) snippetCreate(res http.ResponseWriter, req *http.Request) {
  if req.Method != http.MethodPost {
    res.Header().Set("Allow", http.MethodPost)
    app.clientError(res, http.StatusMethodNotAllowed)
    return
  }

  res.Write([]byte("Create a new snippet..."))
}

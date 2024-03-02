package main

import (
  "errors"
  "fmt"
  "net/http"
  "strconv"
  "strings"
  "unicode/utf8"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
  "github.com/julienschmidt/httprouter"
)

func (app *application) home(res http.ResponseWriter, req *http.Request) {

  // Query the database
  snippets, err := app.snippets.Latest()
  if err != nil {
    app.serverError(res, err)
    return
  }

  data := app.newTemplateData(req)
  data.Snippets = snippets

  app.render(res, http.StatusOK, "home.html", data)
}


func (app *application) snippetView(res http.ResponseWriter, req *http.Request) {
  params := httprouter.ParamsFromContext(req.Context())

  id, err := strconv.Atoi(params.ByName("id"))
  if err != nil || id < 1 {
    app.notFound(res)
    return
  }

  snippet, err := app.snippets.Get(id)
  if err != nil {
    if errors.Is(err, models.ErrNoRecord) {
      app.notFound(res)
    } else {
      app.serverError(res, err)
    }
    return
  }

  data := app.newTemplateData(req)
  data.Snippet = snippet

  app.render(res, http.StatusOK, "view.html", data)
}


func (app *application) snippetCreate(res http.ResponseWriter, req *http.Request) {
  data := app.newTemplateData(req)
  app.render(res, http.StatusOK, "create.html", data)
}


func (app *application) snippetCreatePost(res http.ResponseWriter, req *http.Request) {
  // Parse the request form
  err := req.ParseForm()
  if err != nil {
    app.clientError(res, http.StatusBadRequest)
    return
  }

  // Retrieve field data from form
  title := req.PostForm.Get("title")
  content := req.PostForm.Get("content")
  expires, err := strconv.Atoi(req.PostForm.Get("expires"))
  if err != nil {
    app.clientError(res, http.StatusBadRequest)
    return
  }

  // Validate data
  fieldErrors := make(map[string]string)

  if strings.TrimSpace(title) == "" {
    fieldErrors["title"] = "This field cannot be blank"
  } else if utf8.RuneCountInString(title) > 100 {
    fieldErrors["title"] = "This field cannot be more than 100 characters long."
  }

  if strings.TrimSpace(content) == "" {
    fieldErrors["content"] = "This field cannot be blank"
  }

  if expires != 1 && expires != 7 && expires != 365 {
    fieldErrors["expires"] = "This field must equal 1, 7 or 365"
  }

  // Dump error to plain HTTP response if any
  if len(fieldErrors) > 0 {
    fmt.Fprint(res, fieldErrors)
    return
  }

  // Add form data to database
  id, err := app.snippets.Insert(title, content, expires)
  if err != nil {
    app.serverError(res, err)
  }

  http.Redirect(res, req, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}


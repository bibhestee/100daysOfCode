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

type snippetCreateForm struct {
  Title string
  Content string
  Expires int
  FieldErrors map[string]string
}


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
  data.Form = snippetCreateForm{
    Expires: 365,
  }
  app.render(res, http.StatusOK, "create.html", data)
}


func (app *application) snippetCreatePost(res http.ResponseWriter, req *http.Request) {
  // Parse the request form
  err := req.ParseForm()
  if err != nil {
    app.clientError(res, http.StatusBadRequest)
    return
  }

  // Retrieve expires field from PostForm
  expires, err := strconv.Atoi(req.PostForm.Get("expires"))
  if err != nil {
    app.clientError(res, http.StatusBadRequest)
    return
  }

  form := snippetCreateForm{
    Title: req.PostForm.Get("title"),
    Content: req.PostForm.Get("content"),
    Expires: expires,
    FieldErrors: map[string]string{},
  }

  // Validate data

  if strings.TrimSpace(form.Title) == "" {
    form.FieldErrors["title"] = "This field cannot be blank"
  } else if utf8.RuneCountInString(form.Title) > 100 {
    form.FieldErrors["title"] = "This field cannot be more than 100 characters long."
  }

  if strings.TrimSpace(form.Content) == "" {
    form.FieldErrors["content"] = "This field cannot be blank"
  }

  if form.Expires != 1 && form.Expires != 7 && form.Expires != 365 {
    form.FieldErrors["expires"] = "This field must equal 1, 7 or 365"
  }


  if len(form.FieldErrors) > 0 {
    data := app.newTemplateData(req)
    data.Form = form
    app.render(res, http.StatusUnprocessableEntity, "create.html", data)
    return
  }

  // Add form data to database
  id, err := app.snippets.Insert(form.Title, form.Content, form.Expires)
  if err != nil {
    app.serverError(res, err)
  }

  http.Redirect(res, req, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}


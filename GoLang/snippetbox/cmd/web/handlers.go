package main

import (
  "errors"
  "fmt"
  "net/http"
  "strconv"
  "github.com/bibhestee/100daysOfCode/Golang/snippetbox/internal/validator"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
  "github.com/julienschmidt/httprouter"
)

type snippetCreateForm struct {
  Title string
  Content string
  Expires int
  validator.Validator
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
  }

  // Validate data
  form.CheckField(validator.NotBlank(form.Title), "title", "This field cannot be blank")

  form.CheckField(validator.MaxChars(form.Title, 100), "title", "This field cannot be more than 100 characters long.")
  form.CheckField(validator.NotBlank(form.Content), "content", "This field cannot be blank")
  form.CheckField(validator.PermittedInt(form.Expires, 1, 7, 365), "expires", "This field must equal 1, 7 or 365")


  if !form.Valid() {
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


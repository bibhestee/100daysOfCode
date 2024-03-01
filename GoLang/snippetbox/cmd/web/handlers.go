package main

import (
  "errors"
  "fmt"
  "net/http"
  "strconv"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
)

func (app *application) home(res http.ResponseWriter, req *http.Request) {
  if req.URL.Path != "/" {
    app.notFound(res)
    return
  }

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
  id, err := strconv.Atoi(req.URL.Query().Get("id"))
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
  if req.Method != http.MethodPost {
    res.Header().Set("Allow", http.MethodPost)
    app.clientError(res, http.StatusMethodNotAllowed)
    return
  }

  // Create some variables holding dummy data. We'll remove these later on
  // during the build.
  title := "O snail"
  content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
  expires := 7

  id, err := app.snippets.Insert(title, content, expires)
  if err != nil {
    app.serverError(res, err)
  }

  http.Redirect(res, req, fmt.Sprintf("/snippet/view?id=%d", id), http.StatusSeeOther)
}

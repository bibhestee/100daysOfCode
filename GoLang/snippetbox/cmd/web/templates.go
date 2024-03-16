package main

import (
  "html/template"
  "io/fs"
  "path/filepath"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/ui"
  "time"
)


type templateData struct {
  CurrentYear     int
  Snippet         *models.Snippet
  Snippets        []*models.Snippet
  Form            any
  Flash           string
  IsAuthenticated bool
  CSRFToken       string
  User            *models.User
}


func newTemplateCache() (map[string]*template.Template, error) {
  cache := map[string]*template.Template{}

  pages, err := fs.Glob(ui.Files, "html/pages/*.html")
  if err != nil {
    return nil, err
  }

  for _, page := range pages {
    // Extract filename
    name := filepath.Base(page)

    patterns := []string{
      "html/base.html",
      "html/partials/*.html",
      page,
    }

    ts, err := template.New(name).Funcs(functions).ParseFS(ui.Files, patterns...)
    if err != nil {
      return nil, err
    }

    cache[name] = ts
  }

  return cache, nil

}


func humanDate(t time.Time) string {
  if t.IsZero() {
    return ""
  }

  return t.UTC().Format("02 Jan 2006 at 3:04pm")
}


var functions = template.FuncMap{
  "humanDate": humanDate,
}

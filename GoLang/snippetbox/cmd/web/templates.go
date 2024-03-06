package main

import (
  "html/template"
  "path/filepath"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
  "time"
)


type templateData struct {
  CurrentYear int
  Snippet *models.Snippet
  Snippets []*models.Snippet
  Form any
  Flash string
  isAuthenticated bool
}


func newTemplateCache() (map[string]*template.Template, error) {
  cache := map[string]*template.Template{}

  pages, err := filepath.Glob("./ui/html/pages/*.html")
  if err != nil {
    return nil, err
  }

  for _, page := range pages {
    // Extract filename
    name := filepath.Base(page)

    ts, err := template.New(name).Funcs(functions).ParseFiles("./ui/html/base.html")
    if err != nil {
      return nil, err
    }

    ts, err = ts.ParseGlob("./ui/html/partials/*.html")
    if err != nil {
      return nil, err
    }

    ts, err = ts.ParseFiles(page)
    if err != nil {
      return nil, err
    }

    cache[name] = ts
  }

  return cache, nil

}


func humanDate(t time.Time) string {
  return t.Format("02 Jan 2006 at 3:04pm")
}


var functions = template.FuncMap{
  "humanDate": humanDate,
}

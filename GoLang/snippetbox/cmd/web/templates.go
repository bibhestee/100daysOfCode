package main

import (
  "html/template"
  "path/filepath"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
)


type templateData struct {
  Snippet *models.Snippet
  Snippets []*models.Snippet
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

    ts, err := template.ParseFiles("./ui/html/base.html")
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

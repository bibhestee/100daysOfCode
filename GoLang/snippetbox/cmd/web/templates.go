package main

import (
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/models"
)


type templateData struct {
  Snippet *models.Snippet
  Snippets []*models.Snippet
}

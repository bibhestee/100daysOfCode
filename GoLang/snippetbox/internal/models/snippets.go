package models

import (
  "database/sql"
  "time"
)


type Snippet struct {
  ID int
  Title string
  Content string
  Created time.Time
  Expires time.Time
}

// Define a SnippetModel type which wraps the sql.DB connection pool
type SnippetModel struct {
  DB *sql.DB
}

// Insert new record into database
func (m *SnippetModel) Insert(title string, content string, expires int) (int, error) {
  return 0, nil
}

// Get specific snippet based on ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
  return nil, nil
}

// This will return 10 most recently created snippet
func (m *SnippetModel) Latest() ([]*Snippet, error) {
  return nil, nil
}

package models

import (
  "database/sql"
  "errors"
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

  stmt := `INSERT INTO snippets (title, content, created, expires) VALUES(?, ?, UTC_TIMESTAMP(), DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY))`

  result, err := m.DB.Exec(stmt, title, content, expires)
  if err != nil {
    return 0, err
  }

  id, err := result.LastInsertId()
  if err != nil {
    return 0, nil
  }

  return int(id), nil
}

// Get specific snippet based on ID
func (m *SnippetModel) Get(id int) (*Snippet, error) {
  stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() AND id = ?`

  row := m.DB.QueryRow(stmt, id)

  s := &Snippet{}

  err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
  if err != nil {
    if errors.Is(err, sql.ErrNoRows) {
      return nil, ErrNoRecord
    } else {
      return nil, err
    }
  }
  return s, nil
}

// This will return 10 most recently created snippet
func (m *SnippetModel) Latest() ([]*Snippet, error) {
  return nil, nil
}

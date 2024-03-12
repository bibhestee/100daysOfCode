package models

import (
  "database/sql"
  "errors"
  "time"
)

type SnippetModelInterface interface {
  Insert(title string, content string, expires int) (int, error)
  Get(id int) (*Snippet, error)
  Latest() ([]*Snippet, error)
}

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

  stmt := `SELECT id, title, content, created, expires FROM snippets WHERE expires > UTC_TIMESTAMP() ORDER BY id DESC LIMIT 10`

  rows, err := m.DB.Query(stmt)

  if err != nil {
    return nil, err
  }

  defer rows.Close()

  // Initialize an empty slice to hold the snippet struct
  snippets := []*Snippet{}

  for rows.Next() {
    s := &Snippet{}
    // Use rows.Scan to copy the values for each field in the row to the new Snippet object
    err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
    if err != nil {
      return nil, err
    }

    snippets = append(snippets, s)
  }

  // When the rows.Next() loop has finished we call rows.Err() to retrieve any
  // error that was encountered during the iteration.
  if err = rows.Err(); err != nil {
    return nil, err
  }

  return snippets, nil
}

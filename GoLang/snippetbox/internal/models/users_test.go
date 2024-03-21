package models

import (
  "testing"
  "github.com/bibhestee/100daysOfCode/GoLang/snippetbox/internal/assert"
)

func TestUserModelExists(t *testing.T) {
  tests := []struct {
    name string
    userID int
    want bool
    }{
      {  name: "Valid ID",  userID: 1,  want: true,  },
      {  name: "Zero ID",  userID: 0,  want: false,  },
      {  name: "Non-existent ID",  userID: 2,  want: false,  },
    }

  for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {

      db := newTestDB(t)

      m := UserModel{db}

      exists, err := m.Exists(tt.userID)

      assert.Equal(t, exists, tt.want)
      assert.NilError(t, err)
    })
  }
}

func TestUserGet(t *testing.T) {

  db := newTestDB(t)
  m := UserModel{db}

  t.Run("Valid ID", func(t *testing.T) {

      user, err := m.Get(1)

      assert.Equal(t, user.Email, "alice@example.com")
      assert.IsError(t, err, nil)
    })

  t.Run("Invalid ID",func(t *testing.T) {
    user, err := m.Get(2)

    assert.Equal(t, user, nil)
    assert.IsError(t, err, ErrInvalidCredentials)
  } )
}

package models

import (
  "database/sql"
  "errors"
  "strings"
  "github.com/go-sql-driver/mysql"
  "golang.org/x/crypto/bcrypt"
  "time"
)

type User struct {
  ID              int
  Name            string
  Email           string
  HashedPassword  []byte
  Created         time.Time
}

type UserModel struct {
  DB *sql.DB
}


func (m *UserModel) Insert(name, email, password string) error {
  hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 12)
  if err != nil {
    return err
  }

  stmt := `INSERT INTO users (name, email, hashed_password, created) VALUE(?, ?, ?, UTC_TIMESTAMP())`

  _, err = m.DB.Exec(stmt, name, email, string(hashedPassword))
  if err != nil {
    var mySQLError *mysql.MySQLError

    if errors.As(err, &mySQLError) {
      if mySQLError.Number == 1062 && strings.Contains(mySQLError.Message, "user_uc_email") {
        return ErrDuplicateEmail
      }
    }
    return err
  }
  return nil
}

func (m *UserModel) Authenticate(email, password string) (int, error) {
  return 0, nil
}

func (m *UserModel) Exists(id int) (bool, error) {
  return false, nil
}

package main
import (
  "fmt"
  "time"
)

/*
 * User Struct
 * @username
 * @email
 * @createdAt
 * @updatedAt
 */

type User struct {
  username string
  email string
  createdAt time.Time
  updatedAt time.Time
  }


func (u *User) Details() string {
  return fmt.Sprintf("{username: %s, email: %s, createdAt: %v, updatedAt: %v}", u.username, u.email, u.createdAt, u.updatedAt)
}

func (u *User) Create() {
  u.createdAt = time.Now()
  u.updatedAt = time.Now()
}

func main() {
  newUser := User{
    username: "JohnDoe",
    email: "johndoe@gmail.com",
  }

  newUser.Create()
  fmt.Println(newUser.Details())
}

package main

import (
  "fmt"
  "modules/lib"
)

func main() {
  user := user.UserDetails{
    Name: "John Doe",
    Email: "Johndoe@mail.com",
    Age: 32,
    Password: "Password",
  }

  fmt.Println(user)
}

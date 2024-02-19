package main
import (
  "fmt"
)

/*
 * fullname - get fullname and print the short form
 * @firstname: user firstname
 * @lastname: user lastname
 * Return: fullname short form
 */

func fullname(lastname string, firstname string) string {
  return string(firstname[0]) + "." + lastname
}


func main() {
  // Get input from user
  var (
      first string
      last string
      )

  fmt.Print("Enter your first name: ")
  fmt.Scanln(&first)
  fmt.Print("Enter your last name: ")
  fmt.Scanln(&last)
  // Print out the fullname short form
  fmt.Println("Welcome", fullname(last, first))
}

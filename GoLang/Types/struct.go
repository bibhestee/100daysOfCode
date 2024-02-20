package main
import "fmt"


/* Structs Type
 * Declaration
 * Structs literal
 * Anonymous Structs
 */

func main() {
  // Delaration
  type Person struct {
    name string
    age int
    occupation string
    skills []string
  }

  // Assign value
  person := Person{
    "John Doe",
    28,
    "Software Engineer",
    []string{"Java", ".Net"},
  }

  info := fmt.Sprintf("%s is a %d years old %s with proficiency in %s", person.name, person.age, person.occupation, person.skills)
  fmt.Println(info)

  // Anonymous Struct
  engineer := struct {
    Field string
    YOE int
    }{
    "Software Engineer",
    5,
    }

  fmt.Println(engineer)
}

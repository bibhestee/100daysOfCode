package main
import "fmt"

/*
 * Embedding is not Inheritance
 * Embedding can be used to access methods declared by other methods
 */

// Person
type Person struct {
  Name string
}

// Android
type Android struct {
  Person
  Model string
}

// Person Method - Talk
func (p Person) Talk() {
  fmt.Println("Hi, It's nice to meet you")
}

func main() {
  a := new(Android)
  a.Talk()
}

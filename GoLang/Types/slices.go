package main
import (
  "fmt"
  )

// Slices

func main() {
  // Declaration
  var list []int
  // Declare and initialize
  var list2 = []int{1, 2, 3, 4}
  // Using :=
  list3 := []string{"This", "is", "a", "list"}
  // Using make
  list4 := make([]string, 0, 2)
  // Using empty slice literal
  var list5 = []int{}
  list = append(list, 8, 9, 7)
  fmt.Println("list 1:", list)
  fmt.Println("list 2:", list2)
  fmt.Println("list 3:", list3)
  fmt.Println("list 4:", list4)
  fmt.Println("list 5:", list5)
}

package print // import "my-module/lib"

import "fmt"

// ListInt print list of ints separated by commas
func ListInt (a []int) {
  fmt.Printf("[")
  for _, element := range a {
    if l := len(a); element == a[l - 1] {
      fmt.Printf(`'%v'`, element)
      continue
    }
    fmt.Printf(`'%v', `, element)
  }
  fmt.Printf("]\n")
}

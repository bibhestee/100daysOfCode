package main
import (
  "fmt"
  "math"
  )

/* min - find the smallest number in a list
 * @list: list of integers
 *
 * return: return the smallest number
 */

func min(list []int) (out int) {
  out = math.MaxInt

  for _, v := range list {
    if v < out {
    out = v
    }
  }
  return out
}


func main() {
  list := []int{3, 5, 1, 8, 2, 9}
  fmt.Println("List:", list)
  fmt.Println("min:", min(list))
}

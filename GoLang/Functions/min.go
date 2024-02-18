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
  x := []int{48,96,86,68,57,82,63,70, 37,34,83,27,19,97,9,17}
  list := []int{3, 5, 1, 8, 2, 9}
  fmt.Println("List:", list)
  fmt.Println("min:", min(list))
  fmt.Println("List 2:", x)
  fmt.Println("min:", min(x))
}

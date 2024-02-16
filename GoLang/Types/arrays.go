package main
import "fmt"

// Arrays

func main() {
  // Declaration without specifying length
  var arr1 = [...]int{3, 4, 5}
  fmt.Println(arr1, "length:", len(arr1))

  // Declaration with length
  var arr2 [4]string
  arr2[1] = "item1"
  fmt.Println(arr2[1], "\n")

  // Looping through arrays
  arr2[0], arr2[2], arr2[3] = "item0", "item2", "item3"
  for _, v := range arr2 {
    fmt.Println(v)
  }
}

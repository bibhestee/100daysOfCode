package main
import "fmt"

// Maps - Dictionary in python and Object in Js

func main() {
  // A nil map
  var nilMap map[string]int
  fmt.Println("nilMap:", nilMap)
  fmt.Println("nilMap == nil:", nilMap == nil)
  // Using make
  var makeMap = make(map[string]string)
  fmt.Println("Create map with make:", makeMap)
  // Using :=
  m := map[string]int{
    "hello": 5,
    "world": 2,
    }
  // Use comma ok idiom
  v, ok := m["hello"]
  fmt.Println(`Does key "hello" exist:`, ok, v)
  makeMap["Item1"] = "1"
  fmt.Println("Print map created by make and updated:", makeMap)
  fmt.Println("map items before delete:", m)
  delete(m, "world")
  fmt.Println("map items after delete:", m)
}

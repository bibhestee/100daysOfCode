package main
import "fmt"

/* A program that prints out all the numbers
 * evenly divisible by 3 between 1 and 100. (3,
 * 6, 9, etc.)
*/


func main() {
  for num := 3; num < 100; num += 3 {
    fmt.Println(num)
  }
}

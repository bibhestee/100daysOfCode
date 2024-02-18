package main
import "fmt"

/*
 * swap - swap variable values
 * @x: pointer to the first variable
 * @y: pointer to the second variable
 * Return: nothing
 */

func swap(x *int, y *int) {
  temp := *x
  *x = *y
  *y = temp
}


func main() {
  a, b := 1, 2
  fmt.Println("Before swap")
  fmt.Println("a:", a, "b:", b)
  swap(&a, &b)
  fmt.Println("After swap")
  fmt.Println("a:", a, "b:", b)
}

package main
import (
  "fmt"
  )

/*
 * half - a function which takes an integer and    * halves it and returns true if it was even or
 * false if it was odd.
 * @num: number to be halved
 * return - (value, true/false)
 */

func half(num int) (n int, bol bool) {
  n, rem := num / 2, num % 2
  bol = false

  if rem == 0 {
    bol = true
  }
  return n, bol
}


func main() {
  fmt.Println(half(1))
  fmt.Println(half(2))
}

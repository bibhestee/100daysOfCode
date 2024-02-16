package main
import "fmt"

/*
* Write a program that prints the numbers from 1 * to 100. But for multiples of three print "Fizz" * in-stead of the number and for the multiples of * five print "Buzz". For numbers which are
* multiples of both three and five print
* "FizzBuzz"
*/

func main() {
  for num := 1; num < 100; num++ {
    if num % 3 == 0 && num % 5 == 0 {
      fmt.Println("FizzBuzz")
    } else if num % 3 == 0 {
      fmt.Println("Fizz")
    } else if num % 5 == 0 {
      fmt.Println("Buzz")
    } else {
      fmt.Println(num)
    }
  }
}


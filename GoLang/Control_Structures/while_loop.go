package main
import (
  "fmt"
)

// A guessing game using while loop and if statement

func main() {
  var (
    num int = 5
    guess int
    i int = 5
    )

  for {
    // Confirm if no of guess is not exceeded
    if i == 0 {
      fmt.Println("Oops! You've exceeded the number of guess. Try again later")
      break
    }
    fmt.Print("Guess the hidden number: ")
    fmt.Scanln(&guess)
    // Decrease no of guess after each guess
    i--
    if num == guess {
      fmt.Println("Wow! You got it right!")
      break
    } else {
      fmt.Println("Incorrect! try again. Number of guess left:", i)
    }
    }
}

package main
import (
  "fmt"
  "os"
  )

// Primitive Datatypes in Go

func main() {
  // Int
  var num int = 12
  // String
  var str string = "Hello world!"
  // Bool
  var bol bool = true
  // Float
  var dec float32 = 3.2
  fmt.Println("Integer:", num)
  fmt.Println("String:", str)
  fmt.Println("Boolean:", bol)
  fmt.Println("Float:", dec)
  fmt.Println("Using Os package")
  argLen := len(os.Args)
  args := os.Args
  fmt.Println("Arguments:", args)
  fmt.Println("Arguments length:", argLen)
}

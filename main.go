package main

import (
  "fmt"
  "os"
)

func install(packagename string)  {
  fmt.Println(packagename)
}

func main()  {
  args := os.Args[1:]

  switch args[0] {
  case "install":
    install(args[1])
  default:
    fmt.Println("Not founded!")
  }
}

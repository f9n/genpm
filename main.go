package main

import (
  "fmt"
  "os"
  "strings"
)

func install(packagename string)  {
  fmt.Println(packagename)
  commands := []string{
    "sudo",
    "pacman",
    "-S",
    packagename,
  }
  script := strings.Join(commands, " ")
  fmt.Println(script)
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

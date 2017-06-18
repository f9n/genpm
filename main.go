package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

// Checking tool with which command
func checkTool(tool string) bool {
  cmdOut, err := exec.Command("which", tool).Output()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return false
	} else {
		fmt.Println(string(cmdOut))
		return true
	}
}

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
  if checkTool("pacman") {
    cmd := exec.Command("/bin/bash", "-c", script)
    cmd.Stdout = os.Stdout
    cmd.Stdin = os.Stdin
		cmd.Stderr = os.Stderr
		_ = cmd.Run()
  }
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

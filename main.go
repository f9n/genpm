package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

type Tool struct {
  install string
  remove  string
  search  string
}

var tools = map[string]Tool{
  "pacman": {
    "pacman -S",        // install
    "pacman -R",        // remove
    "pacman -Ss"        // search
  },
  "apt-get": {
    "apt-get install",  // install
    "apt-get remove",   // remove
    "apt-cache search"  // search
  },
}

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

// Searching string with tool
func search(searchstring string)  {
  fmt.Println(searchstring)
  commands := []string{
    "sudo",
    "pacman",
    "-Ss",
    searchstring,
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

// Removing package with tool
func remove(packagename string)  {
  fmt.Println(packagename)
  commands := []string{
    "sudo",
    "pacman",
    "-R",
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

// Installing package with tool
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
  fmt.Println(tools)
  fmt.Println(tools["pacman"])
  fmt.Println(tools["pacman"].install)
  args := os.Args[1:]

  switch args[0] {
  case "install":
    install(args[1])
  case "remove":
    remove(args[1])
  case "search":
    search(args[1])
  default:
    fmt.Println("Not founded!")
  }
}

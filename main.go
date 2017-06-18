package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

const tools = []string{
  "pacman",   // Arch Based
  "apt-get",  // Debian Based
  "dnf",      // Fedora Based
  "yum",      // Rpm Based
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

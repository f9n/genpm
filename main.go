package main

import (
  "fmt"
  "os"
  "os/exec"
  "strings"
)

// contain tool's install, remove, search commands
type Tool struct {
  installCommand string
  removeCommand  string
  searchCommand  string
}

// run script in bash
func (this *Tool) runScript(script string) {
  fmt.Println("[+] Runned Tool runScript method")
  cmd := exec.Command("/bin/bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Stdin = os.Stdin
  cmd.Stderr = os.Stderr
  _ = cmd.Run()
}

// install package with Tool
func (this *Tool) install(packagename string) {
  fmt.Println("[+] Runned Tool install method")
  fmt.Println(packagename)
  script := this.installCommand + " " + packagename
  fmt.Println(script)
  this.runScript(script)
}

// remove package with Tool
func (this *Tool) remove(packagename string) {
  fmt.Println("[+] Runned Tool remove method")
  fmt.Println(packagename)
  script := this.removeCommand + " " + packagename
  fmt.Println(script)
  this.runScript(script)
}

var tools = map[string]Tool{
  "pacman": {
    "sudo pacman -S",        // install
    "sudo pacman -R",        // remove
    "sudo pacman -Ss",       // search
  },
  "apt-get": {
    "sudo apt-get install",  // install
    "sudo apt-get remove",   // remove
    "sudo apt-cache search", // search
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

func main()  {
  fmt.Println(tools)
  fmt.Println(tools["pacman"])
  fmt.Println(tools["pacman"].installCommand)
  args := os.Args[1:]

  switch args[0] {
  case "install":
    tool := tools["pacman"]
    tool.install(args[1])
    // install(args[1])
  case "remove":
    tool := tools["pacman"]
    tool.remove(args[1])
    // remove(args[1])
  case "search":
    search(args[1])
  default:
    fmt.Println("Not founded!")
  }
}

package tool

import (
  "fmt"
  "os"
  "os/exec"
)

// contain tool's install, remove, search commands
type Tool struct {
  InstallCommand string
  RemoveCommand  string
  SearchCommand  string
  UpgradeCommand string
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
func (this *Tool) Install(packagename string) {
  fmt.Println("[+] Runned Tool install method")
  fmt.Println(packagename)
  script := this.InstallCommand + " " + packagename
  fmt.Println(script)
  this.runScript(script)
}

// remove package with Tool
func (this *Tool) Remove(packagename string) {
  fmt.Println("[+] Runned Tool remove method")
  fmt.Println(packagename)
  script := this.RemoveCommand + " " + packagename
  fmt.Println(script)
  this.runScript(script)
}

func (this *Tool) Search(searchstring string) {
  fmt.Println("[+] Runned Tool search method")
  fmt.Println(searchstring)
  script := this.SearchCommand + " " + searchstring
  fmt.Println(script)
  this.runScript(script)
}

func (this *Tool) Upgrade() {
  fmt.Println("[+] Runned Tool upgrade method")
  this.runScript(this.UpgradeCommand)
}

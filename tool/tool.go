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
  fmt.Println("[+] Runned runScript method in Tool Struct")
  cmd := exec.Command("/bin/bash", "-c", script)
  cmd.Stdout = os.Stdout
  cmd.Stdin = os.Stdin
  cmd.Stderr = os.Stderr
  _ = cmd.Run()
}

// install package with Tool
func (this *Tool) Install(packagename string) {
  fmt.Println("[+] Runned Install method in Tool Struct")
  script := this.InstallCommand + " " + packagename
  this.runScript(script)
}

// remove package with Tool
func (this *Tool) Remove(packagename string) {
  fmt.Println("[+] Runned Remove method in Tool Struct")
  script := this.RemoveCommand + " " + packagename
  this.runScript(script)
}

// search package with Tool
func (this *Tool) Search(searchstring string) {
  fmt.Println("[+] Runned Search method in Tool Struct")
  script := this.SearchCommand + " " + searchstring
  this.runScript(script)
}

// upgrade packages with Tool
func (this *Tool) Upgrade() {
  fmt.Println("[+] Runned Tool upgrade method")
  this.runScript(this.UpgradeCommand)
}

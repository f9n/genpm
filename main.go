package main

import (
  "fmt"
  "os"
  "os/exec"
)

// contain tool's install, remove, search commands
type Tool struct {
  installCommand string
  removeCommand  string
  searchCommand  string
  upgradeCommand string
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

func (this *Tool) search(searchstring string) {
  fmt.Println("[+] Runned Tool search method")
  fmt.Println(searchstring)
  script := this.searchCommand + " " + searchstring
  fmt.Println(script)
  this.runScript(script)
}

// https://stackoverflow.com/questions/35343707/linux-apt-get-command-not-found-how-to-install-a-package-in-arch-linux , Thnaks for https://stackoverflow.com/users/523100/czechnology
var tools = map[string]Tool{
  "pacman": {               // Arch
    "sudo pacman -S",         // install
    "sudo pacman -Rs",        // remove
    "sudo pacman -Ss",        // search
    "sudo pacman -Syu",      // upgrade
  },
  "apt": {                  // Debian/Ubuntu
    "sudo apt install",
    "sudo apt remove",
    "sudo apt search",
    "sudo apt update; sudo apt upgrade",
  },
  "dnf": {                  //  RedHat/Fedora
    "sudo dnf install",
    "sudo dnf remove",
    "sudo dnf search",
    "sudo dnf upgrade",
  },
  "zypper": {               // SLES/openSUSE
    "zypper install",
    "zypper remove",
    "zypper search",
    "zypper update",
  },
  "emerge": {               // Gentoo
    "emerge [-a]",
    "emerge -C",
    "emerge -S",
    "emerge -u world",
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
  fmt.Println(tools["pacman"])
  fmt.Println(tools["pacman"].installCommand)
  fmt.Println(tools["pacman"].removeCommand)
  fmt.Println(tools["pacman"].searchCommand)
  fmt.Println(tools["pacman"].upgradeCommand)
  args := os.Args[1:]

  switch args[0] {
  case "install":
    tool := tools["pacman"]
    tool.install(args[1])
  case "remove":
    tool := tools["pacman"]
    tool.remove(args[1])
  case "search":
    tool := tools["pacman"]
    tool.search(args[1])
  default:
    fmt.Println("Not founded!")
  }
}

package main

import (
  "fmt"
  // "os"
  "os/exec"
  "github.com/pleycpl/genpm/tool"
)

var toolsList = []string{
  "pacman",
  "apt",
  "dnf",
  "zypper",
  "emerge",
}

// https://stackoverflow.com/questions/35343707/linux-apt-get-command-not-found-how-to-install-a-package-in-arch-linux , Thnaks for https://stackoverflow.com/users/523100/czechnology
var tools = map[string]tool.Tool{
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

func checkAvailableTools() []string {
  var availableTools []string
  for _, tool := range toolsList {
    if checkTool(tool) {
      availableTools = append(availableTools, tool)
    }
  }
  return availableTools
}

// Checking tool with which command
func checkTool(tool string) bool {
  cmdOut, err := exec.Command("which", tool).Output()
	if err != nil {
		// fmt.Fprintln(os.Stderr, err)
		return false
	} else {
		fmt.Println(string(cmdOut))
		return true
	}
}

func main()  {
  fmt.Println(tools["pacman"])
  fmt.Println(tools["pacman"].InstallCommand)
  fmt.Println(tools["pacman"].RemoveCommand)
  fmt.Println(tools["pacman"].SearchCommand)
  fmt.Println(tools["pacman"].UpgradeCommand)

  available := checkAvailableTools()
  fmt.Println(available)
  /*
  args := os.Args[1:]

  switch args[0] {
  case "install":
    tool := tools["pacman"]
    tool.Install(args[1])
  case "remove":
    tool := tools["pacman"]
    tool.Remove(args[1])
  case "search":
    tool := tools["pacman"]
    tool.Search(args[1])
  case "upgrade":
    tool := tools["pacman"]
    tool.Upgrade()
  default:
    fmt.Println("Not founded!")
  }
  */
}

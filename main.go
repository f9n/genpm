package main

import (
  "fmt"
  "os"
  "github.com/pleycpl/genpm/tool"
  "github.com/pleycpl/genpm/genpm"
)

// https://stackoverflow.com/questions/35343707/linux-apt-get-command-not-found-how-to-install-a-package-in-arch-linux , Thnaks for https://stackoverflow.com/users/523100/czechnology
var tools = []tool.Tool{
  {
    "pacman",
    "sudo pacman -S",         // install
    "sudo pacman -Rs",        // remove
    "sudo pacman -Ss",        // search
    "sudo pacman -Syu",      // upgrade
  },
  /*
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
  },*/
}

func main()  {
  fmt.Println("[+] Runned Main function")
  GenpmInstance := genpm.NewGenpm("/home/chuck/.genpmrc")
  GenpmInstance.Check()

  args := os.Args[1:]
  if len(args) > 0 {
    tool := GenpmInstance.Tool
    switch args[0] {
    case "install":
      tool.Install(args[1])
    case "remove":
      tool.Remove(args[1])
    case "search":
      tool.Search(args[1])
    case "upgrade":
      tool.Upgrade()
    default:
      fmt.Println("Not founded!")
    }
  }
}

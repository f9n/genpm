package genpm

import (
  "fmt"
  "os"
)

var toolsList = []string{
  "pacman",
  "apt",
  "dnf",
  "zypper",
  "emerge",
}

// checking available tools
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

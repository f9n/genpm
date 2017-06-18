package genpm

import (
  "fmt"
  "os"
  "os/exec"
  "io/ioutil"
)

var toolsList = []string{
  "pacman",
  "apt",
  "dnf",
  "zypper",
  "emerge",
}

// Creating .genpmrc file with Genpm struct
type Genpm struct {
  Path string
  Tool string
}

// reload .genpmrc file
func (this *Genpm) reload() {
  file, err := os.Create(this.Path)
  if err != nil {
    fmt.Println("The file doesn't created, ", err)
    return
  }
  defer file.Close()

  available := checkAvailableTools()
  fmt.Println(available)
  file.WriteString(available[0])
  this.Tool = available[0]
}

// checking .genpmrc file
func (this *Genpm) Check() {
  // read the whole file at once
  b, err := ioutil.ReadFile(this.Path)
  this.Tool = string(b)
  if err != nil {
    fmt.Println("The File is not exists!", err)
    this.reload()
  }
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

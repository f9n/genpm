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

// Creating .genpmrc file with Genpm struct
type Genpm struct {
  path string
  tool string
}

// reload .genpmrc file
func (this *Genpm) reload() {
  file, err := os.Create(this.path)
  if err != nil {
    // handle the error here
    return
  }
  defer file.Close()

  available := checkAvailableTools()
  fmt.Println(available)
  file.WriteString(available[0])
}

// checking .genpmrc file
func (this *Genpm) Check() {
  // read the whole file at once
  b, err := ioutil.ReadFile(this.path)
  this.tool = string(b)
  if err != nil {
    fmt.Println("The File is not exists!")
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

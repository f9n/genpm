package genpm

import (
  "fmt"
  "os"
  "os/exec"
  "io/ioutil"
  "github.com/pleycpl/genpm/tool"
)

var allpackagemanagers = []string{
  "pacman",
  "apt",
  "dnf",
  "zypper",
  "emerge",
}

// Creating .genpmrc file with Genpm struct
type Genpm struct {
  Path string
  Tool tool.Tool
}

// reload .genpmrc file
func (this *Genpm) reload() string {
  fmt.Println("[+] Runned reload method in Genpm Struct")
  file, err := os.Create(this.Path)
  if err != nil {
    fmt.Println("The file doesn't created, ", err)
    return ""
  }
  defer file.Close()

  availables := getExistsPmTools()
  file.WriteString(availables[0])
  toolname := availables[0]
  return toolname
}

// checking .genpmrc file
func (this *Genpm) Check() {
  fmt.Println("[+] Runned Check method in Genpm Struct")
  b, err := ioutil.ReadFile(this.Path)
  toolname := string(b)
  if err != nil {
    fmt.Println("The File is not exists!", err)
    toolname = this.reload()
  }
  this.Tool = tool.Tool{
    ToolName: toolname,
    InstallCommand: "echo installcommand",
    RemoveCommand: "echo RemoveCommand",
    SearchCommand: "echo stringcommand",
    UpgradeCommand: "echo UpgradeCommand",
  }
}

// checking available tools
func getExistsPmTools() []string {
  fmt.Println("[+] Runned getExistsPmTools function")
  var tools []string
  for _, packagemanager := range allpackagemanagers {
    if isExistsPm(packagemanager) {
      tools = append(tools, packagemanager)
    }
  }
  return tools
}

// Checking package manager with which command
func isExistsPm(tool string) bool {
  fmt.Println("[+] Runned isExistsPm function")
  _, err := exec.Command("which", tool).Output()
	if err != nil {
		return false
	} else {
		return true
	}
}

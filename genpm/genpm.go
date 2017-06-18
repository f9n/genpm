package genpm

import (
  "fmt"
  "strings"
  "os"
  "os/exec"
  "io/ioutil"
  "encoding/json"
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
  this.Tool = readJsonFile(toolname)
}

func NewGenpm(path string) Genpm {
  return Genpm{
    Path: path,
    Tool: tool.Tool{
      ToolName: "",
      InstallCommand: "",
      RemoveCommand: "",
      SearchCommand: "",
      UpgradeCommand: "",
    },
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

func readJsonFile(filename string) tool.Tool {
  fmt.Println("[+] Runned readJsonFile function")
  jsonFile, err := os.Open(strings.Join([]string{"static/", filename, ".json"}, ""))
  if err != nil {
	   fmt.Println(err)
  }
  defer jsonFile.Close()
  fmt.Println("Successfully Opened xxxx.json: ", filename)

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var toolA tool.Tool
	json.Unmarshal(byteValue, &toolA)
  return toolA
}

package util

import (
  "encoding/json"
  "fmt"
  "io/ioutil"
  "os"
  "os/exec"
  "strings"
  "github.com/pleycpl/genpm/tool"
)


var allpackagemanagers = []string{
  "pacman",
  "apt",
  "dnf",
  "zypper",
  "emerge",
}

// checking available tools
func GetExistsPmTools() []string {
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

func ReadJsonFile(filename string) tool.Tool {
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

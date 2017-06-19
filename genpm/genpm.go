package genpm

import (
  "fmt"
  "os"
  "io/ioutil"
  "github.com/pleycpl/genpm/tool"
  "github.com/pleycpl/genpm/util"
)

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
    fmt.Println("The File doesn't created, ", err)
    return ""
  }
  defer file.Close()
  fmt.Println("The File created")
  availables := util.GetExistsPmTools()
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
  this.Tool = util.ReadJsonFile(toolname)
}

func NewGenpm(path string) Genpm {
  fmt.Println("[+] Runned NewGenpm function")
  GenpmInstance := Genpm{
    Path: path,
    Tool: tool.Tool{
      ToolName: "",
      InstallCommand: "",
      RemoveCommand: "",
      SearchCommand: "",
      UpgradeCommand: "",
    },
  }
  GenpmInstance.Check()
  return GenpmInstance
}

package main

import(
  "fmt"
  "os"
  "flag"
  "dredd"
  log "github.com/cihub/seelog"
)

const(
  VERSION = "0.0.1"
  LOADMON_DB = "loadmon"
  LOADMON_COLLECTION = "loadmon_test"
)

var testConfig = `
<seelog>
<outputs>
<file path="./log/main.log"/>
</outputs>
</seelog>`

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")


func main(){
  fmt.Printf("Version: %v",VERSION)
  mode := "unknown"
 
  Logger,_ := log.LoggerFromConfigAsBytes([]byte(testConfig))
  log.ReplaceLogger(Logger)

  if len(os.Args) > 1 {
    mode = os.Args[1]
  }

  os.Args = append(os.Args[:1], os.Args[2:]...)

  if *cpuprofile != "" {
    fmt.Println("running profiler")
  }
  if *memprofile!= "" {
    fmt.Println("running profiler")
  }

  if mode == "start" {
    fmt.Println("running start")
    dredd.NewMongoConnection()
    defer dredd.MongoConnection.Close()

    dredd.Run()
  }
}

package main

import(
  "fmt"
  "os"
  "flag"
  "listener"
)

const(
  VERSION = "0.0.1"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
var memprofile = flag.String("memprofile", "", "write memory profile to this file")

func main(){
  fmt.Println("Version:",VERSION)
  mode := "unknown"

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
    listener.Run()
  }
}

package main

import(
  "fmt"
  "os"
  "flag"
  "listener"
  "labix.org/v2/mgo"
)

const(
  VERSION = "0.0.1"
  LOADMON_DB = "loadmon"
  LOADMON_COLLECTION = "loadmon_test"
  MONGO_URL = "localhost:27017"
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
    session := setupDB()
    defer session.Close()
    collection := session.DB(LOADMON_DB).C(LOADMON_COLLECTION)

    listener.Run(collection)
  }
}

func setupDB() *mgo.Session{
  fmt.Println("settign up collection")
  session,err := mgo.Dial(MONGO_URL)
  if err != nil {
    panic(err)
  }

  session.SetMode(mgo.Monotonic, true)
  return session
}

package listener

import(
  "fmt"
)

func Run() {
  fmt.Println("running loadmon")
  fmt.Println(Settings)
  listener := TCPListen(Settings.Address,Settings.Port)

  //currentTime := time.Now().UnixNano()
  for {
    m := listener.Receive()
    fmt.Println(m)
  }
}

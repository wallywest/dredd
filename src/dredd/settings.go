package dredd

import(
)

const(
  //defaultPort = 10500
  defaultPort = 10503
  defaultAddress = "0.0.0.0"
)

type ListenerSettings struct {
  Port int
  Address string
  Verbose bool
}
var Settings ListenerSettings = ListenerSettings{}

func init() {
  Settings.Port = defaultPort
  Settings.Address = defaultAddress
  Settings.Verbose = false
}

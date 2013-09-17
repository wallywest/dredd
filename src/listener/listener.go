package listener

import(
  "fmt"
  "labix.org/v2/mgo"
)

func Run(s *mgo.Collection) {
  fmt.Println("running loadmon")
  fmt.Println(Settings)
  listener := TCPListen(Settings.Address,Settings.Port)

  app_id_chan := make(chan loadmonMessage)
  dnis_chan := make(chan loadmonMessage)
  total_chan := make(chan loadmonMessage)

  chan_listeners := make([]chan loadmonMessage,3)
  chan_listeners = append(chan_listeners,app_id_chan,dnis_chan,total_chan)

  aggregator := &Aggregator{
    collection: s,
    message_chan: make(chan loadmonMessage),
    listeners: chan_listeners,
  }

  go aggregator.Listen()

  for {
    m := listener.Receive()
    aggregator.Process(m)
  }
}

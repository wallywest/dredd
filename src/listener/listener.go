package listener

import(
  "fmt"
  log "github.com/cihub/seelog"
)

func Run() {
  fmt.Println("running loadmon")
  fmt.Println(Settings)
  listener := TCPListen(Settings.Address,Settings.Port)

  app_id_chan := make(chan loadmonMessage)
  dnis_chan := make(chan loadmonMessage)
  total_chan := make(chan loadmonMessage)

  chan_listeners := make([]chan loadmonMessage,3)
  chan_listeners = append(chan_listeners,app_id_chan,dnis_chan,total_chan)

  aggregator := &Aggregator{
    message_chan: make(chan loadmonMessage),
    listeners: chan_listeners,
    writer: &MongoConnection,
  }

  go aggregator.Listen()

  for {
    m := listener.Receive()
    log.Info(m)
    //aggregator.Process(m)
  }
}

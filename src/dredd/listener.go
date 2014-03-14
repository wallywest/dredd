package dredd

import(
  "fmt"
  //log "github.com/cihub/seelog"
)

var(
  Sinks = make(map[string]func() interface{})
)

func RegisterSink(name string, factory func() interface{}) {
  Sinks["name"] = factory
}

func Run() {
  fmt.Println("running loadmon")
  fmt.Println(Settings)
  listener := NewTCPListener(Settings.Address,Settings.Port)

  app_id_chan := make(chan loadmonMessage)
  dnis_chan := make(chan loadmonMessage)
  total_chan := make(chan loadmonMessage)

  chan_listeners := make([]chan loadmonMessage,3)
  chan_listeners = append(chan_listeners,app_id_chan,dnis_chan,total_chan)
  loadmon_message_chan := make(chan *loadmonMessage)
  stream := &Parser{loadmon_message_chan: loadmon_message_chan}
  aggregator := NewAggregator()


  for {
    select {
    case message := <-listener.raw_messages:
      go func(m string) {
       stream.Process(m)
      }(message)
    case loadmon_message := <-loadmon_message_chan:
      aggregator.Process(loadmon_message)
    }
  }
}

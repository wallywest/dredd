package dredd

import(
  "strconv"
  "time"
  "fmt"
  log "github.com/cihub/seelog"
)

type Aggregator struct {
  containers []Container
}

func NewAggregator() (a *Aggregator){
  var c []Container
  a = &Aggregator{containers: c}
  return
}

func (a *Aggregator) Process(l *loadmonMessage) {
  fmt.Println("Processing message")
  log.Info(l)

  switch l.messageType {
  case "loadmon_call":
    a.updateTotal(l)
  case "loadmon_dnis":
    a.updateDnis(l)
  case "loadmon_outcome":
    //a.updateAppId(l)
  }
}

func (a *Aggregator) updateDnis(l *loadmonMessage) {
  fmt.Println("updateDnis")
  fmt.Println(l)
}

func (a *Aggregator) updateTotal(l *loadmonMessage) {
  ts,err := strconv.ParseInt(l.fieldMap["timestamp"],10,64)
  if err != nil {
    panic(err)
  }
  t := time.Unix(ts,0)
  hour := strconv.Itoa(t.Hour())
  minute:= strconv.Itoa(t.Hour() * 60 + t.Minute())

  tot,_ := strconv.ParseInt(l.fieldMap["num_total"],10,64)
  inttot,_ := strconv.ParseInt(l.fieldMap["num_inbound"],10,64)
  outtot,_ := strconv.ParseInt(l.fieldMap["num_outbound"],10,64)

  MongoConnection.UpdateTotal(t,hour,minute,tot,inttot,outtot)
}

func (a *Aggregator) updateAppId(l *loadmonMessage) {
  if l.messageType != "loadmon_call" {
    fmt.Println(l)
  }
}

package listener

import(
  "strconv"
  "time"
  "fmt"
  "labix.org/v2/mgo"
  "labix.org/v2/mgo/bson"
)

type Aggregator struct {
  collection *mgo.Collection
  message_chan chan loadmonMessage
  listeners []chan loadmonMessage
}

type CallTotal struct {
  metadata map[string]string
  daily int
  hourly map[string]int
  minute map[string]map[string]int
}


func (a *Aggregator) Listen() {
  for {
    select {
    case message := <- a.message_chan:
      a.Log(message)
      /*go a.updateDnis(message)*/
      //go a.updateTotal(message)
      go a.updateAppId(message)
    }
  }
}

func (a *Aggregator) updateDnis(l loadmonMessage) {
  fmt.Println("updateDnis")
}

func (a *Aggregator) updateTotal(l loadmonMessage) {
  if(l.messageType == "loadmon_call") {
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

    query := bson.M{"day":t.YearDay()}
    change := bson.M{"$inc": bson.M{
      "daily":tot,
      "dailyinbound":inttot,
      "dailyoutbound":outtot,
      "hourly."+hour: tot,
      "minute."+hour+"."+minute: tot,
    }}

    mongo_err := a.collection.Update(query,change)
    if mongo_err != nil {
      panic(mongo_err)
      return
    }
  }
}

func (a *Aggregator) updateAppId(l loadmonMessage) {
  if l.messageType != "loadmon_call" {
  }
}

func (a *Aggregator) Log(l loadmonMessage) {
  fmt.Println(l);
}

func (a *Aggregator) Process(m string) {
  stream := &Stream{data: m}
  stream.Process(a.message_chan)
}

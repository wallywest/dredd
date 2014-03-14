package sink

import(
  "strconv"
  . ""
)

type TotalCallSink struct {
}

func (s *TotalCallSink) Run(){
}

func (s *TotalCallSink) updateTotal(l *loadmonMessage) {
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

func init() {
  RegisterSink("TotalCallSink", func() interface{} {return new(TotalCallSink)})
}

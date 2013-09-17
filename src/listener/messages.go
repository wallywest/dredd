package listener

import(
  "fmt"
)


type LoadmonMessage interface {
  parse([] string) map[string]string
}

type loadmonMessage struct {
  messageType string
  fieldMap map[string]string
}

type loadmonOutcome struct {}
type loadmonDnis struct {}
type loadmonCall struct {}

func(l *loadmonCall) parse(fields []string) map[string]string{
    return map[string]string{
      "cdr": fields[0],
      "timestamp": fields[1],
      "num_total": fields[2],
      "num_inbound": fields[3],
      "num_outbound": fields[4],
    }
}

func(l *loadmonOutcome) parse(fields []string) map[string]string{
    return map[string]string{
      "cdr": fields[0],
      "timestamp": fields[1],
      "dnis": fields[2],
      "outcome": fields[3],
      "app_id": fields[4],
      "job_id": fields[5],
      "count": fields[6],
    }
}

func(l *loadmonDnis) parse(fields []string) map[string]string{
  return map[string]string{
      "cdr": fields[0],
      "timestamp": fields[1],
      "dnis": fields[2],
      "app_id": fields[3],
      "job_id": fields[4],
      "direction": fields[5],
      "num_active": fields[6],
      "num_complete": fields[7],
      "total_time": fields[8],
    }
}

func (m *loadmonMessage) parse(fields []string, s interface{}){
  obj,err := s.(LoadmonMessage)
  if !err {
    fmt.Println("Could not find object")
    return
  }
  m.fieldMap = obj.parse(fields)
}

func (m *loadmonMessage) buildObject(fields []string) (s interface{}) {

  switch m.messageType {
  case "loadmon_call":
    return new(loadmonCall)
  case "loadmon_dnis":
    return new(loadmonDnis)
  case "loadmon_outcome":
    return new(loadmonOutcome)
  }
  return
}

func NewLoadmonMessage(m_type string,fields []string) (l *loadmonMessage){
  l = &loadmonMessage{}
  l.messageType = m_type
  obj := l.buildObject(fields)
  l.parse(fields,obj)
  return l
}

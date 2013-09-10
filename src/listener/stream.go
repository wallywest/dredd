package listener
import(
  "strings"
  "fmt"
  "reflect"
)

type Stream struct {
  data string
}

type LoadmonOutcome struct {
  cdr string
  timestamp string
  dnis string
  outcome string
  app_id string
  job_id string
  count string
}

type LoadmonDnis struct {
  cdr string
  timestamp string
  dnis string
  app_id string
  job_id string
  direction string
  num_active string
  num_complete string
  total_time string
}

type LoadmonCall struct {
  cdr string
  timestamp string
  num_total string
  num_inbound string
  num_outbound string
}



func (s *Stream) Process() {
  split := strings.Split(s.data,"}")
  for _,v := range split[0:len(split)-1] {
    //fmt.Printf("Got record: %v\n",v)
    record := strings.Split(v,"|")
    field_class := record[0]
    //fmt.Printf("Field Class: %v\n",field_class)

    switch field_class {
    case "loadmon_call":
      message := &LoadmonCall{
        cdr: record[1],
        timestamp: record[2],
        num_total: record[3],
        num_inbound: record[4],
        num_outbound: record[5],
      }
      fmt.Println("value:", reflect.ValueOf(message))
      fmt.Println(message)
    case "loadmon_outcome":
      message := LoadmonOutcome{
        cdr: record[1],
        timestamp: record[2],
        dnis: record[3],
        outcome: record[4],
        app_id: record[5],
        job_id: record[6],
        count: record[7],

      }

      fmt.Println("value:", reflect.ValueOf(message))
      fmt.Println(message)
    case "loadmon_dnis":
      message := &LoadmonDnis{
        cdr: record[1],
        timestamp: record[2],
        dnis: record[3],
        app_id: record[4],
        job_id: record[5],
        direction: record[6],
        num_active: record[7],
        num_complete: record[8],
        total_time: record[9],
      }

      fmt.Println("value:", reflect.ValueOf(message))
      fmt.Println(message)
    }
  }
}

package listener
import(
  "strings"
  "fmt"
)

type Stream struct {
  data string
  messages []loadmonMessage
}

func (s *Stream) Process(m chan loadmonMessage) {
  split := strings.Split(s.data,"}")
  for _,v := range split[0:len(split)-1] {
    record := strings.Split(v,"|")
    m_type := record[0]
    fields := record[1:len(record)]

    message := NewLoadmonMessage(m_type,fields)
    m <- *message
    //s.messages = append(s.messages,*message)
  }
  //s.message_chan <- s.messages
}

func (s *Stream) Log() {
  fmt.Printf("Messages length is: %v\n",len(s.messages))
  fmt.Printf("Messages are: %v\n",s.messages)
  for _,message := range s.messages {
    fmt.Println(message.messageType)
    fmt.Println(message.fieldMap)
  }
}

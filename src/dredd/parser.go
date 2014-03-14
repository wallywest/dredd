package dredd
import(
  "strings"
  "fmt"
)

type Parser struct {
  loadmon_message_chan chan *loadmonMessage
  messages []loadmonMessage
}

func (s *Parser) Process(data string) {
  split := strings.Split(data,"}")
  for _,v := range split[0:len(split)-1] {
    record := strings.Split(v,"|")
    m_type := strings.TrimPrefix(record[0],"\"\"")
    fields := record[1:len(record)]
    /*fmt.Println("TYPE IS")*/
    //fmt.Println(m_type)

    message := NewLoadmonMessage(m_type,fields)
    s.loadmon_message_chan <-message
    //s.messages = append(s.messages,*message)
  }
  //s.message_chan <- s.messages
}

func (s *Parser) Log() {
  fmt.Printf("Messages length is: %v\n",len(s.messages))
  fmt.Printf("Messages are: %v\n",s.messages)
  for _,message := range s.messages {
    fmt.Println(message.messageType)
    fmt.Println(message.fieldMap)
  }
}

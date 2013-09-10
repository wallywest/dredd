package listener

import(
  "bufio"
  "net"
  "fmt"
)

const(
  LOADMON_END = "`"
)

type TCPListener struct {
  c_messages chan string
  c_packets chan string
  addr string
  port int
}

func TCPListen(addr string, port int) (listener *TCPListener) {
  listener = &TCPListener{}

  listener.c_messages = make(chan string)
  listener.c_packets= make(chan string)

  listener.addr = addr
  listener.port = port

  go listener.listen()
  go listener.readRawSocket()
  return
}

func (t *TCPListener) listen() {
  for {
    select {
    case message := <-t.c_packets:
      fmt.Println("got packet")
      stream := &Stream{data: message}
      go stream.Process()
    }
  }
}

func (t *TCPListener) readRawSocket() {
  socket, err := net.Listen("tcp",":10503")
  defer socket.Close()
  if err != nil {
    fmt.Println(err)
  }

  for {
    conn, err := socket.Accept()
    reader := bufio.NewReader(conn)

    if err != nil {
      fmt.Printf("Error: %v",err)
    }

    go t.readPackets(conn,reader)
  }
}

func (t *TCPListener) readPackets(conn net.Conn, reader *bufio.Reader) {
  for {
    message, err := reader.ReadString('`')
    if err != nil {
      //fmt.Printf("Error: %v",err)
      return
    }
    t.c_packets<- message
    t.sendAck(conn)
  }
}

func (t *TCPListener) sendAck(conn net.Conn){
  fmt.Println("Writing Ack")
  b := []byte("OK`")
  conn.Write(b)
}

func (t *TCPListener) Receive() string{
  return <-t.c_messages
}

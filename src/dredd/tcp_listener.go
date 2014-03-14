package dredd

import(
  "bufio"
  "net"
  "fmt"
)

const(
  LOADMON_END = "`"
  SOCKET = ":10503"
)

type TCPListener struct {
  raw_messages chan string
  addr string
  port int
}

func NewTCPListener(addr string, port int) (listener *TCPListener) {
  listener = &TCPListener{}

  listener.raw_messages = make(chan string)

  listener.addr = addr
  listener.port = port

  go listener.readSocket()
  return
}

func (t *TCPListener) readSocket() {
  socket, err := net.Listen("tcp",SOCKET)
  defer socket.Close()

  if err != nil {
    fmt.Printf("Error is: %v\n",err)
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
      return
    }
    t.raw_messages <- message
    t.sendAck(conn)
  }
}

func (t *TCPListener) sendAck(conn net.Conn){
  b := []byte("OK`")
  conn.Write(b)
}

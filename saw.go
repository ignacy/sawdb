package main

import (
  "log"
  "net"
  "sync"

  "github.com/ignacy/sawdb/api"
)

type ConnectionManager struct {
  initiated       bool
  ConnectionCount int
  Db              *api.SawDB
}

func Initiate() *ConnectionManager {
  return &ConnectionManager{
    initiated: false,
    Db:        api.NewSawDB(),
  }
}

func (cM *ConnectionManager) Listen(listener net.Listener) {
  log.Println("Waiting for connections")
  var wg sync.WaitGroup

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println("Connection error", err)
      conn.Close()
    } else {
      cM.ConnectionCount++
      log.Printf("%s connected. Active connections: %d", conn.RemoteAddr(), cM.ConnectionCount)
      wg.Add(1)
      go func() {
        cM.handleMessage(conn)
        wg.Done()
      }()

      go func() {
        wg.Wait()
        conn.Close()
      }()

    }
  }
}

func (cM *ConnectionManager) handleMessage(conn net.Conn) (err error) {
  defer func(cM *ConnectionManager) {
    cM.ConnectionCount--
  }(cM)

  messageBuffer := make([]byte, 1024)

  _, err = conn.Read(messageBuffer)
  if err != nil {
    log.Println("Failed to read in a message")
    return
  }

  message := string(messageBuffer)
  log.Printf("Received message: %+v \n", message)

  request, err := api.NewRequest(message)
  if err != nil {
    log.Println("Error while creating DB request ", err)
    return
  }

  v, err := request.Process(*cM.Db)

  log.Printf("Values: %+v \n", cM.Db.Items)

  if err != nil {
    log.Println("Failed request processing", err)
    return
  }

  if v != "" {
    conn.Write([]byte(v))
  }
  conn.Write([]byte("All good action handled"))
  return
}

func main() {
  serverClosed := make(chan bool)

  listener, err := net.Listen("tcp", ":9000")
  if err != nil {
    log.Fatal("Failed to initialize DB instance")
  }

  connManage := Initiate()
  go connManage.Listen(listener)

  <-serverClosed
}

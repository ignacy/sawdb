package main

import (
  "log"
  "net"

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

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println("Connection error", err)
    }

    cM.ConnectionCount++
    log.Printf("%s connected. Active connections: %d", conn.RemoteAddr(), cM.ConnectionCount)

    go cM.handleMessage(conn)
  }
}

func (cM *ConnectionManager) handleMessage(conn net.Conn) {
  messageBuffer := make([]byte, 1024)
  _, err := conn.Read(messageBuffer)
  if err != nil {
    log.Println("Failed to read in a message")
  }

  message := string(messageBuffer)
  log.Println("Received was: ", message)

  request, err := api.NewRequest(message)
  if err != nil {
    log.Println("Error while creating DB request ", err)
  }

  v, err := request.Process(*cM.Db)

  if err != nil {
    log.Println("Failed request processing", err)
  }

  if v != "" {
    conn.Write([]byte(v))
  } else {
    conn.Write([]byte("All good action handled"))
  }

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

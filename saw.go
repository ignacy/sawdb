package main

import (
  "log"
  "net"
)

type ConnectionManager struct {
  initiated bool
}

func Initiate() *ConnectionManager {
  return &ConnectionManager{
    initiated: false,
  }
}

func (cM *ConnectionManager) Listen(listener net.Listener) {
  log.Println("Waiting for connections")

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println("Connection error", err)
    }

    log.Println(conn.RemoteAddr(), " connected")
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

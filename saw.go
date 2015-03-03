package main

import (
  "errors"
  "log"
  "net"
  "strings"

  "github.com/ignacy/sawdb/api"
)

type ConnectionManager struct {
  initiated bool
  db        *api.SawDB
}

func Initiate() *ConnectionManager {
  return &ConnectionManager{
    initiated: false,
    db:        api.NewSawDB(),
  }
}

func handleAction(action string, cM *ConnectionManager) error {
  components := strings.Split(action, "\t")
  if len(components) < 3 {
    return errors.New("Malformed action description")
  }

  if components[0] == "S" {
    cM.db.Put(components[1], components[2])
  }

  value, err := cM.db.Get(components[1])
  if err != nil {
    log.Println("Failed to store record")
  } else {
    log.Println("Stored record: ", value)
  }
  return nil
}

func (cM *ConnectionManager) Listen(listener net.Listener) {
  log.Println("Waiting for connections")

  for {
    conn, err := listener.Accept()
    if err != nil {
      log.Println("Connection error", err)
    }

    log.Println(conn.RemoteAddr(), " connected")

    messageBuffer := make([]byte, 1024)
    _, err = conn.Read(messageBuffer)
    if err != nil {
      log.Println("Failed to read in a message")
    }

    message := string(messageBuffer)
    log.Println("Received was: ", message)

    if err = handleAction(message, cM); err != nil {
      log.Println("Error while handling action ", err)
    }

    conn.Write([]byte("All good action handled \t\r\n"))
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

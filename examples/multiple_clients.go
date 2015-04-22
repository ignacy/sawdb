package main

import (
  "fmt"
  "log"
  "net"
)

const Host = "127.0.0.1:9000"

func main() {
  conn1, err := net.Dial("tcp", Host)
  if err != nil {
    fmt.Println("Could not connect to sawDB")
  }

  defer conn1.Close()

  conn2, err := net.Dial("tcp", Host)
  if err != nil {
    fmt.Println("Could not connect to sawDB")
  }

  defer conn2.Close()

  keyToStore := []byte("S\tfoo\tbar\t")
  go conn1.Write(keyToStore)

  messageBuffer := make([]byte, 1024)
  _, err = conn1.Read(messageBuffer)
  if err != nil {
    log.Println("Failed to read in a message")
  }

  response := string(messageBuffer)
  log.Println("Received response was: ", response)

  getValue := []byte("G\tfoo\t\r\n")
  go conn2.Write(getValue)

  messageBuffer = make([]byte, 1024)
  _, err = conn2.Read(messageBuffer)
  if err != nil {
    log.Println("Failed to read in a message")
  }

  response = string(messageBuffer)
  log.Println("Received response was: ", response)
}

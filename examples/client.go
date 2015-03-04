package main

import (
  "fmt"
  "log"
  "net"
)

const Host = "127.0.0.1:9000"

func main() {
  conn, err := net.Dial("tcp", Host)
  if err != nil {
    fmt.Println("Could not connect to sawDB")
  }

  keyToStore := []byte("S\tkey\tvalue\t")
  conn.Write(keyToStore)

  messageBuffer := make([]byte, 1024)
  _, err = conn.Read(messageBuffer)
  if err != nil {
    log.Println("Failed to read in a message")
  }

  response := string(messageBuffer)
  log.Println("Received response was: ", response)

  getValue := []byte("G\tvalue\t\r\n")
  conn.Write(getValue)

  messageBuffer = make([]byte, 1024)
  _, err = conn.Read(messageBuffer)
  if err != nil {
    log.Println("Failed to read in a message")
  }

  response = string(messageBuffer)
  log.Println("Received response was: ", response)
  conn.Close()

}

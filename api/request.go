package api

import (
  "errors"
  "fmt"
  "strings"
)

type MessageType int

const (
  PUT MessageType = iota
  GET
  DELETE
)

type Request struct {
  Action  MessageType
  command []string
}

func NewRequest(rawMessage string) (*Request, error) {
  components := strings.Split(rawMessage, "\t")
  if len(components) < 3 {
    return nil, errors.New("Malformed action description")
  }

  switch components[0] {
  case "S":
    return &Request{Action: PUT, command: components[1:]}, nil
  case "G":
    return &Request{Action: GET, command: components[1:]}, nil
  case "D":
    return &Request{Action: DELETE, command: components[1:]}, nil
  default:
    return nil, errors.New(fmt.Sprintf("Unsupported action: %s", components[0]))
  }

}

func (r *Request) Process(db SawDB) (value string, err error) {
  switch r.Action {
  default:
    return "", errors.New("This shouldn't happen: usuported action")
  case PUT:
    return "", db.Put(r.command[0], r.command[1])
  case GET:
    return db.Get(r.command[0])
  case DELETE:
    return "", db.Delete(r.command[0])
  }

}

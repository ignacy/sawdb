package api

import (
  "errors"
  "fmt"
)

type SawDB struct {
  Name  string
  Items map[string]string
}

func NewSawDB() *SawDB {
  return &SawDB{
    Name:  "new db instance",
    Items: make(map[string]string),
  }
}

func (db SawDB) Put(key string, value string) (err error) {
  fmt.Printf("Inserting %s under key %s \n", value, key)
  db.Items[key] = value
  if _, ok := db.Items[key]; ok != true {
    return errors.New("Unable to save record into db")
  }
  return nil
}

func (db SawDB) Get(key string) (value string, err error) {
  fmt.Println("Retriving value for key: ", key)
  value, ok := db.Items[key]

  if ok != true {
    return "", errors.New("Record not found")
  }

  return value, nil
}

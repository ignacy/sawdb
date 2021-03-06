package api

import (
  "errors"
  "log"
)

var (
  ErrRecordNotFound = errors.New("Record not found in database")
  ErrSaveFailed     = errors.New("Unable to save record into db")
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
  log.Println("Inserting value, key", value, key)
  db.Items[key] = value
  if _, ok := db.Items[key]; ok != true {
    return ErrSaveFailed
  }
  return nil
}

func (db SawDB) Get(key string) (value string, err error) {
  log.Println("Retriving value for key: ", key)
  value, ok := db.Items[key]

  if ok != true {
    return "", ErrRecordNotFound
  }

  return value, nil
}

func (db SawDB) Delete(key string) (err error) {
  log.Println("Removing record for key:", key)

  _, ok := db.Items[key]

  if ok != true {
    return ErrRecordNotFound
  }

  delete(db.Items, key)

  return nil
}

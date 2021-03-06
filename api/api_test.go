package api

import (
  "testing"
)

func Test_Put(t *testing.T) {
  db := NewSawDB()
  if ok := db.Put("hello", "world"); ok != nil {
    t.Error("Failed to save record to DB")
  } else {
    t.Log("Record saved")
  }
}

func Test_Put_AndThen_Get(t *testing.T) {
  db := NewSawDB()
  db.Put("key", "value")

  value, err := db.Get("key")
  if err != nil {
    t.Error("Failed to retrive record from DB")
  } else if value != "value" {
    t.Error("Retrived unexpected value")
  } else {
    t.Log("Record saved and then retrived.")
  }
}

func Test_Get_NonExistingRecord(t *testing.T) {
  db := NewSawDB()

  _, err := db.Get("NonExistingKey")
  if err == nil {
    t.Error("Expected error - missing value")
  } else {
    t.Log("Returns error when value for key is missing")
  }
}

func Test_Put_ThenDelete_AndThen_Get(t *testing.T) {
  db := NewSawDB()
  db.Put("something", "different")

  if err := db.Delete("something"); err != nil {
    t.Error("Failed to remove record")
  }

  if _, err := db.Get("something"); err == nil {
    t.Error("Expected an error as the record should be missing")
  } else {
    t.Log("Record saved and then removed from DB")
  }

}

func Test_Delete_Non_Existing(t *testing.T) {
  db := NewSawDB()

  err := db.Delete("something")

  if err == nil {
    t.Error("Expected an error as the record should be missing")
  } else {
    t.Log("Error returned when trying to delete non exisitng record")
  }

}

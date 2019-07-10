package models

type Word struct {
  ID string 
  Definitions string
  Master bool `gorm:"default:false"`
}
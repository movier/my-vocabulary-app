package models

type Stem struct {
  ID string
  Definitions string
  Master bool `gorm:"default:false"`
}
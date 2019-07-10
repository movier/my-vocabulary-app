package models

import (
  "github.com/jinzhu/gorm"
)

type Lookup struct {
  gorm.Model
  Usage string
  WordID string
  Word Word
}
package models

type Word struct {
  ID string
  StemID string
  Stem Stem 
  Lookups []Lookup `gorm:"many2many:lookup_words;"`
}
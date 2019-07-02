package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "fmt"
)

type Person struct {
  gorm.Model
  Name string
  Age int
}

var db *gorm.DB

func main() {

  username := "postgres"
  password := "postgres" 
	dbName := "my_database" 
	dbHost := "db" 


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
	}

	db = conn
	db.Debug().AutoMigrate(&Person{})
}

func GetDB() *gorm.DB {
	return db
}
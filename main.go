package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/gorilla/mux"
  "fmt"
  "net/http"
  "encoding/json"
)

type Person struct {
  gorm.Model
  Name string
  Age int
}

func getVocabulary(w http.ResponseWriter, r *http.Request) {
  username := "postgres"
  password := "postgres" 
	dbName := "my_database" 
	dbHost := "db" 

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
  
  db, err := gorm.Open("postgres", dbUri)
	if err != nil {
		fmt.Print(err)
  }
  
  db.Create(&Person{Name: "Peter", Age: 22})

  var person []Person
  db.Find(&person)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func main() {
  router := mux.NewRouter()
  router.HandleFunc("/api/vocabulary", getVocabulary).Methods("GET")

  port := "8080"
  fmt.Println(port)
  err := http.ListenAndServe(":" + port, router)
  if err != nil {
    fmt.Print(err)
  }
}
package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/gorilla/mux"
  "fmt"
  "net/http"
  "encoding/json"
  "app/models"
  "html/template"
)

var db *gorm.DB

func getVocabulary(w http.ResponseWriter, r *http.Request) {

  db.Create(&models.Person{Name: "Johnson", Age: 22})

  var person []models.Person
  db.Find(&person)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(person)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
  p := &models.Person{Name: "title", Age: 22}
  t, _ := template.ParseFiles("upload.html")
  t.Execute(w, p)
}

func main() {
  username := "postgres"
  password := "postgres" 
	dbName := "my_database" 
	dbHost := "db" 


	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(dbUri)

	conn, dbErr := gorm.Open("postgres", dbUri)
	if dbErr != nil {
		fmt.Print(dbErr)
	}
  defer conn.Close()

  db = conn
  db.Debug().AutoMigrate(&models.Person{})

  router := mux.NewRouter()
  router.HandleFunc("/api/vocabulary", getVocabulary).Methods("GET")
  router.HandleFunc("/upload", uploadFile).Methods("GET")

  port := "8080"
  fmt.Println(port)
  err := http.ListenAndServe(":" + port, router)
  if err != nil {
    fmt.Print(err)
  }
}
package main

import (
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  _ "github.com/jinzhu/gorm/dialects/sqlite"
  "github.com/gorilla/mux"
  "fmt"
  "net/http"
  "encoding/json"
  "app/models"
  "html/template"
  "io/ioutil"
  "os"
  "log"
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

func getWord(w http.ResponseWriter, r *http.Request) {
  // lookup := &models.Lookup{
  //   Word: models.Word{ID: "new"},
  //   Usage: "",
  // }
  // db.Create(&lookup)

  var lookups []models.Lookup
  db.Preload("Word").Limit(100).Order(gorm.Expr("random()")).Find(&lookups)

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(lookups)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
  p := &models.Stem{ID: "title", Definitions: "22"}
  t, _ := template.ParseFiles("upload.html")
  t.Execute(w, p)
}

func receiveFile(w http.ResponseWriter, r *http.Request) {
  fmt.Println("File Upload Endpoint Hit")

  // Parse our multipart form, 10 << 20 specifies a maximum
  // upload of 10 MB files.
  r.ParseMultipartForm(10 << 20)
  // FormFile returns the first file for the given key `myFile`
  // it also returns the FileHeader so we can get the Filename,
  // the Header and the size of the file
  file, handler, err := r.FormFile("myFile")
  if err != nil {
    fmt.Println("Error Retrieving the File")
    fmt.Println(err)
    return
  }
  defer file.Close()
  fmt.Printf("Uploaded File: %+v\n", handler.Filename)
  fmt.Printf("File Size: %+v\n", handler.Size)
  fmt.Printf("MIME Header: %+v\n", handler.Header)

  dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
  }
  fmt.Printf("current dir: %+v\n", dir)
  // Create a temporary file within our temp-images directory that follows
  // a particular naming pattern
  tempFile, err := ioutil.TempFile(dir + "/temp-images", "upload-*.db")
  if err != nil {
    fmt.Println("Create Error: %+v\n", err)
  }
  defer tempFile.Close()

  // read all of the contents of our uploaded file into a
  // byte array
  fileBytes, err := ioutil.ReadAll(file)
  if err != nil {
    fmt.Println(err)
  }
  // write this byte array to our temporary file
  tempFile.Write(fileBytes)

  sqlite, err := gorm.Open("sqlite3", tempFile.Name())
  if err != nil {
    fmt.Println("open sqlite error:", err)
  }
  rows, err := sqlite.Raw("select word_key, usage from lookups").Rows()
  if err != nil {
    fmt.Println(err)
  }
  for rows.Next() {
    word_key := ""
    usage := ""
    rows.Scan(&word_key, &usage)
    fmt.Printf("word: %+v, usage: %+v\n", word_key, usage)
    lookup := &models.Lookup{
      Word: models.Word{ID: word_key},
      Usage: usage,
    }
    db.Create(&lookup)
  }
  defer sqlite.Close()

  // return that we have successfully uploaded our file!
  fmt.Fprintf(w, "Successfully Uploaded File\n")
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
  db.Debug().AutoMigrate(&models.Person{}, &models.Lookup{}, &models.Word{})

  router := mux.NewRouter()
  router.HandleFunc("/api/vocabulary", getVocabulary).Methods("GET")
  router.HandleFunc("/api/word", getWord).Methods("GET")
  router.HandleFunc("/upload", uploadFile).Methods("GET")
  router.HandleFunc("/upload", receiveFile).Methods("POST")

  port := "8080"
  fmt.Println(port)
  err := http.ListenAndServe(":" + port, router)
  if err != nil {
    fmt.Print(err)
  }
}
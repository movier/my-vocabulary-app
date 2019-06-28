package main

import (
  "github.com/gorilla/mux"
  "fmt"
  "net/http"
  "encoding/json"
)

type Person struct {
  Name string
  Age int
}

func getVocabulary(w http.ResponseWriter, r *http.Request) {
  d := Person{"Bobb", 24}
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(d)
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
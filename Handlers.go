package main

import (
  "encoding/json"
  "fmt"
  "strconv"
  "net/http"
  "github.com/gorilla/mux"
)

var todos = Todos{
  Todo{Name: "Write presentation", Completed: true},
  Todo{Name: "Host meetup"},
}

func Index(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintln(w, "Welcome")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
  if err := json.NewEncoder(w).Encode(todos); err != nil {
    panic(err)
  }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)
  todoId, err := strconv.ParseInt(vars["todo"], 10, 64)

  if err != nil {
    panic(err)
  }

  if err1 := json.NewEncoder(w).Encode(todos[todoId]); err1 != nil {
    panic(err)
  }
}
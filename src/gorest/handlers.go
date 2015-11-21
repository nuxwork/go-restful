package main

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/gorilla/mux"
)

func Index(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Welcome to gorestful")
}

func NotFound(w http.ResponseWriter, r *http.Request){
    fmt.Fprintln(w, "Erro Not Found")
}

func PostObjects(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    fmt.Fprintln(w, "Post objects: ", name)
}

func GetObjects(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    fmt.Fprintln(w, "Get objects: ", name)
}

func GetObjectsById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    id := vars["id"]
    fmt.Fprintln(w, "Get objects: ", name, ", id = ", id)
}

func PutObjectsById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    id := vars["id"]
    fmt.Fprintln(w, "Put objects: ", name, ", id = ", id)
}

func DeleteObjectsById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    name := vars["name"]
    id := vars["id"]
    fmt.Fprintln(w, "Delete objects: ", name, ", id = ", id)
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}
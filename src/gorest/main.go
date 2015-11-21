package main

import (
    "log"
    "net/http"
)

func main() {

    router := NewRouter("https", "rest.mumu.chat")

    log.Fatal(http.ListenAndServe(":8080", router))
}
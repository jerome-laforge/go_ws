package main

import (
	"github.com/jerome-laforge/go_ws"
	"log"
	"net/http"
)

func main() {
	router := go_ws.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

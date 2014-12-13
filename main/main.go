package main

import (
	"log"
	"net/http"
	"github.com/jerome-laforge/go_ws"
)

func main() {
	router := go_ws.NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}

package main

import (
	"log"
	"net/http"
	"phoenix-data-service/handler"
)

func main() {
	http.HandleFunc("/ping", handler.PingHandler)
	http.HandleFunc("/feed", handler.FeedFileHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

package handler

import (
	"fmt"
	"io"
	"net/http"
	"phoenix-data-service/service"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping")
	w.Write([]byte("ping"))
}

func FeedFileHandler(w http.ResponseWriter, r *http.Request) {
	file, _, err := r.FormFile("myFile")
	defer file.Close()
	if err != nil {
		fmt.Println("Error in reading uploaded file : ", err)
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error in reading uploaded file : ", err)
	}
	contents := string(fileBytes[:])
	service.ReadFeed(contents)
}

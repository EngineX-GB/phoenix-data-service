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
	refreshTimeInDB := service.CheckLoadDate()
	fmt.Print("Maximum refresh time in database is : " + refreshTimeInDB.Format("2006-01-02 15:04:05"))
	service.ReadFeed(contents)
}

package handler

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"phoenix-data-service/service"
	"regexp"
	"time"
)

func PingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ping")
	w.Write([]byte("ping"))
}

func FeedFileHandler(w http.ResponseWriter, r *http.Request) {
	file, metadata, err := r.FormFile("myFile")
	defer file.Close()
	if err != nil {
		fmt.Println("Error in reading uploaded file : ", err)
		return
	}
	dateOnFeedFile, err := checkLoadDateOnFile(metadata.Filename)
	refreshTimeInDB := service.CheckLoadDate()
	if err != nil {
		log.Fatal(err)
	}

	if dateOnFeedFile.Compare(refreshTimeInDB) < 0 {
		log.Println("Date on file is earlier than the latest load date")
		return
	}

	fileBytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error in reading uploaded file : ", err)
		return
	}
	contents := string(fileBytes[:])

	fmt.Print("Maximum refresh time in database is : " + refreshTimeInDB.Format("2006-01-02 15:04:05"))
	service.ReadFeed(contents)
}

func checkLoadDateOnFile(filename string) (time.Time, error) {
	// use a regex to extract the date:
	datePattern := `\d{4}-\d{2}-\d{2}`
	re := regexp.MustCompile(datePattern)
	date := re.FindString(filename)

	if date != "" {
		fmt.Printf("Extracted date: %s\n", date)
		dateObj, err := time.Parse("2006-01-02", date)
		if err != nil {
			return time.Now(), errors.New("error in parsing date")
		}
		return dateObj, nil
	} else {
		return time.Now(), errors.New("no error found in file " + filename)
	}

}

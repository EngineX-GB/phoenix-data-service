package service

import (
	"encoding/csv"
	"log"
	"phoenix-data-service/dao"
	"phoenix-data-service/domain"
	"strconv"
	"strings"
	"time"
)

func ReadFeed(contents string) {
	reader := strings.NewReader(contents)
	csvReader := csv.NewReader(reader)
	csvReader.Comma = '|'
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var clientList []domain.Client

	for _, record := range records {
		var client domain.Client
		client.Username = record[0]
		client.Nationality = record[1]
		client.Location = record[2]
		client.Rating = convertStrToInt(record[3])
		client.Age = convertStrToInt(record[4])
		client.R15 = convertStrToInt(record[5])
		client.R30 = convertStrToInt(record[6])
		client.R45 = convertStrToInt(record[7])
		client.R100 = convertStrToInt(record[8])
		client.R150 = convertStrToInt(record[9])
		client.R200 = convertStrToInt(record[10])
		client.R250 = convertStrToInt(record[11])
		client.R300 = convertStrToInt(record[12])
		client.R350 = convertStrToInt(record[13])
		client.R400 = convertStrToInt(record[14])
		client.RON = convertStrToInt(record[15])
		client.Telephone = record[16]
		client.Url = record[17]
		client.RefreshTime = convertStrToTime(record[18])
		client.UserId = record[19]
		//imageavailable
		client.Region = record[21]
		client.Gender = record[22]
		// member since
		//height
		//dsize
		//eyecolor
		//haircolor

		// add to the client list.
		clientList = append(clientList, client)
	}

	if len(clientList) > 0 {
		log.Println("Client list is " + strconv.Itoa(len(clientList)) + " clients")
		dao.AddClient(clientList)
	}
}

func convertStrToInt(value string) int {
	if value == "Not Specified" {
		return 0
	}
	num, err := strconv.Atoi(value)
	if err != nil {
		log.Fatal("Cannot convert string to int, " + err.Error())
	}
	return num
}

func convertStrToTime(value string) time.Time {
	timeLayout := "2006-01-02 15:04:05"
	refreshTime, err := time.Parse(timeLayout, value)
	if err != nil {
		log.Fatal("Cannot convert string to time, " + err.Error())
	}
	return refreshTime
}

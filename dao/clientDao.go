package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"phoenix-data-service/config"
	"phoenix-data-service/domain"
)

func AddClient(clientList []domain.Client) {
	db := connect()
	for _, client := range clientList {
		_, err := db.Query(InsertClientQuery(), client.Username, client.Nationality, client.Location, client.Rating,
			client.Age, client.R15, client.R30, client.R100, client.R150, client.R200, client.R250, client.R300, client.R350,
			client.R400, client.RON, client.Telephone, client.Url, client.UserId, client.Region)
		if err != nil {
			log.Fatal(err)
		}
	}
	defer db.Close()

	//
	//"rate_30_min, rate_1_hour, rate_1_50_hour, rate_2_hour, rate_2_50_hour, rate_3_hour, rate_3_50_hour, rate_4_hour, rate_overnight," +
	//	"telephone, url_page, user_id, region)" +
	//

}

func connect() *sql.DB {
	db, err := sql.Open(config.GetDriverName(), config.GetDataSourceConfig())
	if err != nil {
		panic(err.Error())
	}
	return db
}

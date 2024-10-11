package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"phoenix-data-service/config"
	"phoenix-data-service/domain"
	"time"
)

func GetMaxLoadDate() time.Time {
	db := connect()
	var refreshTime sql.NullTime
	result := db.QueryRow(GetMaxRefreshTimeQuery())
	err := result.Scan(&refreshTime)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	if refreshTime.Valid {
		return refreshTime.Time
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", "1999-01-01 00:00:00")
	if err != nil {
		log.Fatal(err)
	}
	return parsedTime
}

func AddClient(clientList []domain.Client) {
	db := connect()
	defer db.Close() // Ensure the connection is closed at the end

	// Use a transaction for better performance and rollback capabilities
	tx, err := db.Begin()
	if err != nil {
		log.Fatal("Failed to begin transaction:", err)
	}

	stmt, err := tx.Prepare(InsertClientQuery()) // Prepare the statement once
	if err != nil {
		log.Fatal("Failed to prepare statement:", err)
	}
	defer stmt.Close()

	for _, client := range clientList {
		// Execute the prepared statement with the client's data
		_, err = stmt.Exec(client.Username, client.Nationality, client.Location, client.Rating,
			client.Age, client.R15, client.R30, client.R100, client.R150, client.R200, client.R250, client.R300, client.R350,
			client.R400, client.RON, client.Telephone, client.Url, client.UserId, client.Region, client.RefreshTime)

		if err != nil {
			tx.Rollback() // Rollback on error to prevent partial inserts
			log.Fatalf("Failed to insert client data: %v", err)
		}
	}

	// Commit the transaction if no errors
	if err := tx.Commit(); err != nil {
		log.Fatal("Failed to commit transaction:", err)
	}

	//log.Println("Processing the data in the staging table")

	// Call the stored procedure
	_, err = db.Exec("CALL prc_new_clean_up_data()")
	if err != nil {
		log.Fatal("Failed to call stored procedure:", err)
	}
}

func connect() *sql.DB {
	db, err := sql.Open(config.GetDriverName(), config.GetDataSourceConfig())
	if err != nil {
		panic(err.Error())
	}
	return db
}

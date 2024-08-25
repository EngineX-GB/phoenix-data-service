package dao

func InsertClientQuery() string {
	return "INSERT INTO tbl_client_temp (username, nationality, location, rating, age, rate_15_min, " +
		"rate_30_min, rate_1_hour, rate_1_50_hour, rate_2_hour, rate_2_50_hour, rate_3_hour, rate_3_50_hour, rate_4_hour, rate_overnight," +
		"telephone, url_page, user_id, region, refresh_time)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

func GetMaxRefreshTimeQuery() string {
	return "SELECT MAX(DATE(refresh_time)) FROM tbl_client"
}

package dao

func InsertClientQuery() string {
	return "INSERT INTO tbl_client_temp (username, nationality, location, rating, age, rate_15_min, " +
		"rate_30_min, rate_1_hour, rate_1_50_hour, rate_2_hour, rate_2_50_hour, rate_3_hour, rate_3_50_hour, rate_4_hour, rate_overnight," +
		"telephone, url_page, user_id, region)" +
		" VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
}

//return "INSERT INTO tbl_client_temp (username, nationality, location, rating, age, rate_15_min, " +
//"rate_30_min, rate_1_hour, rate_1_50_hour, rate_2_hour, rate_2_50_hour, rate_3_hour, rate_3_50_hour, rate_4_hour, rate_overnight," +
//"telephone, url_page, refresh_time, user_id, region, gender, member_since, height, dress_size, hair_colour, eye_colour)" +
//" VALUES ("

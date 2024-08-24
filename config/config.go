package config

func GetDataSourceConfig() string {
	return "root:root@tcp(127.0.0.1:3306)/db_phoenix?charset=utf8mb4&parseTime=True&loc=Local"
}

func GetDriverName() string {
	return "mysql"
}

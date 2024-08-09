package database

var connection string

func init() {
	connection = "mysql"
}

func GetConnection() string {
	return connection
}

package infrastructure

var mysql *MySQL

func GoMySQL() {
	mysql = NewMySQL()
}

func GetMySQL() *MySQL {
	return mysql 
}


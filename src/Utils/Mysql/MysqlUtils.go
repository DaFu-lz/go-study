package Mysql

import "database/sql"

var DB  *sql.DB

func ConnectionMysql(DbName string)  {
	var dataSourceName string
	dataSourceName= "root:lfc322325@tcp(localhost:3306)/"+DbName
	DB, err := sql.Open("mysql", dataSourceName)
	defer DB.Close()
}

func ()  {
	
}


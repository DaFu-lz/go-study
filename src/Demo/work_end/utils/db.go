package utils

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	Db  *sql.DB
	err error
)

func init() {
	Db, err = sql.Open("mysql", "root:lfc322325@tcp(localhost:3306)/works")
	if err != nil {
		panic(err.Error())
	}
}

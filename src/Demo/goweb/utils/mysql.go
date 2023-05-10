package utils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var err error

func init() {
	DB, err = sql.Open("mysql", "root:lfc322325@tcp(localhost:3306)/goweb")
	if err != nil {
		fmt.Println("数据库连接失败", err)
	}
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

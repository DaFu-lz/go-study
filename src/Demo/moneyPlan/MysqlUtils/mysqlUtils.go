/*
	使用前要定义DbName
*/
package MysqlUtils

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

var dataSourceName = "root:lfc322325@tcp(localhost:3306)/moneyPlan"

var driverName = "mysql"

type Account struct {
	id        int
	income    bool
	species   string
	price     float32
	detail    string
	tradeTime string
}

/*
	Connect to the database test, return true on success, false otherwise
*/
func Init() bool {
	_, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		return false
	}
	return true
}

/*
	Query today's data
*/
func QueryToday() []Account {
	DB, _ := sql.Open(driverName, dataSourceName)
	sqlStr := "select a.id, a.income, t.species, a.price, a.detail, a.trade_time " +
		"from account a,trade_type t " +
		"where a.trade_id = t.trade_id and to_days(a.trade_time) = to_days(now());"
	res, err := DB.Query(sqlStr)
	if err != nil {
		fmt.Println("数据查询失败")
	}
	defer res.Close()

	var data []Account
	for res.Next() {
		single := Account{}
		err1 := res.Scan(&single.id, &single.income, &single.species, &single.price, &single.detail, &single.tradeTime)
		if err1 != nil {
			fmt.Println("数据查询失败！")
		}
		data = append(data, single)
	}
	return data
}

/*
	Query the data of the past week
*/
func QueryWeek() []Account {
	DB, _ := sql.Open(driverName, dataSourceName)
	sqlStr := "select a.id, a.income, t.species, a.price, a.detail, a.trade_time " +
		"from account a,trade_type t " +
		"where a.trade_id = t.trade_id and DATE_SUB(CURDATE(), INTERVAL 6 DAY) <= date(a.trade_time);"
	res, err := DB.Query(sqlStr)
	if err != nil {
		fmt.Println("数据查询失败")
	}
	defer res.Close()

	var data []Account
	for res.Next() {
		single := Account{}
		err1 := res.Scan(&single.id, &single.income, &single.species, &single.price, &single.detail, &single.tradeTime)
		if err1 != nil {
			fmt.Println("数据查询失败！")
		}
		data = append(data, single)
	}
	return data
}

/*
	Registered income data
*/
func Income(income int, tradeId int, price float32, detail string, tradeTime string) error {
	DB, _ := sql.Open(driverName, dataSourceName)
	sqlStr := "insert into account(income, trade_id, price, detail, trade_time) " +
		"values (?,?,?,?,?);"
	res, err := DB.Exec(sqlStr, income, tradeId, price, detail, tradeTime)

	row, _ := res.RowsAffected()

	fmt.Printf("成功插入了%d行数据\n", row)

	return err
}

/*
	Registered expenditure data

*/
func Expenditure() {

}

//来源：source    生活

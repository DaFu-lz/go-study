package main

import (
	"Demo/moneyPlan/MysqlUtils"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var id int
var income int
var tradeId int
var price float32
var detail string
var tradeTime string

func main() {
	//测试是否能来连接到数据库
	if !MysqlUtils.Init() {
		fmt.Println("未连接到数据库，请联系管理员")
		return
	}

TOP:
	fmt.Println()
	fmt.Println("--------------------收支记录--------------------")
	fmt.Println("                  1 收支明细")
	fmt.Println("                  2 登记收入")
	fmt.Println("				   3 登记支出")
	fmt.Println("				   4 查询详细记录")
	fmt.Println("				   5 退 出\n")
	fmt.Print("                  请选择（1-5）：")

	var key int
	fmt.Scanln(&key)

	switch key {
	case 1:
		fmt.Println("今日收支明细：")
		today := MysqlUtils.QueryToday()
		MysqlUtils.InitData(today)

		fmt.Println("最近七天的收支明细：")
		week := MysqlUtils.QueryWeek()
		MysqlUtils.InitData(week)

	case 2:
		fmt.Println("登记收入")
		fmt.Print("收入/支出(1/0):")
		fmt.Scanln(&income)
		fmt.Print("用途（1：个人 2：家人 3：亲戚朋友 4：其他）:")
		fmt.Scanln(&tradeId)
		fmt.Print("金额:")
		fmt.Scanln(&price)
		fmt.Print("备注:")
		fmt.Scanln(&detail)
		fmt.Print("交易时间（xxxx-xx-xx xx:xx:xx）:")
		fmt.Scanln(&tradeTime)
		result := MysqlUtils.Income(income, tradeId, price, detail, tradeTime)
		if result != nil {
			goto TOP
		}
	case 3:
		fmt.Println("登记支出")
	case 4:
		fmt.Println("查询详细记录")

	case 5:
		fmt.Print("是否确定退出（y/s）:")
		var quit string
		fmt.Scanln(&quit)
		if quit == "y" || quit == "Y" {
			return
		}
	default:
		fmt.Println("无效的输入，请输入（1-5）")
	}
	fmt.Println("有一次寻新欢")
	goto TOP
}

func mainPage() {
	//数据处理
	//记账 生活日常
}

package MysqlUtils

import "fmt"

var id int
var income string
var species string
var price float32
var detail string
var tradeTime string

/*
	Formatted output
*/
func InitData(data []Account) {
	fmt.Println("编号\t\t类型\t\t用途\t\t金额\t\t备注\t\t\t日期")

	for _, v := range data {
		money := v

		id = money.id
		if money.income == false {
			income = "支出"
		} else {
			income = "收入"
		}
		species = money.species
		price = money.price
		detail = money.detail
		tradeTime = money.tradeTime

		fmt.Println(id, "\t", income, "\t", species, "\t", price, "\t", detail, "\t", tradeTime)
	}
	fmt.Println()
}

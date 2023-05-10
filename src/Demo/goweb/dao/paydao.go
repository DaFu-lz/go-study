package dao

import (
	"Demo/goweb/model"
	"Demo/goweb/utils"
	"Utils/TimeUtils"
	"fmt"
)

//Get today's consumption record
func PayToday() (error, []model.ConsumptionRecode) {
	//获取今天消费记录
	sqlStr := "select con.consumption_recode_id,con.income,con.purpose,pay.pay_type,pay.pay_type_card,con.price,con.detail,con.trade_time,con.recode_time" +
		"from consumption_recode con,pay_type pay " +
		"where con.pay_type_id = pay.pay_type_id " +
		" and to_days(con.trade_time) = to_days(now());"

	rows, _ := utils.DB.Query(sqlStr)
	defer rows.Close()

	var results []model.ConsumptionRecode
	for rows.Next() {
		result := model.ConsumptionRecode{}
		err := rows.Scan(&result.ConsumptionRecodeId, &result.Income, &result.Purpose, &result.PayType.PayType, &result.PayType.PayTypeCard, &result.Price, &result.Detail, &result.TradeTime, &result.RecodeTime)
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		results = append(results, result)
	}
	return nil, results
}

//Get the last week's spending record
func PayWeek() (error, []model.ConsumptionRecode) {
	//获取最近一个周的消费记录
	sqlStr := "select con.consumption_recode_id,con.income,con.purpose,pay.pay_type,pay.pay_type_card,con.price,con.detail,con.trade_time,con.recode_time" +
		"from consumption_recode con,pay_type pay " +
		"where con.pay_type_id = pay.pay_type_id " +
		" DATE_SUB(CURDATE(), INTERVAL 6 DAY) <= date(con.trade_time);"

	rows, _ := utils.DB.Query(sqlStr)
	defer rows.Close()

	var results []model.ConsumptionRecode
	for rows.Next() {
		result := model.ConsumptionRecode{}
		err := rows.Scan(&result.ConsumptionRecodeId, &result.Income, &result.Purpose, &result.PayType.PayType, &result.PayType.PayTypeCard, &result.Price, &result.Detail, &result.TradeTime, &result.RecodeTime)
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		results = append(results, result)
	}
	return nil, results
}

//Get the last 30 days of spending
func PayMonth() (error, []model.ConsumptionRecode) {
	//获取最近30天的消费记录
	sqlStr := "select con.consumption_recode_id,con.income,con.purpose,pay.pay_type,pay.pay_type_card,con.price,con.detail,con.trade_time,con.recode_time " +
		"from consumption_recode con,pay_type pay " +
		"where con.pay_type_id = pay.pay_type_id " +
		" and DATE_SUB(CURDATE(), INTERVAL 30 DAY) <= date(con.trade_time) order by trade_time;"

	rows, _ := utils.DB.Query(sqlStr)
	defer rows.Close()

	var results []model.ConsumptionRecode
	for rows.Next() {
		result := model.ConsumptionRecode{}
		err := rows.Scan(&result.ConsumptionRecodeId, &result.Income, &result.Purpose, &result.PayType.PayType, &result.PayType.PayTypeCard, &result.Price, &result.Detail, &result.TradeTime, &result.RecodeTime)
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		results = append(results, result)
	}
	return nil, results
}

//Add consumption record
func AddPayRecode(data model.ConsumptionRecode) error {
	//添加消费记录
	sqlStr := "insert into consumption_recode(income, purpose, pay_type_id, price, detail, trade_time, recode_time) " +
		"values (?,?,?,?,?,?,?);"

	_, err := utils.DB.Exec(sqlStr, data.Income, data.Purpose, data.PayType.PayTypeId, data.Price, data.Detail, data.TradeTime, TimeUtils.GetTime())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil

}

//Access to payment method
func GetPayMethod() (error, []model.PayType) {
	//获取支付方式
	sqlStr := "select * " +
		"from pay_type;"

	rows, _ := utils.DB.Query(sqlStr)
	defer rows.Close()

	var results []model.PayType
	for rows.Next() {
		result := model.PayType{}
		err := rows.Scan(&result.PayTypeId, &result.PayType, &result.PayTypeCard, &result.Balance)
		if err != nil {
			fmt.Println(err)
			return err, nil
		}
		results = append(results, result)
	}
	return nil, results
}

//Query balance
func QueryBalance() ([]model.PayType, float64) {
	//查询余额：balances:明细余额，balance:总余额
	//
	sqlStr := "select sum(balance) " +
		"from pay_type;"

	_, balances := GetPayMethod()
	rows, _ := utils.DB.Query(sqlStr)
	var balance float64
	err := rows.Scan(balance)
	if err != nil {
		fmt.Println(err)
		return balances, 0
	}
	return balances, balance
}

//Balance change
func ChangeBalance(pay_type_id int, amount float64) {
	//余额变更
	sqlStr := "update pay_type " +
		"set balance = ? " +
		"where pay_type_id = ?;"

	_, err := utils.DB.Exec(sqlStr, pay_type_id, amount)
	if err != nil {
		fmt.Println(err)
	}
}

//Check whether the record exists. By spending time and amount to determine,Existence returns true
func CheckRecode(trade_time string, price float64) bool {
	//检查记录是否存在。  通过消费时间和金额来确定,存在返回true
	sqlStr := "select * " +
		"from consumption_recode " +
		"where trade_time = ? " +
		"and price like ?;"
	rows, _ := utils.DB.Query(sqlStr, trade_time, price)
	if rows.Next() {
		return true
	}
	return false
}

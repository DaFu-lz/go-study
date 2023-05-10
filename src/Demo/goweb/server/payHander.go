package server

import (
	"Demo/goweb/dao"
	"Demo/goweb/model"
	"html/template"
	"net/http"
	"strconv"
)

func AddPayPage(w http.ResponseWriter, r *http.Request) {

	var page model.Page
	_, page.PayType = dao.GetPayMethod()

	t := template.Must(template.ParseFiles("src/Demo/goweb/views/pages/addPayPage.html"))
	t.Execute(w, page)

}

func AddPayRecode(w http.ResponseWriter, r *http.Request) {
	//获取网页的数据
	var data model.ConsumptionRecode
	data.Income = r.PostFormValue("income")
	data.Purpose = r.PostFormValue("purpose")
	data.PayType.PayTypeId, _ = strconv.Atoi(r.PostFormValue("pay_type_id"))
	data.Price, _ = strconv.ParseFloat(r.PostFormValue("price"), 64)
	data.Detail = r.PostFormValue("detail")
	data.TradeTime = r.PostFormValue("trade_time")

	//向网页传输的数据
	var page model.Page
	_, page.PayType = dao.GetPayMethod() //获取具体的支付方式

	//检查数据是否已经存在,存在则为true
	recode := dao.CheckRecode(data.TradeTime, data.Price)
	if recode {
		page.Msg = "添加失败,数据已经存在"
	} else {
		err := dao.AddPayRecode(data) //将数据保存到数据库
		if err != nil {
			page.Msg = "添加失败"
		} else {
			page.Msg = "添加成功"
		}
	}

	t := template.Must(template.ParseFiles("src/Demo/goweb/views/pages/addPayPage.html"))
	t.Execute(w, page)
}

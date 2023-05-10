package server

import (
	"Demo/goweb/dao"
	"Demo/goweb/model"
	"fmt"
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("src/Demo/goweb/views/index.html"))
	t.Execute(w, "")
}

func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("输入的用户名和密码：", username, password)

	var Page model.Page

	//用户名和密码
	Page.User, _ = dao.CheckLogin(username, password)

	//消费记录
	_, Page.ConsumptionRecode = dao.PayMonth()

	if Page.User != nil {
		//用户名和密码正确
		fmt.Println(Page.User.Nickname, Page.User.Username, Page.User.Password)
		t := template.Must(template.ParseFiles("src/Demo/goweb/views/pages/main.html"))
		_ = t.Execute(w, Page)
	} else {
		//用户名和密码不正确
		t := template.Must(template.ParseFiles("src/Demo/goweb/views/index.html"))
		_ = t.Execute(w, "用户名或密码错误！")

	}

}

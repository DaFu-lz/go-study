package server

import (
	"fmt"
	"html/template"
	"net/http"

	"Demo/work_end/usedb"
)

//Login 处理用户登录的函数
func Login(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("输入的用户名和密码：", username, password)
	//调用usedb中的方法
	user, _ := usedb.CheckUserNameAndPassWord(username, password)

	if user.ID > 0 {
		//用户名和密码正确

		t := template.Must(template.ParseFiles("src/Demo/work_end/views/pages/loginsuccess.html"))
		t.Execute(w, "")
	} else {
		//用户名和密码不正确
		t := template.Must(template.ParseFiles("src/Demo/work_end/views/pages/login.html"))
		t.Execute(w, "")
	}
}

//Login 处理用户注册的函数
func Register(w http.ResponseWriter, r *http.Request) {
	//获取用户名和密码
	username := r.PostFormValue("username")
	password := r.PostFormValue("password")
	fmt.Println("输入的用户名和密码：", username, password)

	//调用usedb中的方法
	user, _ := usedb.CheckUserName(username)
	fmt.Println("验证：", username, password)
	if username == "" {
		t := template.Must(template.ParseFiles("work_end/views/pages/register.html"))
		t.Execute(w, "")
	} else {
		if user.ID > 0 {
			//用户名存在 不能注册
			t := template.Must(template.ParseFiles("work_end/views/pages/register.html"))
			t.Execute(w, "")
		} else {
			//用户名可用，并将用户保存到数据库
			usedb.SaveUser(username, password)
			t := template.Must(template.ParseFiles("work_end/views/pages/login.html"))
			t.Execute(w, "")
		}
	}

}

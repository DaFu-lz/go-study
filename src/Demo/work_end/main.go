package main

import (
	"Demo/work_end/server"
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("work_end/views/index.html"))
	t.Execute(w, nil)
}

func main() {

	//解析静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/Demo/work_end/views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/Demo/work_end/views/pages/"))))

	//去登录
	http.HandleFunc("/Login", server.Login)
	//去注册
	http.HandleFunc("/register", server.Register)
	//主页
	http.HandleFunc("/", IndexHandler)
	//文章列表 Article list
	//文章详情 //Article details
	//文章评论//Article review
	//用户管理//User management
	//文章管理//Article management
	//评论管理//Comment management

	http.ListenAndServe(":8089", nil)
}

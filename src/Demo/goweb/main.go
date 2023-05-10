package main

/*import (
	"Demo/goweb/dao"
	"fmt"
)
//测试模块
func main()  {
	recode := dao.CheckRecode("2023-04-13 14:59:01", 21.89)
	fmt.Println(recode)
}*/

import (
	"Demo/goweb/server"
	"net/http"
)

func main() {
	//解析静态资源
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("src/Demo/goweb/views/static/"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("src/Demo/goweb/views/pages/"))))

	//登录页面
	http.HandleFunc("/", server.Index)

	//主页面
	http.HandleFunc("/Login", server.Login)

	//添加消费记录页面
	http.HandleFunc("/AddPayPage", server.AddPayPage)
	http.HandleFunc("/AddPayRecode", server.AddPayRecode) //添加消费记录

	http.ListenAndServe(":8080", nil)
}

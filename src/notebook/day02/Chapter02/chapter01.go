package main

import "fmt"

func main() {
	//定义变量，不设置初始值
	var a int
	a = 10

	//定义变量并设置初始值
	var b int = 20

	//批量定义变量，根据情况设置初始值
	var (
		_ string
		c int = 30
	)

	//多个变量同一类型
	var d, e int
	d, e = 40, 50

	//定义变量并赋值，通过数值设置数据类型
	f := 60

	fmt.Printf("a=%d,b=%d,c=%d,d=%d,e=%d,f=%d", a, b, c, d, e, f)
}

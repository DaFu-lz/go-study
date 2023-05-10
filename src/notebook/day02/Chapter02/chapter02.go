package main

import "fmt"

func main() {
	//定义单个常量方式一
	const a int = 10

	//定义单个常量方式二
	const b = 20

	//定义多个常量方式一
	const (
		c int = 30
		d     = "golang"
	)

	//定义多个常量方式二
	const e, f = true, 40

	fmt.Print(a, b, c, d, e, f)
}

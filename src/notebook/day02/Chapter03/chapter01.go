package main

import "fmt"

func main() {
	var x, y = 8, 5
	fmt.Printf("加法：%d\n", x+y)
	fmt.Printf("减法：%d\n", x-y)
	fmt.Printf("乘法：%d\n", x*y)
	fmt.Printf("除法：%d\n", x/y)
	fmt.Printf("求余：%d\n", x%y)
	x++
	fmt.Printf("自增x=：%d\n", x)
	y--
	fmt.Printf("自减y=：%d\n", y)
}

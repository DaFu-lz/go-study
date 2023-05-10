package main

import "fmt"

func main() {
	var x, y = 8, 5
	fmt.Printf("x的值为：%d，y的值为：%d\n", x, y)
	fmt.Printf("x等于y：%v\n", x == y)
	fmt.Printf("x不等于y：%v\n", x != y)
	fmt.Printf("x大于y：%v\n", x > y)
	fmt.Printf("x小于y：%v\n", x < y)
	fmt.Printf("x大于等于y：%v\n", x >= y)
	fmt.Printf("x小于等于y：%v\n", x <= y)
}

package main

import "fmt"

func main() {
	//定义变量name、age、addr
	//用于存储用户输入的数据
	var name, age, addr string

	//输出操作提示
	fmt.Printf("Please enter your first name:\n")
	//存储用户输入的数据
	_, _ = fmt.Scanln(&name)

	//输出操作提示
	fmt.Printf("Please enter you age:\n")
	//存储用户输入的数据
	_, _ = fmt.Scanln(&age)

	//输出操作提示
	fmt.Printf("Please enter you address:\n")
	//存储用户输入的数据
	_, _ = fmt.Scanln(&addr)

	//输出用户输入的所有数据
	fmt.Printf("your first name : %v , age is : %v ,address : %v", name, age, addr)
}

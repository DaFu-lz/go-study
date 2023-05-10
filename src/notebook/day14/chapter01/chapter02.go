package main

import (
	"fmt"
	"os"
)

func main() {

	//获取当前目录
	getwd, _ := os.Getwd()
	fmt.Printf("获取当前目录：%v\n", getwd)
}

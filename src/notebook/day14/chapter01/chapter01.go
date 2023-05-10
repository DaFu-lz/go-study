package main

import (
	"fmt"
	"os"
)

func main() {

	//获取主机名
	hostname, _ := os.Hostname()
	fmt.Printf("获取主机名：%v\n", hostname)

	//获取用户ID
	fmt.Printf("获取用户ID:%v\n", os.Getuid())

	//获取有效用户ID
	fmt.Printf("获取有效用户ID:%v\n", os.Geteuid())

	//获取组ID
	fmt.Printf("获取组ID:%v\n", os.Getgid())

	//获取有效组ID
	fmt.Printf("获取有效组ID:%v\n", os.Getegid())

	//获取进程ID
	fmt.Printf("获取进程ID:%v\n", os.Getpid())

	//获取父进程ID
	fmt.Printf("获取父进程ID:%v\n", os.Getppid())

	//获取某个环境变量的值
	fmt.Printf("获取某个环境变量的值:%v\n", os.Getenv("GOPATH"))

	//设置某个环境变量的值
	err := os.Setenv("TEST", "test")
	fmt.Printf("获取某个环境变量的值:%v\n", err)

	//删除某个环境变量的值
	res1 := os.Unsetenv("TEST")
	fmt.Printf("删除某个环境变量的值:%v\n", res1)

	//获取所有环境变量
	for _, e := range os.Environ() {
		fmt.Printf("环境变量：%v\n", e)
	}
	//获取某个环境变量
	fmt.Printf("获取某个环境变量:%v\n", os.Getenv("GOPATH"))

	//删除所有环境变量
	//os.Clearenv()
}

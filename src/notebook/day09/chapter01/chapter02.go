package main

import "fmt"

//定义接口
type action interface {
	//没有参数也没有返回值
	walk()

	//没有参数，有返回值
	runs() (int, int)

	//有参数，没有返回值
	speak(content string, speed int)

	//有参数也有返回值
	rest(sleepTime int) int
}

//定义结构体
type cat struct {
	name string
}

//定义结构体方法的逻辑功能
func (c *cat) walk() {
	fmt.Printf("%v在散步\n", c.name)
}
func (c *cat) runs() (int, int) {
	fmt.Printf("%v在跑步\n", c.name)
	speed := 10
	time := 1
	return speed, time
}
func (c *cat) speak(content string, speed int) {
	fmt.Printf("%v在说话：%v,语速：%v\n", c.name, content, speed)
}

//func (c *cat) rest(sleepTime int) int {
//	fmt.Printf("%v在休息，入睡时间：%v小时\n", c.name, sleepTime)
//	return sleepTime
//}

func main() {
	//定义接口变量
	var a action
	//结构体实例化
	c := cat{name: "kitty"}
	//结构体实例化变量的指针赋值给接口变量
	a = &c

	//调用接口中的方法
	a.walk()
	speed, time := a.runs()
	fmt.Printf("跑步的速度：%v，跑步时间：%v\n", speed, time)
	a.speak("喵喵", 2)
	sleepTime := a.rest(10)
	fmt.Printf("入睡时间：%v小时\n", sleepTime)
}

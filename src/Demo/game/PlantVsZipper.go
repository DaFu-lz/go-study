package main

import (
	"Utils/MouseUtils"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"time"
	//"time"
)

func main() {
	fmt.Println("start........")

	//当点击了鼠标右键开始跑程序
	MouseUtils.MouseEvent("right")

	//监测是否点击空格键，点了结束程序
	go ending()

	x, y := MouseUtils.GetMousePosition()
	MouseUtils.MoveMousePosition(x, y+200)
	for i := 1; i <= 1000000; i++ {
		fmt.Println(i)
		MouseUtils.MoveMouseRelative(20, 0)
		MouseUtils.ClickLeftMouse()
		time.Sleep(70 * time.Millisecond)
		if i%20 == 0 {
			MouseUtils.MoveMousePosition(x, y)
			MouseUtils.ClickLeftMouse()
			time.Sleep(90 * time.Millisecond)
			MouseUtils.MoveMousePosition(x, y+200)
			MouseUtils.ClickLeftMouse()
		}
	}
	MouseUtils.ClickRightMouse()

	fmt.Println("ending........")
}

/*
	结束程序
*/
func ending() {
	MouseUtils.KeyCodeEvent(robotgo.Space)
	fmt.Println("ending........")
	os.Exit(0)
}

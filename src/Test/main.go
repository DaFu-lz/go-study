package main

import (
	"Utils/MouseUtils"
	"fmt"
	"time"
)

func main() {
	fmt.Println("start........")

	x, y := MouseUtils.GetMousePosition()
	MouseUtils.MoveMousePosition(x, y+200)
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
		MouseUtils.MoveMouseRelative(20, 0)
		MouseUtils.ClickLeftMouse()
		time.Sleep(70 * time.Millisecond)
		if i%20 == 0 {
			MouseUtils.MoveMousePosition(840, 355)
			MouseUtils.ClickLeftMouse()
			time.Sleep(90 * time.Millisecond)
			MouseUtils.MoveMousePosition(849, 555)
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
	MouseUtils.MouseEvent("")
}

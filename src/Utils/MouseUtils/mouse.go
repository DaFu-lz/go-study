package MouseUtils

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

/**
获取鼠标的坐标信息
返回值：x,y
*/
func GetMousePosition() (int, int) {
	x, y := robotgo.GetMousePos()
	fmt.Printf("鼠标的坐标信息是(x:%d,y:%d)\n", x, y)
	return x, y
}

/*
	移动鼠标的位置:将鼠标的位置移动到(x,y)
*/
func MoveMousePosition(x, y int) {
	robotgo.Move(x, y)
	fmt.Printf("鼠标的位置移动到(x:%d,y:%d)的位置\n", x, y)
}

/*
	移动鼠标的位置:将鼠标的位置移动(x,y)
*/
func MoveMouseRelative(x int, y int) {
	robotgo.MoveRelative(x, y)
	if x >= 0 {
		if y >= 0 {
			fmt.Println("鼠标向右移动:", x, "px，向下移动:", y, "px。")
		} else {
			fmt.Println("鼠标向右移动:", x, "px，向上移动:", y, "px。")
		}
	} else {
		if y >= 0 {
			fmt.Println("鼠标向左移动:", x, "px，向下移动:", y, "px。")
		} else {
			fmt.Println("鼠标向左移动:", x, "px，向上移动:", y, "px。")
		}
	}
}

/*
	单击鼠标左键
*/
func ClickLeftMouse() {
	robotgo.Click()
	fmt.Println("点击了鼠标左键")
}

/*
	双击鼠标左键
*/
func ClickDoubleLeftMouse() {
	robotgo.Click()
	robotgo.Click()
	fmt.Println("双击了鼠标左键")
}

/*
	单击鼠标右键
*/
func ClickRightMouse() {
	robotgo.Click("right")
	fmt.Println("点击了鼠标右键")
}

/*
	双击鼠标右键
*/
func ClickDoubleRightMouse() {
	robotgo.Click("right")
	robotgo.Click("right")
	fmt.Println("双击了鼠标右键")
}

/*
	监控鼠标：被点击了返回true
*/
func MouseEvent(mouse string) bool {
	result := hook.AddMouse(mouse)
	fmt.Printf("你点击了鼠标%v键\n", mouse)
	return result
}

/*
	监控键盘：被点击了返回true
*/
func KeyCodeEvent(keycode string) bool {
	result := hook.AddEvent(keycode)
	fmt.Printf("你点击键盘%v键\n", keycode)
	return result
}

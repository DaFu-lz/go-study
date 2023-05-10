package main

import "C"
import (
	"fmt"
	"os"
)

func main() {
	//MouseUtils.ScrollMouse(100,100)
	//MouseUtils.GetMousePosition()

	//fmt.Println(robotgo.CheckMouse(robotgo.Left))
	//fmt.Println(robotgo.CheckMouse("right"))
	//fmt.Println(robotgo.CheckMouse("center"))

	//err := robotgo.ActiveName("chrome")
	//fmt.Println(err)

	//fmt.Println(robotgo.GetPid())
	//name, err := robotgo.FindName(21428)
	//fmt.Println(name)
	//fmt.Println(err)

	//pid := robotgo.GetPid()
	//fmt.Println(pid)
	//ids, err2 := robotgo.FindIds("goland64.exe")
	//fmt.Println(ids)
	//fmt.Println(err2)

	//err := robotgo.ActiveName("goland64.exe")
	//fmt.Println(err)

	process, err := os.StartProcess("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe",
		nil, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	//process, err := os.StartProcess("C:\\Program Files\\Google\\Chrome\\Application\\chrome.exe", nil, nil)
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}
	fmt.Println(process)
	fmt.
}

package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
)

/**
软件打开返回true
pid：父进程id
*/
func checkProcess(pid int) bool {
	_, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println("dead")
		return false
	} else {
		fmt.Println("alive")
		return true
	}
}

/*
	获取进程id
*/
func getPid(name string) []int {
	ids, err := robotgo.FindIds(name)
	if err != nil {
		return nil
	} else {
		return ids
	}
}

/*
	通过进程id打开窗口,
	成功返回true
*/
func openWindow(id int) {
	err := robotgo.ActivePid(id)
	if err != nil {
		fmt.Println("打开进程失败")
	} else {
		fmt.Println("打开进程成功")
	}
}

/*
	获取所有进程名:根据进程id获取进程名
*/
func getAllProcessNames() map[int]string {
	var info = map[int]string{}
	ids, _ := robotgo.Pids()
	for _, v := range ids {
		name, _ := robotgo.FindName(v)
		info[v] = name
	}
	return info
}

/*
	软件定时器，定时打开软件
*/

func main() {
	//names := getAllProcessNames()
	//for name, id := range names {
	//	fmt.Println(name, ":", id)
	//}
	//fmt.Println(len(names))
	//
	//pid := getPid("chrome.exe")
	//fmt.Println("进程id有：", pid)
	//for _, id := range pid {
	//	openWindow(id)
	//}
	//
	//ids, err := robotgo.FindIds("chrome.exe")
	//fmt.Println(ids,err)

	process, err := os.StartProcess("D:\\AppData\\Program Files\\ynote-desktop\\有道云笔记.exe",
		nil, &os.ProcAttr{Files: []*os.File{os.Stdin, os.Stdout, os.Stderr}})
	if err != nil {
		fmt.Println("fail")
	} else {
		fmt.Println("success")
	}
	fmt.Println(process)
}

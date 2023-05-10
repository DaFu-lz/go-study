package main

import (
	"fmt"
	"math/rand"
	"time"
)

var count, trueNum, falseNum int

func letter() (leter1, leter2, leter3, leter4 string) {
	var leter = [...]string{"a", "b", "c", "d", "e", "f", "g",
		"h", "i", "j", "k", "l", "m", "n",
		"o", "p", "q", "r", "s", "t", "u",
		"v", "w", "x", "y", "z"}

	rand.Seed(time.Now().UnixNano())
	var num1 = rand.Intn(26)
	var num2 = rand.Intn(26)
	var num3 = rand.Intn(26)
	var num4 = rand.Intn(26)

	return leter[num1], leter[num2], leter[num3], leter[num4]
}

func judge() (exit bool) {
	leter1, leter2, leter3, leter4 := letter()
	fmt.Println("请输入下面的字母(输入exit退出)：" + leter1 + "--" + leter2 + "--" + leter3 + "--" + leter4)
	var str string
	fmt.Scan(&str)
	var newWord = leter1 + leter2 + leter3 + leter4
	if str == newWord {
		count = count + 1
		trueNum = trueNum + 1
		fmt.Println("true,当前积分：", count, " ，你输对了", trueNum, "个,输错了", falseNum, "个.")
	} else if str == "exit" {
		return true
	} else {
		count = count - 1
		falseNum = falseNum + 1
		fmt.Println("false,当前积分:", count, " ，你输对了", trueNum, "个,输错了", falseNum, "个.")
	}
	return false
}

func main() {
	for {
		var exit = judge()
		if exit {
			break
		}
	}
}

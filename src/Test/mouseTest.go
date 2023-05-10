package main

import (
	"Utils/MouseUtils"
	"fmt"
)

func main() {
	fmt.Println("start")
	MouseUtils.MouseEvent("left")
	fmt.Println("ending")
}

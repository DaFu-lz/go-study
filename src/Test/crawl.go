package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("1")
	os.Exit(1)
	fmt.Println("2")
	os.Exit(1)
	fmt.Println("3")
	fmt.Println("4")
}

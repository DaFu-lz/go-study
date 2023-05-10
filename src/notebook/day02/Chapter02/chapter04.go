package main

import "fmt"

func main() {
	const (
		_          = iota
		KB float64 = 1 << (10 * iota)
		MB
		GB
		TB
	)
	fmt.Printf("B转KB的进制为：%.0f\n", KB)
	fmt.Printf("B转MB的进制为：%.0f\n", MB)
	fmt.Printf("B转GB的进制为：%.0f\n", GB)
	fmt.Printf("B转TB的进制为：%.0f\n", TB)
}

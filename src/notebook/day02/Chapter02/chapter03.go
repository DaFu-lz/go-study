package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c = 10
		d
		e
		f = iota
	)
	fmt.Print(a, b, c, d, e, f)
}

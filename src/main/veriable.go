package main

import "fmt"

func main() {
	const (
		x = 10
		y = 3.14
		z = "Kenny"
	)

	const (
		a = iota
		b = iota
		c = iota
	)

	const (
		d = 1
		e
		f
	)

	const (
		g = iota
		h
		i
	)

	fmt.Println(x)
	fmt.Println(y)
}

package main

import "fmt"

func main()  {
	var x int32 = 10
	var y int16 = 20
	// mismatched types error
	// fmt.Println(x + y)
	fmt.Println(int16(x) + y)
	fmt.Println(x + int32(y))

	var z int32 = 10
	fmt.Println(z + 20)
}


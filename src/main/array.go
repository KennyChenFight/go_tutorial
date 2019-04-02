package main

import "fmt"

func main() {
	var score [10]int
	score[0] = 90
	score[1] = 89
	fmt.Println(score)
	fmt.Println(len(score))

	arr1 := [3]int{1, 2, 3}
	arr2 := [5]int{1, 2, 3}
	arr3 := [...]int{1, 2, 3, 4, 5}
	arr4 := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(arr4)

	arr5 := [...]int{1, 2, 3}
	arr6 := arr5
	fmt.Println(arr5)
	fmt.Println(arr6)
	arr5[0] = 10
	fmt.Println(arr5)
	fmt.Println(arr6)

	var arr7 [2][3]int
	fmt.Println(arr7)
}

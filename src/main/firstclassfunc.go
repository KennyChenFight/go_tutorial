package main

import (
	"fmt"
	"reflect"
)

type BiFunc1 func(int, int) int   // 定義了新型態
type Bifunc2 = func(int, int) int // 型態別名宣告
type Predicate = func(int) bool

func main() {
	MaxDemo()
	fmt.Println("-----------")
	MaxDemo2()
	fmt.Println("-----------")
	MaxDemo3()
	fmt.Println("-----------")
	MaxDemo4()
	fmt.Println("-----------")
	MaxDemo5()
	fmt.Println("-----------")
	FilterDemo()
}

func MaxDemo() {
	maximum := max
	fmt.Println(max(10, 5))
	fmt.Println(reflect.TypeOf(max))
	fmt.Println(maximum(10, 5))
	fmt.Println(reflect.TypeOf(maximum))
}

func MaxDemo2() {
	var maximum func(int, int) int
	fmt.Println(maximum)

	maximum = max
	fmt.Println(maximum(10, 5))
}

func MaxDemo3() {
	var maximum BiFunc1
	fmt.Println(maximum)

	maximum = max
	fmt.Println(maximum(10, 5))
}

func MaxDemo4() {
	var maximum Bifunc2
	fmt.Println(maximum)

	maximum = max
	fmt.Println(maximum(10, 5))
}

func MaxDemo5() {
	var maximum Bifunc2
	fmt.Println(&maximum)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FilterDemo() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(filter(data, greatThan5))
	fmt.Println(filter(data, lessEqualsThan6))
}

func filter(origin []int, predicate Predicate) []int {
	filtered := []int{}
	for _, elem := range origin {
		if predicate(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

func greatThan5(elem int) bool {
	return elem > 5
}

func lessEqualsThan6(elem int) bool {
	return elem <= 6
}

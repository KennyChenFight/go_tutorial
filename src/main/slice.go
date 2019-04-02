package main

import (
	"fmt"
	"reflect"
)

func main() {
	s1 := make([]int, 5)
	s2 := s1
	fmt.Println(s1)
	fmt.Println(s2)
	s1[0] = 1
	fmt.Println(s1)
	fmt.Println(s2)
	s1[1] = 2
	fmt.Println(s1)
	fmt.Println(s2)

	s := []int{1, 2, 3, 4, 5}
	a := [...]int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(s))
	fmt.Println(reflect.TypeOf(a))

	slice := []int{10, 20, 30, 10: 100, 20: 200}
	fmt.Println(slice)

	SliceDemo1()
	fmt.Println("----------")
	SliceDemo2()
	fmt.Println("----------")
	SliceDemo3()
	fmt.Println("----------")
	SliceDemo4()
}

func SliceDemo1() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice := arr[1:4]
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	fmt.Println(slice)
	fmt.Println(arr)

	slice[0] = 20
	fmt.Println(slice)
	fmt.Println(arr)
}

func SliceDemo2() {
	slice := make([]int, 5, 10)
	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
}

func SliceDemo3() {
	arr := [...]int{1, 2, 3, 4, 5}
	slice1 := arr[0:2:4]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))
}

func SliceDemo4() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := arr[:5]
	slice2 := slice1[:3]

	fmt.Println(slice1)
	fmt.Println(slice2)

	slice2[0] = 10
	fmt.Println(arr)
	fmt.Println(slice1)
	fmt.Println(slice2)
}

func AppendDemo() {

}

func CopyDemo() {

}

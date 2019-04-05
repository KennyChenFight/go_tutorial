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
	fmt.Println("----------")
	AppendDemo()
	fmt.Println("----------")
	AppendDemo2()
	fmt.Println("----------")
	AppendDemo3()
	fmt.Println("----------")
	CopyDemo()
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
	arr := [...]int{1, 2, 3, 4, 5}
	slice1 := arr[:2]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := append(slice1, 6)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2[0] = 10
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(arr)
}

func AppendDemo2() {
	arr := [...]int{1, 2, 3, 4, 5}
	slice1 := arr[:]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	slice2 := append(slice1, 6)
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice2[0] = 10
	fmt.Println(slice1)
	fmt.Println(slice2)
	fmt.Println(arr)
}

func AppendDemo3() {
	slice1 := []int{1, 2, 3}
	slice2 := []int{4, 5, 6}
	fmt.Println(append(slice1, slice2...))
}

func CopyDemo() {
	src := []int{1, 2, 3, 4, 5}
	dest := make([]int, len(src), (cap(src)+1)*2)
	fmt.Println(copy(dest, src))
	fmt.Println(src)
	fmt.Println(dest)

	src[0] = 10
	fmt.Println(src)
	fmt.Println(dest)
}

package main

import (
	"errors"
	"fmt"
	"strings"
)

func main() {
	GcdDemo()
	fmt.Println("----------")
	GcdDemo2()
	fmt.Println("----------")
	FirstMatchDemo()
	fmt.Println("----------")
	DivDemo()
	fmt.Println("----------")
	SumDemo()
	fmt.Println("----------")
	SumDemo2()
	fmt.Println("----------")
	PointerDemo()
	fmt.Println("----------")
	PointerDemo2()
}

func GcdDemo() {
	fmt.Printf("Gcd of 10 and 4: %d\n", Gcd(10, 4))
}

func Gcd(m, n int) int {
	if n == 0 {
		return m
	} else {
		return Gcd(n, m%n)
	}
}

func GcdDemo2() {
	fmt.Printf("Gcd of 10 and 4: %d\n", Gcd2(10, 4))
}

func Gcd2(m, n int) (gcd int) {
	if n == 0 {
		gcd = m
	} else {
		gcd = Gcd2(n, m%n)
	}
	return
}

func FirstMatchDemo() {
	names := []string{"Kenny", "Nicole", "Jack"}
	if index, name := FirstMatch(names, "Kenny"); index == -1 {
		fmt.Println("找不到任何東西")
	} else {
		fmt.Printf("在索引 %d 找到 \"%s\"\n", index, name)
	}
}

func FirstMatch(elems []string, substr string) (int, string) {
	for index, elem := range elems {
		if strings.Contains(elem, substr) {
			return index, elem
		}
	}
	return -1, ""
}

func DivDemo() {
	if result, err := Div(10, 5); err == nil {
		fmt.Printf("10 / 5 = %d\n", result)
	} else {
		fmt.Println(err)
	}
}

func Div(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("division by zero")
	}
	return x / y, nil
}

func SumDemo() {
	fmt.Println(Sum(1, 2))
	fmt.Println(Sum(1, 2, 3))
	fmt.Println(Sum(1, 2, 3, 4))
	fmt.Println(Sum(1, 2, 3, 4, 5))
}

func SumDemo2() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Sum(numbers...))
}

func Sum(numbers ...int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func PointerDemo() {
	number := 1
	add1To(number)
	fmt.Println(number)
}

func PointerDemo2() {
	number := 1
	add1To2(&number)
	fmt.Println(number)
}

func add1To(n int) {
	n = n + 1
}

func add1To2(n *int) {
	*n = *n + 1
}

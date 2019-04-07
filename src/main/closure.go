package main

import "fmt"

type Predicate2 = func(int) bool
type Func1 = func(int) int
type Consumer = func(int)
type Getter = func() int
type Setter = func(int)

func main() {
	FilterDemo2()
	fmt.Println("----------")
	FuncADemo()
	fmt.Println("----------")
	ForEachDemo()
	fmt.Println("----------")
	GetterSetterDemo()
}

func FilterDemo2() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(filter2(data, func(elem int) bool {
		return elem > 5
	}))
	fmt.Println(filter2(data, func(elem int) bool {
		return elem <= 6
	}))
}

func filter2(origin []int, predicate Predicate2) []int {
	filtered := []int{}
	for _, elem := range origin {
		if predicate(elem) {
			filtered = append(filtered, elem)
		}
	}
	return filtered
}

func FuncADemo() {
	fmt.Println(funcA()(2))
}

func funcA() Func1 {
	x := 10
	return func(n int) int {
		return x + n
	}
}

func forEach(elems []int, consumer Consumer) {
	for _, elem := range elems {
		consumer(elem)
	}
}

func ForEachDemo() {
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	forEach(numbers, func(elem int) {
		sum += elem
	})
	fmt.Println(sum)
}

func xGetterSetter(x int) (Getter, Setter) {
	fmt.Printf("the parameter:\tx (%p) = %d\n", &x, x)

	getter := func() int {
		fmt.Printf("getter invoked:\tx (%p) = %d\n", &x, x)
		return x
	}
	setter := func(n int) {
		x = n
		fmt.Printf("setter invoked:\tx (%p) = %d\n", &x, x)
	}
	return getter, setter
}

func GetterSetterDemo() {
	getx, setx := xGetterSetter(10)

	fmt.Println(getx())
	setx(20)
	fmt.Println(getx())
}

package main

import "fmt"

type Pig struct {
	name string
}

func main() {
	TypeDemo()
	fmt.Println("----------")
	TypeDemo2()
	fmt.Println("----------")
	TypeDemo3()
	fmt.Println("----------")
	TypeDemo4()
}

func TypeDemo() {
	values := [...]interface{}{
		Pig{"Kenny"},
		Pig{"Nicole"},
	}

	for _, value := range values {
		pig := value.(Pig)
		fmt.Println(pig.name)
	}
}

func TypeDemo2() {
	values := [...]interface{}{
		Pig{"Kenny"},
		Pig{"Nicole"},
		[...]int{1, 2, 3, 4, 5},
		map[string]int{"Kenny": 123456, "Nicole": 456789},
	}
	for _, value := range values {
		if pig, ok := value.(Pig); ok {
			fmt.Println(pig.name)
		}
	}
}

func TypeDemo3() {
	values := [...]interface{}{
		Pig{"Kenny"},
		Pig{"Nicole"},
		[...]int{1, 2, 3, 4, 5},
		map[string]int{"Kenny": 123456, "Nicole": 456789},
		10,
	}

	for _, value := range values {
		if pig, ok := value.(Pig); ok {
			fmt.Println(pig.name)
		} else if arr, ok := value.([5]int); ok {
			fmt.Println(arr)
		} else if passwds, ok := value.(map[string]int); ok {
			fmt.Println(passwds)
		} else {
			fmt.Println("非預期之型態")
		}
	}
}

func TypeDemo4() {
	values := [...]interface{}{
		Pig{"Kenny"},
		Pig{"Nicole"},
		[...]int{1, 2, 3, 4, 5},
		map[string]int{"Kenny": 123456, "Nicole": 456789},
		10,
	}

	for _, value := range values {
		switch v := value.(type) {
		case Pig:
			fmt.Println(v.name)
		case [5]int:
			fmt.Println(v[0])
		case int:
			fmt.Println(v)
		default:
			fmt.Println("非預期之型態")
		}
	}
}

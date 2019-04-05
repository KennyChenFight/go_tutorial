package main

import "fmt"

func main() {
	BreakDemo()
	fmt.Println("----------")
	BreakDemo2()
	fmt.Println("----------")
	ContinueDemo()
	fmt.Println("----------")
	ContinueDemo2()
	fmt.Println("----------")
	GotoDemo()
}

func BreakDemo() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("i = %d\n", i)
	}
}

func BreakDemo2() {
BACK:
	for j := 1; j < 10; j++ {
		for i := 1; i < 10; i++ {
			if i == 5 {
				break BACK
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
		fmt.Println("test")
	}
}

func ContinueDemo() {
	for i := 1; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("i = %d\n", i)
	}
}

func ContinueDemo2() {
BACK:
	for j := 1; j < 10; j++ {
		for i := 1; i < 10; i++ {
			if i == 5 {
				continue BACK
			}
			fmt.Printf("i = %d, j = %d\n", i, j)
		}
		fmt.Println("test")
	}
}

func GotoDemo() {
	var input int

RETRY:
	fmt.Printf("輸入數字")
	fmt.Scanf("%d", &input)

	if input == 0 {
		fmt.Println("除數不得為0")
		goto RETRY
	}
	fmt.Printf("100 / %d = %f\n", input, 100/float32(input))
}

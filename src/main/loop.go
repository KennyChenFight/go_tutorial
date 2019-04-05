package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	LoopDemo()
	fmt.Println("-----------")
	LoopDemo2()
	fmt.Println("-----------")
	LoopDemo3()
	fmt.Println("-----------")
	LoopDemo4()
	fmt.Println("-----------")
	LoopDemo5()
	fmt.Println("-----------")
	LoopDemo6()
}

func LoopDemo() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func LoopDemo2() {
	for j := 1; j < 10; j++ {
		for i := 2; i < 10; i++ {
			fmt.Printf("%d * %d = %2d ", i, j, i*j)
		}
		fmt.Println()
	}
}

func LoopDemo3() {
	for i, j := 0, 0; i < 10; i, j = i+1, j+1 {
		fmt.Printf("%d * %d = %2d\n", i, j, i*j)
	}
}

func LoopDemo4() {
	foo(1)
	multiplication_table()
}

func foo(i int) {
	for ; i < 10; i++ {
		fmt.Println(i)
	}
}

func multiplication_table() {
	for i, j := 2, 1; j < 10; {
		fmt.Printf("%d * %d = %2d ", i, j, i*j)
		if i == 9 {
			fmt.Println()
			j++
			i = (j+1)/j + 1
		} else {
			i++
		}
	}
}

func LoopDemo5() {
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
}

func LoopDemo6() {
	for {
		number := random(1, 10)
		fmt.Println(number)
		if number == 5 {
			break
		}
		time.Sleep(time.Second)
	}
	fmt.Println("I his 5....Orz")
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

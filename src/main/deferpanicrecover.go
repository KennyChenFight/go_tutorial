package main

import (
	"fmt"
	"os"
)

func main() {
	DeferDemo()
	fmt.Println("----------")
	DeferDemo2()
	fmt.Println("----------")
	DeferDemo3()
	fmt.Println("----------")
	DeferDemo4()
	fmt.Println("----------")
	DeferDemo5()
	fmt.Println("----------")
	DeferDemo6()
	fmt.Println("----------")
	DeferDemo7()
}

func DeferDemo() {
	defer deferredFunc()
	fmt.Println("Hello, 世界")
}

func DeferDemo2() {
	defer deferredFunc1()
	defer deferredFunc2()
	fmt.Println("Hello, 世界")
}

func DeferDemo3() {
	defer fmt.Println("deffered 1")
	defer fmt.Println("deffered 2")
	fmt.Println("Hello, 世界")
}

func DeferDemo4() {
	i := 10
	defer fmt.Println(i)
	i++
	fmt.Println(i)
}

func deferredFunc() {
	fmt.Println("deferredFunc")
}

func deferredFunc1() {
	fmt.Println("deferredFunc1")
}

func deferredFunc2() {
	fmt.Println("deferredFunc2")
}

func DeferDemo5() {
	f, err := os.Open("test.txt")

	defer func() { // 延遲執行，而且函式return前一定會執行
		if f != nil {
			f.Close()
		}
	}()

	if err != nil {
		fmt.Println(err)
	} else {
		b1 := make([]byte, 20)
		n1, err := f.Read(b1)
		if err != nil {
			fmt.Printf("%d bytes: %s\n", n1, string(b1))
			// 處理讀取的內容
		} else {
			fmt.Printf(string(b1))
		}
	}
}

func DeferDemo6() {
	f, err := os.Open("test2.txt")

	defer func() {
		if f != nil {
			f.Close()
		}
	}()

	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)

	check(err)

	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func DeferDemo7() {
	f, err := os.Open("test2.txt")

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err) // 假設是頂層的UI介面，以自己的方式呈現
		}
		if f != nil {
			if err := f.Close(); err != nil {
				panic(err) // 再拋出panic
			}
		}
	}()

	check(err)

	b1 := make([]byte, 5)
	n1, err := f.Read(b1)

	check(err)

	fmt.Printf("%d bytes: %s\n", n1, string(b1))
}

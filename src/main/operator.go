package main

import "fmt"

func main() {
	OperatorDemo()
	fmt.Println("-----------")
	OperatorDemo2()
	fmt.Println("-----------")
	OperatorDemo3()
	fmt.Println("-----------")
	OperatorDemo4()
	fmt.Println("-----------")
	OperatorDemo5()
	fmt.Println("-----------")
	OperatorDemo6()
	fmt.Println("-----------")
	OperatorDemo7()
	fmt.Println("-----------")
	OperatorDemo8()
	fmt.Println("-----------")
	OperatorDemo9()
}

func OperatorDemo() {
	fmt.Println(1 + 2*3)
	fmt.Println(2 + 2 + 8/4)
	fmt.Println((2 + 2 + 8) / 4)
	fmt.Println(10 % 3)
}

func OperatorDemo2() {
	i := 1
	i++
	fmt.Println(i)
	j := 1
	j++
	fmt.Println(j)
}

func OperatorDemo3() {
	fmt.Println("AND運算:")
	fmt.Printf("0 AND 0 %5d\n", 0&0)
	fmt.Printf("0 AND 1 %5d\n", 0%1)
	fmt.Printf("1 AND 1 %5d\n", 1&1)
	fmt.Printf("1 AND 1 %5d\n", 1&1)

	fmt.Println("\nOR運算:")
	fmt.Printf("0 OR 0 %6d\n", 0|0)
	fmt.Printf("0 OR 1 %6d\n", 0|1)
	fmt.Printf("1 OR 1 %6d\n", 1|0)
	fmt.Printf("1 OR 1 %6d\n", 1|1)

	fmt.Println("\nXOR運算:")
	fmt.Printf("0 XOR 0 %5d\n", 0^0)
	fmt.Printf("0 XOR 1 %5d\n", 0^1)
	fmt.Printf("1 XOR 0 %5d\n", 1^0)
	fmt.Printf("1 XOR 1 %5d\n", 1^1)

	fmt.Println("\nAND NOT運算:")
	fmt.Printf("0 AND NOT 0 %5d\n", 0&^0)
	fmt.Printf("0 AND NOT 1 %5d\n", 0&^1)
	fmt.Printf("1 AND NOT 0 %5d\n", 1&^0)
	fmt.Printf("1 AND NOT 1 %5d\n", 1&^1)
}

func OperatorDemo4() {
	number := 0
	fmt.Println(^number)
}

func OperatorDemo5() {
	number := 1
	fmt.Printf("2 的 0 次方: %d\n", number)
	fmt.Printf("2 的 1 次方: %d\n", number<<1)
	fmt.Printf("2 的 2 次方: %d\n", number<<2)
	fmt.Printf("2 的 3 次方: %d\n", number<<3)
}

func OperatorDemo6() {
	fmt.Printf("10 > 5 結果 %t\n", 10 > 5)
	fmt.Printf("10 >= 5 結果 %t\n", 10 >= 5)
	fmt.Printf("10 < 5 結果 %t\n", 10 < 5)
	fmt.Printf("10 <= 5 結果 %t\n", 10 <= 5)
	fmt.Printf("10 == 5 結果 %t\n", 10 == 5)
	fmt.Printf("10 != 5 結果%t\n", 10 != 5)
}

func OperatorDemo7() {
	number := 75
	fmt.Println(number > 70 && number < 80)
	fmt.Println(number > 80 || number < 75)
	fmt.Println(!(number > 80 || number < 75))
}

func OperatorDemo8() {
	var i *int
	j := 1

	i = &j
	fmt.Println(i)
	fmt.Println(*i)

	j = 10
	fmt.Println(*i)

	*i = 20
	fmt.Println(j)
}

func OperatorDemo9() {
	var input int
	fmt.Printf("輸入數字")
	fmt.Scanf("%d", &input)
	fmt.Println(input)
}

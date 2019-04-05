package main

import "fmt"

func main() {
	ConditionalDemo()
	fmt.Println("-----------")
	ConditionalDemo2()
	fmt.Println("-----------")
	ConditionalDemo3()
	fmt.Println("-----------")
	ConditionalDemo4()
	fmt.Println("-----------")
	ConditionalDemo5()
	fmt.Println("-----------")
	ConditionalDemo6()
}

func ConditionalDemo() {
	input := 10
	remain := input % 2
	if remain == 1 {
		fmt.Printf("%d 為奇數\n", input)
	} else {
		fmt.Printf("%d 為偶數\n", input)
	}
}

func ConditionalDemo2() {
	input := 10
	if remain := input % 2; remain == 1 {
		fmt.Printf("%d 為奇數\n", input)
	} else {
		fmt.Printf("%d 為偶數\n", input)
	}
}

func ConditionalDemo3() {
	var level rune
	if score := 88; score >= 90 {
		level = 'A'
	} else if score >= 80 && score < 90 {
		level = 'B'
	} else if score >= 70 && score < 80 {
		level = 'C'
	} else if score >= 60 && score < 70 {
		level = 'E'
	}
	fmt.Printf("得分等級: %c\n", level)
}

func ConditionalDemo4() {
	var level rune
	score := 88

	switch score / 10 {
	case 10, 9:
		level = 'A'
	case 8:
		level = 'B'
	case 7:
		level = 'C'
	case 6:
		level = 'D'
	default:
		level = 'E'
	}
	fmt.Printf("得分等級: %c\n", level)
}

func ConditionalDemo5() {
	var level rune

	switch score := 100; score / 10 {
	case 10:
		fmt.Println("滿分喔!")
		fallthrough
	case 9:
		level = 'A'
	case 8:
		level = 'B'
	case 7:
		level = 'C'
	case 6:
		level = 'D'
	default:
		level = 'E'
	}
	fmt.Printf("得分等級: %c\n", level)
}

func ConditionalDemo6() {
	var level rune
	score := 88
	switch {
	case score >= 90:
		level = 'A'
	case score >= 80 && score < 90:
		level = 'B'
	case score >= 70 && score < 80:
		level = 'C'
	case score >= 60 && score < 70:
		level = 'D'
	default:
		level = 'E'
	}
	fmt.Printf("得分等級: %c\n", level)
}

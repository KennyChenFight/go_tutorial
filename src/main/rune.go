package main

import "fmt"
import "strings"

func main() {
	text := "Go語言"
	cs := []rune(text)
	fmt.Printf("%c\n", cs[2])
	fmt.Println(len(cs))

	text2 := "Go語言"
	fmt.Printf("%d\n", strings.Index(text2, "語"))

	text3 := "Go語言"
	for index, runeValue := range text3 {
		fmt.Printf("%#U 位元起始位置 %d\n", runeValue, index)
	}
}

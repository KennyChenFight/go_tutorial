package main

import "fmt"

func main() {
	text1 := "Go語言"
	text2 := "Cool"
	var text3 string
	fmt.Println(text1 + text2)
	fmt.Printf("%q\n", text3)
	fmt.Println(text1 > text2)

	text4 := "Go語言" +
		"Cool"
	fmt.Println(text4)

	text5 := `Go語言
			  Cool`
	fmt.Printf("%q\n", text5)

	text6 := `Go語言\nCool`
	fmt.Println(text6)

	text7 := "Go語言"
	for i := 0; i < len(text7); i++ {
		fmt.Printf("%x ", text7[i])
	}

	text8 := "Go語言"
	bs := []byte(text8)
	bs[0] = 103
	text9 := string(bs)
	fmt.Println(text8)
	fmt.Println(text9)

	text10 := "\x47\x6f\xe8\xaa\x9e\xe8\xa8\x80"
	fmt.Println(text10)
}

package main

import "fmt"

func main() {
	arr1 := [2][3]int{[3]int{1, 2, 3}, [3]int{4, 5, 6}}
	fmt.Println(arr1)

	arr2 := [...][3]int{[...]int{1, 2, 3}, [...]int{4, 5, 6}}
	fmt.Println(arr2)

	arr3 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr3)

	arr4 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr4)

	arr5 := [...]int{1, 2, 3}
	for i := 0; i < len(arr5); i++ {
		fmt.Printf("%d\n", arr5[i])
	}
	for index, element := range arr5 {
		fmt.Printf("%d: %d\n", index, element)
	}
	for _, element := range arr5 {
		fmt.Printf("%d\n", element)
	}
}

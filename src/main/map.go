package main

import (
	"fmt"
	"sort"
)

func main() {
	MapDemo()
	fmt.Println("-----------")
	MapDemo2()
	fmt.Println("-----------")
	MapDemo3()
	fmt.Println("-----------")
	MapDemo4()
	fmt.Println("-----------")
	MapDemo5()
	fmt.Println("-----------")
	MapDemo6()
	fmt.Println("-----------")
	MapDemo7()
	fmt.Println("-----------")
	MapDemo8()
	fmt.Println("-----------")
	MapDemo9()
	fmt.Println("-----------")
	MapDemo10()
	fmt.Println("-----------")
	MapDemo11()
}

func MapDemo() {
	passwords := make(map[string]int)
	fmt.Println(passwords)
	fmt.Println(len(passwords))

	passwords["Kenny"] = 123
	passwords["Nicole"] = 456
	fmt.Println(passwords)
	fmt.Println(len(passwords))
	fmt.Println(passwords["Kenny"])
	fmt.Println(passwords["Nicole"])
}

func MapDemo2() {
	passwords := map[string]int{
		"Kenny":  123,
		"Nicole": 456,
	}

	fmt.Println(passwords)
	fmt.Println(len(passwords))
	fmt.Println(passwords["Kenny"])
	fmt.Println(passwords["Nicole"])
}

func MapDemo3() {
	passwds1 := map[string]int{"Kenny": 123}
	passwds2 := passwds1

	fmt.Println(passwds1)

	passwds2["Nicole"] = 456
	fmt.Println(passwds2)
}

func MapDemo4() {
	passwds := map[string]int{"Kenny": 123}

	v, exists := passwds["Nicole"]
	fmt.Printf("%d %t\n", v, exists)

	passwds["Nicole"] = 456
	v, exists = passwds["Nicole"]
	fmt.Printf("%d %t\n", v, exists)
}

func MapDemo5() {
	passwds := map[string]int{"Kenny": 123}
	name := "Kenny"

	_, exists := passwds[name]
	if exists {
		fmt.Printf("%s's password is %d\n", name, passwds[name])
	} else {
		fmt.Printf("No password for %s\n", name)
	}
}

func MapDemo6() {
	passwds := map[string]int{"Kenny": 123}
	fmt.Println(passwds)

	delete(passwds, "Kenny")
	fmt.Println(passwds)
}

func MapDemo7() {
	passwords := map[string]int{
		"Kenny":  123,
		"Nicole": 456,
	}

	for name, passwd := range passwords {
		fmt.Printf("%s : %d\n", name, passwd)
	}
}

func MapDemo8() {
	passwords := map[string]int{
		"Kenny":  123,
		"Nicole": 456,
	}

	for name := range passwords {
		fmt.Printf("%s\n", name)
	}
}

func MapDemo9() {
	passwords := map[string]int{
		"Kenny":  123,
		"Nicole": 456,
	}

	for _, passwd := range passwords {
		fmt.Printf("%d\n", passwd)
	}
}

func MapDemo10() {
	passwords := map[string]int{
		"Kenny":  123,
		"Nicole": 456,
	}

	fmt.Println(keys(passwords))
	fmt.Println(values(passwords))
}

func keys(m map[string]int) []string {
	ks := make([]string, 0, len(m))
	for k := range m {
		ks = append(ks, k)
	}
	return ks
}

func values(m map[string]int) []int {
	vs := make([]int, 0, len(m))
	for _, v := range m {
		vs = append(vs, v)
	}
	return vs
}

func MapDemo11() {
	passwords := map[string]int{
		"Kenny":  123,
		"Nicole": 456,
		"Jack":   789,
	}

	keys := make([]string, 0, len(passwords))
	for key := range passwords {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		fmt.Printf("%s : %d\n", key, passwords[key])
	}
}

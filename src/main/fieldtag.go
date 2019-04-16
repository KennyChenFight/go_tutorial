package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
)

func main() {
	FieldTagDemo()
	fmt.Println("---------")
	FieldTagDemo2()
	fmt.Println("---------")
	FieldTagDemo3()
	fmt.Println("---------")
	FieldTagDemo4()
}

type Customer struct {
	Name string
	City string
}

func ToJSON(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var b []string
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		b = append(b, fmt.Sprintf(`"%s": "%s"`, f.Name, v.FieldByName(f.Name)))
	}

	return fmt.Sprintf("{%s}", strings.Join(b, ","))
}

func FieldTagDemo() {
	cust := Customer{"Kenny", "Taichung"}
	fmt.Println(ToJSON(cust))
}

type User struct {
	Name string `name`
	City string `city`
}

func ToJSON_2(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var b []string
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		b = append(b, fmt.Sprintf(`"%s": "%s"`, f.Tag, v.FieldByName(f.Name)))
	}

	return fmt.Sprintf("{%s}", strings.Join(b, ","))
}

func FieldTagDemo2() {
	user := User{"Kenny", "Taichung"}
	fmt.Println(ToJSON_2(user))
}

type Person struct {
	Name string `json:"name"`
	City string `json:"city"`
}

func ToJSON_3(obj interface{}) string {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	var b []string
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fv, _ := f.Tag.Lookup("json")
		b = append(b, fmt.Sprintf(`"%s": "%s"`, fv, v.FieldByName(f.Name)))
	}

	return fmt.Sprintf("{%s}", strings.Join(b, ","))
}

func FieldTagDemo3() {
	person := Person{"Kenny", "Taichung"}
	fmt.Println(ToJSON_3(person))
}

func FieldTagDemo4() {
	person := Person{"Kenny", "Taichung"}
	b, _ := json.Marshal(person)
	fmt.Println(string(b))
}

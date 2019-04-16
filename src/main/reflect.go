package main

import (
	"errors"
	"fmt"
	"reflect"
)

func main() {
	ReflectDemo()
	fmt.Println("----------")
	ReflectDemo2()
	fmt.Println("----------")
	ReflectDemo3()
	fmt.Println("----------")
	ReflectDemo4()
	fmt.Println("----------")
	ReflectDemo5()
	fmt.Println("----------")
	ReflectDemo6()
}

type Savings interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
}

type Account struct {
	id      string
	name    string
	balance float64
}

func (account *Account) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("必須存入正數")
	}
	account.balance += amount
	return nil
}

func (account *Account) Withdraw(amount float64) error {
	if amount > account.balance {
		return errors.New("餘額不足")
	}
	account.balance -= amount
	return nil
}

func ReflectDemo() {
	account := Account{"X123", "Kenny Chen", 1000}
	t := reflect.TypeOf(account)
	fmt.Println(t.Kind())
	fmt.Println(t.String())

	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

func ReflectDemo2() {
	var savings Savings = &Account{"X123", "Kenny Chen", 1000}
	t := reflect.TypeOf(savings)

	for i, n := 0, t.NumMethod(); i < n; i++ {
		f := t.Method(i)
		fmt.Println(f.Name, f.Type)
	}

	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}

	fmt.Println(t.Kind())
	fmt.Println(t.String())
	for i, n := 0, t.NumField(); i < n; i++ {
		f := t.Field(i)
		fmt.Println(f.Name, f.Type)
	}
}

type Duck struct {
	name string
}

func ReflectDemo3() {
	values := [...]interface{}{
		Duck{"Kenny"},
		Duck{"Nicole"},
		[...]int{1, 2, 3, 4, 5},
		map[string]int{"Kenny": 123456, "Nicole": 456789},
		10,
	}

	for _, value := range values {
		switch t := reflect.TypeOf(value); t.Kind() {
		case reflect.Struct:
			fmt.Println("It's a struct.")
		case reflect.Array:
			fmt.Println("It's a array.")
		case reflect.Map:
			fmt.Println("It's a map.")
		case reflect.Int:
			fmt.Println("It's as integer.")
		default:
			fmt.Println("非預期之型態")
		}
	}
}

func ReflectDemo4() {
	x := 10
	vx := reflect.ValueOf(x)
	fmt.Printf("x = %d\n", vx.Int())

	account := Account{"X123", "Kenny Chen", 1000}
	vacct := reflect.ValueOf(account)
	fmt.Printf("id = %s\n", vacct.FieldByName("id").String())
	fmt.Printf("name = %s\n", vacct.FieldByName("name").String())
	fmt.Printf("balance = %.2f\n", vacct.FieldByName("balance").Float())
}

func ReflectDemo5() {
	x := 10
	vx := reflect.ValueOf(&x)
	fmt.Printf("x = %d\n", vx.Elem().Int())

	account := &Account{"X123", "Kenny Chen", 1000}
	vacct := reflect.ValueOf(account).Elem()
	fmt.Printf("id = %s\n", vacct.FieldByName("id").String())
	fmt.Printf("name = %s\n", vacct.FieldByName("name").String())
	fmt.Printf("balance = %.2f\n", vacct.FieldByName("balance").Float())
}

func ReflectDemo6() {
	x := 10
	vx := reflect.ValueOf(&x).Elem()
	fmt.Printf("x = %d\n", vx)

	vx.SetInt(20)
	fmt.Printf("x = %d\n", vx)
}

package main

import (
	"errors"
	"fmt"
)

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

func (account *Account) ToString() string {
	return fmt.Sprintf("Account{%s %s %.2f}",
		account.id, account.name, account.balance)
}

type Point2 struct {
	x, y int
}

func (point *Point2) ToString() string {
	return fmt.Sprintf("Point{%d %d}", point.x, point.y)
}

func main() {
	AccountDemo()
	fmt.Println("----------")
	PointDemo()
	fmt.Println("----------")
	AccountDemo2()
	fmt.Println("----------")
	AccountDemo3()
	fmt.Println("----------")
	MethodDemo()
	fmt.Println("----------")
	MethodDemo2()
	fmt.Println("----------")
}

func AccountDemo() {
	account := &Account{"1234-5678", "Kenny Chen", 1000}
	account.Deposit(500)
	account.Withdraw(200)
	fmt.Println(account.ToString())
}

func PointDemo() {
	point := &Point2{10, 20}
	fmt.Println(point.ToString())
}

func AccountDemo2() {
	deposit := (*Account).Deposit
	withdraw := (*Account).Withdraw
	toString := (*Account).ToString

	account := &Account{"1234-5678", "Kenny Chen", 1000}
	deposit(account, 500)
	withdraw(account, 200)
	fmt.Println(toString(account))

	account2 := &Account{"5678-1234", "Nicole Chen", 500}
	deposit(account2, 250)
	withdraw(account2, 150)
	fmt.Println(toString(account2))
}

func AccountDemo3() {
	account1 := &Account{"1234-5678", "Kenny Chen", 1000}
	acct1Deposit := account1.Deposit
	acct1Withdraw := account1.Withdraw
	acct1ToString := account1.ToString

	acct1Deposit(500)
	acct1Withdraw(200)
	fmt.Println(acct1ToString())

	account2 := &Account{"5678-1234", "Nicole Chen", 500}
	acct2Deposit := account2.Deposit
	acct2Withdraw := account2.Withdraw
	acct2ToString := account2.ToString

	acct2Deposit(250)
	acct2Withdraw(150)
	fmt.Println(acct2ToString())
}

type IntList []int
type Funcint func(int)

func (lt IntList) ForEach(f Funcint) {
	for _, ele := range lt {
		f(ele)
	}
}

func MethodDemo() {
	var lt IntList = []int{10, 20, 30, 40, 50}
	lt.ForEach(func(ele int) {
		fmt.Println(ele)
	})
}

type Int int
type FuncInt func(Int)

func (number Int) Times(f FuncInt) error {
	if number < 0 {
		return errors.New("必須是正數")
	}

	var i Int
	for i = 0; i < number; i++ {
		f(i)
	}
	return nil
}

func MethodDemo2() {
	var x Int = 10
	x.Times(func(n Int) {
		fmt.Println(n)
	})
}

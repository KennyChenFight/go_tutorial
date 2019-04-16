package main

import (
	"errors"
	"fmt"
)

type Account2 struct {
	id      string
	name    string
	balance float64
}

func (account *Account2) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("必須存入正數")
	}
	account.balance += amount
	return nil
}

func (account *Account2) Withdraw(amount float64) error {
	if amount > account.balance {
		return errors.New("餘額不足")
	}
	account.balance -= amount
	return nil
}

type CheckingAccount struct {
	Account2
	balance        float64
	overdraftlimit float64
}

func main() {
	inheritDemo()
	fmt.Println("----------")
	inheritDemo2()
	fmt.Println("----------")
	inheritDemo3()
	fmt.Println("----------")
	inheritDemo4()
	fmt.Println("----------")
}

func inheritDemo() {
	account := CheckingAccount{Account2{"1234-5678", "Kenny Chen", 1000}, 3000, 30000}

	fmt.Println(account)
	fmt.Println(account.id)
	fmt.Println(account.name)
	fmt.Println(account.overdraftlimit)
}

func inheritDemo2() {
	account := CheckingAccount{Account2{"1234-5678", "Kenny Chen", 1000}, 3000, 30000}

	fmt.Println(account)
	fmt.Println(account.Account2.id)
	fmt.Println(account.Account2.name)
	fmt.Println(account.Account2.balance)
	fmt.Println(account.overdraftlimit)
}

func inheritDemo3() {
	account := CheckingAccount{Account2{"1234-5678", "Kenny Chen", 1000}, 3000, 30000}

	fmt.Println(account.balance)
	fmt.Println(account.Account2.balance)
}

func inheritDemo4() {
	account := CheckingAccount{Account2{"1234-5678", "Kenny Chen", 1000}, 3000, 30000}
	account.Deposit(2000)
	account.Withdraw(500)
	fmt.Println(account)
}

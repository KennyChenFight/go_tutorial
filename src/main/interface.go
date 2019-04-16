package main

import (
	"errors"
	"fmt"
)

type Savings interface {
	Deposit(amount float64) error
	Withdraw(amount float64) error
}

type Account3 struct {
	id      string
	name    string
	balance float64
}

func (account *Account3) Deposit(amount float64) error {
	if amount <= 0 {
		return errors.New("必須存入正數")
	}
	account.balance += amount
	return nil
}

func (account *Account3) Withdraw(amount float64) error {
	if amount > account.balance {
		return errors.New("餘額不足")
	}
	account.balance -= amount
	return nil
}

func (account *Account3) String() string {
	return fmt.Sprintf("Account(id = %s, name = %s, balance = %.2f)",
		account.id, account.name, account.balance)
}

type CheckingAccount2 struct {
	Account3
	overdraftlimit float64
}

func (account *CheckingAccount2) Withdraw(amount float64) error {
	if amount > account.balance+account.overdraftlimit {
		return errors.New("超出信用額度")
	}
	account.balance -= amount
	return nil
}

func (account *CheckingAccount2) String() string {
	return fmt.Sprintf("CheckingAccount(id = %s, name = %s, balance = %.2f, overdraftlimit = %.2f)",
		account.id, account.name, account.balance, account.overdraftlimit)
}

func Withdraw(savings Savings) {
	if err := savings.Withdraw(500); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(savings)
	}
}

func main() {
	SavingsDemo()
	fmt.Println("----------")
	ArrayDemo()
	fmt.Println("----------")
	DuckDemo()
	fmt.Println("----------")
	DuckDemo2()
	fmt.Println("----------")
	StringDemo()
}

func SavingsDemo() {
	account1 := Account3{"1234-5678", "Kenny Chen", 1000}
	account2 := CheckingAccount2{Account3{"1234-5678", "Kenny Chen", 1000}, 30000}
	Withdraw(&account1)
	Withdraw(&account2)
}

func ArrayDemo() {
	savingsArray := [...]Savings{
		&Account3{"1234-5678", "Kenny Chen", 1000},
		&CheckingAccount2{Account3{"1234-5678", "Kenny Chen", 1000}, 30000},
	}

	for _, savings := range savingsArray {
		fmt.Println(savings)
	}

	savingsSlice := []Savings{
		&Account3{"1234-5678", "Kenny Chen", 1000},
		&CheckingAccount2{Account3{"1234-5678", "Kenny Chen", 1000}, 30000},
	}

	for _, savings := range savingsSlice {
		fmt.Println(savings)
	}
}

type Duck struct {
}

func (duck *Duck) Deposit(amount float64) error {
	fmt.Println("我是一隻鴨子，我沒帳戶")
	return nil
}

func (duck *Duck) Withdraw(amount float64) error {
	fmt.Println("我是一隻鴨子，我沒錢")
	return nil
}

func DuckDemo() {
	duckArray := [...]Savings{
		&Duck{},
		&Duck{},
	}
	for _, duck := range duckArray {
		duck.Deposit(1000)
	}

	duckSlice := []Savings{
		&Duck{},
		&Duck{},
	}
	for _, duck := range duckSlice {
		duck.Withdraw(500)
	}
}

func DuckDemo2() {
	instances := []interface{}{
		&Duck{},
		[...]int{1, 2, 3, 4, 5},
		map[string]int{"Kenny": 123456, "Nicole": 54321},
	}

	for _, instance := range instances {
		fmt.Println(instance)
	}
}

func StringDemo() {
	account1 := Account3{"1234-5678", "Kenny Chen", 1000}
	account2 := CheckingAccount2{Account3{"1234-5678", "Kenny Chen", 1000}, 30000}

	fmt.Println(&account1)
	fmt.Println(&account2)
}

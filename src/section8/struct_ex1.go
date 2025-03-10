package main

import "fmt"

type Account2 struct {
	number   string
	balance  float64
	interest float64
}

// 생성자 패턴이다. 아래와 같은 함수가
func NewAccount2(number string, balance float64, interest float64) *Account2 {
	return &Account2{number, balance, interest}
}

func main() {
	// 구조체 생성자 패턴 예제

	// 예제 1
	kim := Account2{number: "245-901", balance: 10000000, interest: 0.015}

	var lee *Account2 = new(Account2)
	lee.number = "245-902"
	lee.balance = 1300000
	lee.interest = 0.025

	fmt.Println("ex 1 : ", kim)
	fmt.Println("ex 1 : ", lee)

	park := NewAccount2("245-903", 17000000, 0.04)
	fmt.Println("ex 2 : ", park)

}

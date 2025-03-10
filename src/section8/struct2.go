package main

import "fmt"

type Account1 struct {
	number   string
	balance  float64
	interest float64
}

func (a Account1) Calculate() float64 {
	return a.balance + (a.balance * a.interest)
}

func main() {
	// 다양한 선언 방법
	// 선언 방법 1
	var kim *Account1 = new(Account1)
	kim.number = "245-901"
	kim.balance = 10000000
	kim.interest = 0.015

	// 선언 방법 2
	hong := &Account1{number: "245-902", balance: 15000000, interest: 0.04}

	// 선언 방법 3
	lee := new(Account1)
	lee.number = "245-903"
	lee.balance = 13000000
	lee.interest = 0.025

	fmt.Println("ex 1 : ", kim)
	fmt.Println("ex 1 : ", hong)
	fmt.Println("ex 1 : ", lee)

	fmt.Println()

	fmt.Printf("ex 2 : %#v\n", kim)
	fmt.Printf("ex 2 : %#v\n", hong)
	fmt.Printf("ex 2 : %#v\n", lee)

	fmt.Println()

	// 예제 2
	fmt.Println("ex 3 : ", int(kim.Calculate()))
	fmt.Println("ex 3 : ", int(hong.Calculate()))
	fmt.Println("ex 3 : ", int(lee.Calculate()))

}

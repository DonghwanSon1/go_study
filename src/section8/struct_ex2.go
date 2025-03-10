package main

import "fmt"

type Account3 struct {
	number   string
	balance  float64
	interest float64
}

func CalculateD(a Account3) {
	a.balance = a.balance + (a.balance * a.interest)
}

func CalculateP(a *Account3) {
	a.balance = a.balance + (a.balance * a.interest)
}

func main() {
	// 예제 1
	kim := Account3{"245-901", 10000000, 0.015}
	lee := Account3{"245-902", 20000000, 0.025}

	fmt.Println("ex 1 : ", kim)
	fmt.Println("ex 1 : ", lee)

	fmt.Println()

	CalculateD(kim)
	CalculateP(&lee)

	fmt.Println("ex 2 : ", int(kim.balance))
	fmt.Println("ex 2 : ", int(lee.balance))
}

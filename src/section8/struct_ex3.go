package main

import "fmt"

type Account4 struct {
	number   string
	balance  float64
	interest float64
}

func (a Account4) CalculateD(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func (a *Account4) CalculateF(bonus float64) {
	a.balance = a.balance + (a.balance * a.interest) + bonus
}

func main() {
	// 정리 : 구조체 인스턴스 값 변경 시 -> 포인터 전달, 보통의 경우 -> 값 전달
	// 예제 1
	kim := Account4{"245-901", 10000000, 0.015}
	lee := Account4{"245-902", 20000000, 0.025}

	fmt.Println("ex 1 : ", kim)
	fmt.Println("ex 1 : ", lee)

	fmt.Println()

	lee.CalculateD(1000000)
	kim.CalculateF(1000000)

	fmt.Println("ex 2 : ", int(lee.balance))
	fmt.Println("ex 2 : ", int(kim.balance))
}

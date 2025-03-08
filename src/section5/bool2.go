package main

import "fmt"

func main() {
	// 예제 1 - 논리 연산자
	fmt.Println("ex 1 : ", true && true)   // true
	fmt.Println("ex 1 : ", true && false)  // false
	fmt.Println("ex 1 : ", false && false) // false
	fmt.Println("ex 1 : ", true || true)   // true
	fmt.Println("ex 1 : ", true || false)  // true
	fmt.Println("ex 1 : ", false || false) // false
	fmt.Println("ex 1 : ", !true)          // false
	fmt.Println("ex 1 : ", !false)         // true

	fmt.Println()

	// 예제 2 - 비교 연산자
	num1 := 15
	num2 := 37
	fmt.Println("ex 2 : ", num1 < num2)  // true
	fmt.Println("ex 2 : ", num1 > num2)  // false
	fmt.Println("ex 2 : ", num1 >= num2) // false
	fmt.Println("ex 2 : ", num1 <= num2) // true
	fmt.Println("ex 2 : ", num1 == num2) // false
	fmt.Println("ex 2 : ", num1 != num2) // true
}

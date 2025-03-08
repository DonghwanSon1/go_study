package main

import "fmt"

func main() {
	// 데이터 타입 : 숫자형
	// 정수, 실수, 복소수
	// 32bit, 64bit, unsigned(양수)
	// 정수 : 8진수(0), 16진수(0x), 10진수

	// 참고 : byte = uint8(0 to 255)
	// 참고 : rune = int32(-2147483648 to 2147483647)

	var num1 int = 17
	var num2 int = -68
	var num3 int = 0631
	var num4 int = 0x32fa2c75

	fmt.Println("ex 1 : ", num1)
	fmt.Println("ex 1 : ", num2)
	fmt.Println("ex 1 : ", num3)
	fmt.Println("ex 1 : ", num4)

}

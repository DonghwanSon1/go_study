package main

import "fmt"

func main() {
	var a int = 50
	b := 70

	// 예제 1
	if a >= 65 {
		fmt.Println("65 이상입니다.")
	} else {
		fmt.Println("64 미만 입니다.")
	}

	if b >= 70 {
		fmt.Println("70 이상입니다.")
	} else {
		fmt.Println("70 미만입니다.")
	}

	/**
	에러 발생
	else 문도 바로 옆에 {}를 넣어줘야한다.
	엔터 후 {} 하면 에러 발생한다.
	*/

}

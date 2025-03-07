package main

import "fmt"

func main() {
	// 짧은 선언 -> 특징 = 함수 안에서만 사용 (전역 X), 선언 후 똑같이 다시 할당 하면 예외 발생
	// 주로 제한된 범위의 함수내에서 사용 할 경우 코드 가독성을 높일 수 있음.
	shortVar1 := 3
	// shortVar1 := 10 // 해당 처럼 한번 더 선언 하면 안됨
	shortVar2 := "Test"
	shortVar3 := false

	fmt.Println("shortVar1 : ", shortVar1, "shortVar2 : ", shortVar2, "shortVar3 : ", shortVar3)

	if shortVar4 := false; shortVar4 {
		fmt.Println("shortVar4 : ", shortVar4)
	}
}

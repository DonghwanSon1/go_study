package main

import "fmt"

func main() {
	// Bool(Boolean) : 참과 거짓
	// 조건부 논리 연산자랑 주로 사용 : !(not), ||(or), &&(and)
	// 암묵적 형변환 X : 0,1 이것들이 True, False 가 아니다. (0, nil -> false X)

	// 예제 1
	var b1 bool = true
	var b2 bool = false
	b3 := true

	fmt.Println("ex 1 : ", b1)
	fmt.Println("ex 2 : ", b2)
	fmt.Println("ex 3 : ", b3)
	fmt.Println("ex 4 : ", b3 == b3)
	fmt.Println("ex 5 : ", b1 && b2)
	fmt.Println("ex 6 : ", b1 || b2)
	fmt.Println("ex 7 : ", !b1)

	/**
	에러 발생
	if b4 := 1; b4 {
		fmt.Println("ex 8 : true!")
	}
	1 로 if 문을 하면 안됨, 무조건 false, true (boolean 값으로만) 가능 - 따라서 에러 발생한다.

	*/
}
